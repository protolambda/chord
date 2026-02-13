// Package script provides HTML scripting elements.
package script

import (
	"github.com/protolambda/chord/core/attrib"
	"github.com/protolambda/chord/core/elem"
)

// Script creates a script element.
func Script(attrs ...attrib.Node) elem.Scope { return elem.New("script", attrs...) }

// Noscript creates a noscript element.
func Noscript(attrs ...attrib.Node) elem.Scope { return elem.New("noscript", attrs...) }

// Template creates a template element.
func Template(attrs ...attrib.Node) elem.Scope { return elem.New("template", attrs...) }

// Canvas creates a canvas element.
func Canvas(attrs ...attrib.Node) elem.Scope { return elem.New("canvas", attrs...) }
