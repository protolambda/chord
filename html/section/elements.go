// Package section provides HTML sectioning elements.
package section

import "github.com/protolambda/chord/core"

// Body creates a body element.
func Body(opts ...core.Node) core.Node { return core.Element("body", opts...) }

// Article creates an article element.
func Article(opts ...core.Node) core.Node { return core.Element("article", opts...) }

// Section creates a section element.
func Section(opts ...core.Node) core.Node { return core.Element("section", opts...) }

// Nav creates a nav element.
func Nav(opts ...core.Node) core.Node { return core.Element("nav", opts...) }

// Aside creates an aside element.
func Aside(opts ...core.Node) core.Node { return core.Element("aside", opts...) }

// H1 creates an h1 element.
func H1(opts ...core.Node) core.Node { return core.Element("h1", opts...) }

// H2 creates an h2 element.
func H2(opts ...core.Node) core.Node { return core.Element("h2", opts...) }

// H3 creates an h3 element.
func H3(opts ...core.Node) core.Node { return core.Element("h3", opts...) }

// H4 creates an h4 element.
func H4(opts ...core.Node) core.Node { return core.Element("h4", opts...) }

// H5 creates an h5 element.
func H5(opts ...core.Node) core.Node { return core.Element("h5", opts...) }

// H6 creates an h6 element.
func H6(opts ...core.Node) core.Node { return core.Element("h6", opts...) }

// Hgroup creates an hgroup element.
func Hgroup(opts ...core.Node) core.Node { return core.Element("hgroup", opts...) }

// Header creates a header element.
func Header(opts ...core.Node) core.Node { return core.Element("header", opts...) }

// Footer creates a footer element.
func Footer(opts ...core.Node) core.Node { return core.Element("footer", opts...) }

// Address creates an address element.
func Address(opts ...core.Node) core.Node { return core.Element("address", opts...) }

// Main creates a main element.
func Main(opts ...core.Node) core.Node { return core.Element("main", opts...) }
