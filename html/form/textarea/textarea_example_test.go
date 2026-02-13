package textarea_test

import (
	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/html/form/textarea"
	"github.com/protolambda/chord/html/text"
)

func ExampleTextarea() {
	core.Dump(textarea.Textarea(textarea.Name("message"))(text.Text("Default text")))
	// Output: <textarea name="message">Default text</textarea>
}

func ExampleTextarea_sized() {
	core.Dump(textarea.Textarea(textarea.Rows("5"), textarea.Cols("40")))
	// Output: <textarea rows="5" cols="40"></textarea>
}

func ExampleTextarea_placeholder() {
	core.Dump(textarea.Textarea(textarea.Placeholder("Enter your message...")))
	// Output: <textarea placeholder="Enter your message..."></textarea>
}

func ExampleTextarea_required() {
	core.Dump(textarea.Textarea(textarea.Name("bio"), textarea.Required()))
	// Output: <textarea name="bio" required></textarea>
}
