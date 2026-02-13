package bs

import (
	"context"

	"github.com/protolambda/chord/core/attrib"
	"github.com/protolambda/chord/core/elem"
	"github.com/protolambda/chord/html/attr"
	"github.com/protolambda/chord/html/group/div"
	"github.com/protolambda/chord/html/section"
	"github.com/protolambda/chord/html/text"
)

// Card represents a Bootstrap card component with named fields for clarity.
type Card struct {
	Header elem.Node   // optional card-header content
	Img    elem.Node   // optional card-img-top
	Body   elem.Node   // card-body content (use CardBody for structured body)
	Footer elem.Node   // optional card-footer content
	Attrs  attrib.Node // additional attributes
}

func (Card) ChordNode() {}

// Eval builds the card structure with proper Bootstrap markup.
func (c Card) Eval(ctx context.Context) (elem.Obj, error) {
	var cardAttrs []attrib.Node
	if c.Attrs != nil {
		cardAttrs = append(cardAttrs, c.Attrs)
	}
	cardAttrs = append(cardAttrs, attr.Class("card"))

	var children []elem.Node
	if c.Header != nil {
		children = append(children, div.Div(attr.Class("card-header"))(c.Header))
	}
	if c.Img != nil {
		children = append(children, c.Img)
	}
	if c.Body != nil {
		children = append(children, div.Div(attr.Class("card-body"))(c.Body))
	}
	if c.Footer != nil {
		children = append(children, div.Div(attr.Class("card-footer"))(c.Footer))
	}

	return div.Div(cardAttrs...)(children...).Eval(ctx)
}

// CardBody represents a structured Bootstrap card body.
type CardBody struct {
	Title    elem.Node // card-title (typically h5)
	Subtitle elem.Node // card-subtitle (typically h6)
	Text     elem.Node // card-text content
	Content  elem.Node // additional body content
}

func (CardBody) ChordNode() {}

// Eval builds the card body structure as a fragment.
func (cb CardBody) Eval(ctx context.Context) (elem.Obj, error) {
	var children []elem.Node

	if cb.Title != nil {
		children = append(children, section.H5(attr.Class("card-title"))(cb.Title))
	}
	if cb.Subtitle != nil {
		children = append(children, section.H6(attr.Class("card-subtitle mb-2 text-body-secondary"))(cb.Subtitle))
	}
	if cb.Text != nil {
		children = append(children, text.P(attr.Class("card-text"))(cb.Text))
	}
	if cb.Content != nil {
		children = append(children, cb.Content)
	}

	return elem.Bundle(children).Eval(ctx)
}

// CardImgOverlay creates a card image overlay wrapper.
func CardImgOverlay(attrs ...attrib.Node) elem.Scope {
	return div.Div(attrib.Cons(attr.Class("card-img-overlay"), attrs...))
}
