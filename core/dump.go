package core

import (
	"context"
	"fmt"
	"strings"

	"github.com/protolambda/chord/core/elem"
)

// Dump prints the element to stdout,
// for convenience in testing and debugging.
func Dump(v elem.Node, opts ...Option) {
	DumpCtx(context.Background(), v, opts...)
}

// DumpCtx prints the element (evaluated with ctx) to stdout,
// for convenience in testing and debugging.
func DumpCtx(ctx context.Context, v elem.Node, opts ...Option) {
	var out strings.Builder
	err := Render(ctx, v, &out, opts...)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println(out.String())
}
