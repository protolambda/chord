package input_test

import (
	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/html/form/input"
)

func ExampleInput() {
	core.Dump(input.Input(input.Type("text"), input.Name("username")))
	// Output: <input type="text" name="username"/>
}

func ExampleInput_checkbox() {
	core.Dump(input.Input(input.Type("checkbox"), input.Name("agree"), input.Checked()))
	// Output: <input type="checkbox" name="agree" checked/>
}

func ExampleInput_password() {
	core.Dump(input.Input(input.Type("password"), input.Name("pwd"), input.Required()))
	// Output: <input type="password" name="pwd" required/>
}

func ExampleInput_number() {
	core.Dump(input.Input(input.Type("number"), input.Min("0"), input.Max("100"), input.Step("5")))
	// Output: <input type="number" min="0" max="100" step="5"/>
}

func ExampleInput_placeholder() {
	core.Dump(input.Input(input.Type("email"), input.Placeholder("you@example.com")))
	// Output: <input type="email" placeholder="you@example.com"/>
}

func ExampleInput_disabled() {
	core.Dump(input.Input(input.Type("text"), input.Disabled()))
	// Output: <input type="text" disabled/>
}

func ExampleInput_readonly() {
	core.Dump(input.Input(input.Type("text"), input.Value("fixed"), input.Readonly()))
	// Output: <input type="text" value="fixed" readonly/>
}
