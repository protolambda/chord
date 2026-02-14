package div_test

import (
	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/core/attr"
	"github.com/protolambda/chord/html/group/div"
	"github.com/protolambda/chord/html/text"
)

func ExampleDiv() {
	core.Dump(div.Div()(text.Text("Hello")))
	// Output: <div>Hello</div>
}

func ExampleDiv_styled() {
	core.Dump(div.Div(attr.Class("container"), attr.ID("main"))(text.Text("Content")))
	// Output: <div class="container" id="main">Content</div>
}

func ExampleDiv_nested() {
	core.Dump(div.Div(attr.Class("outer"))(
		div.Div(attr.Class("inner"))(text.Text("Nested")),
	))
	// Output: <div class="outer"><div class="inner">Nested</div></div>
}
