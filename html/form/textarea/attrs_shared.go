package textarea

import (
	"html"

	"github.com/protolambda/chord/core/attrib"
)

// Name sets the name attribute.
func Name(v string) attrib.Node { return attrib.KV("name", v) }

// Value sets the value attribute.
func Value(v string) attrib.Node { return attrib.KV("value", v) }

// Placeholder sets the placeholder attribute.
func Placeholder(v string) attrib.Node { return attrib.KV("placeholder", html.EscapeString(v)) }

// Required sets the required boolean attribute.
func Required() attrib.Node { return attrib.Bool("required") }

// Disabled sets the disabled boolean attribute.
func Disabled() attrib.Node { return attrib.Bool("disabled") }

// Readonly sets the readonly boolean attribute.
func Readonly() attrib.Node { return attrib.Bool("readonly") }

// Autofocus sets the autofocus boolean attribute.
func Autofocus() attrib.Node { return attrib.Bool("autofocus") }

// Autocomplete sets the autocomplete attribute.
func Autocomplete(v string) attrib.Node { return attrib.KV("autocomplete", v) }
