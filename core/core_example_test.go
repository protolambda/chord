package core

import (
	"github.com/protolambda/chord/core/attrib"
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
	Dump(elem.Void("input", attrib.KV("type", "checkbox"), attrib.Bool("checked")))
	// Output: <input type="checkbox" checked/>
}

func ExampleKV() {
	Dump(elem.New("div", attrib.KV("data-x", "123")))
	// Output: <div data-x="123"></div>
}

func ExampleComment() {
	Dump(elem.New("div")(elem.Comment("Hello world")))
	// Output: <div><!-- Hello world --></div>
}

func ExampleWithIndent() {
	Dump(
		elem.New("div", attrib.KV("class", "outer"), attrib.KV("id", "root"))(
			elem.New("div", attrib.KV("class", "middle"))(
				elem.New("div", attrib.KV("class", "inner"))(
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
