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
	"slices"
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
	// Doctype delivers a document type declaration with the given name. It
	// only occurs outside elements.
	Doctype(name string) error
	// Close ends the most recently opened element.
	Close() error
}

// Element describes an opened element.
type Element struct {
	Tag  string
	Void bool
	// Literal is set for HTML script and style elements: their text
	// children must be written verbatim, not escaped. The walker has checked
	// that the text cannot end the element early, nor an enclosing element
	// whose content is text (such as noscript or textarea). It is not set
	// where parsers do not create the element and read its text as markup:
	// inside svg and math, for style inside select, and after a frameset.
	Literal bool
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
// ctx is the root of every scope: the attributes and children of an object
// are evaluated with the object's [elem.Obj.Context] when it derives from ctx
// (checked with a marker that Walk adds to ctx), and with the context of the
// enclosing scope otherwise, so that a Context recorded in another walk is
// never used. A Context that is already done, which
// its node released when Eval returned, provides its values only, with the
// cancellation of the parent's context. Only ctx stops the walk: it is
// checked before each element is opened, so canceling it stops the walk even
// below a node that detached its context from ctx's cancellation.
// Errors are returned with the location of the failure, as a prefix
// formatted like "at html[0]/body[1]/form#login[0]/[2]": opened elements as
// tag, optional simple id, and child index; a trailing bare index is the
// child that was being evaluated.
func Walk(ctx context.Context, node elem.Node, r Receiver) error {
	w := &walker{
		mark:   new(walkMark),
		r:      r,
		frames: []frame{{inChildren: true}},
		seen:   make(map[string]int),
	}
	// Mark the context of this walk, so that scope contexts can be told apart
	// from contexts recorded outside of it (see enter).
	ctx = context.WithValue(ctx, walkKey{}, w.mark)
	w.root, w.ctx = ctx, ctx
	if err := w.node(node); err != nil {
		return fmt.Errorf("at %s: %w", w.location(), err)
	}
	return nil
}

type walker struct {
	root context.Context // the context of the walk
	ctx  context.Context // the context of the current scope
	// mark is the value of the walk marker in root. It is not the walker
	// itself, so that a context that application code retains does not keep
	// the walker and its receiver alive.
	mark   *walkMark
	r      Receiver
	frames []frame // frames[0] is the virtual root
	attrs  []Attribute
	seen   map[string]int
	// frameset is set once a frameset element was opened: parsers ignore
	// script and style start tags inside and after it, and read their text
	// as markup, so text is no longer written literally.
	frameset bool
}

type frame struct {
	tag   string
	id    string
	index int // index within the parent's normalized children
	next  int // index of the next normalized child
	// inChildren is true while children are evaluated, so that a failure
	// is located at the pending child rather than at the element itself.
	inChildren bool
	// foreign is set inside svg and math, where script and style are
	// ordinary elements.
	foreign bool
	// inSelect is set inside select, where older parsers ignore a style
	// start tag and read its text as markup.
	inSelect bool
	// inText is set inside an element whose content is text (see
	// textContainer).
	inText bool
	// rawText is "script" or "style" for an HTML element whose text is
	// written literally, and "" otherwise.
	rawText string
	// ends are the lowercase tags of the text containers around a literal
	// element: literal text must not contain their end tags either.
	ends []string
	// tail holds the end of the literal content written so far, so that a
	// sequence split over adjacent text nodes is still found.
	tail string
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
	if n == nil {
		return elem.ErrNilNode
	}
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
		if obj.Children == nil {
			return nil
		}
		outer := w.enter(obj.Context)
		for child := range obj.Children {
			if err := w.node(child); err != nil {
				return err
			}
		}
		w.ctx = outer
		return nil
	case elem.KindText:
		if err := w.literal(obj.Data, true); err != nil {
			return err
		}
		return w.content(w.r.Text(obj.Data))
	case elem.KindRaw:
		if err := w.literal(obj.Data, false); err != nil {
			return err
		}
		return w.content(w.r.Raw(obj.Data))
	case elem.KindComment:
		w.top().tail = ""
		return w.content(w.r.Comment(obj.Data))
	case elem.KindDoctype:
		if len(w.frames) > 1 {
			return fmt.Errorf("doctype inside element %q", w.top().tag)
		}
		return w.content(w.r.Doctype(obj.Data))
	case elem.KindElement:
		return w.element(obj)
	default:
		return fmt.Errorf("unsupported %s object", kind)
	}
}

// literal checks content of a script or style element, which is written
// verbatim: text must not end the element early. Raw content is trusted and
// only tracked, since it may start a sequence that the next text completes.
func (w *walker) literal(v string, check bool) error {
	f := w.top()
	if f.rawText == "" {
		return nil
	}
	joined := f.tail + v
	if check {
		if err := checkLiteral(f.rawText, f.ends, joined); err != nil {
			return err
		}
	}
	f.tail = joined[max(0, len(joined)-maxTail):]
	return nil
}

// maxTail is the length of the longest checked sequence ("</noscript", as
// long as "</textarea" and "</noframes") minus one.
const maxTail = len("</noscript") - 1

// checkLiteral checks the literal text v of a tag element (script or style)
// inside the text containers ends. The text must not contain the end tag of
// the element or of a container, and not "<!--" when the element or a
// container is a script, where it enters the escape states that change
// which end tag closes the script.
func checkLiteral(tag string, ends []string, v string) error {
	if strings.Contains(v, "<!--") && (tag == "script" || slices.Contains(ends, "script")) {
		return fmt.Errorf("%w: %s text contains %q", elem.ErrUnsafeText, tag, "<!--")
	}
	for i := 0; ; {
		j := strings.Index(v[i:], "</")
		if j < 0 {
			return nil
		}
		i += j + 2
		if end := endTagAt(v[i:], tag, ends); end != "" {
			return fmt.Errorf("%w: %s text contains %q", elem.ErrUnsafeText, tag, v[i-2:i+len(end)])
		}
	}
}

// endTagAt returns the lowercase tag, or the first of ends, that s starts
// with, ignoring ASCII case, and "" when there is none.
func endTagAt(s, tag string, ends []string) string {
	for _, end := range append([]string{tag}, ends...) {
		if len(s) >= len(end) && asciiEqualFold(s[:len(end)], end) {
			return end
		}
	}
	return ""
}

func (w *walker) content(err error) error {
	if err != nil {
		return err
	}
	w.top().next++
	return nil
}

// walkKey is the context key of the walk marker.
type walkKey struct{}

// walkMark is the value of the walk marker, unique per walk. It is not
// zero-sized, since pointers to distinct zero-sized values may be equal.
type walkMark struct{ _ byte }

// enter switches to the scope context ctx, when it is set and derives from
// the context of this walk, and returns the context to restore afterwards.
// Any other context was recorded by an evaluation outside of this walk, of
// an object that was retained: using it would carry the values of another
// render into this one, or fail on its canceled context. After an error the
// walk stops, so the context is only restored on success.
//
// A scope context that is already done was typically canceled by its node
// when Eval returned (a timeout for the node's own work, released with
// defer cancel()): its values are still used, but with the cancellation of
// the enclosing scope, so that the subtree is not evaluated with a context
// that its ancestor released.
func (w *walker) enter(ctx context.Context) (outer context.Context) {
	outer = w.ctx
	if ctx == nil || ctx.Value(walkKey{}) != any(w.mark) {
		return outer
	}
	if ctx.Err() != nil {
		ctx = scopeValues{Context: outer, values: context.WithoutCancel(ctx)}
	}
	w.ctx = ctx
	return outer
}

// scopeValues is a context with the values of a done scope context and the
// deadline and cancellation of its enclosing scope.
type scopeValues struct {
	context.Context // the enclosing scope
	// values is the done scope context without its cancellation. A lookup
	// ends at its WithoutCancel layer with nil for the context package's
	// private cancellation key.
	values context.Context
}

// Value returns the value for key of the scope context, and otherwise that of
// the enclosing scope. The fallback is for the private cancellation key of
// the context package, which context.Cause and derived cancel contexts look
// up: they find the enclosing scope. For other keys it changes nothing, since
// the scope context derives from the enclosing scope, except that a value
// the node explicitly set to nil shows the enclosing scope's value.
func (c scopeValues) Value(key any) any {
	if v := c.values.Value(key); v != nil {
		return v
	}
	return c.Context.Value(key)
}

func (w *walker) element(obj elem.Obj) error {
	// Only the context of the walk stops it. A scope context that ends
	// during the walk (a timeout of a node) is left to the nodes that use
	// it: static content below it is still delivered.
	if err := w.root.Err(); err != nil {
		return err
	}
	outer := w.enter(obj.Context)
	parent := w.top()
	parent.tail = ""
	if asciiEqualFold(obj.Tag, "frameset") {
		w.frameset = true
	}
	f := frame{
		tag:      obj.Tag,
		index:    parent.next,
		foreign:  parent.foreign || isForeignRoot(obj.Tag),
		inSelect: parent.inSelect || asciiEqualFold(obj.Tag, "select"),
		inText:   parent.inText || textContainer(obj.Tag) != "",
	}
	if !f.foreign && !obj.Void && !w.frameset {
		f.rawText = rawTextTag(obj.Tag)
		if f.rawText == "style" && parent.inSelect {
			f.rawText = ""
		}
		if f.rawText != "" && parent.inText {
			f.ends = w.textContainers()
		}
	}
	w.frames = append(w.frames, f)

	attrs, err := w.attributes(obj.Attribs)
	if err != nil {
		return err
	}
	w.top().id = SimpleID(attrs)
	if err := w.r.Open(Element{Tag: obj.Tag, Void: obj.Void, Literal: f.rawText != "", Attrs: attrs}); err != nil {
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
	w.ctx = outer
	return nil
}

// isForeignRoot reports whether tag starts SVG or MathML content.
func isForeignRoot(tag string) bool {
	return asciiEqualFold(tag, "svg") || asciiEqualFold(tag, "math")
}

// rawTextTag returns the lowercase tag of the HTML elements whose text is
// written literally, and "" for other elements.
func rawTextTag(tag string) string {
	switch {
	case asciiEqualFold(tag, "script"):
		return "script"
	case asciiEqualFold(tag, "style"):
		return "style"
	}
	return ""
}

// textContainer returns the lowercase tag of the HTML elements whose content
// the parser reads as text up to their end tag, and "" for other elements:
// script, style, textarea, title, xmp, iframe, noembed, noframes, and
// noscript (with scripting enabled). Nothing ends the content of plaintext,
// so text inside it cannot break out.
func textContainer(tag string) string {
	for _, c := range [...]string{"script", "style", "textarea", "title", "xmp", "iframe", "noembed", "noframes", "noscript"} {
		if asciiEqualFold(tag, c) {
			return c
		}
	}
	return ""
}

// textContainers returns the lowercase tags of the text containers among the
// open elements.
func (w *walker) textContainers() []string {
	var out []string
	for _, f := range w.frames[1:] {
		if c := textContainer(f.tag); c != "" {
			out = append(out, c)
		}
	}
	return out
}

// asciiEqualFold compares a with the lowercase ASCII string lower, ignoring
// ASCII case only, as HTML does for tag names. Unicode folding would be
// wrong here: strings.EqualFold matches "\u017fcript" (long s) to "script",
// a tag that browsers treat as an unknown element.
func asciiEqualFold(a, lower string) bool {
	if len(a) != len(lower) {
		return false
	}
	for i := 0; i < len(a); i++ {
		c := a[i]
		if 'A' <= c && c <= 'Z' {
			c += 'a' - 'A'
		}
		if c != lower[i] {
			return false
		}
	}
	return true
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
	if a == nil {
		return fmt.Errorf("attribute: %w", attr.ErrNilNode)
	}
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
		if obj.Sub == nil {
			return nil
		}
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
