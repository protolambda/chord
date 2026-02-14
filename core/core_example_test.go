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
	Dump(elem.New("div")(elem.Raw(`<span>Hello World</span>`)))
	// Output: <div><span>Hello World</span></div>
}

func ExampleVoid() {
	Dump(elem.Void("hr"))
	// Output: <hr/>
}

func ExampleBundle() {
	Dump(elem.Bundle{elem.Raw("123"), elem.Void("br"), elem.Raw("hello world")})
	// Output: 123<br/>hello world
}

func ExampleCons() {
	Dump(elem.Cons(elem.New("h1")(elem.Raw("Title")), elem.New("p")(elem.Raw("Content"))))
	// Output: <h1>Title</h1><p>Content</p>
}

func ExampleCons_single() {
	Dump(elem.Cons(elem.New("div")(elem.Raw("only one"))))
	// Output: <div>only one</div>
}

func ExampleBool() {
	Dump(elem.Void("input", attr.KV("type", "checkbox"), attr.Bool("checked")))
	// Output: <input type="checkbox" checked/>
}

func ExampleKV() {
	Dump(elem.New("div", attr.KV("data-x", "123")))
	// Output: <div data-x="123"></div>
}

func ExampleComment() {
	Dump(elem.New("div")(elem.Comment("Hello world")))
	// Output: <div><!-- Hello world --></div>
}

func ExampleWithIndent() {
	Dump(
		elem.New("div", attr.KV("class", "outer"), attr.KV("id", "root"))(
			elem.New("div", attr.KV("class", "middle"))(
				elem.New("div", attr.KV("class", "inner"))(
					elem.Raw("Hello"),
				),
				elem.New("span")(elem.Raw("World")),
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
	Dump(elem.New("div", attr.Class("container")))
	// Output: <div class="container"></div>
}

func ExampleClass_multiple() {
	Dump(elem.New("div", attr.Class("foo"), attr.Class("bar")))
	// Output: <div class="foo bar"></div>
}

func ExampleID() {
	Dump(elem.New("div", attr.ID("main")))
	// Output: <div id="main"></div>
}

func ExampleStyle() {
	Dump(elem.New("div", attr.Style("color:red")))
	// Output: <div style="color:red"></div>
}

func ExampleStyle_multiple() {
	Dump(elem.New("div", attr.Style("color:red"), attr.Style("font-size:12px")))
	// Output: <div style="color:red;font-size:12px"></div>
}

func ExampleData() {
	Dump(elem.New("div", attr.Data("id", "123"), attr.Data("name", "test")))
	// Output: <div data-id="123" data-name="test"></div>
}

func ExampleHidden() {
	Dump(elem.New("div", attr.Hidden()))
	// Output: <div hidden></div>
}

func ExampleTitle() {
	Dump(elem.New("div", attr.Title("Tooltip text")))
	// Output: <div title="Tooltip text"></div>
}

func ExampleTabindex() {
	Dump(elem.New("div", attr.Tabindex("0")))
	// Output: <div tabindex="0"></div>
}

func ExampleLang() {
	Dump(elem.New("div", attr.Lang("en")))
	// Output: <div lang="en"></div>
}
