// Package edit provides HTML edit elements (ins, del).
package edit

import (
	"github.com/protolambda/chord/core/attr"
	"github.com/protolambda/chord/core/elem"
)

// Ins creates an ins element.
func Ins(attrs ...attr.Node) elem.Scope { return elem.Name("ins").New(attrs...) }

// Del creates a del element.
func Del(attrs ...attr.Node) elem.Scope { return elem.Name("del").New(attrs...) }
