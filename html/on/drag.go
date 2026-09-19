package on

import (
	"github.com/protolambda/chord/core/attr"
)

// Drag sets the ondrag event handler.
func Drag(v string) attr.Node { return attr.Name("ondrag").Value(v) }

// DragStart sets the ondragstart event handler.
func DragStart(v string) attr.Node { return attr.Name("ondragstart").Value(v) }

// DragEnd sets the ondragend event handler.
func DragEnd(v string) attr.Node { return attr.Name("ondragend").Value(v) }

// DragEnter sets the ondragenter event handler.
func DragEnter(v string) attr.Node { return attr.Name("ondragenter").Value(v) }

// DragLeave sets the ondragleave event handler.
func DragLeave(v string) attr.Node { return attr.Name("ondragleave").Value(v) }

// DragOver sets the ondragover event handler.
func DragOver(v string) attr.Node { return attr.Name("ondragover").Value(v) }

// Drop sets the ondrop event handler.
func Drop(v string) attr.Node { return attr.Name("ondrop").Value(v) }
