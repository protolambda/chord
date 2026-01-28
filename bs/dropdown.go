package bs

import (
	"context"

	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/html/attr"
	"github.com/protolambda/chord/html/form/button"
	"github.com/protolambda/chord/html/group/div"
	"github.com/protolambda/chord/html/group/list"
	"github.com/protolambda/chord/html/section"
	"github.com/protolambda/chord/html/text"
)

// Dropdown represents a Bootstrap dropdown component.
type Dropdown struct {
	Toggle core.Node   // toggle button content
	Items  []core.Node // dropdown items
	Attrs  core.Node
}

// Eval builds the dropdown structure with proper Bootstrap markup.
func (d Dropdown) Eval(ctx context.Context) (core.Obj, error) {
	var children []core.Node
	if d.Attrs != nil {
		children = append(children, d.Attrs)
	}
	children = append(children, attr.Class("dropdown"))

	toggleBtn := button.Button(
		attr.Class("btn btn-secondary dropdown-toggle"),
		button.Type(button.TypeButton),
		attr.Data("bs-toggle", "dropdown"),
		d.Toggle,
	)

	menu := list.UL(
		core.Compose(attr.Class("dropdown-menu"), d.Items...),
	)

	children = append(children, toggleBtn, menu)
	return div.Div(children...).Eval(ctx)
}

// DropdownItem creates a Bootstrap dropdown item.
func DropdownItem(opts ...core.Node) core.Node {
	return list.LI(
		text.A(core.Compose(attr.Class("dropdown-item"), opts...)),
	)
}

// DropdownDivider creates a Bootstrap dropdown divider.
func DropdownDivider() core.Node {
	return list.LI(text.HR(attr.Class("dropdown-divider")))
}

// DropdownHeader creates a Bootstrap dropdown header.
func DropdownHeader(opts ...core.Node) core.Node {
	return list.LI(
		section.H6(core.Compose(attr.Class("dropdown-header"), opts...)),
	)
}
