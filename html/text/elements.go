// Package text provides HTML text-level semantic elements.
package text

import "github.com/protolambda/chord/core"

// A creates an a (anchor) element.
func A(opts ...core.Node) core.Node { return core.Element("a", opts...) }

// EM creates an em element.
func EM(opts ...core.Node) core.Node { return core.Element("em", opts...) }

// Strong creates a strong element.
func Strong(opts ...core.Node) core.Node { return core.Element("strong", opts...) }

// Small creates a small element.
func Small(opts ...core.Node) core.Node { return core.Element("small", opts...) }

// S creates an s element.
func S(opts ...core.Node) core.Node { return core.Element("s", opts...) }

// Cite creates a cite element.
func Cite(opts ...core.Node) core.Node { return core.Element("cite", opts...) }

// Q creates a q element.
func Q(opts ...core.Node) core.Node { return core.Element("q", opts...) }

// Dfn creates a dfn element.
func Dfn(opts ...core.Node) core.Node { return core.Element("dfn", opts...) }

// Abbr creates an abbr element.
func Abbr(opts ...core.Node) core.Node { return core.Element("abbr", opts...) }

// Ruby creates a ruby element.
func Ruby(opts ...core.Node) core.Node { return core.Element("ruby", opts...) }

// RT creates an rt element.
func RT(opts ...core.Node) core.Node { return core.Element("rt", opts...) }

// RP creates an rp element.
func RP(opts ...core.Node) core.Node { return core.Element("rp", opts...) }

// Data creates a data element.
func Data(opts ...core.Node) core.Node { return core.Element("data", opts...) }

// Time creates a time element.
func Time(opts ...core.Node) core.Node { return core.Element("time", opts...) }

// Code creates a code element.
func Code(opts ...core.Node) core.Node { return core.Element("code", opts...) }

// Var creates a var element.
func Var(opts ...core.Node) core.Node { return core.Element("var", opts...) }

// Samp creates a samp element.
func Samp(opts ...core.Node) core.Node { return core.Element("samp", opts...) }

// Kbd creates a kbd element.
func Kbd(opts ...core.Node) core.Node { return core.Element("kbd", opts...) }

// Sub creates a sub element.
func Sub(opts ...core.Node) core.Node { return core.Element("sub", opts...) }

// Sup creates a sup element.
func Sup(opts ...core.Node) core.Node { return core.Element("sup", opts...) }

// I creates an i element.
func I(opts ...core.Node) core.Node { return core.Element("i", opts...) }

// B creates a b element.
func B(opts ...core.Node) core.Node { return core.Element("b", opts...) }

// U creates a u element.
func U(opts ...core.Node) core.Node { return core.Element("u", opts...) }

// Mark creates a mark element.
func Mark(opts ...core.Node) core.Node { return core.Element("mark", opts...) }

// Bdi creates a bdi element.
func Bdi(opts ...core.Node) core.Node { return core.Element("bdi", opts...) }

// Bdo creates a bdo element.
func Bdo(opts ...core.Node) core.Node { return core.Element("bdo", opts...) }

// Span creates a span element.
func Span(opts ...core.Node) core.Node { return core.Element("span", opts...) }

// BR creates a br element (void).
func BR(opts ...core.Node) core.Node { return core.VoidElement("br", opts...) }

// WBR creates a wbr element (void).
func WBR(opts ...core.Node) core.Node { return core.VoidElement("wbr", opts...) }

// P creates a p element.
func P(opts ...core.Node) core.Node { return core.Element("p", opts...) }

// HR creates an hr element (void).
func HR(opts ...core.Node) core.Node { return core.VoidElement("hr", opts...) }

// Pre creates a pre element.
func Pre(opts ...core.Node) core.Node { return core.Element("pre", opts...) }

// Blockquote creates a blockquote element.
func Blockquote(opts ...core.Node) core.Node { return core.Element("blockquote", opts...) }
