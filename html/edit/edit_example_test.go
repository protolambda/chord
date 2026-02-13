package edit_test

import (
	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/html/edit"
	"github.com/protolambda/chord/html/text"
)

func ExampleIns() {
	core.Dump(edit.Ins()(text.Text("new text")))
	// Output: <ins>new text</ins>
}

func ExampleDel() {
	core.Dump(edit.Del()(text.Text("removed")))
	// Output: <del>removed</del>
}

func Example_revision() {
	core.Dump(text.P()(edit.Del()(text.Text("old")), text.Text(" "), edit.Ins()(text.Text("new"))))
	// Output: <p><del>old</del> <ins>new</ins></p>
}
