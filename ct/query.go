package ct

import (
	"errors"
	"fmt"
	"strings"
)

// Query is an immutable, descriptive element predicate.
//
// Matching may fail with an error for semantic queries that follow
// references; such an error is reported as [ErrQuery] and never converted
// into a match or a non-match. The zero Query is invalid.
type Query struct {
	desc  string
	match func(n Node) (bool, error)
}

func (q Query) String() string {
	return q.desc
}

// Match reports whether the node matches. Non-element nodes never match.
func (q Query) Match(n Node) (bool, error) {
	if q.match == nil {
		return false, errors.New("zero query")
	}
	if n.Kind() != KindElement {
		return false, nil
	}
	return q.match(n)
}

// And requires this query and all others to match.
func (q Query) And(others ...Query) Query {
	if len(others) == 0 {
		return q
	}
	all := append([]Query{q}, others...)
	return Query{
		desc: "and(" + describeAll(all) + ")",
		match: func(n Node) (bool, error) {
			for _, q := range all {
				ok, err := q.Match(n)
				if err != nil || !ok {
					return false, err
				}
			}
			return true, nil
		},
	}
}

// Or requires this query or any of the others to match.
func (q Query) Or(others ...Query) Query {
	if len(others) == 0 {
		return q
	}
	all := append([]Query{q}, others...)
	return Query{
		desc: "or(" + describeAll(all) + ")",
		match: func(n Node) (bool, error) {
			for _, q := range all {
				ok, err := q.Match(n)
				if err != nil || ok {
					return ok, err
				}
			}
			return false, nil
		},
	}
}

// Not inverts a successful match. A query error stays an error.
func (q Query) Not() Query {
	return Query{
		desc: "not(" + q.desc + ")",
		match: func(n Node) (bool, error) {
			ok, err := q.Match(n)
			if err != nil {
				return false, err
			}
			return !ok, nil
		},
	}
}

func describeAll(queries []Query) string {
	parts := make([]string, len(queries))
	for i, q := range queries {
		parts[i] = q.String()
	}
	return strings.Join(parts, ", ")
}

// Custom creates a query from a description and a predicate over elements.
// The predicate only receives element nodes.
func Custom(desc string, match func(n Node) (bool, error)) Query {
	return Query{desc: desc, match: match}
}

// Tag matches elements with the given tag, case-insensitively.
func Tag(name string) Query {
	name = strings.ToLower(name)
	return Query{
		desc:  fmt.Sprintf("tag(%q)", name),
		match: func(n Node) (bool, error) { return n.Tag() == name, nil },
	}
}

// ID matches elements with exactly the given id attribute.
func ID(value string) Query {
	return Query{
		desc:  fmt.Sprintf("id(%q)", value),
		match: func(n Node) (bool, error) { return attrEquals(n, "id", value), nil },
	}
}

// Class matches elements whose class attribute contains the given token.
func Class(token string) Query {
	return Query{
		desc: fmt.Sprintf("class(%q)", token),
		match: func(n Node) (bool, error) {
			classes, ok := n.Attr("class")
			if !ok {
				return false, nil
			}
			for c := range strings.FieldsSeq(classes) {
				if c == token {
					return true, nil
				}
			}
			return false, nil
		},
	}
}

// Attr matches elements whose attribute has exactly the given value.
// A boolean attribute has the empty value.
func Attr(name, value string) Query {
	return Query{
		desc:  fmt.Sprintf("attr(%q, %q)", name, value),
		match: func(n Node) (bool, error) { return attrEquals(n, name, value), nil },
	}
}

// AttrMatches matches elements whose attribute value satisfies the matcher.
func AttrMatches(name string, value ValueMatch) Query {
	return Query{
		desc: fmt.Sprintf("attr(%q, %s)", name, value),
		match: func(n Node) (bool, error) {
			v, ok := n.Attr(name)
			return ok && value.Match(v), nil
		},
	}
}

// HasAttr matches elements that have the attribute, with any value.
func HasAttr(name string) Query {
	return Query{
		desc: fmt.Sprintf("hasAttr(%q)", name),
		match: func(n Node) (bool, error) {
			_, ok := n.Attr(name)
			return ok, nil
		},
	}
}

// TestID matches elements with exactly the given data-testid attribute.
func TestID(value string) Query {
	return Query{
		desc:  fmt.Sprintf("testID(%q)", value),
		match: func(n Node) (bool, error) { return attrEquals(n, "data-testid", value), nil },
	}
}

// HasChild matches elements with a direct child element matching the query.
func HasChild(query Query) Query {
	return Query{
		desc: "hasChild(" + query.desc + ")",
		match: func(n Node) (bool, error) {
			for c := range n.Children() {
				ok, err := query.Match(c)
				if err != nil || ok {
					return ok, err
				}
			}
			return false, nil
		},
	}
}

// HasDescendant matches elements with any descendant element matching the query.
func HasDescendant(query Query) Query {
	return Query{
		desc: "hasDescendant(" + query.desc + ")",
		match: func(n Node) (bool, error) {
			for d := range descendants(n) {
				ok, err := query.Match(d)
				if err != nil || ok {
					return ok, err
				}
			}
			return false, nil
		},
	}
}

func attrEquals(n Node, name, value string) bool {
	v, ok := n.Attr(name)
	return ok && v == value
}
