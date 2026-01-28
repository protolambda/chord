package on

import "github.com/protolambda/chord/core"

// Toggle sets the ontoggle event handler.
func Toggle(v string) core.Node { return core.Attribute("ontoggle", v) }

// BeforeToggle sets the onbeforetoggle event handler.
func BeforeToggle(v string) core.Node { return core.Attribute("onbeforetoggle", v) }
