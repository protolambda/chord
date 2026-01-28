package bs

import (
	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/html/attr"
	"github.com/protolambda/chord/html/form/button"
	"github.com/protolambda/chord/html/group/div"
)

// Alert creates a basic Bootstrap alert.
func Alert(opts ...core.Node) core.Node {
	return div.Div(core.Compose(attr.Class("alert"), opts...))
}

// AlertPrimary creates a primary Bootstrap alert.
func AlertPrimary(opts ...core.Node) core.Node {
	return div.Div(core.Compose(attr.Class("alert alert-primary"), opts...))
}

// AlertSecondary creates a secondary Bootstrap alert.
func AlertSecondary(opts ...core.Node) core.Node {
	return div.Div(core.Compose(attr.Class("alert alert-secondary"), opts...))
}

// AlertSuccess creates a success Bootstrap alert.
func AlertSuccess(opts ...core.Node) core.Node {
	return div.Div(core.Compose(attr.Class("alert alert-success"), opts...))
}

// AlertDanger creates a danger Bootstrap alert.
func AlertDanger(opts ...core.Node) core.Node {
	return div.Div(core.Compose(attr.Class("alert alert-danger"), opts...))
}

// AlertWarning creates a warning Bootstrap alert.
func AlertWarning(opts ...core.Node) core.Node {
	return div.Div(core.Compose(attr.Class("alert alert-warning"), opts...))
}

// AlertInfo creates an info Bootstrap alert.
func AlertInfo(opts ...core.Node) core.Node {
	return div.Div(core.Compose(attr.Class("alert alert-info"), opts...))
}

// AlertDismissible creates a dismissible Bootstrap alert.
func AlertDismissible(opts ...core.Node) core.Node {
	return div.Div(core.Compose(
		attr.Class("alert alert-dismissible fade show"),
		button.Button(
			attr.Class("btn-close"),
			button.Type(button.TypeButton),
			attr.Data("bs-dismiss", "alert"),
		),
	), core.Bundle(opts...))
}
