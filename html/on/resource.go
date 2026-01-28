package on

import "github.com/protolambda/chord/core"

// Load sets the onload event handler.
func Load(v string) core.Node { return core.Attribute("onload", v) }

// Error sets the onerror event handler.
func Error(v string) core.Node { return core.Attribute("onerror", v) }

// Abort sets the onabort event handler.
func Abort(v string) core.Node { return core.Attribute("onabort", v) }
