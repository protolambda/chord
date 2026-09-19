// Package meta provides HTML document metadata elements.
package meta

import (
	"github.com/protolambda/chord/core/attr"
	"github.com/protolambda/chord/core/elem"
)

// HTML creates an html element.
func HTML(attrs ...attr.Node) elem.Scope { return elem.Name("html").New(attrs...) }

// Head creates a head element.
func Head(attrs ...attr.Node) elem.Scope { return elem.Name("head").New(attrs...) }

// Title creates a title element.
func Title(attrs ...attr.Node) elem.Scope { return elem.Name("title").New(attrs...) }

// Base creates a base element (void).
func Base(attrs ...attr.Node) elem.Node { return elem.Name("base").Void(attrs...) }

// Link creates a link element (void).
func Link(attrs ...attr.Node) elem.Node { return elem.Name("link").Void(attrs...) }

// Meta creates a meta element (void).
func Meta(attrs ...attr.Node) elem.Node { return elem.Name("meta").Void(attrs...) }

// Style creates a style element.
func Style(attrs ...attr.Node) elem.Scope { return elem.Name("style").New(attrs...) }
