// Package bs provides Bootstrap 5.3 components and utilities.
// Reference: https://getbootstrap.com/docs/5.3/
package bs

import (
	"github.com/protolambda/chord/core/attr"
	"github.com/protolambda/chord/core/elem"
	"github.com/protolambda/chord/html/form/button"
	"github.com/protolambda/chord/html/group/div"
)

// Btn creates a basic Bootstrap button.
func Btn(attrs ...attr.Node) elem.Scope {
	return button.Button(attr.Cons(attr.Class("btn"), attrs...))
}

// BtnPrimary creates a primary Bootstrap button.
func BtnPrimary(attrs ...attr.Node) elem.Scope {
	return button.Button(attr.Cons(attr.Class("btn btn-primary"), attrs...))
}

// BtnSecondary creates a secondary Bootstrap button.
func BtnSecondary(attrs ...attr.Node) elem.Scope {
	return button.Button(attr.Cons(attr.Class("btn btn-secondary"), attrs...))
}

// BtnSuccess creates a success Bootstrap button.
func BtnSuccess(attrs ...attr.Node) elem.Scope {
	return button.Button(attr.Cons(attr.Class("btn btn-success"), attrs...))
}

// BtnDanger creates a danger Bootstrap button.
func BtnDanger(attrs ...attr.Node) elem.Scope {
	return button.Button(attr.Cons(attr.Class("btn btn-danger"), attrs...))
}

// BtnWarning creates a warning Bootstrap button.
func BtnWarning(attrs ...attr.Node) elem.Scope {
	return button.Button(attr.Cons(attr.Class("btn btn-warning"), attrs...))
}

// BtnInfo creates an info Bootstrap button.
func BtnInfo(attrs ...attr.Node) elem.Scope {
	return button.Button(attr.Cons(attr.Class("btn btn-info"), attrs...))
}

// BtnLight creates a light Bootstrap button.
func BtnLight(attrs ...attr.Node) elem.Scope {
	return button.Button(attr.Cons(attr.Class("btn btn-light"), attrs...))
}

// BtnDark creates a dark Bootstrap button.
func BtnDark(attrs ...attr.Node) elem.Scope {
	return button.Button(attr.Cons(attr.Class("btn btn-dark"), attrs...))
}

// BtnLink creates a link-styled Bootstrap button.
func BtnLink(attrs ...attr.Node) elem.Scope {
	return button.Button(attr.Cons(attr.Class("btn btn-link"), attrs...))
}

// BtnOutlinePrimary creates an outline primary Bootstrap button.
func BtnOutlinePrimary(attrs ...attr.Node) elem.Scope {
	return button.Button(attr.Cons(attr.Class("btn btn-outline-primary"), attrs...))
}

// BtnOutlineSecondary creates an outline secondary Bootstrap button.
func BtnOutlineSecondary(attrs ...attr.Node) elem.Scope {
	return button.Button(attr.Cons(attr.Class("btn btn-outline-secondary"), attrs...))
}

// BtnGroup creates a Bootstrap button group wrapper.
func BtnGroup(attrs ...attr.Node) elem.Scope {
	return div.Div(attr.Cons(attr.Class("btn-group"), attrs...))
}
