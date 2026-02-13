package interactive

import (
	"github.com/protolambda/chord/core/attrib"
)

// Open sets the open boolean attribute.
func Open() attrib.Node { return attrib.Bool("open") }
