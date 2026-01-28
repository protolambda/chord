package util

import (
	"context"
	"strings"

	"github.com/protolambda/chord/core"
)

// Fn is a type of node that generates its output dynamically.
type Fn func(ctx context.Context) (core.Node, error)

var _ core.Node = Fn(nil)

func (fn Fn) Eval(ctx context.Context) (core.Obj, error) {
	out, err := fn(ctx)
	if err != nil {
		return core.Obj{}, err
	}
	return out.Eval(ctx)
}

// If as a function makes the node conditional,
// and substitutes it with a no-op node if false.
func If(cond bool, node core.Node) core.Node {
	if cond {
		return node
	} else {
		return core.Noop()
	}
}

type FallbackFn func(ctx context.Context, err error) core.Node

type fallback struct {
	Inner core.Node
	OnErr FallbackFn
}

func (e fallback) Eval(ctx context.Context) (core.Obj, error) {
	var out strings.Builder
	err := core.Render(ctx, e.Inner, &out)
	if err != nil {
		// If we have an error, wrap it, to make it a non-error.
		// This strategy can be used to "guard" a sub-graph by showing
		// an in-place HTML error if it fails to render.
		alt := e.OnErr(ctx, err)
		return alt.Eval(ctx)
	}
	return core.Raw(out.String()).Eval(ctx)
}

// Fallback renders a node, and can fall back the rendering to the given function,
// in case of an error during rendering.
// This can be used to continue page-rendering with a sub-section of the page
// falling back to some visual error alternative, like a user-warning.
func Fallback(node core.Node, onErr FallbackFn) core.Node {
	return fallback{Inner: node, OnErr: onErr}
}
