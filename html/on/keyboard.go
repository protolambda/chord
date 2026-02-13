package on

import (
	"github.com/protolambda/chord/core/attrib"
)

// KeyDown sets the onkeydown event handler.
func KeyDown(v string) attrib.Node { return attrib.KV("onkeydown", v) }

// KeyUp sets the onkeyup event handler.
func KeyUp(v string) attrib.Node { return attrib.KV("onkeyup", v) }

// KeyPress sets the onkeypress event handler.
func KeyPress(v string) attrib.Node { return attrib.KV("onkeypress", v) }
