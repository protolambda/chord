package on

import (
	"github.com/protolambda/chord/core/attr"
)

// Drag sets the ondrag event handler.
func Drag(v string) attr.Node { return attr.KV("ondrag", v) }

// DragStart sets the ondragstart event handler.
func DragStart(v string) attr.Node { return attr.KV("ondragstart", v) }

// DragEnd sets the ondragend event handler.
func DragEnd(v string) attr.Node { return attr.KV("ondragend", v) }

// DragEnter sets the ondragenter event handler.
func DragEnter(v string) attr.Node { return attr.KV("ondragenter", v) }

// DragLeave sets the ondragleave event handler.
func DragLeave(v string) attr.Node { return attr.KV("ondragleave", v) }

// DragOver sets the ondragover event handler.
func DragOver(v string) attr.Node { return attr.KV("ondragover", v) }

// Drop sets the ondrop event handler.
func Drop(v string) attr.Node { return attr.KV("ondrop", v) }
