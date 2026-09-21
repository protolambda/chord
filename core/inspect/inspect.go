// Package inspect builds renderer-equivalent snapshots of evaluated element trees.
//
// [Build] evaluates a node graph exactly once, through the same normalized
// walk that [github.com/protolambda/chord/core.Render] uses, and returns a
// detached read-only [Document]. The document can be traversed repeatedly
// without re-running application code, and every [Node] implements
// [elem.Node] so a snapshot can be rendered again as static content.
//
// Snapshots keep logical values: text is unescaped, raw content is one opaque
// node, and attribute values are reported together with their [attr.Kind].
// This package has no opinion on testing; see the ct package for queries and
// assertions built on top of it.
package inspect

import (
	"context"
	"fmt"
	"iter"
	"strings"

	"github.com/protolambda/chord/core/attr"
	"github.com/protolambda/chord/core/elem"
	"github.com/protolambda/chord/core/internal/walk"
)

// Build evaluates source into a detached document. Application code reached
// through the node graph runs exactly once. On failure no document is
// returned; the error carries the location of the failing node.
func Build(ctx context.Context, source elem.Node) (*Document, error) {
	root := &Node{kind: elem.KindFragment}
	b := &builder{stack: []*Node{root}}
	if err := walk.Walk(ctx, source, b); err != nil {
		return nil, fmt.Errorf("inspect: %w", err)
	}
	return &Document{root: root}, nil
}

// Document is a completed, read-only snapshot.
type Document struct {
	root *Node
}

// Root returns the virtual fragment root that holds the top-level nodes.
func (d *Document) Root() *Node {
	return d.root
}

// Attribute is a normalized attribute of an element node.
type Attribute struct {
	Key string
	// Val is logical for [attr.KindValue], output-ready for [attr.KindRawValue],
	// and empty for [attr.KindBool].
	Val  string
	Kind attr.Kind
}

// Node is a read-only node of a [Document]. The root is a fragment; other
// nodes are elements, text, raw content, or comments.
type Node struct {
	kind     elem.Kind
	tag      string
	data     string
	void     bool
	attrs    []Attribute
	parent   *Node
	index    int
	children []*Node
}

var _ elem.Node = (*Node)(nil)

// Kind returns the node kind. It is never [elem.KindNoop].
func (n *Node) Kind() elem.Kind { return n.kind }

// Tag returns the element tag, or "" for non-element nodes.
func (n *Node) Tag() string { return n.tag }

// Data returns logical text, raw HTML, or comment content, depending on Kind.
func (n *Node) Data() string { return n.data }

// Void reports whether the node is a self-closing element.
func (n *Node) Void() bool { return n.void }

// Attr returns the attribute with the given name.
func (n *Node) Attr(name string) (Attribute, bool) {
	for _, a := range n.attrs {
		if a.Key == name {
			return a, true
		}
	}
	return Attribute{}, false
}

// Attrs iterates the attributes in output order. The iterator is restartable.
func (n *Node) Attrs() iter.Seq[Attribute] {
	return func(yield func(Attribute) bool) {
		for _, a := range n.attrs {
			if !yield(a) {
				return
			}
		}
	}
}

// Parent returns the parent node, or nil for the root.
func (n *Node) Parent() *Node { return n.parent }

// Index returns the position of the node among its parent's children.
func (n *Node) Index() int { return n.index }

// Children iterates the direct children in order. The iterator is restartable.
func (n *Node) Children() iter.Seq[*Node] {
	return func(yield func(*Node) bool) {
		for _, c := range n.children {
			if !yield(c) {
				return
			}
		}
	}
}

// Descendants iterates all descendants in document (depth-first, pre-order)
// order, excluding n itself. The iterator is restartable.
func (n *Node) Descendants() iter.Seq[*Node] {
	return func(yield func(*Node) bool) {
		n.descend(yield)
	}
}

func (n *Node) descend(yield func(*Node) bool) bool {
	for _, c := range n.children {
		if !yield(c) {
			return false
		}
		if !c.descend(yield) {
			return false
		}
	}
	return true
}

// TextContent concatenates the logical text of all text nodes in the subtree.
// Raw content and comments are excluded, and whitespace is kept as is.
func (n *Node) TextContent() string {
	var b strings.Builder
	n.appendText(&b)
	return b.String()
}

func (n *Node) appendText(b *strings.Builder) {
	if n.kind == elem.KindText {
		b.WriteString(n.data)
	}
	for _, c := range n.children {
		c.appendText(b)
	}
}

// Location returns the structural location of the node, for diagnostics.
func (n *Node) Location() Location {
	var steps []Step
	for cur := n; cur.parent != nil; cur = cur.parent {
		var id string
		if cur.kind == elem.KindElement {
			id = cur.simpleID()
		}
		steps = append(steps, Step{Tag: cur.tag, ID: id, Index: cur.index})
	}
	for i, j := 0, len(steps)-1; i < j; i, j = i+1, j-1 {
		steps[i], steps[j] = steps[j], steps[i]
	}
	return Location{steps: steps}
}

func (n *Node) simpleID() string {
	attrs := make([]walk.Attribute, 0, len(n.attrs))
	for _, a := range n.attrs {
		attrs = append(attrs, walk.Attribute{Key: a.Key, Val: a.Val, Kind: a.Kind})
	}
	return walk.SimpleID(attrs)
}

// Eval returns the node as a static element object, so a snapshot can be
// rendered again without re-running the application code that produced it.
func (n *Node) Eval(context.Context) (elem.Obj, error) {
	switch n.kind {
	case elem.KindFragment:
		return elem.Obj{Kind: elem.KindFragment, Children: n.childNodes()}, nil
	case elem.KindElement:
		obj := elem.Obj{Kind: elem.KindElement, Tag: n.tag, Void: n.void, Attribs: n.attribSeq()}
		if !n.void {
			obj.Children = n.childNodes()
		}
		return obj, nil
	default:
		return elem.Obj{Kind: n.kind, Data: n.data}, nil
	}
}

func (n *Node) childNodes() iter.Seq[elem.Node] {
	return func(yield func(elem.Node) bool) {
		for _, c := range n.children {
			if !yield(c) {
				return
			}
		}
	}
}

func (n *Node) attribSeq() attr.Seq {
	return func(yield func(attr.Node) bool) {
		for _, a := range n.attrs {
			if !yield(attr.Obj{Kind: a.Kind, Key: a.Key, Val: a.Val}) {
				return
			}
		}
	}
}

// Location is an immutable structural position in a document: the sequence of
// child indices from the root, annotated with tags and simple ids. It is a
// diagnostic value, not a selector.
type Location struct {
	steps []Step
}

// Step is one level of a [Location].
type Step struct {
	// Tag is the element tag, or "" for text, raw, and comment nodes.
	Tag string
	// ID is the element id when it is a simple token, for readability only.
	ID string
	// Index is the position among the parent's children.
	Index int
}

// Steps returns a copy of the steps from the root down to the node.
func (l Location) Steps() []Step {
	return append([]Step(nil), l.steps...)
}

// String formats the location like "html[0]/body[1]/form#login[0]/[2]".
// The root has an empty location.
func (l Location) String() string {
	var b strings.Builder
	for i, s := range l.steps {
		if i > 0 {
			b.WriteByte('/')
		}
		b.WriteString(s.Tag)
		if s.ID != "" {
			b.WriteByte('#')
			b.WriteString(s.ID)
		}
		fmt.Fprintf(&b, "[%d]", s.Index)
	}
	return b.String()
}

// builder is the snapshot receiver.
type builder struct {
	stack []*Node
}

var _ walk.Receiver = (*builder)(nil)

func (b *builder) add(n *Node) *Node {
	parent := b.stack[len(b.stack)-1]
	n.parent = parent
	n.index = len(parent.children)
	parent.children = append(parent.children, n)
	return n
}

func (b *builder) Open(el walk.Element) error {
	n := &Node{kind: elem.KindElement, tag: el.Tag, void: el.Void}
	if len(el.Attrs) > 0 {
		n.attrs = make([]Attribute, len(el.Attrs))
		for i, a := range el.Attrs {
			n.attrs[i] = Attribute{Key: a.Key, Val: a.Val, Kind: a.Kind}
		}
	}
	b.stack = append(b.stack, b.add(n))
	return nil
}

func (b *builder) Text(v string) error {
	b.add(&Node{kind: elem.KindText, data: v})
	return nil
}

func (b *builder) Raw(v string) error {
	b.add(&Node{kind: elem.KindRaw, data: v})
	return nil
}

func (b *builder) Comment(v string) error {
	b.add(&Node{kind: elem.KindComment, data: v})
	return nil
}

func (b *builder) Close() error {
	b.stack = b.stack[:len(b.stack)-1]
	return nil
}
