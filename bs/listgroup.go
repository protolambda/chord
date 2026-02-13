package bs

import (
	"github.com/protolambda/chord/core/attrib"
	"github.com/protolambda/chord/core/elem"
	"github.com/protolambda/chord/html/attr"
	"github.com/protolambda/chord/html/group/list"
)

// ListGroup creates a Bootstrap list group.
func ListGroup(attrs ...attrib.Node) elem.Scope {
	return list.UL(attrib.Cons(attr.Class("list-group"), attrs...))
}

// ListGroupItem creates a Bootstrap list group item.
func ListGroupItem(attrs ...attrib.Node) elem.Scope {
	return list.LI(attrib.Cons(attr.Class("list-group-item"), attrs...))
}
