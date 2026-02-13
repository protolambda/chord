package list

import (
	"github.com/protolambda/chord/core/attrib"
)

// Reversed sets the reversed boolean attribute for ol elements.
func Reversed() attrib.Node { return attrib.Bool("reversed") }

// Start sets the start attribute for ol elements.
func Start(v string) attrib.Node { return attrib.KV("start", v) }

// Type sets the type attribute for ol elements.
func Type(v string) attrib.Node { return attrib.KV("type", v) }
