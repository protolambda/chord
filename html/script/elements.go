// Package script provides HTML scripting elements.
package script

import "github.com/protolambda/chord/core"

// Script creates a script element.
func Script(opts ...core.Node) core.Node { return core.Element("script", opts...) }

// Noscript creates a noscript element.
func Noscript(opts ...core.Node) core.Node { return core.Element("noscript", opts...) }

// Template creates a template element.
func Template(opts ...core.Node) core.Node { return core.Element("template", opts...) }

// Canvas creates a canvas element.
func Canvas(opts ...core.Node) core.Node { return core.Element("canvas", opts...) }
