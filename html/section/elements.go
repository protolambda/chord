// Package section provides HTML sectioning elements.
package section

import (
	"github.com/protolambda/chord/core/attr"
	"github.com/protolambda/chord/core/elem"
)

// Body creates a body element.
func Body(attrs ...attr.Node) elem.Scope { return elem.New("body", attrs...) }

// Article creates an article element.
func Article(attrs ...attr.Node) elem.Scope { return elem.New("article", attrs...) }

// Section creates a section element.
func Section(attrs ...attr.Node) elem.Scope { return elem.New("section", attrs...) }

// Nav creates a nav element.
func Nav(attrs ...attr.Node) elem.Scope { return elem.New("nav", attrs...) }

// Aside creates an aside element.
func Aside(attrs ...attr.Node) elem.Scope { return elem.New("aside", attrs...) }

// H1 creates an h1 element.
func H1(attrs ...attr.Node) elem.Scope { return elem.New("h1", attrs...) }

// H2 creates an h2 element.
func H2(attrs ...attr.Node) elem.Scope { return elem.New("h2", attrs...) }

// H3 creates an h3 element.
func H3(attrs ...attr.Node) elem.Scope { return elem.New("h3", attrs...) }

// H4 creates an h4 element.
func H4(attrs ...attr.Node) elem.Scope { return elem.New("h4", attrs...) }

// H5 creates an h5 element.
func H5(attrs ...attr.Node) elem.Scope { return elem.New("h5", attrs...) }

// H6 creates an h6 element.
func H6(attrs ...attr.Node) elem.Scope { return elem.New("h6", attrs...) }

// Hgroup creates an hgroup element.
func Hgroup(attrs ...attr.Node) elem.Scope { return elem.New("hgroup", attrs...) }

// Header creates a header element.
func Header(attrs ...attr.Node) elem.Scope { return elem.New("header", attrs...) }

// Footer creates a footer element.
func Footer(attrs ...attr.Node) elem.Scope { return elem.New("footer", attrs...) }

// Address creates an address element.
func Address(attrs ...attr.Node) elem.Scope { return elem.New("address", attrs...) }

// Main creates a main element.
func Main(attrs ...attr.Node) elem.Scope { return elem.New("main", attrs...) }
