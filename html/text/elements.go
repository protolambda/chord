// Package text provides HTML text-level semantic elements.
package text

import (
	"github.com/protolambda/chord/core/attrib"
	"github.com/protolambda/chord/core/elem"
)

// A creates an a (anchor) element.
func A(attrs ...attrib.Node) elem.Scope { return elem.New("a", attrs...) }

// EM creates an em element.
func EM(attrs ...attrib.Node) elem.Scope { return elem.New("em", attrs...) }

// Strong creates a strong element.
func Strong(attrs ...attrib.Node) elem.Scope { return elem.New("strong", attrs...) }

// Small creates a small element.
func Small(attrs ...attrib.Node) elem.Scope { return elem.New("small", attrs...) }

// S creates an s element.
func S(attrs ...attrib.Node) elem.Scope { return elem.New("s", attrs...) }

// Cite creates a cite element.
func Cite(attrs ...attrib.Node) elem.Scope { return elem.New("cite", attrs...) }

// Q creates a q element.
func Q(attrs ...attrib.Node) elem.Scope { return elem.New("q", attrs...) }

// Dfn creates a dfn element.
func Dfn(attrs ...attrib.Node) elem.Scope { return elem.New("dfn", attrs...) }

// Abbr creates an abbr element.
func Abbr(attrs ...attrib.Node) elem.Scope { return elem.New("abbr", attrs...) }

// Ruby creates a ruby element.
func Ruby(attrs ...attrib.Node) elem.Scope { return elem.New("ruby", attrs...) }

// RT creates an rt element.
func RT(attrs ...attrib.Node) elem.Scope { return elem.New("rt", attrs...) }

// RP creates an rp element.
func RP(attrs ...attrib.Node) elem.Scope { return elem.New("rp", attrs...) }

// Data creates a data element.
func Data(attrs ...attrib.Node) elem.Scope { return elem.New("data", attrs...) }

// Time creates a time element.
func Time(attrs ...attrib.Node) elem.Scope { return elem.New("time", attrs...) }

// Code creates a code element.
func Code(attrs ...attrib.Node) elem.Scope { return elem.New("code", attrs...) }

// Var creates a var element.
func Var(attrs ...attrib.Node) elem.Scope { return elem.New("var", attrs...) }

// Samp creates a samp element.
func Samp(attrs ...attrib.Node) elem.Scope { return elem.New("samp", attrs...) }

// Kbd creates a kbd element.
func Kbd(attrs ...attrib.Node) elem.Scope { return elem.New("kbd", attrs...) }

// Sub creates a sub element.
func Sub(attrs ...attrib.Node) elem.Scope { return elem.New("sub", attrs...) }

// Sup creates a sup element.
func Sup(attrs ...attrib.Node) elem.Scope { return elem.New("sup", attrs...) }

// I creates an i element.
func I(attrs ...attrib.Node) elem.Scope { return elem.New("i", attrs...) }

// B creates a b element.
func B(attrs ...attrib.Node) elem.Scope { return elem.New("b", attrs...) }

// U creates a u element.
func U(attrs ...attrib.Node) elem.Scope { return elem.New("u", attrs...) }

// Mark creates a mark element.
func Mark(attrs ...attrib.Node) elem.Scope { return elem.New("mark", attrs...) }

// Bdi creates a bdi element.
func Bdi(attrs ...attrib.Node) elem.Scope { return elem.New("bdi", attrs...) }

// Bdo creates a bdo element.
func Bdo(attrs ...attrib.Node) elem.Scope { return elem.New("bdo", attrs...) }

// Span creates a span element.
func Span(attrs ...attrib.Node) elem.Scope { return elem.New("span", attrs...) }

// BR creates a br element (void).
func BR(attrs ...attrib.Node) elem.Node { return elem.Void("br", attrs...) }

// WBR creates a wbr element (void).
func WBR(attrs ...attrib.Node) elem.Node { return elem.Void("wbr", attrs...) }

// P creates a p element.
func P(attrs ...attrib.Node) elem.Scope { return elem.New("p", attrs...) }

// HR creates an hr element (void).
func HR(attrs ...attrib.Node) elem.Node { return elem.Void("hr", attrs...) }

// Pre creates a pre element.
func Pre(attrs ...attrib.Node) elem.Scope { return elem.New("pre", attrs...) }

// Blockquote creates a blockquote element.
func Blockquote(attrs ...attrib.Node) elem.Scope { return elem.New("blockquote", attrs...) }
