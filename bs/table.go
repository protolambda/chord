package bs

import (
	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/html/attr"
	"github.com/protolambda/chord/html/group/div"
	"github.com/protolambda/chord/html/table"
)

// Table creates a Bootstrap table.
func Table(opts ...core.Node) core.Node {
	return table.Table(core.Compose(attr.Class("table"), opts...))
}

// TableStriped creates a striped Bootstrap table.
func TableStriped(opts ...core.Node) core.Node {
	return table.Table(core.Compose(attr.Class("table table-striped"), opts...))
}

// TableBordered creates a bordered Bootstrap table.
func TableBordered(opts ...core.Node) core.Node {
	return table.Table(core.Compose(attr.Class("table table-bordered"), opts...))
}

// TableHover creates a hover Bootstrap table.
func TableHover(opts ...core.Node) core.Node {
	return table.Table(core.Compose(attr.Class("table table-hover"), opts...))
}

// TableResponsive creates a responsive table wrapper.
func TableResponsive(opts ...core.Node) core.Node {
	return div.Div(core.Compose(attr.Class("table-responsive"), opts...))
}
