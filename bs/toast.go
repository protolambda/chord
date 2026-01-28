package bs

import (
	"context"

	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/html/attr"
	"github.com/protolambda/chord/html/form/button"
	"github.com/protolambda/chord/html/group/div"
)

// Toast represents a Bootstrap toast component.
type Toast struct {
	Header core.Node // toast-header content
	Body   core.Node // toast-body content
	Attrs  core.Node
}

// Eval builds the toast structure with proper Bootstrap markup.
func (t Toast) Eval(ctx context.Context) (core.Obj, error) {
	var children []core.Node
	if t.Attrs != nil {
		children = append(children, t.Attrs)
	}
	children = append(children, attr.Class("toast"))

	if t.Header != nil {
		header := div.Div(
			attr.Class("toast-header"),
			t.Header,
			button.Button(
				attr.Class("btn-close"),
				button.Type(button.TypeButton),
				attr.Data("bs-dismiss", "toast"),
			),
		)
		children = append(children, header)
	}

	if t.Body != nil {
		children = append(children, div.Div(attr.Class("toast-body"), t.Body))
	}

	return div.Div(children...).Eval(ctx)
}
