// Package input provides the HTML input element and its attributes.
package input

import (
	"github.com/protolambda/chord/core/attr"
	"github.com/protolambda/chord/core/elem"
)

// Input creates an input element (void).
func Input(attrs ...attr.Node) elem.Node { return elem.Name("input").Void(attrs...) }
