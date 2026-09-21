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
	default:
		return fmt.Sprintf("kind(%d)", uint8(k))
	}
}

// Obj is the evaluated representation of an element.
//
// Each [Kind] uses a subset of the fields; the other fields must stay zero:
//   - [KindNoop]: nothing.
//   - [KindFragment]: Children.
//   - [KindElement]: Tag, Attribs, Void, and Children (only when not Void).
//   - [KindText], [KindRaw], [KindComment]: Data.
//
// For compatibility with literals written before Kind existed, an Obj with a
// zero Kind is interpreted as an element when Tag is set, and as a fragment
// when only Children is set. See [Obj.Validate]. Prefer the constructors in
// this package over literals.
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
}

// Obj can be used as a Node itself
var _ Node = Obj{}

func (o Obj) Eval(ctx context.Context) (Obj, error) {
	return o, nil
}

// Validate returns the effective kind of the object, applying the legacy
// zero-kind inference described on [Obj], and reports an error wrapping
// [ErrInvalidObj] when the populated fields are inconsistent with that kind.
func (o Obj) Validate() (Kind, error) {
	kind := o.Kind
	if kind == KindNoop {
		switch {
		case o.Tag != "":
			kind = KindElement
		case o.Children != nil:
			kind = KindFragment
		}
	}
	switch kind {
	case KindNoop:
		if o.Data != "" || o.Attribs != nil {
			return kind, fmt.Errorf("%w: noop with data or attributes", ErrInvalidObj)
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
	case KindText, KindRaw, KindComment:
		if o.Tag != "" || o.Void || o.Attribs != nil || o.Children != nil {
			return kind, fmt.Errorf("%w: %s with tag, void, attributes, or children", ErrInvalidObj, kind)
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

// Text creates a text node. The text is HTML-escaped when rendered.
func Text(v string) Node {
	return Obj{Kind: KindText, Data: v}
}

// Raw creates an element that renders the given string directly (no escaping).
func Raw(v string) Node {
	return Obj{Kind: KindRaw, Data: v}
}

// Comment creates an HTML comment. The content is HTML-escaped when rendered,
// so it cannot terminate the comment early.
func Comment(content string) Node {
	return Obj{Kind: KindComment, Data: content}
}
