package form_test

import (
	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/html/form"
	"github.com/protolambda/chord/html/form/button"
	"github.com/protolambda/chord/html/form/input"
	"github.com/protolambda/chord/html/text"
)

func ExampleForm() {
	core.Dump(form.Form(form.Action("/submit"), form.Method("post"),
		input.Input(input.Type("text"), input.Name("name")),
		button.Button(button.Type(button.TypeSubmit), text.Text("Submit")),
	))
	// Output: <form action="/submit" method="post"><input type="text" name="name"/><button type="submit">Submit</button></form>
}

func ExampleFieldset() {
	core.Dump(form.Fieldset(
		form.Legend(text.Text("Personal Info")),
		input.Input(input.Type("text"), input.Name("name")),
	))
	// Output: <fieldset><legend>Personal Info</legend><input type="text" name="name"/></fieldset>
}

func ExampleForm_get() {
	core.Dump(form.Form(form.Action("/search"), form.Method("get")))
	// Output: <form action="/search" method="get"></form>
}
