// Package button provides the HTML button element and its attributes.
package button

import (
	"github.com/protolambda/chord/core/attr"
	"github.com/protolambda/chord/core/elem"
)

// Button creates a button element.
func Button(attrs ...attr.Node) elem.Scope { return elem.Name("button").New(attrs...) }
