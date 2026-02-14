package label

import (
	"github.com/protolambda/chord/core/attr"
)

// For sets the for attribute.
func For(v string) attr.Node { return attr.KV("for", v) }
