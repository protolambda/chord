package ct

import (
	"context"
	"fmt"
	"iter"
)

// NodeKind classifies a [Node] independently of its source.
type NodeKind uint8

const (
	// KindDocument is the root of a document or fragment.
	KindDocument NodeKind = iota
	// KindElement is an element with a tag, attributes, and children.
	KindElement
	// KindText is logical (unescaped) text.
	KindText
	// KindRaw is trusted raw HTML from a direct Chord view. Parsers never produce it.
	KindRaw
	// KindComment is a comment.
	KindComment
	// KindDoctype is a document type declaration, from a parser or from
	// [github.com/protolambda/chord/core/elem.Doctype] in a view.
	KindDoctype
)

func (k NodeKind) String() string {
	switch k {
	case KindDocument:
		return "document"
	case KindElement:
		return "element"
	case KindText:
		return "text"
	case KindRaw:
		return "raw"
	case KindComment:
		return "comment"
	case KindDoctype:
		return "doctype"
	default:
		return fmt.Sprintf("kind(%d)", uint8(k))
	}
}

// Source produces a queryable document. Implementations adapt a tree
// representation such as an evaluated Chord view or a parsed HTML page.
type Source interface {
	// String names the source for diagnostics, e.g. "chord view".
	String() string
	// Load produces the document. A Subject calls it at most once.
	Load(ctx context.Context) (Document, error)
}

// Document is a loaded, read-only tree.
type Document interface {
	// Root returns the document root, a node of KindDocument.
	Root() Node
}

// Node is a read-only view of one node in a [Document].
//
// Implementations must be comparable values that are equal exactly when they
// identify the same node, for example a struct wrapping a pointer. Nodes and
// their iterators stay valid for the lifetime of the document, and iterators
// are restartable.
type Node interface {
	Kind() NodeKind
	// Tag returns the lowercase element tag, or "" for other kinds.
	Tag() string
	// Data returns text, raw HTML, comment content, or the doctype name,
	// depending on Kind.
	Data() string
	// Attr returns the logical attribute value, as an HTML parser reads it
	// from the page: character references are decoded, also in raw
	// (output-ready) values of a view, and line breaks are LF. Boolean
	// attributes report an empty value and true.
	Attr(name string) (value string, ok bool)
	// Attrs iterates attributes as name/value pairs in source order.
	Attrs() iter.Seq2[string, string]
	// Parent returns the parent node, or nil at the root.
	Parent() Node
	// Children iterates the direct children in order.
	Children() iter.Seq[Node]
	// Location describes the node position for diagnostics.
	Location() string
}

// BoolAttrNode is implemented by nodes whose source tells a boolean
// attribute, written without a value, apart from an attribute with an empty
// value. [View] nodes implement it. Parsed documents do not: the HTML parser
// gives both the empty value, and so does the DOM.
//
// Queries follow the DOM and treat both alike; diagnostics show a boolean
// attribute as a bare name and an empty value as name="".
type BoolAttrNode interface {
	Node
	// BoolAttr reports whether the attribute is present without a value.
	BoolAttr(name string) bool
}

// descendants iterates all descendants of n in document order, excluding n.
func descendants(n Node) iter.Seq[Node] {
	return func(yield func(Node) bool) {
		descend(n, yield)
	}
}

func descend(n Node, yield func(Node) bool) bool {
	for c := range n.Children() {
		if !yield(c) {
			return false
		}
		if !descend(c, yield) {
			return false
		}
	}
	return true
}

// root returns the root of the tree containing n.
func root(n Node) Node {
	for {
		p := n.Parent()
		if p == nil {
			return n
		}
		n = p
	}
}
