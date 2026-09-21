package ct

import (
	"context"
	"iter"
	"strings"

	"github.com/protolambda/chord/core/attr"
	"github.com/protolambda/chord/core/elem"
	"github.com/protolambda/chord/core/inspect"
)

// View creates a subject for a Chord view. The view is evaluated at most
// once, through [inspect.Build], with the context of the first check.
// Application code in the view therefore runs exactly once per subject.
func View(node elem.Node, opts ...Option) *Subject {
	return From(viewSource{node: node}, opts...)
}

type viewSource struct {
	node elem.Node
}

var _ Source = viewSource{}

func (viewSource) String() string {
	return "chord view"
}

func (s viewSource) Load(ctx context.Context) (Document, error) {
	doc, err := inspect.Build(ctx, s.node)
	if err != nil {
		return nil, err
	}
	return viewDocument{doc: doc}, nil
}

type viewDocument struct {
	doc *inspect.Document
}

func (d viewDocument) Root() Node {
	return viewNode{n: d.doc.Root()}
}

// viewNode adapts an inspect node. It is a comparable value wrapping a pointer.
type viewNode struct {
	n *inspect.Node
}

var _ Node = viewNode{}

func (v viewNode) Kind() NodeKind {
	switch v.n.Kind() {
	case elem.KindElement:
		return KindElement
	case elem.KindText:
		return KindText
	case elem.KindRaw:
		return KindRaw
	case elem.KindComment:
		return KindComment
	default:
		return KindDocument
	}
}

func (v viewNode) Tag() string {
	return strings.ToLower(v.n.Tag())
}

func (v viewNode) Data() string {
	return v.n.Data()
}

func (v viewNode) Attr(name string) (string, bool) {
	a, ok := v.n.Attr(name)
	if !ok {
		return "", false
	}
	if a.Kind == attr.KindBool {
		return "", true
	}
	return a.Val, true
}

func (v viewNode) Attrs() iter.Seq2[string, string] {
	return func(yield func(string, string) bool) {
		for a := range v.n.Attrs() {
			val := a.Val
			if a.Kind == attr.KindBool {
				val = ""
			}
			if !yield(a.Key, val) {
				return
			}
		}
	}
}

func (v viewNode) Parent() Node {
	p := v.n.Parent()
	if p == nil {
		return nil
	}
	return viewNode{n: p}
}

func (v viewNode) Children() iter.Seq[Node] {
	return func(yield func(Node) bool) {
		for c := range v.n.Children() {
			if !yield(viewNode{n: c}) {
				return
			}
		}
	}
}

func (v viewNode) Location() string {
	return v.n.Location().String()
}
