package on

import (
	"github.com/protolambda/chord/core/attr"
)

// Load sets the onload event handler.
func Load(v string) attr.Node { return attr.KV("onload", v) }

// Error sets the onerror event handler.
func Error(v string) attr.Node { return attr.KV("onerror", v) }

// Abort sets the onabort event handler.
func Abort(v string) attr.Node { return attr.KV("onabort", v) }
