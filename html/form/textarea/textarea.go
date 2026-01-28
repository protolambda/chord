// Package textarea provides the HTML textarea element and its attributes.
package textarea

import "github.com/protolambda/chord/core"

// Textarea creates a textarea element.
func Textarea(opts ...core.Node) core.Node { return core.Element("textarea", opts...) }
