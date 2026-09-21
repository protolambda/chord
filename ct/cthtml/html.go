// Package cthtml creates ct subjects from HTML text, parsed with
// golang.org/x/net/html.
//
// Parsing happens on the first check or load of the subject, and at most
// once. The parser repairs malformed input rather than rejecting it, and
// applies HTML tree construction rules: a page gets html, head, and body
// elements, entities are decoded, and some elements may be inserted or
// moved. A successful parse therefore tests the DOM implied by the bytes,
// not the validity or exact bytes of the HTML.
//
// Fragments need a context tag because parsing is context sensitive; for
// example table rows only parse inside a table context:
//
//	page := cthtml.PageString(body)
//	row := cthtml.FragmentString("tbody", partial)
//
// Parsed documents never contain raw nodes, so queries over raw provenance
// only apply to direct views.
package cthtml

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"iter"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"

	"github.com/protolambda/chord/ct"
)

// Page creates a subject for a complete HTML document read from r.
// The reader is consumed once, when the subject first loads.
func Page(r io.Reader, opts ...ct.Option) *ct.Subject {
	return ct.From(&source{desc: "html page", reader: r}, opts...)
}

// PageBytes creates a subject for a complete HTML document. The bytes are
// copied, so later changes to data do not affect the subject.
func PageBytes(data []byte, opts ...ct.Option) *ct.Subject {
	return Page(bytes.NewReader(bytes.Clone(data)), opts...)
}

// PageString creates a subject for a complete HTML document.
func PageString(data string, opts ...ct.Option) *ct.Subject {
	return Page(strings.NewReader(data), opts...)
}

// Fragment creates a subject for HTML parsed as the content of an element
// with the given tag, for example "div", "tbody", or "select".
// The reader is consumed once, when the subject first loads.
func Fragment(contextTag string, r io.Reader, opts ...ct.Option) *ct.Subject {
	return ct.From(&source{desc: fmt.Sprintf("html fragment in %s", contextTag), reader: r, context: contextTag}, opts...)
}

// FragmentBytes is [Fragment] for bytes, which are copied.
func FragmentBytes(contextTag string, data []byte, opts ...ct.Option) *ct.Subject {
	return Fragment(contextTag, bytes.NewReader(bytes.Clone(data)), opts...)
}

// FragmentString is [Fragment] for a string.
func FragmentString(contextTag, data string, opts ...ct.Option) *ct.Subject {
	return Fragment(contextTag, strings.NewReader(data), opts...)
}

type source struct {
	desc    string
	reader  io.Reader
	context string // fragment context tag; "" for a page
}

var _ ct.Source = (*source)(nil)

func (s *source) String() string {
	return s.desc
}

func (s *source) Load(context.Context) (ct.Document, error) {
	r := s.reader
	s.reader = nil // consumed once
	if r == nil {
		return nil, fmt.Errorf("html input already consumed")
	}
	if s.context == "" {
		doc, err := html.Parse(r)
		if err != nil {
			return nil, fmt.Errorf("parse html page: %w", err)
		}
		return document{root: doc}, nil
	}
	ctx := &html.Node{
		Type:     html.ElementNode,
		Data:     s.context,
		DataAtom: atom.Lookup([]byte(s.context)),
	}
	nodes, err := html.ParseFragment(r, ctx)
	if err != nil {
		return nil, fmt.Errorf("parse html fragment in %s: %w", s.context, err)
	}
	root := &html.Node{Type: html.DocumentNode}
	for _, n := range nodes {
		root.AppendChild(n)
	}
	return document{root: root}, nil
}

// Document is the parsed tree. It exposes the underlying x/net/html root for
// custom assertions.
type Document interface {
	ct.Document
	// HTML returns the parsed root node. Do not mutate it while the subject
	// is in use.
	HTML() *html.Node
}

type document struct {
	root *html.Node
}

var _ Document = document{}

func (d document) Root() ct.Node {
	return node{n: d.root}
}

func (d document) HTML() *html.Node {
	return d.root
}

// node adapts an x/net/html node. It is a comparable value wrapping a pointer.
type node struct {
	n *html.Node
}

var _ ct.Node = node{}

func (v node) Kind() ct.NodeKind {
	switch v.n.Type {
	case html.ElementNode:
		return ct.KindElement
	case html.TextNode:
		return ct.KindText
	case html.CommentNode:
		return ct.KindComment
	case html.DoctypeNode:
		return ct.KindDoctype
	default:
		return ct.KindDocument
	}
}

func (v node) Tag() string {
	if v.n.Type != html.ElementNode {
		return ""
	}
	return strings.ToLower(v.n.Data)
}

func (v node) Data() string {
	if v.n.Type == html.ElementNode {
		return ""
	}
	return v.n.Data
}

func (v node) Attr(name string) (string, bool) {
	for _, a := range v.n.Attr {
		if a.Key == name {
			return a.Val, true
		}
	}
	return "", false
}

func (v node) Attrs() iter.Seq2[string, string] {
	return func(yield func(string, string) bool) {
		for _, a := range v.n.Attr {
			if !yield(a.Key, a.Val) {
				return
			}
		}
	}
}

func (v node) Parent() ct.Node {
	if v.n.Parent == nil {
		return nil
	}
	return node{n: v.n.Parent}
}

func (v node) Children() iter.Seq[ct.Node] {
	return func(yield func(ct.Node) bool) {
		for c := v.n.FirstChild; c != nil; c = c.NextSibling {
			if !yield(node{n: c}) {
				return
			}
		}
	}
}

// Location formats the path from the root like "html[0]/body[1]/form#login[0]/[2]".
func (v node) Location() string {
	var steps []string
	for cur := v.n; cur.Parent != nil; cur = cur.Parent {
		index := 0
		for s := cur.PrevSibling; s != nil; s = s.PrevSibling {
			index++
		}
		step := ""
		if cur.Type == html.ElementNode {
			step = strings.ToLower(cur.Data)
			if id := simpleID(cur); id != "" {
				step += "#" + id
			}
		}
		steps = append(steps, fmt.Sprintf("%s[%d]", step, index))
	}
	for i, j := 0, len(steps)-1; i < j; i, j = i+1, j-1 {
		steps[i], steps[j] = steps[j], steps[i]
	}
	return strings.Join(steps, "/")
}

func simpleID(n *html.Node) string {
	for _, a := range n.Attr {
		if a.Key != "id" || a.Val == "" {
			continue
		}
		for _, c := range a.Val {
			switch {
			case c >= 'a' && c <= 'z', c >= 'A' && c <= 'Z', c >= '0' && c <= '9':
			case c == '-', c == '_', c == '.', c == ':':
			default:
				return ""
			}
		}
		return a.Val
	}
	return ""
}
