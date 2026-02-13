package elem

import (
	"context"
	"html"
	"iter"
	"slices"

	"github.com/protolambda/chord/core/attrib"
)

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
	Attribs attrib.Seq
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

// New creates a non-void element constructor (e.g. <div>scope content</div>).
// Returns a Scope: call it with child elements to produce an Elem.
func New(name string, attrs ...attrib.Node) Scope {
	return func(children ...Node) Node {
		return Obj{
			Tag:      name,
			Attribs:  attrib.Seq(slices.Values(attrs)),
			Children: slices.Values(children),
		}
	}
}

// Void creates a self-closing void element (e.g. <br/>, <input/>).
func Void(name string, attrs ...attrib.Node) Node {
	return Obj{
		Tag:     name,
		Void:    true,
		Attribs: attrib.Seq(slices.Values(attrs)),
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
