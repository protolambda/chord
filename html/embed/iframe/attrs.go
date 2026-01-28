package iframe

import "github.com/protolambda/chord/core"

// Src sets the src attribute.
func Src(v string) core.Node { return core.Attribute("src", v) }

// Srcdoc sets the srcdoc attribute.
func Srcdoc(v string) core.Node { return core.Attribute("srcdoc", v) }

// Name sets the name attribute.
func Name(v string) core.Node { return core.Attribute("name", v) }

// Width sets the width attribute.
func Width(v string) core.Node { return core.Attribute("width", v) }

// Height sets the height attribute.
func Height(v string) core.Node { return core.Attribute("height", v) }

// Loading sets the loading attribute.
func Loading(v string) core.Node { return core.Attribute("loading", v) }

// Sandbox sets the sandbox attribute.
func Sandbox(v string) core.Node { return core.Attribute("sandbox", v) }

// Allow sets the allow attribute.
func Allow(v string) core.Node { return core.Attribute("allow", v) }

// Allowfullscreen sets the allowfullscreen boolean attribute.
func Allowfullscreen() core.Node { return core.BoolAttribute("allowfullscreen") }

// Referrerpolicy sets the referrerpolicy attribute.
func Referrerpolicy(v string) core.Node { return core.Attribute("referrerpolicy", v) }
