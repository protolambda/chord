package on

import (
	"github.com/protolambda/chord/core/attr"
)

// Copy sets the oncopy event handler.
func Copy(v string) attr.Node { return attr.Name("oncopy").Value(v) }

// Cut sets the oncut event handler.
func Cut(v string) attr.Node { return attr.Name("oncut").Value(v) }

// Paste sets the onpaste event handler.
func Paste(v string) attr.Node { return attr.Name("onpaste").Value(v) }
