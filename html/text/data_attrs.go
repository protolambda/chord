package text

import (
	"github.com/protolambda/chord/core/attrib"
)

// Value sets the value attribute for data elements.
func Value(v string) attrib.Node { return attrib.KV("value", v) }
