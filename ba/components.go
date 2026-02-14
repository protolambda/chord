package ba

import (
	"github.com/protolambda/chord/core/attr"
)

// Tooltip creates tooltip attributes for an element.
func Tooltip(opts ...attr.Node) attr.Node {
	return attr.Cons(attr.Data("bs-toggle", "tooltip"), opts...)
}

// Popover creates popover attributes for an element.
func Popover(opts ...attr.Node) attr.Node {
	return attr.Cons(attr.Data("bs-toggle", "popover"), opts...)
}
