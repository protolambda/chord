// Package script provides HTML scripting elements.
package script

import (
	"github.com/protolambda/chord/core/attr"
	"github.com/protolambda/chord/core/elem"
)

// Script creates a script element.
func Script(attrs ...attr.Node) elem.Scope { return elem.New("script", attrs...) }

// Noscript creates a noscript element.
func Noscript(attrs ...attr.Node) elem.Scope { return elem.New("noscript", attrs...) }

// Template creates a template element.
func Template(attrs ...attr.Node) elem.Scope { return elem.New("template", attrs...) }

// Canvas creates a canvas element.
func Canvas(attrs ...attr.Node) elem.Scope { return elem.New("canvas", attrs...) }
