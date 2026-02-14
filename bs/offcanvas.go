package bs

import (
	"context"

	"github.com/protolambda/chord/core/attr"
	"github.com/protolambda/chord/core/elem"
	"github.com/protolambda/chord/html/form/button"
	"github.com/protolambda/chord/html/group/div"
	"github.com/protolambda/chord/html/section"
)

// Offcanvas represents a Bootstrap offcanvas component.
type Offcanvas struct {
	ID        string    // required for targeting
	Title     elem.Node // offcanvas-title
	Body      elem.Node // offcanvas-body
	Placement string    // "start", "end", "top", "bottom"
	Backdrop  bool      // show backdrop
	Scroll    bool      // allow body scroll
	Attrs     attr.Node
}

func (Offcanvas) ChordNode() {}

// Eval builds the offcanvas structure with proper Bootstrap markup.
func (o Offcanvas) Eval(ctx context.Context) (elem.Obj, error) {
	placement := o.Placement
	if placement == "" {
		placement = "start"
	}

	var attrs []attr.Node
	if o.Attrs != nil {
		attrs = append(attrs, o.Attrs)
	}
	attrs = append(attrs,
		attr.Class("offcanvas offcanvas-"+placement),
		attr.ID(o.ID),
		attr.Tabindex("-1"),
	)

	if o.Scroll {
		attrs = append(attrs, attr.Data("bs-scroll", "true"))
	}
	if !o.Backdrop {
		attrs = append(attrs, attr.Data("bs-backdrop", "false"))
	}

	header := div.Div(attr.Class("offcanvas-header"))(
		section.H5(attr.Class("offcanvas-title"))(o.Title),
		button.Button(
			attr.Class("btn-close"),
			button.Type(button.TypeButton),
			attr.Data("bs-dismiss", "offcanvas"),
		)(),
	)

	body := div.Div(attr.Class("offcanvas-body"))(o.Body)

	return div.Div(attrs...)(header, body).Eval(ctx)
}

// OffcanvasTrigger creates a button that triggers an offcanvas.
func OffcanvasTrigger(targetID string, attrs ...attr.Node) elem.Scope {
	return button.Button(attr.Cons(
		button.Type(button.TypeButton),
		attr.Data("bs-toggle", "offcanvas"),
		attr.Data("bs-target", "#"+targetID),
		attr.Bundle(attrs),
	))
}
