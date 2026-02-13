package core_test

import (
	"context"
	"errors"

	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/core/attrib"
	"github.com/protolambda/chord/core/elem"
	"github.com/protolambda/chord/html/text"
)

// Fallback() catches render errors and substitutes alternative content.
// Useful for graceful degradation of page sections.
func Example_fallback() {
	unreliable := elem.Fn(func(ctx context.Context) (elem.Node, error) {
		return nil, errors.New("database unavailable")
	})

	core.Dump(core.Fallback(
		elem.New("div")(unreliable),
		func(ctx context.Context, err error) elem.Node {
			return elem.New("div", attrib.KV("class", "error"))(
				text.Text("Failed to load content"))
		},
	))
	// Output: <div class="error">Failed to load content</div>
}
