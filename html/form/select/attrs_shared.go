package selectel

import (
	"github.com/protolambda/chord/core/attr"
)

// Name sets the name attribute.
func Name(v string) attr.Node { return attr.Name("name").Value(v) }

// Value sets the value attribute.
func Value(v string) attr.Node { return attr.Name("value").Value(v) }

// Required sets the required boolean attribute.
func Required() attr.Node { return attr.Name("required").Bool() }

// Disabled sets the disabled boolean attribute.
func Disabled() attr.Node { return attr.Name("disabled").Bool() }

// Autofocus sets the autofocus boolean attribute.
func Autofocus() attr.Node { return attr.Name("autofocus").Bool() }

// Size sets the size attribute.
func Size(v string) attr.Node { return attr.Name("size").Value(v) }
