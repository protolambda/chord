package core

import (
	"context"
	"strings"

	"github.com/protolambda/chord/core/elem"
)

// FallbackFn is called when a render error occurs to provide alternative content.
type FallbackFn func(ctx context.Context, err error) elem.Node

type fallback struct {
	inner elem.Node
	onErr FallbackFn
}

func (f fallback) Eval(ctx context.Context) (elem.Obj, error) {
	var out strings.Builder
	err := Render(ctx, f.inner, &out)
	if err != nil {
		alt := f.onErr(ctx, err)
		return alt.Eval(ctx)
	}
	return elem.Raw(out.String()).Eval(ctx)
}

// Fallback renders an element node, and can fall back the rendering to the given function,
// in case of an error during rendering.
func Fallback(node elem.Node, onErr FallbackFn) elem.Node {
	return fallback{inner: node, onErr: onErr}
}
