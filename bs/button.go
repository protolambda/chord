// Package bs provides Bootstrap 5.3 components and utilities.
// Reference: https://getbootstrap.com/docs/5.3/
package bs

import (
	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/html/attr"
	"github.com/protolambda/chord/html/form/button"
	"github.com/protolambda/chord/html/group/div"
)

// Btn creates a basic Bootstrap button.
func Btn(opts ...core.Node) core.Node {
	return button.Button(core.Compose(attr.Class("btn"), opts...))
}

// BtnPrimary creates a primary Bootstrap button.
func BtnPrimary(opts ...core.Node) core.Node {
	return button.Button(core.Compose(attr.Class("btn btn-primary"), opts...))
}

// BtnSecondary creates a secondary Bootstrap button.
func BtnSecondary(opts ...core.Node) core.Node {
	return button.Button(core.Compose(attr.Class("btn btn-secondary"), opts...))
}

// BtnSuccess creates a success Bootstrap button.
func BtnSuccess(opts ...core.Node) core.Node {
	return button.Button(core.Compose(attr.Class("btn btn-success"), opts...))
}

// BtnDanger creates a danger Bootstrap button.
func BtnDanger(opts ...core.Node) core.Node {
	return button.Button(core.Compose(attr.Class("btn btn-danger"), opts...))
}

// BtnWarning creates a warning Bootstrap button.
func BtnWarning(opts ...core.Node) core.Node {
	return button.Button(core.Compose(attr.Class("btn btn-warning"), opts...))
}

// BtnInfo creates an info Bootstrap button.
func BtnInfo(opts ...core.Node) core.Node {
	return button.Button(core.Compose(attr.Class("btn btn-info"), opts...))
}

// BtnLight creates a light Bootstrap button.
func BtnLight(opts ...core.Node) core.Node {
	return button.Button(core.Compose(attr.Class("btn btn-light"), opts...))
}

// BtnDark creates a dark Bootstrap button.
func BtnDark(opts ...core.Node) core.Node {
	return button.Button(core.Compose(attr.Class("btn btn-dark"), opts...))
}

// BtnLink creates a link-styled Bootstrap button.
func BtnLink(opts ...core.Node) core.Node {
	return button.Button(core.Compose(attr.Class("btn btn-link"), opts...))
}

// BtnOutlinePrimary creates an outline primary Bootstrap button.
func BtnOutlinePrimary(opts ...core.Node) core.Node {
	return button.Button(core.Compose(attr.Class("btn btn-outline-primary"), opts...))
}

// BtnOutlineSecondary creates an outline secondary Bootstrap button.
func BtnOutlineSecondary(opts ...core.Node) core.Node {
	return button.Button(core.Compose(attr.Class("btn btn-outline-secondary"), opts...))
}

// BtnGroup creates a Bootstrap button group wrapper.
func BtnGroup(opts ...core.Node) core.Node {
	return div.Div(core.Compose(attr.Class("btn-group"), opts...))
}
