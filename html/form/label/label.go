// Package label provides the HTML label element and its attributes.
package label

import "github.com/protolambda/chord/core"

// Label creates a label element.
func Label(opts ...core.Node) core.Node { return core.Element("label", opts...) }
