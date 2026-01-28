package bs

import (
	"context"

	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/html/attr"
	"github.com/protolambda/chord/html/group/div"
	"github.com/protolambda/chord/html/section"
	"github.com/protolambda/chord/html/text"
)

// Card represents a Bootstrap card component with named fields for clarity.
type Card struct {
	Header core.Node // optional card-header content
	Img    core.Node // optional card-img-top
	Body   core.Node // card-body content (use CardBody for structured body)
	Footer core.Node // optional card-footer content
	Attrs  core.Node // additional attributes
}

// Eval builds the card structure with proper Bootstrap markup.
func (c Card) Eval(ctx context.Context) (core.Obj, error) {
	var children []core.Node

	if c.Attrs != nil {
		children = append(children, c.Attrs)
	}

	children = append(children, attr.Class("card"))

	if c.Header != nil {
		children = append(children, div.Div(attr.Class("card-header"), c.Header))
	}
	if c.Img != nil {
		children = append(children, c.Img)
	}
	if c.Body != nil {
		children = append(children, div.Div(attr.Class("card-body"), c.Body))
	}
	if c.Footer != nil {
		children = append(children, div.Div(attr.Class("card-footer"), c.Footer))
	}

	return div.Div(children...).Eval(ctx)
}

// CardBody represents a structured Bootstrap card body.
type CardBody struct {
	Title    core.Node // card-title (typically h5)
	Subtitle core.Node // card-subtitle (typically h6)
	Text     core.Node // card-text content
	Content  core.Node // additional body content
	Attrs    core.Node // additional attributes
}

// Eval builds the card body structure.
func (cb CardBody) Eval(ctx context.Context) (core.Obj, error) {
	var children []core.Node

	if cb.Attrs != nil {
		children = append(children, cb.Attrs)
	}

	if cb.Title != nil {
		children = append(children, section.H5(attr.Class("card-title"), cb.Title))
	}
	if cb.Subtitle != nil {
		children = append(children, section.H6(attr.Class("card-subtitle mb-2 text-body-secondary"), cb.Subtitle))
	}
	if cb.Text != nil {
		children = append(children, text.P(attr.Class("card-text"), cb.Text))
	}
	if cb.Content != nil {
		children = append(children, cb.Content)
	}

	return core.Bundle(children...).Eval(ctx)
}

// CardImgOverlay creates a card image overlay wrapper.
func CardImgOverlay(opts ...core.Node) core.Node {
	return div.Div(core.Compose(attr.Class("card-img-overlay"), opts...))
}
