package label

import "github.com/protolambda/chord/core"

// For sets the for attribute.
func For(v string) core.Node { return core.Attribute("for", v) }
