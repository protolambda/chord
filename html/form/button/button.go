// Package button provides the HTML button element and its attributes.
package button

import "github.com/protolambda/chord/core"

// Button creates a button element.
func Button(opts ...core.Node) core.Node { return core.Element("button", opts...) }
