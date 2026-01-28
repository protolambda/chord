package bs

import (
	"context"

	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/html/attr"
	"github.com/protolambda/chord/html/form/button"
	"github.com/protolambda/chord/html/group/div"
	"github.com/protolambda/chord/html/section"
)

// Modal represents a Bootstrap modal component.
type Modal struct {
	ID         string    // required for targeting
	Title      core.Node // modal-title content
	Body       core.Node // modal-body content
	Footer     core.Node // modal-footer content (typically buttons)
	Size       string    // "", "sm", "lg", "xl"
	Centered   bool      // vertically centered
	Scrollable bool      // scrollable body
	Static     bool      // static backdrop
	Attrs      core.Node // additional attributes
}

// Eval builds the modal structure with proper Bootstrap markup.
func (m Modal) Eval(ctx context.Context) (core.Obj, error) {
	dialogClass := "modal-dialog"
	if m.Size != "" {
		dialogClass += " modal-" + m.Size
	}
	if m.Centered {
		dialogClass += " modal-dialog-centered"
	}
	if m.Scrollable {
		dialogClass += " modal-dialog-scrollable"
	}

	var modalAttrs []core.Node
	if m.Attrs != nil {
		modalAttrs = append(modalAttrs, m.Attrs)
	}
	modalAttrs = append(modalAttrs,
		attr.Class("modal fade"),
		attr.ID(m.ID),
		attr.Tabindex("-1"),
	)
	if m.Static {
		modalAttrs = append(modalAttrs, attr.Data("bs-backdrop", "static"))
	}

	var headerContent []core.Node
	if m.Title != nil {
		headerContent = append(headerContent, section.H1(attr.Class("modal-title fs-5"), m.Title))
	}
	headerContent = append(headerContent, button.Button(
		attr.Class("btn-close"),
		button.Type(button.TypeButton),
		attr.Data("bs-dismiss", "modal"),
	))

	var contentChildren []core.Node
	contentChildren = append(contentChildren,
		attr.Class("modal-content"),
		div.Div(core.Compose(attr.Class("modal-header"), headerContent...)),
	)
	if m.Body != nil {
		contentChildren = append(contentChildren, div.Div(attr.Class("modal-body"), m.Body))
	}
	if m.Footer != nil {
		contentChildren = append(contentChildren, div.Div(attr.Class("modal-footer"), m.Footer))
	}

	modalContent := div.Div(contentChildren...)
	modalDialog := div.Div(attr.Class(dialogClass), modalContent)

	return div.Div(append(modalAttrs, modalDialog)...).Eval(ctx)
}

// ModalTrigger creates a button that triggers a modal.
func ModalTrigger(targetID string, opts ...core.Node) core.Node {
	return button.Button(core.Compose(
		button.Type(button.TypeButton),
		attr.Data("bs-toggle", "modal"),
		attr.Data("bs-target", "#"+targetID),
	), core.Bundle(opts...))
}
