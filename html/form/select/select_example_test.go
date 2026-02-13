package selectel_test

import (
	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/core/attrib"
	selectel "github.com/protolambda/chord/html/form/select"
	"github.com/protolambda/chord/html/text"
)

func ExampleSelect() {
	core.Dump(selectel.Select()(
		selectel.Option()(text.Text("Option 1")),
		selectel.Option()(text.Text("Option 2")),
	))
	// Output: <select><option>Option 1</option><option>Option 2</option></select>
}

func ExampleSelect_named() {
	core.Dump(selectel.Select(selectel.Name("country"))(
		selectel.Option(selectel.Value("us"))(text.Text("USA")),
		selectel.Option(selectel.Value("uk"))(text.Text("UK")),
	))
	// Output: <select name="country"><option value="us">USA</option><option value="uk">UK</option></select>
}

func ExampleOptgroup() {
	core.Dump(selectel.Select()(
		selectel.Optgroup(attrib.KV("label", "Fruits"))(
			selectel.Option()(text.Text("Apple")),
			selectel.Option()(text.Text("Banana")),
		),
	))
	// Output: <select><optgroup label="Fruits"><option>Apple</option><option>Banana</option></optgroup></select>
}

func ExampleDatalist() {
	core.Dump(selectel.Datalist(attrib.KV("id", "browsers"))(
		selectel.Option(selectel.Value("Chrome")),
		selectel.Option(selectel.Value("Firefox")),
	))
	// Output: <datalist id="browsers"><option value="Chrome"></option><option value="Firefox"></option></datalist>
}

func ExampleSelect_multiple() {
	core.Dump(selectel.Select(selectel.Multiple()))
	// Output: <select multiple></select>
}
