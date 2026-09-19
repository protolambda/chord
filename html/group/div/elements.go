// Package div provides general grouping elements like div, p, hr, pre, and blockquote.
package div

import (
	"github.com/protolambda/chord/core/attr"
	"github.com/protolambda/chord/core/elem"
)

// Div creates a div element.
func Div(attrs ...attr.Node) elem.Scope { return elem.Name("div").New(attrs...) }
