package on

import (
	"github.com/protolambda/chord/core/attrib"
)

// Copy sets the oncopy event handler.
func Copy(v string) attrib.Node { return attrib.KV("oncopy", v) }

// Cut sets the oncut event handler.
func Cut(v string) attrib.Node { return attrib.KV("oncut", v) }

// Paste sets the onpaste event handler.
func Paste(v string) attrib.Node { return attrib.KV("onpaste", v) }
