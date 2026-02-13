// Package webcomp provides HTML web component elements.
package webcomp

import (
	"github.com/protolambda/chord/core/attrib"
	"github.com/protolambda/chord/core/elem"
)

// Slot creates a slot element.
func Slot(attrs ...attrib.Node) elem.Scope { return elem.New("slot", attrs...) }
