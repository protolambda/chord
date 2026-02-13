// Package meta provides HTML document metadata elements.
package meta

import (
	"github.com/protolambda/chord/core/attrib"
	"github.com/protolambda/chord/core/elem"
)

// HTML creates an html element.
func HTML(attrs ...attrib.Node) elem.Scope { return elem.New("html", attrs...) }

// Head creates a head element.
func Head(attrs ...attrib.Node) elem.Scope { return elem.New("head", attrs...) }

// Title creates a title element.
func Title(attrs ...attrib.Node) elem.Scope { return elem.New("title", attrs...) }

// Base creates a base element (void).
func Base(attrs ...attrib.Node) elem.Node { return elem.Void("base", attrs...) }

// Link creates a link element (void).
func Link(attrs ...attrib.Node) elem.Node { return elem.Void("link", attrs...) }

// Meta creates a meta element (void).
func Meta(attrs ...attrib.Node) elem.Node { return elem.Void("meta", attrs...) }

// Style creates a style element.
func Style(attrs ...attrib.Node) elem.Scope { return elem.New("style", attrs...) }
