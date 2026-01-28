package textarea

import "github.com/protolambda/chord/core"

// Name sets the name attribute.
func Name(v string) core.Node { return core.Attribute("name", v) }

// Value sets the value attribute.
func Value(v string) core.Node { return core.Attribute("value", v) }

// Placeholder sets the placeholder attribute.
func Placeholder(v string) core.Node { return core.Attribute("placeholder", v) }

// Required sets the required boolean attribute.
func Required() core.Node { return core.BoolAttribute("required") }

// Disabled sets the disabled boolean attribute.
func Disabled() core.Node { return core.BoolAttribute("disabled") }

// Readonly sets the readonly boolean attribute.
func Readonly() core.Node { return core.BoolAttribute("readonly") }

// Autofocus sets the autofocus boolean attribute.
func Autofocus() core.Node { return core.BoolAttribute("autofocus") }

// Autocomplete sets the autocomplete attribute.
func Autocomplete(v string) core.Node { return core.Attribute("autocomplete", v) }
