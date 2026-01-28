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

func ExampleBoolAttribute() {
	Dump(VoidElement("input", Attribute("type", "checkbox"), BoolAttribute("checked")))
	// Output: <input type="checkbox" checked/>
}

func ExampleAttribute() {
	Dump(Element("div", Attribute("data-x", "123")))
	// Output: <div data-x="123"></div>
}
