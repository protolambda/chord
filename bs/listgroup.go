package bs

import (
	"github.com/protolambda/chord/core/attr"
	"github.com/protolambda/chord/core/elem"
	"github.com/protolambda/chord/html/group/list"
)

// ListGroup creates a Bootstrap list group.
func ListGroup(attrs ...attr.Node) elem.Scope {
	return list.UL(attr.Cons(attr.Class("list-group"), attrs...))
}

// ListGroupItem creates a Bootstrap list group item.
func ListGroupItem(attrs ...attr.Node) elem.Scope {
	return list.LI(attr.Cons(attr.Class("list-group-item"), attrs...))
}
