package text

import "github.com/protolambda/chord/core"

// Value sets the value attribute for data elements.
func Value(v string) core.Node { return core.Attribute("value", v) }
