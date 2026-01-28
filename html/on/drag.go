package on

import "github.com/protolambda/chord/core"

// Drag sets the ondrag event handler.
func Drag(v string) core.Node { return core.Attribute("ondrag", v) }

// DragStart sets the ondragstart event handler.
func DragStart(v string) core.Node { return core.Attribute("ondragstart", v) }

// DragEnd sets the ondragend event handler.
func DragEnd(v string) core.Node { return core.Attribute("ondragend", v) }

// DragEnter sets the ondragenter event handler.
func DragEnter(v string) core.Node { return core.Attribute("ondragenter", v) }

// DragLeave sets the ondragleave event handler.
func DragLeave(v string) core.Node { return core.Attribute("ondragleave", v) }

// DragOver sets the ondragover event handler.
func DragOver(v string) core.Node { return core.Attribute("ondragover", v) }

// Drop sets the ondrop event handler.
func Drop(v string) core.Node { return core.Attribute("ondrop", v) }
