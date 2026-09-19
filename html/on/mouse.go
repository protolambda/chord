// Package on provides DOM event handler attributes.
package on

import (
	"github.com/protolambda/chord/core/attr"
)

// Click sets the onclick event handler.
func Click(v string) attr.Node { return attr.Name("onclick").Value(v) }

// DblClick sets the ondblclick event handler.
func DblClick(v string) attr.Node { return attr.Name("ondblclick").Value(v) }

// MouseDown sets the onmousedown event handler.
func MouseDown(v string) attr.Node { return attr.Name("onmousedown").Value(v) }

// MouseUp sets the onmouseup event handler.
func MouseUp(v string) attr.Node { return attr.Name("onmouseup").Value(v) }

// MouseMove sets the onmousemove event handler.
func MouseMove(v string) attr.Node { return attr.Name("onmousemove").Value(v) }

// MouseOver sets the onmouseover event handler.
func MouseOver(v string) attr.Node { return attr.Name("onmouseover").Value(v) }

// MouseOut sets the onmouseout event handler.
func MouseOut(v string) attr.Node { return attr.Name("onmouseout").Value(v) }

// MouseEnter sets the onmouseenter event handler.
func MouseEnter(v string) attr.Node { return attr.Name("onmouseenter").Value(v) }

// MouseLeave sets the onmouseleave event handler.
func MouseLeave(v string) attr.Node { return attr.Name("onmouseleave").Value(v) }

// ContextMenu sets the oncontextmenu event handler.
func ContextMenu(v string) attr.Node { return attr.Name("oncontextmenu").Value(v) }
