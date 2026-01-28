package list

import "github.com/protolambda/chord/core"

// Reversed sets the reversed boolean attribute for ol elements.
func Reversed() core.Node { return core.BoolAttribute("reversed") }

// Start sets the start attribute for ol elements.
func Start(v string) core.Node { return core.Attribute("start", v) }

// Type sets the type attribute for ol elements.
func Type(v string) core.Node { return core.Attribute("type", v) }
