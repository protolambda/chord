package selectel

import "github.com/protolambda/chord/core"

// Selected sets the selected boolean attribute.
func Selected() core.Node { return core.BoolAttribute("selected") }

// Multiple sets the multiple boolean attribute.
func Multiple() core.Node { return core.BoolAttribute("multiple") }

// Label sets the label attribute for optgroup/option elements.
func Label(v string) core.Node { return core.Attribute("label", v) }
