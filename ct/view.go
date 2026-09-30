package ct

import (
	"context"
	"iter"
	"strings"

	"golang.org/x/net/html"

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

var (
	_ Node         = viewNode{}
	_ BoolAttrNode = viewNode{}
)

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
	case elem.KindDoctype:
		return KindDoctype
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
	return logicalValue(a), true
}

func (v viewNode) Attrs() iter.Seq2[string, string] {
	return func(yield func(string, string) bool) {
		for a := range v.n.Attrs() {
			if !yield(a.Key, logicalValue(a)) {
				return
			}
		}
	}
}

func (v viewNode) BoolAttr(name string) bool {
	a, ok := v.n.Attr(name)
	return ok && a.Kind == attr.KindBool
}

// logicalValue returns the value of an attribute as a parser reads it from
// the rendered output: empty for a boolean attribute, decoded for a raw
// (output-ready) value, and with line breaks and NUL characters normalized,
// so that views and parsed pages agree.
func logicalValue(a inspect.Attribute) string {
	switch a.Kind {
	case attr.KindBool:
		return ""
	case attr.KindRawValue:
		return decodeRawValue(a.Val)
	default:
		// The renderer escapes the markup characters of a logical value, which
		// the parser decodes again; it only normalizes the input.
		if strings.ContainsAny(a.Val, "\r\x00") {
			return inputNormalizer.Replace(a.Val)
		}
		return a.Val
	}
}

// inputNormalizer changes text as the HTML parser does in attribute values:
// CR LF and CR become LF, and NUL becomes U+FFFD.
var inputNormalizer = strings.NewReplacer("\r\n", "\n", "\r", "\n", "\x00", "\uFFFD")

// decodeRawValue decodes an output-ready attribute value with the HTML
// tokenizer, as the renderer writes it: in double quotes. This applies the
// attribute rules for character references and newline normalization. A raw
// value that contains a double quote breaks out of the attribute; the result
// is then the part a parser reads as the value.
func decodeRawValue(v string) string {
	if !strings.ContainsAny(v, "&\r\x00\"") {
		return v
	}
	z := html.NewTokenizer(strings.NewReader(`<a v="` + v + `">`))
	if z.Next() != html.StartTagToken {
		return v
	}
	_, val, _ := z.TagAttr()
	return string(val)
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
