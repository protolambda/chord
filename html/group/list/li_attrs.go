package list

import (
	"github.com/protolambda/chord/core/attr"
)

// Value sets the value attribute for li elements.
func Value(v string) attr.Node { return attr.KV("value", v) }
