// Package on provides DOM event handler attributes.
package on

import (
	"github.com/protolambda/chord/core/attrib"
)

// Click sets the onclick event handler.
func Click(v string) attrib.Node { return attrib.KV("onclick", v) }

// DblClick sets the ondblclick event handler.
func DblClick(v string) attrib.Node { return attrib.KV("ondblclick", v) }

// MouseDown sets the onmousedown event handler.
func MouseDown(v string) attrib.Node { return attrib.KV("onmousedown", v) }

// MouseUp sets the onmouseup event handler.
func MouseUp(v string) attrib.Node { return attrib.KV("onmouseup", v) }

// MouseMove sets the onmousemove event handler.
func MouseMove(v string) attrib.Node { return attrib.KV("onmousemove", v) }

// MouseOver sets the onmouseover event handler.
func MouseOver(v string) attrib.Node { return attrib.KV("onmouseover", v) }

// MouseOut sets the onmouseout event handler.
func MouseOut(v string) attrib.Node { return attrib.KV("onmouseout", v) }

// MouseEnter sets the onmouseenter event handler.
func MouseEnter(v string) attrib.Node { return attrib.KV("onmouseenter", v) }

// MouseLeave sets the onmouseleave event handler.
func MouseLeave(v string) attrib.Node { return attrib.KV("onmouseleave", v) }

// ContextMenu sets the oncontextmenu event handler.
func ContextMenu(v string) attrib.Node { return attrib.KV("oncontextmenu", v) }
