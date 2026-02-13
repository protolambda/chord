package bs

import (
	"github.com/protolambda/chord/core/attrib"
	"github.com/protolambda/chord/core/elem"
	"github.com/protolambda/chord/html/attr"
	"github.com/protolambda/chord/html/group/div"
	"github.com/protolambda/chord/html/group/list"
	"github.com/protolambda/chord/html/section"
)

// Breadcrumb creates a Bootstrap breadcrumb container.
func Breadcrumb(attrs ...attrib.Node) elem.Scope {
	return func(children ...elem.Node) elem.Node {
		return section.Nav()(
			list.OL(attrib.Cons(attr.Class("breadcrumb"), attrs...))(children...),
		)
	}
}

// BreadcrumbItem creates a Bootstrap breadcrumb item.
func BreadcrumbItem(attrs ...attrib.Node) elem.Scope {
	return list.LI(attrib.Cons(attr.Class("breadcrumb-item"), attrs...))
}

// Carousel creates a Bootstrap carousel.
func Carousel(attrs ...attrib.Node) elem.Scope {
	return div.Div(attrib.Cons(attr.Class("carousel slide"), attrs...))
}

// Pagination creates a Bootstrap pagination container.
func Pagination(attrs ...attrib.Node) elem.Scope {
	return func(children ...elem.Node) elem.Node {
		return section.Nav()(
			list.UL(attrib.Cons(attr.Class("pagination"), attrs...))(children...),
		)
	}
}

// PageItem creates a Bootstrap pagination item.
func PageItem(attrs ...attrib.Node) elem.Scope {
	return list.LI(attrib.Cons(attr.Class("page-item"), attrs...))
}

// PageLink creates a Bootstrap pagination link.
func PageLink(attrs ...attrib.Node) elem.Scope {
	return list.LI(attrib.Cons(attr.Class("page-link"), attrs...))
}

// Progress creates a Bootstrap progress container.
func Progress(attrs ...attrib.Node) elem.Scope {
	return div.Div(attrib.Cons(attr.Class("progress"), attrs...))
}

// ProgressBar creates a Bootstrap progress bar.
func ProgressBar(attrs ...attrib.Node) elem.Scope {
	return div.Div(attrib.Cons(attr.Class("progress-bar"), attrs...))
}

// Spinner creates a Bootstrap spinner.
func Spinner(attrs ...attrib.Node) elem.Scope {
	return div.Div(attrib.Cons(attr.Class("spinner-border"), attrs...))
}

// Collapse creates a Bootstrap collapse container.
func Collapse(attrs ...attrib.Node) elem.Scope {
	return div.Div(attrib.Cons(attr.Class("collapse"), attrs...))
}
