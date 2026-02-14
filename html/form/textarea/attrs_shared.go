package textarea

import (
	"html"

	"github.com/protolambda/chord/core/attr"
)

// Name sets the name attribute.
func Name(v string) attr.Node { return attr.KV("name", v) }

// Value sets the value attribute.
func Value(v string) attr.Node { return attr.KV("value", v) }

// Placeholder sets the placeholder attribute.
func Placeholder(v string) attr.Node { return attr.KV("placeholder", html.EscapeString(v)) }

// Required sets the required boolean attribute.
func Required() attr.Node { return attr.Bool("required") }

// Disabled sets the disabled boolean attribute.
func Disabled() attr.Node { return attr.Bool("disabled") }

// Readonly sets the readonly boolean attribute.
func Readonly() attr.Node { return attr.Bool("readonly") }

// Autofocus sets the autofocus boolean attribute.
func Autofocus() attr.Node { return attr.Bool("autofocus") }

// Autocomplete sets the autocomplete attribute.
func Autocomplete(v string) attr.Node { return attr.KV("autocomplete", v) }
