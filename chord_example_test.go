package chord

import (
	"context"
	"errors"

	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/input"
)

func ExampleDiv_minimal() {
	core.Dump(Div(Text("hello world")))
	// Output: <div>hello world</div>
}

func ExampleDiv_basics() {
	// multiple classes are allowed!
	core.Dump(Div(ID("my-div"), Class("foo"), Class("bar")))
	// Output: <div id="my-div" class="foo bar"></div>
}

func ExampleDiv_duplicate() {
	// normal attributes should not be duplicated!
	core.Dump(Div(ID("my-div"), ID("alt-id")))
	// Output: ERROR: failed to render: duplicate attribute "id"
}

func ExampleDiv_fallback() {
	core.Dump(Fallback(Div(Class("wrapper"), Fn(func(ctx context.Context) (core.Node, error) {
		return nil, errors.New("silly error")
	})), func(ctx context.Context, err error) core.Node {
		return Div(Class("err"), Text("gotcha: "+err.Error()))
	}))
	// Output: <div class="err">gotcha: failed to render: failed to open: silly error</div>
}

func ExampleDiv_bundle() {
	core.Dump(core.Bundle(Div(ID("a")), Div(ID("b"))))
	// Output: <div id="a"></div><div id="b"></div>
}

func ExampleInput() {
	core.Dump(Input(input.Type("checkbox"), input.Checked()))
	// Output: <input type="checkbox" checked/>
}

func ExampleDiv_full() {
	core.Dump(
		Div(
			Class("outer"),
			ID("x"),
			Input(
				Class("foo"),
				Class("foo2"),
				core.Bundle( // combine options as one
					Class("bar"),
					Class("baz"),
				),
				input.Size("123"),
				// hx.Swap("")
			),
			If(false, Class("excluded")),
			If(true, Class("included")),
			Text("interesting"),
		),
	)
	// Output:
	// <div class="outer included" id="x"><input class="foo foo2 bar baz" size="123"/>interesting</div>
}

func ExampleDiv_styled() {
	core.Dump(Div(Style("color:blue"), Style("background:#222"), Text("styled")))
	// Output: <div style="color:blue;background:#222">styled</div>
}
