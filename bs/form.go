package bs

import (
	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/html/attr"
	"github.com/protolambda/chord/html/form/label"
	"github.com/protolambda/chord/html/group/div"
	"github.com/protolambda/chord/html/text"
)

// FormGroup creates a Bootstrap form group (mb-3 div wrapper).
func FormGroup(opts ...core.Node) core.Node {
	return div.Div(core.Compose(attr.Class("mb-3"), opts...))
}

// FormLabel creates a Bootstrap form label.
func FormLabel(opts ...core.Node) core.Node {
	return label.Label(core.Compose(attr.Class("form-label"), opts...))
}

// FormCheck creates a Bootstrap form check wrapper.
func FormCheck(opts ...core.Node) core.Node {
	return div.Div(core.Compose(attr.Class("form-check"), opts...))
}

// FormCheckLabel creates a Bootstrap form check label.
func FormCheckLabel(opts ...core.Node) core.Node {
	return label.Label(core.Compose(attr.Class("form-check-label"), opts...))
}

// FormText creates a Bootstrap form text helper.
func FormText(opts ...core.Node) core.Node {
	return div.Div(core.Compose(attr.Class("form-text"), opts...))
}

// InputGroup creates a Bootstrap input group.
func InputGroup(opts ...core.Node) core.Node {
	return div.Div(core.Compose(attr.Class("input-group"), opts...))
}

// InputGroupText creates a Bootstrap input group text addon.
func InputGroupText(opts ...core.Node) core.Node {
	return text.Span(core.Compose(attr.Class("input-group-text"), opts...))
}
