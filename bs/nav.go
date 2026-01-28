package bs

import (
	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/html/attr"
	"github.com/protolambda/chord/html/form/button"
	"github.com/protolambda/chord/html/group/div"
	"github.com/protolambda/chord/html/group/list"
	"github.com/protolambda/chord/html/section"
	"github.com/protolambda/chord/html/text"
)

// Nav creates a Bootstrap nav container.
func Nav(opts ...core.Node) core.Node {
	return list.UL(core.Compose(attr.Class("nav"), opts...))
}

// NavItem creates a Bootstrap nav item.
func NavItem(opts ...core.Node) core.Node {
	return list.LI(core.Compose(attr.Class("nav-item"), opts...))
}

// NavLink creates a Bootstrap nav link.
func NavLink(opts ...core.Node) core.Node {
	return text.A(core.Compose(attr.Class("nav-link"), opts...))
}

// Navbar creates a Bootstrap navbar.
func Navbar(opts ...core.Node) core.Node {
	return section.Nav(core.Compose(attr.Class("navbar"), opts...))
}

// NavbarBrand creates a Bootstrap navbar brand.
func NavbarBrand(opts ...core.Node) core.Node {
	return text.A(core.Compose(attr.Class("navbar-brand"), opts...))
}

// NavbarToggler creates a Bootstrap navbar toggler button.
func NavbarToggler(opts ...core.Node) core.Node {
	return button.Button(core.Compose(
		attr.Class("navbar-toggler"),
		button.Type(button.TypeButton),
		attr.Data("bs-toggle", "collapse"),
		text.Span(attr.Class("navbar-toggler-icon")),
	), core.Bundle(opts...))
}

// NavbarCollapse creates a Bootstrap navbar collapse container.
func NavbarCollapse(opts ...core.Node) core.Node {
	return div.Div(core.Compose(attr.Class("collapse navbar-collapse"), opts...))
}

// NavbarNav creates a Bootstrap navbar nav container.
func NavbarNav(opts ...core.Node) core.Node {
	return list.UL(core.Compose(attr.Class("navbar-nav"), opts...))
}
