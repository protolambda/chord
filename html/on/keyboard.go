package on

import (
	"github.com/protolambda/chord/core/attr"
)

// KeyDown sets the onkeydown event handler.
func KeyDown(v string) attr.Node { return attr.Name("onkeydown").Value(v) }

// KeyUp sets the onkeyup event handler.
func KeyUp(v string) attr.Node { return attr.Name("onkeyup").Value(v) }

// KeyPress sets the onkeypress event handler.
func KeyPress(v string) attr.Node { return attr.Name("onkeypress").Value(v) }
