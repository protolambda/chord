// Package edit provides HTML edit elements (ins, del).
package edit

import (
	"github.com/protolambda/chord/core/attrib"
	"github.com/protolambda/chord/core/elem"
)

// Ins creates an ins element.
func Ins(attrs ...attrib.Node) elem.Scope { return elem.New("ins", attrs...) }

// Del creates a del element.
func Del(attrs ...attrib.Node) elem.Scope { return elem.New("del", attrs...) }
