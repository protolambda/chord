// Package section provides HTML sectioning elements.
package section

import (
	"github.com/protolambda/chord/core/attr"
	"github.com/protolambda/chord/core/elem"
)

// Body creates a body element.
func Body(attrs ...attr.Node) elem.Scope { return elem.Name("body").New(attrs...) }

// Article creates an article element.
func Article(attrs ...attr.Node) elem.Scope { return elem.Name("article").New(attrs...) }

// Section creates a section element.
func Section(attrs ...attr.Node) elem.Scope { return elem.Name("section").New(attrs...) }

// Nav creates a nav element.
func Nav(attrs ...attr.Node) elem.Scope { return elem.Name("nav").New(attrs...) }

// Aside creates an aside element.
func Aside(attrs ...attr.Node) elem.Scope { return elem.Name("aside").New(attrs...) }

// H1 creates an h1 element.
func H1(attrs ...attr.Node) elem.Scope { return elem.Name("h1").New(attrs...) }

// H2 creates an h2 element.
func H2(attrs ...attr.Node) elem.Scope { return elem.Name("h2").New(attrs...) }

// H3 creates an h3 element.
func H3(attrs ...attr.Node) elem.Scope { return elem.Name("h3").New(attrs...) }

// H4 creates an h4 element.
func H4(attrs ...attr.Node) elem.Scope { return elem.Name("h4").New(attrs...) }

// H5 creates an h5 element.
func H5(attrs ...attr.Node) elem.Scope { return elem.Name("h5").New(attrs...) }

// H6 creates an h6 element.
func H6(attrs ...attr.Node) elem.Scope { return elem.Name("h6").New(attrs...) }

// Hgroup creates an hgroup element.
func Hgroup(attrs ...attr.Node) elem.Scope { return elem.Name("hgroup").New(attrs...) }

// Header creates a header element.
func Header(attrs ...attr.Node) elem.Scope { return elem.Name("header").New(attrs...) }

// Footer creates a footer element.
func Footer(attrs ...attr.Node) elem.Scope { return elem.Name("footer").New(attrs...) }

// Address creates an address element.
func Address(attrs ...attr.Node) elem.Scope { return elem.Name("address").New(attrs...) }

// Main creates a main element.
func Main(attrs ...attr.Node) elem.Scope { return elem.Name("main").New(attrs...) }
