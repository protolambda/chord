// Package text provides HTML text-level semantic elements.
package text

import (
	"github.com/protolambda/chord/core/attr"
	"github.com/protolambda/chord/core/elem"
)

// A creates an a (anchor) element.
func A(attrs ...attr.Node) elem.Scope { return elem.Name("a").New(attrs...) }

// EM creates an em element.
func EM(attrs ...attr.Node) elem.Scope { return elem.Name("em").New(attrs...) }

// Strong creates a strong element.
func Strong(attrs ...attr.Node) elem.Scope { return elem.Name("strong").New(attrs...) }

// Small creates a small element.
func Small(attrs ...attr.Node) elem.Scope { return elem.Name("small").New(attrs...) }

// S creates an s element.
func S(attrs ...attr.Node) elem.Scope { return elem.Name("s").New(attrs...) }

// Cite creates a cite element.
func Cite(attrs ...attr.Node) elem.Scope { return elem.Name("cite").New(attrs...) }

// Q creates a q element.
func Q(attrs ...attr.Node) elem.Scope { return elem.Name("q").New(attrs...) }

// Dfn creates a dfn element.
func Dfn(attrs ...attr.Node) elem.Scope { return elem.Name("dfn").New(attrs...) }

// Abbr creates an abbr element.
func Abbr(attrs ...attr.Node) elem.Scope { return elem.Name("abbr").New(attrs...) }

// Ruby creates a ruby element.
func Ruby(attrs ...attr.Node) elem.Scope { return elem.Name("ruby").New(attrs...) }

// RT creates an rt element.
func RT(attrs ...attr.Node) elem.Scope { return elem.Name("rt").New(attrs...) }

// RP creates an rp element.
func RP(attrs ...attr.Node) elem.Scope { return elem.Name("rp").New(attrs...) }

// Data creates a data element.
func Data(attrs ...attr.Node) elem.Scope { return elem.Name("data").New(attrs...) }

// Time creates a time element.
func Time(attrs ...attr.Node) elem.Scope { return elem.Name("time").New(attrs...) }

// Code creates a code element.
func Code(attrs ...attr.Node) elem.Scope { return elem.Name("code").New(attrs...) }

// Var creates a var element.
func Var(attrs ...attr.Node) elem.Scope { return elem.Name("var").New(attrs...) }

// Samp creates a samp element.
func Samp(attrs ...attr.Node) elem.Scope { return elem.Name("samp").New(attrs...) }

// Kbd creates a kbd element.
func Kbd(attrs ...attr.Node) elem.Scope { return elem.Name("kbd").New(attrs...) }

// Sub creates a sub element.
func Sub(attrs ...attr.Node) elem.Scope { return elem.Name("sub").New(attrs...) }

// Sup creates a sup element.
func Sup(attrs ...attr.Node) elem.Scope { return elem.Name("sup").New(attrs...) }

// I creates an i element.
func I(attrs ...attr.Node) elem.Scope { return elem.Name("i").New(attrs...) }

// B creates a b element.
func B(attrs ...attr.Node) elem.Scope { return elem.Name("b").New(attrs...) }

// U creates a u element.
func U(attrs ...attr.Node) elem.Scope { return elem.Name("u").New(attrs...) }

// Mark creates a mark element.
func Mark(attrs ...attr.Node) elem.Scope { return elem.Name("mark").New(attrs...) }

// Bdi creates a bdi element.
func Bdi(attrs ...attr.Node) elem.Scope { return elem.Name("bdi").New(attrs...) }

// Bdo creates a bdo element.
func Bdo(attrs ...attr.Node) elem.Scope { return elem.Name("bdo").New(attrs...) }

// Span creates a span element.
func Span(attrs ...attr.Node) elem.Scope { return elem.Name("span").New(attrs...) }

// BR creates a br element (void).
func BR(attrs ...attr.Node) elem.Node { return elem.Name("br").Void(attrs...) }

// WBR creates a wbr element (void).
func WBR(attrs ...attr.Node) elem.Node { return elem.Name("wbr").Void(attrs...) }

// P creates a p element.
func P(attrs ...attr.Node) elem.Scope { return elem.Name("p").New(attrs...) }

// HR creates an hr element (void).
func HR(attrs ...attr.Node) elem.Node { return elem.Name("hr").Void(attrs...) }

// Pre creates a pre element.
func Pre(attrs ...attr.Node) elem.Scope { return elem.Name("pre").New(attrs...) }

// Blockquote creates a blockquote element.
func Blockquote(attrs ...attr.Node) elem.Scope { return elem.Name("blockquote").New(attrs...) }
