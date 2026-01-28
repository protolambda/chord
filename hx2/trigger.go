package hx2

import "github.com/protolambda/chord/core"

// Trigger sets the hx-trigger attribute to specify what triggers the request.
func Trigger(v string) core.Node { return core.Attribute("hx-trigger", v) }
