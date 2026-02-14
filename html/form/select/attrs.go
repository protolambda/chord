package selectel

import (
	"github.com/protolambda/chord/core/attr"
)

// Selected sets the selected boolean attribute.
func Selected() attr.Node { return attr.Bool("selected") }

// Multiple sets the multiple boolean attribute.
func Multiple() attr.Node { return attr.Bool("multiple") }

// Label sets the label attribute for optgroup/option elements.
func Label(v string) attr.Node { return attr.KV("label", v) }
