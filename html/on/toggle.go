package on

import (
	"github.com/protolambda/chord/core/attrib"
)

// Toggle sets the ontoggle event handler.
func Toggle(v string) attrib.Node { return attrib.KV("ontoggle", v) }

// BeforeToggle sets the onbeforetoggle event handler.
func BeforeToggle(v string) attrib.Node { return attrib.KV("onbeforetoggle", v) }
