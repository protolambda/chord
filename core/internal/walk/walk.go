// Package walk evaluates a node graph exactly once, in depth-first order, and
// delivers it to a [Receiver] as a balanced event stream.
//
// The stream is strict and sequential: every successful Open is followed by
// exactly one Close once traversal completes, content events apply to the
// most recently opened element, fragments and bundles never produce events,
// and attributes are complete and normalized before Open. Events never carry
// a location; the walker only uses locations to annotate errors.
//
// After an error, traversal stops immediately. Opened elements are not closed
// in that case, and the receiver's partial state is the caller's problem.
//
// This package is internal: the receiver contract is the seam shared by the
// renderer and the inspection snapshot, and is not a public API.
package walk

import (
	"context"
	"fmt"
	"html"
	"strings"

	"github.com/protolambda/chord/core/attr"
	"github.com/protolambda/chord/core/elem"
)

// Receiver consumes the event stream produced by [Walk].
type Receiver interface {
	// Open starts an element. Void elements are also closed by a following
	// Close call, without content events in between.
	Open(el Element) error
	// Text delivers logical text that must be escaped for HTML output.
	Text(v string) error
	// Raw delivers trusted, output-ready HTML.
	Raw(v string) error
	// Comment delivers logical comment content.
	Comment(v string) error
	// Close ends the most recently opened element.
	Close() error
}

// Element describes an opened element.
type Element struct {
	Tag  string
	Void bool
	// Attrs is flattened, deduplicated, and merged. The slice is borrowed
	// and only valid during the Open call; copy it to retain it.
	Attrs []Attribute
}

// Attribute is a normalized attribute: never a noop or a bundle.
type Attribute struct {
	Key string
	// Val is logical for [attr.KindValue], output-ready for [attr.KindRawValue],
	// and empty for [attr.KindBool].
	Val  string
	Kind attr.Kind
}

// Walk evaluates node and its descendants, delivering events to r.
// Application code reached through the node graph runs exactly once.
// The context is checked before each element is opened.
// Errors are returned with the location of the failure, as a prefix
// formatted like "at html[0]/body[1]/form#login[0]/[2]": opened elements as
// tag, optional simple id, and child index; a trailing bare index is the
// child that was being evaluated.
func Walk(ctx context.Context, node elem.Node, r Receiver) error {
	w := &walker{
		ctx:    ctx,
		r:      r,
		frames: []frame{{inChildren: true}},
		seen:   make(map[string]int),
	}
	if err := w.node(node); err != nil {
		return fmt.Errorf("at %s: %w", w.location(), err)
	}
	return nil
}

type walker struct {
	ctx    context.Context
	r      Receiver
	frames []frame // frames[0] is the virtual root
	attrs  []Attribute
	seen   map[string]int
}

type frame struct {
	tag   string
	id    string
	index int // index within the parent's normalized children
	next  int // index of the next normalized child
	// inChildren is true while children are evaluated, so that a failure
	// is located at the pending child rather than at the element itself.
	inChildren bool
}

func (w *walker) top() *frame {
	return &w.frames[len(w.frames)-1]
}

func (w *walker) location() string {
	var b strings.Builder
	for _, f := range w.frames[1:] {
		if b.Len() > 0 {
			b.WriteByte('/')
		}
		b.WriteString(f.tag)
		if f.id != "" {
			b.WriteByte('#')
			b.WriteString(f.id)
		}
		fmt.Fprintf(&b, "[%d]", f.index)
	}
	if top := w.top(); top.inChildren {
		if b.Len() > 0 {
			b.WriteByte('/')
		}
		fmt.Fprintf(&b, "[%d]", top.next)
	}
	return b.String()
}

func (w *walker) node(n elem.Node) error {
	obj, err := n.Eval(w.ctx)
	if err != nil {
		return err
	}
	kind, err := obj.Validate()
	if err != nil {
		return err
	}
	switch kind {
	case elem.KindNoop:
		return nil
	case elem.KindFragment:
		for child := range obj.Children {
			if err := w.node(child); err != nil {
				return err
			}
		}
		return nil
	case elem.KindText:
		return w.content(w.r.Text(obj.Data))
	case elem.KindRaw:
		return w.content(w.r.Raw(obj.Data))
	case elem.KindComment:
		return w.content(w.r.Comment(obj.Data))
	case elem.KindElement:
		return w.element(obj)
	default:
		return fmt.Errorf("unsupported %s object", kind)
	}
}

func (w *walker) content(err error) error {
	if err != nil {
		return err
	}
	w.top().next++
	return nil
}

func (w *walker) element(obj elem.Obj) error {
	if err := w.ctx.Err(); err != nil {
		return err
	}
	w.frames = append(w.frames, frame{tag: obj.Tag, index: w.top().next})

	attrs, err := w.attributes(obj.Attribs)
	if err != nil {
		return err
	}
	w.top().id = SimpleID(attrs)
	if err := w.r.Open(Element{Tag: obj.Tag, Void: obj.Void, Attrs: attrs}); err != nil {
		return err
	}

	if obj.Children != nil {
		w.top().inChildren = true
		for child := range obj.Children {
			if err := w.node(child); err != nil {
				return err
			}
		}
		w.top().inChildren = false
	}

	if err := w.r.Close(); err != nil {
		return err
	}
	w.frames = w.frames[:len(w.frames)-1]
	w.top().next++
	return nil
}

// attributes flattens the attribute sequence into the reusable buffer.
func (w *walker) attributes(seq attr.Seq) ([]Attribute, error) {
	w.attrs = w.attrs[:0]
	clear(w.seen)
	if seq == nil {
		return nil, nil
	}
	for a := range seq {
		if err := w.attribute(a); err != nil {
			return nil, err
		}
	}
	return w.attrs, nil
}

func (w *walker) attribute(a attr.Node) error {
	obj, err := a.Eval(w.ctx)
	if err != nil {
		return fmt.Errorf("attribute: %w", err)
	}
	kind, err := obj.Validate()
	if err != nil {
		return err
	}
	switch kind {
	case attr.KindNoop:
		return nil
	case attr.KindBundle:
		for sub := range obj.Sub {
			if err := w.attribute(sub); err != nil {
				return err
			}
		}
		return nil
	}

	part := Attribute{Key: obj.Key, Val: obj.Val, Kind: kind}
	if i, seen := w.seen[obj.Key]; seen {
		switch obj.Key {
		case "class":
			return mergeValue(&w.attrs[i], part, " ")
		case "style":
			return mergeValue(&w.attrs[i], part, ";")
		default:
			return fmt.Errorf("duplicate attribute %q", obj.Key)
		}
	}
	w.seen[obj.Key] = len(w.attrs)
	w.attrs = append(w.attrs, part)
	return nil
}

// mergeValue appends part to dst, separated by sep.
// When one side is logical and the other raw, the merged value becomes raw
// (output-ready) and the logical side is escaped first, so that the rendered
// bytes equal the concatenation of the individually rendered parts.
func mergeValue(dst *Attribute, part Attribute, sep string) error {
	if dst.Kind == attr.KindBool || part.Kind == attr.KindBool {
		return fmt.Errorf("duplicate attribute %q", part.Key)
	}
	if dst.Kind != part.Kind {
		if dst.Kind == attr.KindValue {
			dst.Val = html.EscapeString(dst.Val)
		}
		if part.Kind == attr.KindValue {
			part.Val = html.EscapeString(part.Val)
		}
		dst.Kind = attr.KindRawValue
	}
	dst.Val += sep + part.Val
	return nil
}

// SimpleID returns the id attribute value when it is a plain token that is
// safe to show in a location, and "" otherwise.
func SimpleID(attrs []Attribute) string {
	for _, a := range attrs {
		if a.Key != "id" || a.Kind == attr.KindBool || a.Val == "" {
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
