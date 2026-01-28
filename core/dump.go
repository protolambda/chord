package core

import (
	"context"
	"fmt"
	"strings"
)

// Dump prints the node to stdout,
// for convenience in testing and debugging.
func Dump(v Node) {
	DumpCtx(context.Background(), v)
}

// DumpCtx prints the node (evaluated with ctx) to stdout,
// for convenience in testing and debugging.
func DumpCtx(ctx context.Context, v Node) {
	var out strings.Builder
	err := Render(ctx, v, &out)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println(out.String())
}
