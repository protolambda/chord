package interactive_test

import (
	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/core/attr"
	"github.com/protolambda/chord/html/interactive"
	"github.com/protolambda/chord/html/text"
)

func ExampleDetails() {
	core.Dump(interactive.Details()(
		interactive.Summary()(text.Text("Click to expand")),
		text.P()(text.Text("Hidden content")),
	))
	// Output: <details><summary>Click to expand</summary><p>Hidden content</p></details>
}

func ExampleDetails_open() {
	core.Dump(interactive.Details(interactive.Open())(
		interactive.Summary()(text.Text("Expanded")),
	))
	// Output: <details open><summary>Expanded</summary></details>
}

func ExampleDialog() {
	core.Dump(interactive.Dialog(attr.ID("modal"))(text.P()(text.Text("Dialog content"))))
	// Output: <dialog id="modal"><p>Dialog content</p></dialog>
}

func ExampleDialog_open() {
	core.Dump(interactive.Dialog(interactive.Open()))
	// Output: <dialog open></dialog>
}
