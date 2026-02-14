package on

import (
	"github.com/protolambda/chord/core/attr"
)

// Copy sets the oncopy event handler.
func Copy(v string) attr.Node { return attr.KV("oncopy", v) }

// Cut sets the oncut event handler.
func Cut(v string) attr.Node { return attr.KV("oncut", v) }

// Paste sets the onpaste event handler.
func Paste(v string) attr.Node { return attr.KV("onpaste", v) }
