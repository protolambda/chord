package list_test

import (
	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/html/group/list"
	"github.com/protolambda/chord/html/text"
)

func ExampleUL() {
	core.Dump(list.UL(
		list.LI(text.Text("Item 1")),
		list.LI(text.Text("Item 2")),
	))
	// Output: <ul><li>Item 1</li><li>Item 2</li></ul>
}

func ExampleOL() {
	core.Dump(list.OL(
		list.LI(text.Text("First")),
		list.LI(text.Text("Second")),
	))
	// Output: <ol><li>First</li><li>Second</li></ol>
}

func ExampleOL_start() {
	core.Dump(list.OL(list.Start("5"), list.LI(text.Text("Item"))))
	// Output: <ol start="5"><li>Item</li></ol>
}

func ExampleMenu() {
	core.Dump(list.Menu(list.LI(text.Text("Action"))))
	// Output: <menu><li>Action</li></menu>
}

func ExampleLI_value() {
	core.Dump(list.LI(list.Value("3"), text.Text("Third")))
	// Output: <li value="3">Third</li>
}
