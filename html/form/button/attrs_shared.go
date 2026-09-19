package button

import (
	"github.com/protolambda/chord/core/attr"
)

// Name sets the name attribute.
func Name(v string) attr.Node { return attr.Name("name").Value(v) }

// Value sets the value attribute.
func Value(v string) attr.Node { return attr.Name("value").Value(v) }

// Disabled sets the disabled boolean attribute.
func Disabled() attr.Node { return attr.Name("disabled").Bool() }

// Autofocus sets the autofocus boolean attribute.
func Autofocus() attr.Node { return attr.Name("autofocus").Bool() }
