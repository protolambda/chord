// Package meta provides HTML document metadata elements.
package meta

import "github.com/protolambda/chord/core"

// HTML creates an html element.
func HTML(opts ...core.Node) core.Node { return core.Element("html", opts...) }

// Head creates a head element.
func Head(opts ...core.Node) core.Node { return core.Element("head", opts...) }

// Title creates a title element.
func Title(opts ...core.Node) core.Node { return core.Element("title", opts...) }

// Base creates a base element (void).
func Base(opts ...core.Node) core.Node { return core.VoidElement("base", opts...) }

// Link creates a link element (void).
func Link(opts ...core.Node) core.Node { return core.VoidElement("link", opts...) }

// Meta creates a meta element (void).
func Meta(opts ...core.Node) core.Node { return core.VoidElement("meta", opts...) }

// Style creates a style element.
func Style(opts ...core.Node) core.Node { return core.Element("style", opts...) }
