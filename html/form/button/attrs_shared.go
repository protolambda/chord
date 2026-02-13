package button

import (
	"github.com/protolambda/chord/core/attrib"
)

// Name sets the name attribute.
func Name(v string) attrib.Node { return attrib.KV("name", v) }

// Value sets the value attribute.
func Value(v string) attrib.Node { return attrib.KV("value", v) }

// Disabled sets the disabled boolean attribute.
func Disabled() attrib.Node { return attrib.Bool("disabled") }

// Autofocus sets the autofocus boolean attribute.
func Autofocus() attrib.Node { return attrib.Bool("autofocus") }
