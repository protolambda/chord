package ba

import (
	"github.com/protolambda/chord/core/attrib"
	"github.com/protolambda/chord/html/attr"
)

// FormControl creates a Bootstrap form-control class attribute.
func FormControl(opts ...attrib.Node) attrib.Node {
	return attrib.Cons(attr.Class("form-control"), opts...)
}

// FormSelect creates a Bootstrap form-select class attribute.
func FormSelect(opts ...attrib.Node) attrib.Node {
	return attrib.Cons(attr.Class("form-select"), opts...)
}

// FormCheckInput creates a Bootstrap form-check-input class attribute.
func FormCheckInput(opts ...attrib.Node) attrib.Node {
	return attrib.Cons(attr.Class("form-check-input"), opts...)
}
