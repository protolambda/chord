package label_test

import (
	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/html/form/input"
	"github.com/protolambda/chord/html/form/label"
	"github.com/protolambda/chord/html/text"
)

func ExampleLabel() {
	core.Dump(label.Label(label.For("email"))(text.Text("Email:")))
	// Output: <label for="email">Email:</label>
}

func ExampleLabel_wrapped() {
	core.Dump(label.Label()(
		text.Text("Name: "),
		input.Input(input.Type("text"), input.Name("name")),
	))
	// Output: <label>Name: <input type="text" name="name"/></label>
}
