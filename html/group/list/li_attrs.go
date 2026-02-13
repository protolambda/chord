package list

import (
	"github.com/protolambda/chord/core/attrib"
)

// Value sets the value attribute for li elements.
func Value(v string) attrib.Node { return attrib.KV("value", v) }
