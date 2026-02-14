package on

import (
	"github.com/protolambda/chord/core/attr"
)

// Toggle sets the ontoggle event handler.
func Toggle(v string) attr.Node { return attr.KV("ontoggle", v) }

// BeforeToggle sets the onbeforetoggle event handler.
func BeforeToggle(v string) attr.Node { return attr.KV("onbeforetoggle", v) }
