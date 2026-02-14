// Package list provides HTML list elements (ol, ul, menu, li).
package list

import (
	"github.com/protolambda/chord/core/attr"
	"github.com/protolambda/chord/core/elem"
)

// OL creates an ol element.
func OL(attrs ...attr.Node) elem.Scope { return elem.New("ol", attrs...) }

// UL creates a ul element.
func UL(attrs ...attr.Node) elem.Scope { return elem.New("ul", attrs...) }

// Menu creates a menu element.
func Menu(attrs ...attr.Node) elem.Scope { return elem.New("menu", attrs...) }

// LI creates an li element.
func LI(attrs ...attr.Node) elem.Scope { return elem.New("li", attrs...) }
