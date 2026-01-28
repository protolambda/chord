// Package edit provides HTML edit elements (ins, del).
package edit

import "github.com/protolambda/chord/core"

// Ins creates an ins element.
func Ins(opts ...core.Node) core.Node { return core.Element("ins", opts...) }

// Del creates a del element.
func Del(opts ...core.Node) core.Node { return core.Element("del", opts...) }
