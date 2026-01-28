package bs

import (
	"context"

	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/html/attr"
	"github.com/protolambda/chord/html/form/button"
	"github.com/protolambda/chord/html/group/div"
	"github.com/protolambda/chord/html/section"
)

// Offcanvas represents a Bootstrap offcanvas component.
type Offcanvas struct {
	ID        string    // required for targeting
	Title     core.Node // offcanvas-title
	Body      core.Node // offcanvas-body
	Placement string    // "start", "end", "top", "bottom"
	Backdrop  bool      // show backdrop
	Scroll    bool      // allow body scroll
	Attrs     core.Node
}

// Eval builds the offcanvas structure with proper Bootstrap markup.
func (o Offcanvas) Eval(ctx context.Context) (core.Obj, error) {
	placement := o.Placement
	if placement == "" {
		placement = "start"
	}

	var children []core.Node
	if o.Attrs != nil {
		children = append(children, o.Attrs)
	}
	children = append(children,
		attr.Class("offcanvas offcanvas-"+placement),
		attr.ID(o.ID),
		attr.Tabindex("-1"),
	)

	if o.Scroll {
		children = append(children, attr.Data("bs-scroll", "true"))
	}
	if !o.Backdrop {
		children = append(children, attr.Data("bs-backdrop", "false"))
	}

	header := div.Div(
		attr.Class("offcanvas-header"),
		section.H5(attr.Class("offcanvas-title"), o.Title),
		button.Button(
			attr.Class("btn-close"),
			button.Type(button.TypeButton),
			attr.Data("bs-dismiss", "offcanvas"),
		),
	)

	body := div.Div(attr.Class("offcanvas-body"), o.Body)

	children = append(children, header, body)
	return div.Div(children...).Eval(ctx)
}

// OffcanvasTrigger creates a button that triggers an offcanvas.
func OffcanvasTrigger(targetID string, opts ...core.Node) core.Node {
	return button.Button(core.Compose(
		button.Type(button.TypeButton),
		attr.Data("bs-toggle", "offcanvas"),
		attr.Data("bs-target", "#"+targetID),
	), core.Bundle(opts...))
}
