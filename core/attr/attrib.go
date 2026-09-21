package attr

import (
	"context"
	"errors"
	"fmt"
	"iter"
	"unicode"
	"unicode/utf8"
)

// ErrInvalidName indicates that an attribute name is not safe HTML syntax.
var ErrInvalidName = errors.New("invalid attribute name")

// ErrInvalidObj indicates that an [Obj] literal combines fields that do not
// belong to its [Kind].
var ErrInvalidObj = errors.New("invalid attribute object")

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

// Kind identifies what an evaluated [Obj] represents.
type Kind uint8

const (
	// KindNoop produces no output. It is the zero value, so an empty Obj is a no-op.
	KindNoop Kind = iota
	// KindBundle is a sequence of attributes applied to the same element.
	KindBundle
	// KindValue is a key with a logical value. The renderer HTML-escapes the value on output.
	KindValue
	// KindRawValue is a key with a trusted, output-ready value that is written verbatim.
	KindRawValue
	// KindBool is a boolean attribute: presence only, no value.
	KindBool
)

func (k Kind) String() string {
	switch k {
	case KindNoop:
		return "noop"
	case KindBundle:
		return "bundle"
	case KindValue:
		return "value"
	case KindRawValue:
		return "raw value"
	case KindBool:
		return "bool"
	default:
		return fmt.Sprintf("kind(%d)", uint8(k))
	}
}

// Obj is the evaluated representation of an attribute.
//
// Each [Kind] uses a subset of the fields; the other fields must stay zero:
//   - [KindNoop]: nothing.
//   - [KindBundle]: Sub.
//   - [KindValue], [KindRawValue]: Key and Val.
//   - [KindBool]: Key.
//
// For compatibility, an Obj with a zero Kind and a Sub sequence is a bundle.
// An Obj with a zero Kind and a Key is rejected by [Obj.Validate], because
// whether its value is escaped cannot be known. Prefer the constructors in
// this package over literals.
type Obj struct {
	// Kind selects the interpretation of the other fields.
	Kind Kind
	// Key is the attribute name.
	Key string
	// Val is the logical value for KindValue, or the output-ready value for KindRawValue.
	Val string
	// Sub holds the attributes of a bundle.
	Sub iter.Seq[Node]
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
	if kind == KindNoop && o.Sub != nil {
		kind = KindBundle
	}
	switch kind {
	case KindNoop:
		if o.Key != "" || o.Val != "" {
			return kind, fmt.Errorf("%w: key %q without kind", ErrInvalidObj, o.Key)
		}
	case KindBundle:
		if o.Key != "" || o.Val != "" {
			return kind, fmt.Errorf("%w: bundle with key or value", ErrInvalidObj)
		}
	case KindValue, KindRawValue:
		if o.Key == "" {
			return kind, fmt.Errorf("%w: %s without key", ErrInvalidObj, kind)
		}
		if o.Sub != nil {
			return kind, fmt.Errorf("%w: %s %q with sub-attributes", ErrInvalidObj, kind, o.Key)
		}
	case KindBool:
		if o.Key == "" {
			return kind, fmt.Errorf("%w: bool without key", ErrInvalidObj)
		}
		if o.Val != "" || o.Sub != nil {
			return kind, fmt.Errorf("%w: bool %q with value or sub-attributes", ErrInvalidObj, o.Key)
		}
	default:
		return kind, fmt.Errorf("%w: unknown %s", ErrInvalidObj, kind)
	}
	return kind, nil
}

type Node interface {
	Eval(ctx context.Context) (Obj, error)
}

type invalidName string

func (n invalidName) Eval(context.Context) (Obj, error) {
	return Obj{}, fmt.Errorf("%w %q", ErrInvalidName, string(n))
}

// KV creates an attribute node from a runtime name and value.
// It validates the name during construction. The value is HTML-escaped when rendered.
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

// Value creates an attribute with this trusted name and a logical value,
// which is HTML-escaped when rendered.
func (n Name) Value(v string) Node {
	return Obj{Kind: KindValue, Key: string(n), Val: v}
}

// Raw creates an attribute with this trusted name and an output-ready value.
// The caller must ensure that v cannot break out of the quoted HTML attribute.
func (n Name) Raw(v string) Node {
	return Obj{Kind: KindRawValue, Key: string(n), Val: v}
}

// Bool creates a boolean attribute with this trusted name.
func (n Name) Bool() Node {
	return Obj{Kind: KindBool, Key: string(n)}
}
