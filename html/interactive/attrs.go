package interactive

import (
	"github.com/protolambda/chord/core/attr"
)

// Open sets the open boolean attribute.
func Open() attr.Node { return attr.Name("open").Bool() }
