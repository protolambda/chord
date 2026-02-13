// Package input provides the HTML input element and its attributes.
package input

import (
	"github.com/protolambda/chord/core/attrib"
	"github.com/protolambda/chord/core/elem"
)

// Input creates an input element (void).
func Input(attrs ...attrib.Node) elem.Node { return elem.Void("input", attrs...) }
