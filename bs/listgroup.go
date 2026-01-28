package bs

import (
	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/html/attr"
	"github.com/protolambda/chord/html/group/list"
)

// ListGroup creates a Bootstrap list group.
func ListGroup(opts ...core.Node) core.Node {
	return list.UL(core.Compose(attr.Class("list-group"), opts...))
}

// ListGroupItem creates a Bootstrap list group item.
func ListGroupItem(opts ...core.Node) core.Node {
	return list.LI(core.Compose(attr.Class("list-group-item"), opts...))
}
