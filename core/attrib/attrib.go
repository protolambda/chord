package attrib

import (
	"context"
	"iter"
)

// Obj is the evaluated representation of an attribute.
type Obj struct {
	// Key is the attribute key. Empty for bundle/noop.
	Key string
	// Val is the attribute value. Assumed to be html-escaped already.
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

// KV creates an attribute node with the given key-value pair.
func KV(k, v string) Node {
	return Obj{Key: k, Val: v}
}

// Bool creates a boolean attribute (presence-only, no value).
func Bool(k string) Node {
	return Obj{Key: k, Bool: true}
}
