// Package webcomp provides HTML web component elements.
package webcomp

import (
	"github.com/protolambda/chord/core/attr"
	"github.com/protolambda/chord/core/elem"
)

// Slot creates a slot element.
func Slot(attrs ...attr.Node) elem.Scope { return elem.Name("slot").New(attrs...) }
