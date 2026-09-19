package elem

import (
	"context"
	"html"
	"iter"
	"slices"
	"unicode"
	"unicode/utf8"

	"github.com/protolambda/chord/core/attr"
)

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

// Obj is the evaluated representation of an element.
type Obj struct {
	// Tag is the element tag name. Empty for raw content or bundles.
	Tag string
	// Raw is pre-rendered HTML content (used for raw/comment nodes).
	// When set, Tag is a sentinel (e.g. "RAW", "COMMENT") to avoid squashing.
	Raw string
	// Void indicates a self-closing element (e.g. <br/>, <input/>).
	Void bool
	// Attribs holds the element's attributes.
	Attribs attr.Seq
	// Children holds the element's child elements.
	Children iter.Seq[Node]
}

// Obj can be used as a Node itself
var _ Node = Obj{}

func (o Obj) Eval(ctx context.Context) (Obj, error) {
	return o, nil
}

type Node interface {
	Eval(ctx context.Context) (Obj, error)
}

// New creates a non-void element constructor for this trusted tag.
// Call the returned Scope with child elements to produce a Node.
func (n Name) New(attrs ...attr.Node) Scope {
	return func(children ...Node) Node {
		return Obj{
			Tag:      string(n),
			Attribs:  attr.Seq(slices.Values(attrs)),
			Children: slices.Values(children),
		}
	}
}

// Void creates a self-closing element for this trusted tag.
func (n Name) Void(attrs ...attr.Node) Node {
	return Obj{
		Tag:     string(n),
		Void:    true,
		Attribs: attr.Seq(slices.Values(attrs)),
	}
}

// Raw creates an element that renders the given string directly (no escaping).
func Raw(v string) Node {
	return Obj{
		Tag: "RAW",
		Raw: v,
	}
}

// Comment creates an HTML comment with string-escaped text content.
func Comment(content string) Node {
	return Obj{
		Tag: "COMMENT",
		Raw: "<!-- " + html.EscapeString(content) + " -->",
	}
}
