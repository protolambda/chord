package selectel

import (
	"github.com/protolambda/chord/core/attr"
)

// Selected sets the selected boolean attribute.
func Selected() attr.Node { return attr.Name("selected").Bool() }

// Multiple sets the multiple boolean attribute.
func Multiple() attr.Node { return attr.Name("multiple").Bool() }

// Label sets the label attribute for optgroup/option elements.
func Label(v string) attr.Node { return attr.Name("label").Value(v) }
