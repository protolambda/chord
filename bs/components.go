package bs

import (
	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/html/attr"
	"github.com/protolambda/chord/html/group/div"
	"github.com/protolambda/chord/html/group/list"
	"github.com/protolambda/chord/html/section"
	"github.com/protolambda/chord/html/text"
)

// Breadcrumb creates a Bootstrap breadcrumb container.
func Breadcrumb(items ...core.Node) core.Node {
	return section.Nav(
		list.OL(core.Compose(attr.Class("breadcrumb"), items...)),
	)
}

// BreadcrumbItem creates a Bootstrap breadcrumb item.
func BreadcrumbItem(opts ...core.Node) core.Node {
	return list.LI(core.Compose(attr.Class("breadcrumb-item"), opts...))
}

// Carousel creates a Bootstrap carousel.
func Carousel(opts ...core.Node) core.Node {
	return div.Div(core.Compose(attr.Class("carousel slide"), opts...))
}

// Pagination creates a Bootstrap pagination container.
func Pagination(opts ...core.Node) core.Node {
	return section.Nav(
		list.UL(core.Compose(attr.Class("pagination"), opts...)),
	)
}

// PageItem creates a Bootstrap pagination item.
func PageItem(opts ...core.Node) core.Node {
	return list.LI(core.Compose(attr.Class("page-item"), opts...))
}

// PageLink creates a Bootstrap pagination link.
func PageLink(opts ...core.Node) core.Node {
	return text.A(core.Compose(attr.Class("page-link"), opts...))
}

// Progress creates a Bootstrap progress container.
func Progress(opts ...core.Node) core.Node {
	return div.Div(core.Compose(attr.Class("progress"), opts...))
}

// ProgressBar creates a Bootstrap progress bar.
func ProgressBar(opts ...core.Node) core.Node {
	return div.Div(core.Compose(attr.Class("progress-bar"), opts...))
}

// Spinner creates a Bootstrap spinner.
func Spinner(opts ...core.Node) core.Node {
	return div.Div(core.Compose(attr.Class("spinner-border"), opts...))
}

// Collapse creates a Bootstrap collapse container.
func Collapse(opts ...core.Node) core.Node {
	return div.Div(core.Compose(attr.Class("collapse"), opts...))
}
