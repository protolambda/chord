package bs

import (
	"github.com/protolambda/chord/core/attrib"
	"github.com/protolambda/chord/core/elem"
	"github.com/protolambda/chord/html/attr"
	"github.com/protolambda/chord/html/form/button"
	"github.com/protolambda/chord/html/group/div"
	"github.com/protolambda/chord/html/group/list"
	"github.com/protolambda/chord/html/section"
	"github.com/protolambda/chord/html/text"
)

// Nav creates a Bootstrap nav container.
func Nav(attrs ...attrib.Node) elem.Scope {
	return list.UL(attrib.Cons(attr.Class("nav"), attrs...))
}

// NavItem creates a Bootstrap nav item.
func NavItem(attrs ...attrib.Node) elem.Scope {
	return list.LI(attrib.Cons(attr.Class("nav-item"), attrs...))
}

// NavLink creates a Bootstrap nav link.
func NavLink(attrs ...attrib.Node) elem.Scope {
	return text.A(attrib.Cons(attr.Class("nav-link"), attrs...))
}

// Navbar creates a Bootstrap navbar.
func Navbar(attrs ...attrib.Node) elem.Scope {
	return section.Nav(attrib.Cons(attr.Class("navbar"), attrs...))
}

// NavbarBrand creates a Bootstrap navbar brand.
func NavbarBrand(attrs ...attrib.Node) elem.Scope {
	return text.A(attrib.Cons(attr.Class("navbar-brand"), attrs...))
}

// NavbarToggler creates a Bootstrap navbar toggler button.
// The toggler icon span is automatically included as a child.
func NavbarToggler(attrs ...attrib.Node) elem.Node {
	return button.Button(attrib.Cons(
		attr.Class("navbar-toggler"),
		button.Type(button.TypeButton),
		attr.Data("bs-toggle", "collapse"),
		attrib.Bundle(attrs),
	))(text.Span(attr.Class("navbar-toggler-icon"))())
}

// NavbarCollapse creates a Bootstrap navbar collapse container.
func NavbarCollapse(attrs ...attrib.Node) elem.Scope {
	return div.Div(attrib.Cons(attr.Class("collapse navbar-collapse"), attrs...))
}

// NavbarNav creates a Bootstrap navbar nav container.
func NavbarNav(attrs ...attrib.Node) elem.Scope {
	return list.UL(attrib.Cons(attr.Class("navbar-nav"), attrs...))
}
