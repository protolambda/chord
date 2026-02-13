package bs

import (
	"github.com/protolambda/chord/core/attrib"
	"github.com/protolambda/chord/core/elem"
	"github.com/protolambda/chord/html/attr"
	"github.com/protolambda/chord/html/form/label"
	"github.com/protolambda/chord/html/group/div"
	"github.com/protolambda/chord/html/text"
)

// FormGroup creates a Bootstrap form group (mb-3 div wrapper).
func FormGroup(attrs ...attrib.Node) elem.Scope {
	return div.Div(attrib.Cons(attr.Class("mb-3"), attrs...))
}

// FormLabel creates a Bootstrap form label.
func FormLabel(attrs ...attrib.Node) elem.Scope {
	return label.Label(attrib.Cons(attr.Class("form-label"), attrs...))
}

// FormCheck creates a Bootstrap form check wrapper.
func FormCheck(attrs ...attrib.Node) elem.Scope {
	return div.Div(attrib.Cons(attr.Class("form-check"), attrs...))
}

// FormCheckLabel creates a Bootstrap form check label.
func FormCheckLabel(attrs ...attrib.Node) elem.Scope {
	return label.Label(attrib.Cons(attr.Class("form-check-label"), attrs...))
}

// FormText creates a Bootstrap form text helper.
func FormText(attrs ...attrib.Node) elem.Scope {
	return div.Div(attrib.Cons(attr.Class("form-text"), attrs...))
}

// InputGroup creates a Bootstrap input group.
func InputGroup(attrs ...attrib.Node) elem.Scope {
	return div.Div(attrib.Cons(attr.Class("input-group"), attrs...))
}

// InputGroupText creates a Bootstrap input group text addon.
func InputGroupText(attrs ...attrib.Node) elem.Scope {
	return text.Span(attrib.Cons(attr.Class("input-group-text"), attrs...))
}
