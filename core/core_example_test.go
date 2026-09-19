package core

import (
	"github.com/protolambda/chord/core/attr"
	"github.com/protolambda/chord/core/elem"
)

func ExampleNoop() {
	Dump(elem.Noop())
	// Output:
}

func ExampleRaw() {
	Dump(elem.Name("div").New()(elem.Raw(`<span>Hello World</span>`)))
	// Output: <div><span>Hello World</span></div>
}

func ExampleBundle() {
	Dump(elem.Bundle{elem.Raw("123"), elem.Name("br").Void(), elem.Raw("hello world")})
	// Output: 123<br/>hello world
}

func ExampleCons() {
	Dump(elem.Cons(elem.Name("h1").New()(elem.Raw("Title")), elem.Name("p").New()(elem.Raw("Content"))))
	// Output: <h1>Title</h1><p>Content</p>
}

func ExampleCons_single() {
	Dump(elem.Cons(elem.Name("div").New()(elem.Raw("only one"))))
	// Output: <div>only one</div>
}

func ExampleBool() {
	Dump(elem.Name("input").Void(attr.KV("type", "checkbox"), attr.Bool("checked")))
	// Output: <input type="checkbox" checked/>
}

func ExampleKV() {
	Dump(elem.Name("div").New(attr.KV("data-x", "123")))
	// Output: <div data-x="123"></div>
}

func ExampleComment() {
	Dump(elem.Name("div").New()(elem.Comment("Hello world")))
	// Output: <div><!-- Hello world --></div>
}

func ExampleWithIndent() {
	Dump(
		elem.Name("div").New(attr.KV("class", "outer"), attr.KV("id", "root"))(
			elem.Name("div").New(attr.KV("class", "middle"))(
				elem.Name("div").New(attr.KV("class", "inner"))(
					elem.Raw("Hello"),
				),
				elem.Name("span").New()(elem.Raw("World")),
			),
		),
		WithIndent(),
	)
	// Output:
	// <div class="outer" id="root">
	//   <div class="middle">
	//     <div class="inner">
	//       Hello
	//     </div>
	//     <span>
	//       World
	//     </span>
	//   </div>
	// </div>
}

func ExampleClass() {
	Dump(elem.Name("div").New(attr.Class("container")))
	// Output: <div class="container"></div>
}

func ExampleClass_multiple() {
	Dump(elem.Name("div").New(attr.Class("foo"), attr.Class("bar")))
	// Output: <div class="foo bar"></div>
}

func ExampleID() {
	Dump(elem.Name("div").New(attr.ID("main")))
	// Output: <div id="main"></div>
}

func ExampleStyle() {
	Dump(elem.Name("div").New(attr.Style("color:red")))
	// Output: <div style="color:red"></div>
}

func ExampleStyle_multiple() {
	Dump(elem.Name("div").New(attr.Style("color:red"), attr.Style("font-size:12px")))
	// Output: <div style="color:red;font-size:12px"></div>
}

func ExampleData() {
	Dump(elem.Name("div").New(attr.Data("id", "123"), attr.Data("name", "test")))
	// Output: <div data-id="123" data-name="test"></div>
}

func ExampleHidden() {
	Dump(elem.Name("div").New(attr.Hidden()))
	// Output: <div hidden></div>
}

func ExampleTitle() {
	Dump(elem.Name("div").New(attr.Title("Tooltip text")))
	// Output: <div title="Tooltip text"></div>
}

func ExampleTabindex() {
	Dump(elem.Name("div").New(attr.Tabindex("0")))
	// Output: <div tabindex="0"></div>
}

func ExampleLang() {
	Dump(elem.Name("div").New(attr.Lang("en")))
	// Output: <div lang="en"></div>
}
