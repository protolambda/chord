// Package list provides HTML list elements (ol, ul, menu, li).
package list

import "github.com/protolambda/chord/core"

// OL creates an ol element.
func OL(opts ...core.Node) core.Node { return core.Element("ol", opts...) }

// UL creates a ul element.
func UL(opts ...core.Node) core.Node { return core.Element("ul", opts...) }

// Menu creates a menu element.
func Menu(opts ...core.Node) core.Node { return core.Element("menu", opts...) }

// LI creates an li element.
func LI(opts ...core.Node) core.Node { return core.Element("li", opts...) }
