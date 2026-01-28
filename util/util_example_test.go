package util_test

import (
	"context"

	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/html/attr"
	"github.com/protolambda/chord/html/group/div"
	"github.com/protolambda/chord/html/text"
	"github.com/protolambda/chord/util"
)

func ExampleIf_true() {
	core.Dump(div.Div(util.If(true, attr.Class("active"))))
	// Output: <div class="active"></div>
}

func ExampleIf_false() {
	core.Dump(div.Div(util.If(false, attr.Class("hidden"))))
	// Output: <div></div>
}

func ExampleFn() {
	core.Dump(util.Fn(func(ctx context.Context) (core.Node, error) {
		return text.Text("Dynamic content"), nil
	}))
	// Output: Dynamic content
}

func Example_conditional() {
	isAdmin := true
	core.Dump(div.Div(
		text.Text("Welcome"),
		util.If(isAdmin, text.Span(attr.Class("badge"), text.Text("Admin"))),
	))
	// Output: <div>Welcome<span class="badge">Admin</span></div>
}
