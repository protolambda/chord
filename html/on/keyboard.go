package on

import "github.com/protolambda/chord/core"

// KeyDown sets the onkeydown event handler.
func KeyDown(v string) core.Node { return core.Attribute("onkeydown", v) }

// KeyUp sets the onkeyup event handler.
func KeyUp(v string) core.Node { return core.Attribute("onkeyup", v) }

// KeyPress sets the onkeypress event handler.
func KeyPress(v string) core.Node { return core.Attribute("onkeypress", v) }
