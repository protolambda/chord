// Package form provides HTML form structure elements.
package form

import (
	"github.com/protolambda/chord/core/attrib"
	"github.com/protolambda/chord/core/elem"
)

// Form creates a form element.
func Form(attrs ...attrib.Node) elem.Scope { return elem.New("form", attrs...) }

// Fieldset creates a fieldset element.
func Fieldset(attrs ...attrib.Node) elem.Scope { return elem.New("fieldset", attrs...) }

// Legend creates a legend element.
func Legend(attrs ...attrib.Node) elem.Scope { return elem.New("legend", attrs...) }
