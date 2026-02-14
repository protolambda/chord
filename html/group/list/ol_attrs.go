package list

import (
	"github.com/protolambda/chord/core/attr"
)

// Reversed sets the reversed boolean attribute for ol elements.
func Reversed() attr.Node { return attr.Bool("reversed") }

// Start sets the start attribute for ol elements.
func Start(v string) attr.Node { return attr.KV("start", v) }

// Type sets the type attribute for ol elements.
func Type(v string) attr.Node { return attr.KV("type", v) }
