// Package on provides DOM event handler attributes.
package on

import (
	"github.com/protolambda/chord/core/attr"
)

// Click sets the onclick event handler.
func Click(v string) attr.Node { return attr.KV("onclick", v) }

// DblClick sets the ondblclick event handler.
func DblClick(v string) attr.Node { return attr.KV("ondblclick", v) }

// MouseDown sets the onmousedown event handler.
func MouseDown(v string) attr.Node { return attr.KV("onmousedown", v) }

// MouseUp sets the onmouseup event handler.
func MouseUp(v string) attr.Node { return attr.KV("onmouseup", v) }

// MouseMove sets the onmousemove event handler.
func MouseMove(v string) attr.Node { return attr.KV("onmousemove", v) }

// MouseOver sets the onmouseover event handler.
func MouseOver(v string) attr.Node { return attr.KV("onmouseover", v) }

// MouseOut sets the onmouseout event handler.
func MouseOut(v string) attr.Node { return attr.KV("onmouseout", v) }

// MouseEnter sets the onmouseenter event handler.
func MouseEnter(v string) attr.Node { return attr.KV("onmouseenter", v) }

// MouseLeave sets the onmouseleave event handler.
func MouseLeave(v string) attr.Node { return attr.KV("onmouseleave", v) }

// ContextMenu sets the oncontextmenu event handler.
func ContextMenu(v string) attr.Node { return attr.KV("oncontextmenu", v) }
