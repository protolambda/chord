package button_test

import (
	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/html/form/button"
	"github.com/protolambda/chord/html/text"
)

func ExampleButton() {
	core.Dump(button.Button()(text.Text("Click me")))
	// Output: <button>Click me</button>
}

func ExampleButton_submit() {
	core.Dump(button.Button(button.Type(button.TypeSubmit))(text.Text("Submit")))
	// Output: <button type="submit">Submit</button>
}

func ExampleButton_reset() {
	core.Dump(button.Button(button.Type(button.TypeReset))(text.Text("Reset")))
	// Output: <button type="reset">Reset</button>
}

func ExampleButton_disabled() {
	core.Dump(button.Button(button.Disabled())(text.Text("Disabled")))
	// Output: <button disabled>Disabled</button>
}
