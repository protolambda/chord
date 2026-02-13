package bs

import (
	"github.com/protolambda/chord/core/attrib"
	"github.com/protolambda/chord/core/elem"
	"github.com/protolambda/chord/html/attr"
	"github.com/protolambda/chord/html/text"
)

// Badge creates a basic Bootstrap badge.
func Badge(attrs ...attrib.Node) elem.Scope {
	return text.Span(attrib.Cons(attr.Class("badge"), attrs...))
}

// BadgePrimary creates a primary Bootstrap badge.
func BadgePrimary(attrs ...attrib.Node) elem.Scope {
	return text.Span(attrib.Cons(attr.Class("badge text-bg-primary"), attrs...))
}

// BadgeSecondary creates a secondary Bootstrap badge.
func BadgeSecondary(attrs ...attrib.Node) elem.Scope {
	return text.Span(attrib.Cons(attr.Class("badge text-bg-secondary"), attrs...))
}

// BadgeSuccess creates a success Bootstrap badge.
func BadgeSuccess(attrs ...attrib.Node) elem.Scope {
	return text.Span(attrib.Cons(attr.Class("badge text-bg-success"), attrs...))
}

// BadgeDanger creates a danger Bootstrap badge.
func BadgeDanger(attrs ...attrib.Node) elem.Scope {
	return text.Span(attrib.Cons(attr.Class("badge text-bg-danger"), attrs...))
}

// BadgeWarning creates a warning Bootstrap badge.
func BadgeWarning(attrs ...attrib.Node) elem.Scope {
	return text.Span(attrib.Cons(attr.Class("badge text-bg-warning"), attrs...))
}

// BadgeInfo creates an info Bootstrap badge.
func BadgeInfo(attrs ...attrib.Node) elem.Scope {
	return text.Span(attrib.Cons(attr.Class("badge text-bg-info"), attrs...))
}
