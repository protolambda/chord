package attr_test

import (
	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/html/attr"
	"github.com/protolambda/chord/html/group/div"
)

func ExampleClass() {
	core.Dump(div.Div(attr.Class("container")))
	// Output: <div class="container"></div>
}

func ExampleClass_multiple() {
	core.Dump(div.Div(attr.Class("foo"), attr.Class("bar")))
	// Output: <div class="foo bar"></div>
}

func ExampleID() {
	core.Dump(div.Div(attr.ID("main")))
	// Output: <div id="main"></div>
}

func ExampleStyle() {
	core.Dump(div.Div(attr.Style("color:red")))
	// Output: <div style="color:red"></div>
}

func ExampleStyle_multiple() {
	core.Dump(div.Div(attr.Style("color:red"), attr.Style("font-size:12px")))
	// Output: <div style="color:red;font-size:12px"></div>
}

func ExampleData() {
	core.Dump(div.Div(attr.Data("id", "123"), attr.Data("name", "test")))
	// Output: <div data-id="123" data-name="test"></div>
}

func ExampleHidden() {
	core.Dump(div.Div(attr.Hidden()))
	// Output: <div hidden></div>
}

func ExampleTitle() {
	core.Dump(div.Div(attr.Title("Tooltip text")))
	// Output: <div title="Tooltip text"></div>
}

func ExampleTabindex() {
	core.Dump(div.Div(attr.Tabindex("0")))
	// Output: <div tabindex="0"></div>
}

func ExampleLang() {
	core.Dump(div.Div(attr.Lang("en")))
	// Output: <div lang="en"></div>
}
