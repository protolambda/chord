package hx2

import (
	"github.com/protolambda/chord/core/attr"
)

// Trigger sets the hx-trigger attribute to specify what triggers the request.
func Trigger(v string) attr.Node { return attr.KV("hx-trigger", v) }
