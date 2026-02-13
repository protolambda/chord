package hx2

import (
	"github.com/protolambda/chord/core/attrib"
)

// Trigger sets the hx-trigger attribute to specify what triggers the request.
func Trigger(v string) attrib.Node { return attrib.KV("hx-trigger", v) }
