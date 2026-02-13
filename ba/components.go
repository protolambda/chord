package ba

import (
	"github.com/protolambda/chord/core/attrib"
	"github.com/protolambda/chord/html/attr"
)

// Tooltip creates tooltip attributes for an element.
func Tooltip(opts ...attrib.Node) attrib.Node {
	return attrib.Cons(attr.Data("bs-toggle", "tooltip"), opts...)
}

// Popover creates popover attributes for an element.
func Popover(opts ...attrib.Node) attrib.Node {
	return attrib.Cons(attr.Data("bs-toggle", "popover"), opts...)
}
