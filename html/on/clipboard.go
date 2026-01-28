package on

import "github.com/protolambda/chord/core"

// Copy sets the oncopy event handler.
func Copy(v string) core.Node { return core.Attribute("oncopy", v) }

// Cut sets the oncut event handler.
func Cut(v string) core.Node { return core.Attribute("oncut", v) }

// Paste sets the onpaste event handler.
func Paste(v string) core.Node { return core.Attribute("onpaste", v) }
