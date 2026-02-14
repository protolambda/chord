package bs

import (
	"github.com/protolambda/chord/core/attr"
	"github.com/protolambda/chord/core/elem"
	"github.com/protolambda/chord/html/text"
)

// Badge creates a basic Bootstrap badge.
func Badge(attrs ...attr.Node) elem.Scope {
	return text.Span(attr.Cons(attr.Class("badge"), attrs...))
}

// BadgePrimary creates a primary Bootstrap badge.
func BadgePrimary(attrs ...attr.Node) elem.Scope {
	return text.Span(attr.Cons(attr.Class("badge text-bg-primary"), attrs...))
}

// BadgeSecondary creates a secondary Bootstrap badge.
func BadgeSecondary(attrs ...attr.Node) elem.Scope {
	return text.Span(attr.Cons(attr.Class("badge text-bg-secondary"), attrs...))
}

// BadgeSuccess creates a success Bootstrap badge.
func BadgeSuccess(attrs ...attr.Node) elem.Scope {
	return text.Span(attr.Cons(attr.Class("badge text-bg-success"), attrs...))
}

// BadgeDanger creates a danger Bootstrap badge.
func BadgeDanger(attrs ...attr.Node) elem.Scope {
	return text.Span(attr.Cons(attr.Class("badge text-bg-danger"), attrs...))
}

// BadgeWarning creates a warning Bootstrap badge.
func BadgeWarning(attrs ...attr.Node) elem.Scope {
	return text.Span(attr.Cons(attr.Class("badge text-bg-warning"), attrs...))
}

// BadgeInfo creates an info Bootstrap badge.
func BadgeInfo(attrs ...attr.Node) elem.Scope {
	return text.Span(attr.Cons(attr.Class("badge text-bg-info"), attrs...))
}
