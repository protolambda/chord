// Package form provides HTML form structure elements.
package form

import (
	"github.com/protolambda/chord/core/attr"
	"github.com/protolambda/chord/core/elem"
)

// Form creates a form element.
func Form(attrs ...attr.Node) elem.Scope { return elem.Name("form").New(attrs...) }

// Fieldset creates a fieldset element.
func Fieldset(attrs ...attr.Node) elem.Scope { return elem.Name("fieldset").New(attrs...) }

// Legend creates a legend element.
func Legend(attrs ...attr.Node) elem.Scope { return elem.Name("legend").New(attrs...) }
