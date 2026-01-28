package iframe_test

import (
	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/html/embed/iframe"
)

func ExampleIframe() {
	core.Dump(iframe.Iframe(iframe.Src("https://example.com")))
	// Output: <iframe src="https://example.com"></iframe>
}

func ExampleIframe_sized() {
	core.Dump(iframe.Iframe(iframe.Src("/embed"), iframe.Width("600"), iframe.Height("400")))
	// Output: <iframe src="/embed" width="600" height="400"></iframe>
}

func ExampleIframe_sandbox() {
	core.Dump(iframe.Iframe(iframe.Src("/app"), iframe.Sandbox("allow-scripts")))
	// Output: <iframe src="/app" sandbox="allow-scripts"></iframe>
}
