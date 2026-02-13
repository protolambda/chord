// Package div provides general grouping elements like div, p, hr, pre, and blockquote.
package div

import (
	"github.com/protolambda/chord/core/attrib"
	"github.com/protolambda/chord/core/elem"
)

// Div creates a div element.
func Div(attrs ...attrib.Node) elem.Scope { return elem.New("div", attrs...) }
