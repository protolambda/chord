// Package textarea provides the HTML textarea element and its attributes.
package textarea

import (
	"github.com/protolambda/chord/core/attrib"
	"github.com/protolambda/chord/core/elem"
)

// Textarea creates a textarea element.
func Textarea(attrs ...attrib.Node) elem.Scope { return elem.New("textarea", attrs...) }
