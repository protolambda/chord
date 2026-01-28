package bs

import (
	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/html/attr"
	"github.com/protolambda/chord/html/text"
)

// Badge creates a basic Bootstrap badge.
func Badge(opts ...core.Node) core.Node {
	return text.Span(core.Compose(attr.Class("badge"), opts...))
}

// BadgePrimary creates a primary Bootstrap badge.
func BadgePrimary(opts ...core.Node) core.Node {
	return text.Span(core.Compose(attr.Class("badge text-bg-primary"), opts...))
}

// BadgeSecondary creates a secondary Bootstrap badge.
func BadgeSecondary(opts ...core.Node) core.Node {
	return text.Span(core.Compose(attr.Class("badge text-bg-secondary"), opts...))
}

// BadgeSuccess creates a success Bootstrap badge.
func BadgeSuccess(opts ...core.Node) core.Node {
	return text.Span(core.Compose(attr.Class("badge text-bg-success"), opts...))
}

// BadgeDanger creates a danger Bootstrap badge.
func BadgeDanger(opts ...core.Node) core.Node {
	return text.Span(core.Compose(attr.Class("badge text-bg-danger"), opts...))
}

// BadgeWarning creates a warning Bootstrap badge.
func BadgeWarning(opts ...core.Node) core.Node {
	return text.Span(core.Compose(attr.Class("badge text-bg-warning"), opts...))
}

// BadgeInfo creates an info Bootstrap badge.
func BadgeInfo(opts ...core.Node) core.Node {
	return text.Span(core.Compose(attr.Class("badge text-bg-info"), opts...))
}
