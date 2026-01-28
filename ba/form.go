package ba

import (
	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/html/attr"
)

// FormControl creates a Bootstrap form-control class attribute.
func FormControl(opts ...core.Node) core.Node {
	return core.Compose(attr.Class("form-control"), opts...)
}

// FormSelect creates a Bootstrap form-select class attribute.
func FormSelect(opts ...core.Node) core.Node {
	return core.Compose(attr.Class("form-select"), opts...)
}

// FormCheckInput creates a Bootstrap form-check-input class attribute.
func FormCheckInput(opts ...core.Node) core.Node {
	return core.Compose(attr.Class("form-check-input"), opts...)
}
