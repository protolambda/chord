// Package on provides DOM event handler attributes.
package on

import "github.com/protolambda/chord/core"

// Click sets the onclick event handler.
func Click(v string) core.Node { return core.Attribute("onclick", v) }

// DblClick sets the ondblclick event handler.
func DblClick(v string) core.Node { return core.Attribute("ondblclick", v) }

// MouseDown sets the onmousedown event handler.
func MouseDown(v string) core.Node { return core.Attribute("onmousedown", v) }

// MouseUp sets the onmouseup event handler.
func MouseUp(v string) core.Node { return core.Attribute("onmouseup", v) }

// MouseMove sets the onmousemove event handler.
func MouseMove(v string) core.Node { return core.Attribute("onmousemove", v) }

// MouseOver sets the onmouseover event handler.
func MouseOver(v string) core.Node { return core.Attribute("onmouseover", v) }

// MouseOut sets the onmouseout event handler.
func MouseOut(v string) core.Node { return core.Attribute("onmouseout", v) }

// MouseEnter sets the onmouseenter event handler.
func MouseEnter(v string) core.Node { return core.Attribute("onmouseenter", v) }

// MouseLeave sets the onmouseleave event handler.
func MouseLeave(v string) core.Node { return core.Attribute("onmouseleave", v) }

// ContextMenu sets the oncontextmenu event handler.
func ContextMenu(v string) core.Node { return core.Attribute("oncontextmenu", v) }
