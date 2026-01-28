package form

import "github.com/protolambda/chord/core"

// Action sets the action attribute.
func Action(v string) core.Node { return core.Attribute("action", v) }

// Target sets the target attribute.
func Target(v string) core.Node { return core.Attribute("target", v) }

// Novalidate sets the novalidate boolean attribute.
func Novalidate() core.Node { return core.BoolAttribute("novalidate") }

// Accept sets the accept attribute.
func Accept(v string) core.Node { return core.Attribute("accept", v) }

// Acceptcharset sets the accept-charset attribute.
func Acceptcharset(v string) core.Node { return core.Attribute("accept-charset", v) }
