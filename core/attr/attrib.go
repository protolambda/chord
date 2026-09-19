package attr

import (
	"context"
	"errors"
	"fmt"
	"html"
	"iter"
	"unicode"
	"unicode/utf8"
)

// ErrInvalidName indicates that an attribute name is not safe HTML syntax.
var ErrInvalidName = errors.New("invalid attribute name")

// Name is a trusted, statically known HTML attribute name.
// Converting runtime input to Name bypasses attribute-name validation.
type Name string

// ParseName validates a runtime HTML attribute name.
func ParseName(v string) (out Name, ok bool) {
	if v == "" || !utf8.ValidString(v) {
		return "", false
	}
	for _, r := range v {
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

// Obj is the evaluated representation of an attribute.
type Obj struct {
	// Key is the attribute key. Empty for bundle/noop.
	Key string
	// Val is the output-ready attribute value, either escaped or trusted raw text.
	Val string
	// Bool indicates a boolean attribute (no value, just presence).
	Bool bool
	// Sub holds sub-attributes for bundles.
	Sub iter.Seq[Node]
}

// Obj can be used as a Node itself
var _ Node = Obj{}

func (o Obj) Eval(ctx context.Context) (Obj, error) {
	return o, nil
}

type Node interface {
	Eval(ctx context.Context) (Obj, error)
}

type invalidName string

func (n invalidName) Eval(context.Context) (Obj, error) {
	return Obj{}, fmt.Errorf("%w %q", ErrInvalidName, string(n))
}

// KV creates an attribute node from a runtime name and value.
// It validates the name and HTML-escapes the value once during construction.
func KV(k, v string) Node {
	name, ok := ParseName(k)
	if !ok {
		return invalidName(k)
	}
	return name.Value(v)
}

// Bool creates a boolean attribute from a runtime name.
// It validates the name during construction.
func Bool(k string) Node {
	name, ok := ParseName(k)
	if !ok {
		return invalidName(k)
	}
	return name.Bool()
}

// Value creates an attribute with this trusted name and an HTML-escaped value.
func (n Name) Value(v string) Node {
	return Obj{Key: string(n), Val: html.EscapeString(v)}
}

// Raw creates an attribute with this trusted name and an output-ready value.
// The caller must ensure that v cannot break out of the quoted HTML attribute.
func (n Name) Raw(v string) Node {
	return Obj{Key: string(n), Val: v}
}

// Bool creates a boolean attribute with this trusted name.
func (n Name) Bool() Node {
	return Obj{Key: string(n), Bool: true}
}
