package selectel

import (
	"github.com/protolambda/chord/core/attrib"
)

// Selected sets the selected boolean attribute.
func Selected() attrib.Node { return attrib.Bool("selected") }

// Multiple sets the multiple boolean attribute.
func Multiple() attrib.Node { return attrib.Bool("multiple") }

// Label sets the label attribute for optgroup/option elements.
func Label(v string) attrib.Node { return attrib.KV("label", v) }
