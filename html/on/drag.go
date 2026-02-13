package on

import (
	"github.com/protolambda/chord/core/attrib"
)

// Drag sets the ondrag event handler.
func Drag(v string) attrib.Node { return attrib.KV("ondrag", v) }

// DragStart sets the ondragstart event handler.
func DragStart(v string) attrib.Node { return attrib.KV("ondragstart", v) }

// DragEnd sets the ondragend event handler.
func DragEnd(v string) attrib.Node { return attrib.KV("ondragend", v) }

// DragEnter sets the ondragenter event handler.
func DragEnter(v string) attrib.Node { return attrib.KV("ondragenter", v) }

// DragLeave sets the ondragleave event handler.
func DragLeave(v string) attrib.Node { return attrib.KV("ondragleave", v) }

// DragOver sets the ondragover event handler.
func DragOver(v string) attrib.Node { return attrib.KV("ondragover", v) }

// Drop sets the ondrop event handler.
func Drop(v string) attrib.Node { return attrib.KV("ondrop", v) }
