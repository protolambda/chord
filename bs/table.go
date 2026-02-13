package bs

import (
	"github.com/protolambda/chord/core/attrib"
	"github.com/protolambda/chord/core/elem"
	"github.com/protolambda/chord/html/attr"
	"github.com/protolambda/chord/html/group/div"
	"github.com/protolambda/chord/html/table"
)

// Table creates a Bootstrap table.
func Table(attrs ...attrib.Node) elem.Scope {
	return table.Table(attrib.Cons(attr.Class("table"), attrs...))
}

// TableStriped creates a striped Bootstrap table.
func TableStriped(attrs ...attrib.Node) elem.Scope {
	return table.Table(attrib.Cons(attr.Class("table table-striped"), attrs...))
}

// TableBordered creates a bordered Bootstrap table.
func TableBordered(attrs ...attrib.Node) elem.Scope {
	return table.Table(attrib.Cons(attr.Class("table table-bordered"), attrs...))
}

// TableHover creates a hover Bootstrap table.
func TableHover(attrs ...attrib.Node) elem.Scope {
	return table.Table(attrib.Cons(attr.Class("table table-hover"), attrs...))
}

// TableResponsive creates a responsive table wrapper.
func TableResponsive(attrs ...attrib.Node) elem.Scope {
	return div.Div(attrib.Cons(attr.Class("table-responsive"), attrs...))
}
