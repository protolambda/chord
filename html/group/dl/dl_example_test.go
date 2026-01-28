package dl_test

import (
	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/html/group/dl"
	"github.com/protolambda/chord/html/text"
)

func ExampleDL() {
	core.Dump(dl.DL(
		dl.DT(text.Text("Term")),
		dl.DD(text.Text("Definition")),
	))
	// Output: <dl><dt>Term</dt><dd>Definition</dd></dl>
}

func ExampleDL_multiple() {
	core.Dump(dl.DL(
		dl.DT(text.Text("HTML")),
		dl.DD(text.Text("HyperText Markup Language")),
		dl.DT(text.Text("CSS")),
		dl.DD(text.Text("Cascading Style Sheets")),
	))
	// Output: <dl><dt>HTML</dt><dd>HyperText Markup Language</dd><dt>CSS</dt><dd>Cascading Style Sheets</dd></dl>
}
