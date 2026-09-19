package list

import (
	"github.com/protolambda/chord/core/attr"
)

// Reversed sets the reversed boolean attribute for ol elements.
func Reversed() attr.Node { return attr.Name("reversed").Bool() }

// Start sets the start attribute for ol elements.
func Start(v string) attr.Node { return attr.Name("start").Value(v) }

// Type sets the type attribute for ol elements.
func Type(v string) attr.Node { return attr.Name("type").Value(v) }
