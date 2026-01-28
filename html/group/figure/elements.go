// Package figure provides HTML figure elements (figure, figcaption).
package figure

import "github.com/protolambda/chord/core"

// Figure creates a figure element.
func Figure(opts ...core.Node) core.Node { return core.Element("figure", opts...) }

// Figcaption creates a figcaption element.
func Figcaption(opts ...core.Node) core.Node { return core.Element("figcaption", opts...) }
