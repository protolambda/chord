package on

import (
	"github.com/protolambda/chord/core/attrib"
)

// Load sets the onload event handler.
func Load(v string) attrib.Node { return attrib.KV("onload", v) }

// Error sets the onerror event handler.
func Error(v string) attrib.Node { return attrib.KV("onerror", v) }

// Abort sets the onabort event handler.
func Abort(v string) attrib.Node { return attrib.KV("onabort", v) }
