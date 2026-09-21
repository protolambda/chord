package core

import (
	"context"

	"github.com/protolambda/chord/core/elem"
	"github.com/protolambda/chord/core/inspect"
)

// FallbackFn is called when evaluating the primary content fails, to provide
// alternative content. The error carries the location of the failure.
type FallbackFn func(ctx context.Context, err error) elem.Node

type fallback struct {
	inner elem.Node
	onErr FallbackFn
}

func (f fallback) Eval(ctx context.Context) (elem.Obj, error) {
	doc, err := inspect.Build(ctx, f.inner)
	if err != nil {
		return f.onErr(ctx, err).Eval(ctx)
	}
	return doc.Root().Eval(ctx)
}

// Fallback evaluates node transactionally: the whole subtree is evaluated into
// a buffer first, and only if that succeeds is it rendered. If evaluation
// fails, the subtree is discarded and onErr provides the content instead.
//
// Only evaluation failures trigger the fallback. Output failures happen after
// evaluation and are reported as render errors, since a broken writer cannot
// be repaired by alternative markup. The primary content runs at most once.
func Fallback(node elem.Node, onErr FallbackFn) elem.Node {
	return fallback{inner: node, onErr: onErr}
}
