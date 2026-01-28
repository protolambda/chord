// Package webcomp provides HTML web component elements.
package webcomp

import "github.com/protolambda/chord/core"

// Slot creates a slot element.
func Slot(opts ...core.Node) core.Node { return core.Element("slot", opts...) }
