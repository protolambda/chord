package hx1

import "github.com/protolambda/chord/core"

// Indicator sets the hx-indicator attribute to specify loading indicator element.
func Indicator(v string) core.Node { return core.Attribute("hx-indicator", v) }
