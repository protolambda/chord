// Package button provides the HTML button element and its attributes.
package button

import (
	"github.com/protolambda/chord/core/attrib"
	"github.com/protolambda/chord/core/elem"
)

// Button creates a button element.
func Button(attrs ...attrib.Node) elem.Scope { return elem.New("button", attrs...) }
