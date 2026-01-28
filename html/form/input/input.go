// Package input provides the HTML input element and its attributes.
package input

import "github.com/protolambda/chord/core"

// Input creates an input element (void).
func Input(opts ...core.Node) core.Node { return core.VoidElement("input", opts...) }
