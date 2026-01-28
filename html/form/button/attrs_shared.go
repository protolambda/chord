package button

import "github.com/protolambda/chord/core"

// Name sets the name attribute.
func Name(v string) core.Node { return core.Attribute("name", v) }

// Value sets the value attribute.
func Value(v string) core.Node { return core.Attribute("value", v) }

// Disabled sets the disabled boolean attribute.
func Disabled() core.Node { return core.BoolAttribute("disabled") }

// Autofocus sets the autofocus boolean attribute.
func Autofocus() core.Node { return core.BoolAttribute("autofocus") }
