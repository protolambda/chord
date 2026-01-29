package core

func ExampleNoop() {
	Dump(Noop())
	// Output:
}

func ExampleRaw() {
	Dump(Element("div", Raw(`<span>Hello World</span>`)))
	// Output: <div><span>Hello World</span></div>
}

func ExampleVoidElement() {
	Dump(VoidElement("hr"))
	// Output: <hr/>
}

func ExampleBundle() {
	Dump(Bundle(Raw("123"), VoidElement("br"), Raw("hello world")))
	// Output: 123<br/>hello world
}

func ExampleCompose() {
	Dump(Compose(Element("h1", Raw("Title")), Element("p", Raw("Content"))))
	// Output: <h1>Title</h1><p>Content</p>
}

func ExampleCompose_single() {
	Dump(Compose(Element("div", Raw("only one"))))
	// Output: <div>only one</div>
}

func ExampleBoolAttribute() {
	Dump(VoidElement("input", Attribute("type", "checkbox"), BoolAttribute("checked")))
	// Output: <input type="checkbox" checked/>
}

func ExampleAttribute() {
	Dump(Element("div", Attribute("data-x", "123")))
	// Output: <div data-x="123"></div>
}

func ExampleComment() {
	Dump(Element("div", Comment("Hello world")))
	// Output: <div><!-- Hello world --></div>
}

func ExampleWithIndent() {
	Dump(
		Element("div", Attribute("class", "outer"), Attribute("id", "root"),
			Element("div", Attribute("class", "middle"),
				Element("div", Attribute("class", "inner"),
					Raw("Hello"),
				),
				Element("span", Raw("World")),
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
