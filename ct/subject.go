package ct

import (
	"context"
	"fmt"
	"sync"

	"github.com/protolambda/mustbe/assertion"
)

// Subject is a lazily loaded, cached source of nodes. It is an assertion
// that the source loads successfully.
//
// The first check or load stores the result, including a failure, for the
// lifetime of the subject. Use one subject per test or subtest.
type Subject struct {
	source Source
	opts   options

	once sync.Once
	doc  Document
	err  error
}

var _ assertion.Assertion = (*Subject)(nil)

// Option configures a [Subject].
type Option func(*options)

type options struct {
	redact        func(tag, key string) bool
	redactContent []Query
}

// WithRedact hides the values of additional attributes in diagnostics.
// The predicate receives the lowercase element tag and attribute name.
// Values that look like secrets are redacted regardless of this option; see
// the package documentation.
func WithRedact(fn func(tag, key string) bool) Option {
	return func(o *options) {
		o.redact = fn
	}
}

// WithRedactContent marks the elements that match the query as sensitive
// in diagnostics, in addition to the elements that look sensitive (see the
// package documentation): the text, raw content and comments inside them,
// and the value and content attributes of the elements and their
// descendants, are shown as [redacted]. A query that fails counts as a
// match. Repeated options add up.
//
// Redaction only applies to failure messages. Values that a test reads,
// such as [Selection.Text], and the expected values that a test states are
// never redacted.
func WithRedactContent(query Query) Option {
	return func(o *options) {
		o.redactContent = append(o.redactContent, query)
	}
}

// From creates a subject for any [Source].
func From(source Source, opts ...Option) *Subject {
	s := &Subject{source: source}
	for _, opt := range opts {
		opt(&s.opts)
	}
	return s
}

// String describes the assertion: the source loads.
func (s *Subject) String() string {
	return fmt.Sprintf("loads(%s)", s.source)
}

// Check loads the subject, if not loaded yet, and reports a load failure.
func (s *Subject) Check(ctx context.Context) error {
	_, err := s.Load(ctx)
	return err
}

// Load returns the loaded document. The source is loaded at most once, with
// the context of the first call; later calls return the same result.
// Prefer assertions over direct document access; Load exists for custom
// assertions that need the tree.
func (s *Subject) Load(ctx context.Context) (Document, error) {
	s.once.Do(func() {
		doc, err := s.source.Load(ctx)
		if err != nil {
			s.err = fmt.Errorf("%w: %s: %w", ErrLoad, s.source, err)
			return
		}
		s.doc = doc
	})
	return s.doc, s.err
}

// Find selects the elements in the document that match every query.
// The document is not loaded until the selection is checked.
func (s *Subject) Find(query Query, more ...Query) *Selection {
	return &Selection{subject: s, query: query.And(more...)}
}
