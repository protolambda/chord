package elem

import (
	"context"
	"errors"
	"fmt"
	"iter"
	"slices"
	"unicode"
	"unicode/utf8"

	"github.com/protolambda/chord/core/attr"
)

// ErrInvalidObj indicates that an [Obj] literal combines fields that do not
// belong to its [Kind].
var ErrInvalidObj = errors.New("invalid element object")

// ErrUnsafeText indicates text inside an HTML script or style element that
// could end the element early: "</script" or "<!--" in a script, "</style"
// in a style element, or the end tag of an enclosing element whose content
// is text, such as "</noscript" or "</textarea" (ASCII case-insensitive).
// Such text is written literally, so it is refused rather than escaped. See
// [Text].
var ErrUnsafeText = errors.New("unsafe script or style text")

// ErrNilNode indicates a nil [Node] where content was expected: a nil child,
// a nil [Fn] or [Scope], or an Fn that returned a nil node without an error.
// Use [Noop] for "no content".
var ErrNilNode = errors.New("nil element node")

// Name is a trusted, statically known HTML element tag name.
// Converting runtime input to Name bypasses tag-name validation.
type Name string

// ParseName validates a runtime HTML element tag name.
// The first character must be an ASCII letter. Subsequent characters may
// include Unicode, but not characters that can reshape HTML tag syntax.
func ParseName(v string) (out Name, ok bool) {
	if v == "" || !utf8.ValidString(v) || !asciiLetter(v[0]) {
		return "", false
	}
	for _, r := range v[1:] {
		if unicode.IsSpace(r) || unicode.IsControl(r) {
			return "", false
		}
		switch r {
		case '"', '\'', '<', '>', '/', '=':
			return "", false
		}
	}
	return Name(v), true
}

func asciiLetter(c byte) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z'
}

// Kind identifies what an evaluated [Obj] represents.
type Kind uint8

const (
	// KindNoop produces no output. It is the zero value, so an empty Obj is a no-op.
	KindNoop Kind = iota
	// KindFragment is a sequence of children rendered adjacent to each other, without a wrapper.
	KindFragment
	// KindElement is an HTML element with a tag, attributes, and (unless void) children.
	KindElement
	// KindText is logical text. The renderer HTML-escapes it on output.
	KindText
	// KindRaw is trusted, output-ready HTML that is written verbatim.
	KindRaw
	// KindComment is logical comment content. The renderer wraps and escapes it on output.
	KindComment
	// KindDoctype is a document type declaration. Data is the doctype name,
	// which must be "html": the renderer writes <!DOCTYPE html>.
	KindDoctype
)

func (k Kind) String() string {
	switch k {
	case KindNoop:
		return "noop"
	case KindFragment:
		return "fragment"
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

// Obj is the evaluated representation of an element.
//
// Each [Kind] uses a subset of the fields; the other fields must stay zero:
//   - [KindNoop]: nothing.
//   - [KindFragment]: Children and Context.
//   - [KindElement]: Tag, Attribs, Void, Children (only when not Void), and Context.
//   - [KindText], [KindRaw], [KindComment]: Data.
//   - [KindDoctype]: Data, which must be "html".
//
// For compatibility with literals written before Kind existed, an Obj with a
// zero Kind is interpreted as an element when Tag is set, and as a fragment
// when only Children is set. See [Obj.Validate]. Prefer the constructors in
// this package over literals.
//
// # Context
//
// Evaluation is lazy: [Node.Eval] returns an Obj whose Attribs and Children
// are evaluated later, when the renderer reaches them, after Eval has
// returned. Context carries the context for that later evaluation, so that a
// node can scope a context to its subtree:
//
//	func (n withUser) Eval(ctx context.Context) (elem.Obj, error) {
//		return n.inner.Eval(context.WithValue(ctx, userKey{}, n.user))
//	}
//
// The render's context (the ctx passed to the renderer) is the root of every
// scope. The attributes and children of an Obj are evaluated with:
//
//  1. its Context, when that derives from the render's context: it is the ctx
//     that Eval received, or a context made from it with the context package;
//  2. otherwise the context of the enclosing scope, which for the root is the
//     render's context.
//
// So Context can only narrow the render's context (add values or a
// deadline), never replace it. A Context that does not derive from the
// current render, such as the one an Obj recorded in an earlier render and
// kept since, or context.Background(), is ignored: values of one render
// (e.g. of one HTTP request) never reach another, and a canceled context of
// an earlier render does not fail a later one. Retain nodes rather than
// evaluated objects: a retained Obj loses the context that its node derived.
//
// Who sets Context:
//   - [Obj.Eval], [Bundle] and [Seq] record the ctx they are evaluated with,
//     on the Obj they return, so a node that delegates to them passes its
//     context on.
//   - A node that returns an Obj literal with attributes or children should
//     set Context itself, to ctx or a context derived from it, or return the
//     literal's Eval(ctx). A nil Context means the enclosing scope.
//
// Only the render's context stops the render. A node may derive a context
// with a timeout for its own work and release it when Eval returns:
//
//	func (n withAccount) Eval(ctx context.Context) (elem.Obj, error) {
//		ctx, cancel := context.WithTimeout(ctx, time.Second)
//		defer cancel()
//		account, err := n.store.Load(ctx, n.id)
//		if err != nil {
//			return elem.Obj{}, err
//		}
//		return n.inner.Eval(context.WithValue(ctx, accountKey{}, account))
//	}
//
// A Context that is already done when the renderer reaches the object, as
// here, still provides its values to the subtree, but with the cancellation
// of the enclosing scope instead of its own. A Context that ends later,
// while the subtree renders, is seen canceled by the nodes below it; the
// render itself goes on.
type Obj struct {
	// Kind selects the interpretation of the other fields.
	Kind Kind
	// Tag is the element tag name.
	Tag string
	// Data is logical text, trusted raw HTML, or logical comment content, depending on Kind.
	Data string
	// Void indicates a self-closing element (e.g. <br/>, <input/>).
	Void bool
	// Attribs holds the element's attributes.
	Attribs attr.Seq
	// Children holds the child elements of a fragment or non-void element.
	Children iter.Seq[Node]
	// Context is the context for evaluating the attributes and children of a
	// fragment or element, recorded when the object was evaluated. The
	// renderer uses it only when it derives from the context of the current
	// render, and the enclosing scope's context otherwise (also when nil).
	// See the Context section above.
	Context context.Context
}

// Obj can be used as a Node itself
var _ Node = Obj{}

// Eval returns the object. A fragment or element without a Context records
// ctx as the Context of the returned copy, so that its attributes and
// children are evaluated with the context that evaluated it.
//
// The receiver is a value on purpose: Eval changes its own copy of o, never
// the Obj it is called on. An Obj that is a node of a tree, such as one that
// an element constructor returned and that every render reuses, keeps a nil
// Context, and each render records its own context on the copy that it
// renders. An Obj that already has a Context keeps it; the renderer decides
// whether to use it (see the Context section of [Obj]).
func (o Obj) Eval(ctx context.Context) (Obj, error) {
	if o.Context == nil {
		switch o.kind() {
		case KindFragment, KindElement:
			// Sets the copy's Context only (value receiver, see above).
			o.Context = ctx
		}
	}
	return o, nil
}

// kind returns the effective kind, applying the legacy zero-kind inference.
func (o Obj) kind() Kind {
	if o.Kind != KindNoop {
		return o.Kind
	}
	switch {
	case o.Tag != "":
		return KindElement
	case o.Children != nil:
		return KindFragment
	}
	return KindNoop
}

// Validate returns the effective kind of the object, applying the legacy
// zero-kind inference described on [Obj], and reports an error wrapping
// [ErrInvalidObj] when the populated fields are inconsistent with that kind.
func (o Obj) Validate() (Kind, error) {
	kind := o.kind()
	switch kind {
	case KindNoop:
		if o.Data != "" || o.Attribs != nil || o.Context != nil {
			return kind, fmt.Errorf("%w: noop with data, attributes, or context", ErrInvalidObj)
		}
	case KindFragment:
		if o.Tag != "" || o.Data != "" || o.Void || o.Attribs != nil {
			return kind, fmt.Errorf("%w: fragment with tag, data, void, or attributes", ErrInvalidObj)
		}
	case KindElement:
		if o.Tag == "" {
			return kind, fmt.Errorf("%w: element without tag", ErrInvalidObj)
		}
		if o.Data != "" {
			return kind, fmt.Errorf("%w: element %q with data", ErrInvalidObj, o.Tag)
		}
		if o.Void && o.Children != nil {
			return kind, fmt.Errorf("%w: void element %q with children", ErrInvalidObj, o.Tag)
		}
	case KindText, KindRaw, KindComment, KindDoctype:
		if o.Tag != "" || o.Void || o.Attribs != nil || o.Children != nil || o.Context != nil {
			return kind, fmt.Errorf("%w: %s with tag, void, attributes, children, or context", ErrInvalidObj, kind)
		}
		if kind == KindDoctype && o.Data != "html" {
			return kind, fmt.Errorf("%w: doctype %q, only \"html\" is supported", ErrInvalidObj, o.Data)
		}
	default:
		return kind, fmt.Errorf("%w: unknown %s", ErrInvalidObj, kind)
	}
	return kind, nil
}

type Node interface {
	Eval(ctx context.Context) (Obj, error)
}

// New creates a non-void element constructor for this trusted tag.
// Call the returned Scope with child elements to produce a Node.
func (n Name) New(attrs ...attr.Node) Scope {
	return func(children ...Node) Node {
		return Obj{
			Kind:     KindElement,
			Tag:      string(n),
			Attribs:  attr.Seq(slices.Values(attrs)),
			Children: slices.Values(children),
		}
	}
}

// Void creates a self-closing element for this trusted tag.
func (n Name) Void(attrs ...attr.Node) Node {
	return Obj{
		Kind:    KindElement,
		Tag:     string(n),
		Void:    true,
		Attribs: attr.Seq(slices.Values(attrs)),
	}
}

// Text creates a text node. The text is HTML-escaped when rendered, except in
// HTML script and style elements, whose content the browser does not
// unescape: there the text is written literally, and rendering fails with
// [ErrUnsafeText] when it could end the element early, or an enclosing
// element whose content is text (such as noscript or textarea). Where
// parsers do not create the script or style element and read its text as
// markup (inside svg and math, for a style inside select, and after a
// frameset), the text stays escaped. Text in a script is code, and text in a
// style is CSS for the whole page, not escaped data: never put untrusted text
// there.
func Text(v string) Node {
	return Obj{Kind: KindText, Data: v}
}

// Raw creates an element that renders the given string directly (no escaping).
func Raw(v string) Node {
	return Obj{Kind: KindRaw, Data: v}
}

// Doctype creates the document type declaration of HTML, <!DOCTYPE html>,
// to put before the html element of a page. Unlike an equivalent [Raw]
// node, it is a node of its own kind, [KindDoctype], which inspection and
// tests see as a doctype. It must not be placed inside an element.
func Doctype() Node {
	return Obj{Kind: KindDoctype, Data: "html"}
}

// Comment creates an HTML comment. The content is HTML-escaped when rendered,
// so it cannot terminate the comment early.
func Comment(content string) Node {
	return Obj{Kind: KindComment, Data: content}
}
