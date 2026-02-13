package selectel

import (
	"github.com/protolambda/chord/core/attrib"
)

// Name sets the name attribute.
func Name(v string) attrib.Node { return attrib.KV("name", v) }

// Value sets the value attribute.
func Value(v string) attrib.Node { return attrib.KV("value", v) }

// Required sets the required boolean attribute.
func Required() attrib.Node { return attrib.Bool("required") }

// Disabled sets the disabled boolean attribute.
func Disabled() attrib.Node { return attrib.Bool("disabled") }

// Autofocus sets the autofocus boolean attribute.
func Autofocus() attrib.Node { return attrib.Bool("autofocus") }

// Size sets the size attribute.
func Size(v string) attrib.Node { return attrib.KV("size", v) }
