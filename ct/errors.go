package ct

import "errors"

// Sentinel errors classify assertion failures so that tests of test code can
// distinguish operational failures from mismatches with errors.Is.
var (
	// ErrLoad indicates that the subject could not be evaluated, parsed, or captured.
	ErrLoad = errors.New("subject failed to load")
	// ErrQuery indicates that a query could not be evaluated, for example
	// because a semantic computation failed. It is never treated as absence.
	ErrQuery = errors.New("query failed")
	// ErrCount indicates that a selection matched an unexpected number of nodes.
	ErrCount = errors.New("unexpected number of matches")
	// ErrMismatch indicates that selected nodes did not have an expected property.
	ErrMismatch = errors.New("selected nodes do not match")
	// ErrRule indicates that a document rule was violated.
	ErrRule = errors.New("document rule violated")
)
