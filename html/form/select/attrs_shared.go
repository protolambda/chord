package selectel

import (
	"github.com/protolambda/chord/core/attr"
)

// Name sets the name attribute.
func Name(v string) attr.Node { return attr.KV("name", v) }

// Value sets the value attribute.
func Value(v string) attr.Node { return attr.KV("value", v) }

// Required sets the required boolean attribute.
func Required() attr.Node { return attr.Bool("required") }

// Disabled sets the disabled boolean attribute.
func Disabled() attr.Node { return attr.Bool("disabled") }

// Autofocus sets the autofocus boolean attribute.
func Autofocus() attr.Node { return attr.Bool("autofocus") }

// Size sets the size attribute.
func Size(v string) attr.Node { return attr.KV("size", v) }
