// Package script provides HTML scripting elements.
package script

import (
	"github.com/protolambda/chord/core/attr"
	"github.com/protolambda/chord/core/elem"
)

// Script creates a script element.
func Script(attrs ...attr.Node) elem.Scope { return elem.Name("script").New(attrs...) }

// Noscript creates a noscript element.
func Noscript(attrs ...attr.Node) elem.Scope { return elem.Name("noscript").New(attrs...) }

// Template creates a template element.
func Template(attrs ...attr.Node) elem.Scope { return elem.Name("template").New(attrs...) }

// Canvas creates a canvas element.
func Canvas(attrs ...attr.Node) elem.Scope { return elem.Name("canvas").New(attrs...) }
