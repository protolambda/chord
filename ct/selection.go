package ct

import (
	"context"
	"fmt"
	"strings"

	"github.com/protolambda/mustbe/assertion"
)

// Selection binds queries to a scope of a [Subject]. It is an assertion that
// exactly one element matches. Other cardinalities and properties are
// available as assertion-returning methods.
//
// Selections are immutable descriptions; matching happens when an assertion
// derived from them is checked.
type Selection struct {
	subject *Subject
	parent  *Selection // nil means the document root
	query   Query
	// pick narrows the parent's matches to one position instead of
	// searching descendants; index counts from the end when negative.
	pick  bool
	index int
}

var _ assertion.Assertion = (*Selection)(nil)

// Find selects the descendants of every matched element that match every query.
func (s *Selection) Find(query Query, more ...Query) *Selection {
	return &Selection{subject: s.subject, parent: s, query: query.And(more...)}
}

// First narrows the selection to its first match, in document order.
func (s *Selection) First() *Selection {
	return &Selection{subject: s.subject, parent: s, query: Query{desc: "first"}, pick: true, index: 0}
}

// Last narrows the selection to its last match, in document order.
func (s *Selection) Last() *Selection {
	return &Selection{subject: s.subject, parent: s, query: Query{desc: "last"}, pick: true, index: -1}
}

// Nth narrows the selection to the match at the zero-based position, in
// document order. A negative position counts from the end. An out-of-range
// position matches nothing.
func (s *Selection) Nth(i int) *Selection {
	return &Selection{subject: s.subject, parent: s, query: Query{desc: fmt.Sprintf("nth(%d)", i)}, pick: true, index: i}
}

// String describes the default assertion: exactly one match.
func (s *Selection) String() string {
	return "exactlyOne(" + s.path() + ")"
}

// Check asserts that exactly one element matches.
func (s *Selection) Check(ctx context.Context) error {
	return s.countAssertion("exactlyOne", "exactly one", func(n int) bool { return n == 1 }).Check(ctx)
}

// None asserts that no element matches. The subject must still load.
func (s *Selection) None() assertion.Assertion {
	return s.countAssertion("none", "no", func(n int) bool { return n == 0 })
}

// Any asserts that at least one element matches.
func (s *Selection) Any() assertion.Assertion {
	return s.countAssertion("any", "at least one", func(n int) bool { return n > 0 })
}

// Count asserts that exactly n elements match.
func (s *Selection) Count(n int) assertion.Assertion {
	return s.countAssertion(fmt.Sprintf("count(%d)", n), fmt.Sprintf("exactly %d", n), func(got int) bool { return got == n })
}

// AtLeast asserts that n or more elements match.
func (s *Selection) AtLeast(n int) assertion.Assertion {
	return s.countAssertion(fmt.Sprintf("atLeast(%d)", n), fmt.Sprintf("at least %d", n), func(got int) bool { return got >= n })
}

// AtMost asserts that n or fewer elements match.
func (s *Selection) AtMost(n int) assertion.Assertion {
	return s.countAssertion(fmt.Sprintf("atMost(%d)", n), fmt.Sprintf("at most %d", n), func(got int) bool { return got <= n })
}

// Matches asserts that exactly one element matches the selection, and that
// this element also matches every given query.
func (s *Selection) Matches(queries ...Query) assertion.Assertion {
	return matchesAssertion{sel: s, queries: queries}
}

// Each asserts that at least one element matches the selection, and that
// every matched element also matches every given query.
func (s *Selection) Each(queries ...Query) assertion.Assertion {
	return eachAssertion{sel: s, queries: queries}
}

// InOrder asserts that exactly as many elements match as there are queries,
// and that the i-th matched element, in document order, matches the i-th
// query.
func (s *Selection) InOrder(queries ...Query) assertion.Assertion {
	return inOrderAssertion{sel: s, queries: queries}
}

// Texts asserts that the matched elements, in document order, have exactly
// the given subtree texts. Whitespace is collapsed before comparison.
func (s *Selection) Texts(values ...string) assertion.Assertion {
	return sequenceAssertion{
		sel:    s,
		name:   "texts",
		what:   "texts",
		values: values,
		get:    func(n Node) (string, bool) { return textContent(n), true },
	}
}

// AttrValues asserts that the matched elements, in document order, have
// exactly the given values for the attribute. A boolean attribute has the
// empty value; a missing attribute never matches.
func (s *Selection) AttrValues(name string, values ...string) assertion.Assertion {
	return sequenceAssertion{
		sel:    s,
		name:   fmt.Sprintf("attrValues(%q)", name),
		what:   name + " values",
		values: values,
		get:    func(n Node) (string, bool) { return n.Attr(name) },
	}
}

// Nodes loads the subject and returns the matched elements in document
// order. It exists for custom assertions; prefer the assertion methods.
func (s *Selection) Nodes(ctx context.Context) ([]Node, error) {
	return s.nodes(ctx)
}

// path describes the selection as a chain of queries, e.g.
// `label("Email") within tag("form") in chord view`.
func (s *Selection) path() string {
	switch {
	case s.parent == nil:
		return s.query.String() + " in " + s.subject.source.String()
	case s.pick:
		return s.query.String() + " of " + s.parent.path()
	default:
		return s.query.String() + " within " + s.parent.path()
	}
}

// scope resolves the nodes searched by this selection.
func (s *Selection) scope(ctx context.Context) ([]Node, error) {
	doc, err := s.subject.Load(ctx)
	if err != nil {
		return nil, err
	}
	if s.parent == nil {
		return []Node{doc.Root()}, nil
	}
	return s.parent.nodes(ctx)
}

// nodes resolves the matched elements in document order, without duplicates.
func (s *Selection) nodes(ctx context.Context) ([]Node, error) {
	scope, err := s.scope(ctx)
	if err != nil {
		return nil, err
	}
	if s.pick {
		i := s.index
		if i < 0 {
			i += len(scope)
		}
		if i < 0 || i >= len(scope) {
			return nil, nil
		}
		return []Node{scope[i]}, nil
	}
	seen := make(map[Node]struct{})
	var out []Node
	for _, sc := range scope {
		for n := range descendants(sc) {
			if n.Kind() != KindElement {
				continue
			}
			if _, dup := seen[n]; dup {
				continue
			}
			ok, err := s.query.Match(n)
			if err != nil {
				return nil, fmt.Errorf("%w: %s at %s: %w", ErrQuery, s.query, locationOf(n), err)
			}
			if ok {
				seen[n] = struct{}{}
				out = append(out, n)
			}
		}
	}
	return out, nil
}

func (s *Selection) countAssertion(name, want string, ok func(int) bool) countAssertion {
	return countAssertion{sel: s, name: name, want: want, ok: ok}
}

// countAssertion asserts a cardinality of a selection.
type countAssertion struct {
	sel  *Selection
	name string
	want string
	ok   func(n int) bool
}

var _ assertion.Assertion = countAssertion{}

func (a countAssertion) String() string {
	return a.name + "(" + a.sel.path() + ")"
}

func (a countAssertion) Check(ctx context.Context) error {
	nodes, err := a.sel.nodes(ctx)
	if err != nil {
		return err
	}
	if a.ok(len(nodes)) {
		return nil
	}
	return a.sel.countFailure(ctx, a.want, nodes)
}

// countFailure builds the diagnostic for an unexpected number of matches.
func (s *Selection) countFailure(ctx context.Context, want string, nodes []Node) error {
	var b strings.Builder
	fmt.Fprintf(&b, "expected %s match of %s, found %d", want, s.path(), len(nodes))
	if len(nodes) > 0 {
		b.WriteString("\nmatches:\n")
		b.WriteString(listNodes(nodes, s.subject.opts.redact))
	}
	if scope, err := s.scope(ctx); err == nil {
		if s.parent != nil {
			fmt.Fprintf(&b, "\nscope %s matched %d", s.parent.query, len(scope))
		}
		if len(scope) > 0 && len(nodes) == 0 {
			b.WriteString("\nscope outline:\n")
			b.WriteString(outline(scope, s.subject.opts.redact))
		}
	}
	return fmt.Errorf("%w: %s", ErrCount, b.String())
}

// matchesAssertion asserts that the single selected node matches more queries.
type matchesAssertion struct {
	sel     *Selection
	queries []Query
}

var _ assertion.Assertion = matchesAssertion{}

func (a matchesAssertion) String() string {
	return fmt.Sprintf("matches(%s, %s)", a.sel.path(), describeAll(a.queries))
}

func (a matchesAssertion) Check(ctx context.Context) error {
	nodes, err := a.sel.nodes(ctx)
	if err != nil {
		return err
	}
	if len(nodes) != 1 {
		return a.sel.countFailure(ctx, "exactly one", nodes)
	}
	n := nodes[0]
	q, err := firstMismatch(n, a.queries)
	if err != nil {
		return err
	}
	if q != nil {
		return fmt.Errorf("%w: expected %s to match %s\nnode:\n%s",
			ErrMismatch, describeNode(n, a.sel.subject.opts.redact), q, outline([]Node{n}, a.sel.subject.opts.redact))
	}
	return nil
}

// sequenceAssertion compares one extracted value per matched node, in
// document order, against an expected sequence.
type sequenceAssertion struct {
	sel    *Selection
	name   string
	what   string
	values []string
	get    func(n Node) (string, bool)
}

var _ assertion.Assertion = sequenceAssertion{}

func (a sequenceAssertion) String() string {
	return fmt.Sprintf("%s(%s, %q)", a.name, a.sel.path(), a.values)
}

func (a sequenceAssertion) Check(ctx context.Context) error {
	nodes, err := a.sel.nodes(ctx)
	if err != nil {
		return err
	}
	got := make([]string, len(nodes))
	equal := len(nodes) == len(a.values)
	for i, n := range nodes {
		v, ok := a.get(n)
		if !ok {
			v = "(absent)"
		}
		got[i] = v
		if !ok || i >= len(a.values) || v != a.values[i] {
			equal = false
		}
	}
	if equal {
		return nil
	}
	return fmt.Errorf("%w: expected %s of %s to be %q, got %q\nmatches:\n%s",
		ErrMismatch, a.what, a.sel.path(), a.values, got, listNodes(nodes, a.sel.subject.opts.redact))
}

// eachAssertion asserts that every matched node matches more queries.
type eachAssertion struct {
	sel     *Selection
	queries []Query
}

var _ assertion.Assertion = eachAssertion{}

func (a eachAssertion) String() string {
	return fmt.Sprintf("each(%s, %s)", a.sel.path(), describeAll(a.queries))
}

func (a eachAssertion) Check(ctx context.Context) error {
	nodes, err := a.sel.nodes(ctx)
	if err != nil {
		return err
	}
	if len(nodes) == 0 {
		return a.sel.countFailure(ctx, "at least one", nodes)
	}
	for i, n := range nodes {
		q, err := firstMismatch(n, a.queries)
		if err != nil {
			return err
		}
		if q != nil {
			return fmt.Errorf("%w: expected every match of %s to match %s, match %d does not\nmatches:\n%s",
				ErrMismatch, a.sel.path(), q, i, listNodes(nodes, a.sel.subject.opts.redact))
		}
	}
	return nil
}

// inOrderAssertion asserts that the i-th matched node matches the i-th query.
type inOrderAssertion struct {
	sel     *Selection
	queries []Query
}

var _ assertion.Assertion = inOrderAssertion{}

func (a inOrderAssertion) String() string {
	return fmt.Sprintf("inOrder(%s, %s)", a.sel.path(), describeAll(a.queries))
}

func (a inOrderAssertion) Check(ctx context.Context) error {
	nodes, err := a.sel.nodes(ctx)
	if err != nil {
		return err
	}
	if len(nodes) != len(a.queries) {
		return a.sel.countFailure(ctx, fmt.Sprintf("exactly %d", len(a.queries)), nodes)
	}
	for i, n := range nodes {
		q, err := firstMismatch(n, a.queries[i:i+1])
		if err != nil {
			return err
		}
		if q != nil {
			return fmt.Errorf("%w: expected match %d of %s to match %s\nmatches:\n%s",
				ErrMismatch, i, a.sel.path(), q, listNodes(nodes, a.sel.subject.opts.redact))
		}
	}
	return nil
}

// firstMismatch returns the first query the node does not match, or nil.
func firstMismatch(n Node, queries []Query) (*Query, error) {
	for i := range queries {
		ok, err := queries[i].Match(n)
		if err != nil {
			return nil, fmt.Errorf("%w: %s at %s: %w", ErrQuery, queries[i], locationOf(n), err)
		}
		if !ok {
			return &queries[i], nil
		}
	}
	return nil, nil
}
