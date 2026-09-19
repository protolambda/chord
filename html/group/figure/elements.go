// Package figure provides HTML figure elements (figure, figcaption).
package figure

import (
	"github.com/protolambda/chord/core/attr"
	"github.com/protolambda/chord/core/elem"
)

// Figure creates a figure element.
func Figure(attrs ...attr.Node) elem.Scope { return elem.Name("figure").New(attrs...) }

// Figcaption creates a figcaption element.
func Figcaption(attrs ...attr.Node) elem.Scope { return elem.Name("figcaption").New(attrs...) }
