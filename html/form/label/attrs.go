package label

import (
	"github.com/protolambda/chord/core/attrib"
)

// For sets the for attribute.
func For(v string) attrib.Node { return attrib.KV("for", v) }
