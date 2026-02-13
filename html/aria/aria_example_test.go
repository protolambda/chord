package aria_test

import (
	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/html/aria"
	"github.com/protolambda/chord/html/form/button"
	"github.com/protolambda/chord/html/group/div"
	"github.com/protolambda/chord/html/text"
)

func ExampleLabel() {
	core.Dump(button.Button(aria.Label("Close dialog"))(text.Text("X")))
	// Output: <button aria-label="Close dialog">X</button>
}

func ExampleHidden() {
	core.Dump(div.Div(aria.Hidden("true")))
	// Output: <div aria-hidden="true"></div>
}

func ExampleExpanded() {
	core.Dump(button.Button(aria.Expanded("false"))(text.Text("Menu")))
	// Output: <button aria-expanded="false">Menu</button>
}

func ExampleRole() {
	core.Dump(div.Div(aria.Role("alert"))(text.Text("Error!")))
	// Output: <div role="alert">Error!</div>
}

func ExampleDescribedby() {
	core.Dump(div.Div(aria.Describedby("desc-1")))
	// Output: <div aria-describedby="desc-1"></div>
}
