package list

import "github.com/protolambda/chord/core"

// Value sets the value attribute for li elements.
func Value(v string) core.Node { return core.Attribute("value", v) }
