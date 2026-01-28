package ba

import (
	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/html/attr"
)

// Tooltip creates tooltip attributes for an element.
func Tooltip(opts ...core.Node) core.Node {
	return core.Compose(attr.Data("bs-toggle", "tooltip"), opts...)
}

// Popover creates popover attributes for an element.
func Popover(opts ...core.Node) core.Node {
	return core.Compose(attr.Data("bs-toggle", "popover"), opts...)
}
