package bs

import (
	"context"
	"fmt"

	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/html/attr"
	"github.com/protolambda/chord/html/form/button"
	"github.com/protolambda/chord/html/group/div"
	"github.com/protolambda/chord/html/section"
)

// Accordion represents a Bootstrap accordion component.
type Accordion struct {
	ID         string          // required for collapse targeting
	Items      []AccordionItem // accordion items
	Flush      bool            // flush style (no borders)
	AlwaysOpen bool            // allow multiple open
	Attrs      core.Node
}

// AccordionItem represents a single accordion item.
type AccordionItem struct {
	Header core.Node // accordion-header content
	Body   core.Node // accordion-body content
	Show   bool      // initially expanded
}

// Eval builds the accordion structure with proper Bootstrap markup.
func (a Accordion) Eval(ctx context.Context) (core.Obj, error) {
	accordionClass := "accordion"
	if a.Flush {
		accordionClass += " accordion-flush"
	}

	var children []core.Node
	if a.Attrs != nil {
		children = append(children, a.Attrs)
	}
	children = append(children, attr.Class(accordionClass), attr.ID(a.ID))

	for i, item := range a.Items {
		itemID := fmt.Sprintf("%s-item-%d", a.ID, i)
		collapseID := fmt.Sprintf("%s-collapse-%d", a.ID, i)

		buttonClass := "accordion-button"
		collapseClass := "accordion-collapse collapse"
		if item.Show {
			collapseClass += " show"
		} else {
			buttonClass += " collapsed"
		}

		var buttonAttrs []core.Node
		buttonAttrs = append(buttonAttrs,
			attr.Class(buttonClass),
			button.Type(button.TypeButton),
			attr.Data("bs-toggle", "collapse"),
			attr.Data("bs-target", "#"+collapseID),
		)

		header := section.H2(
			attr.Class("accordion-header"),
			button.Button(append(buttonAttrs, item.Header)...),
		)

		var collapseAttrs []core.Node
		collapseAttrs = append(collapseAttrs,
			attr.ID(collapseID),
			attr.Class(collapseClass),
		)
		if !a.AlwaysOpen {
			collapseAttrs = append(collapseAttrs, attr.Data("bs-parent", "#"+a.ID))
		}

		collapse := div.Div(
			append(collapseAttrs, div.Div(attr.Class("accordion-body"), item.Body))...,
		)

		children = append(children, div.Div(
			attr.Class("accordion-item"),
			attr.ID(itemID),
			header,
			collapse,
		))
	}

	return div.Div(children...).Eval(ctx)
}
