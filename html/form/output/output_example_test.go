package output_test

import (
	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/core/attrib"
	"github.com/protolambda/chord/html/form/output"
	"github.com/protolambda/chord/html/text"
)

func ExampleOutput() {
	core.Dump(output.Output(output.For("a b"), attrib.KV("name", "result"))(text.Text("0")))
	// Output: <output for="a b" name="result">0</output>
}

func ExampleProgress() {
	core.Dump(output.Progress(output.Value("70"), output.Max("100")))
	// Output: <progress value="70" max="100"></progress>
}

func ExampleMeter() {
	core.Dump(output.Meter(output.Value("0.6"), output.Min("0"), output.Max("1")))
	// Output: <meter value="0.6" min="0" max="1"></meter>
}

func ExampleProgress_indeterminate() {
	core.Dump(output.Progress()(text.Text("Loading...")))
	// Output: <progress>Loading...</progress>
}
