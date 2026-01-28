// Package form provides HTML form structure elements.
package form

import "github.com/protolambda/chord/core"

// Form creates a form element.
func Form(opts ...core.Node) core.Node { return core.Element("form", opts...) }

// Fieldset creates a fieldset element.
func Fieldset(opts ...core.Node) core.Node { return core.Element("fieldset", opts...) }

// Legend creates a legend element.
func Legend(opts ...core.Node) core.Node { return core.Element("legend", opts...) }
