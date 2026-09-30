package script_test

import (
	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/core/attr"
	"github.com/protolambda/chord/html/script"
	"github.com/protolambda/chord/html/text"
)

func ExampleScript() {
	core.Dump(script.Script(script.Src("/app.js")))
	// Output: <script src="/app.js"></script>
}

func ExampleScript_inline() {
	core.Dump(script.Script()(text.Text("if (a < b) console.log('hello');")))
	// Output: <script>if (a < b) console.log('hello');</script>
}

func ExampleInline() {
	core.Dump(script.Inline(`import { start } from "/app.js"; start();`, script.Type("module")))
	// Output: <script type="module">import { start } from "/app.js"; start();</script>
}

func ExampleInline_unsafe() {
	core.Dump(script.Inline(`document.write("</script>");`))
	// Output: ERROR: render: at script[0]/[0]: unsafe script or style text: script text contains "</script"
}

func ExampleScript_module() {
	core.Dump(script.Script(script.Type("module"), script.Src("/mod.js")))
	// Output: <script type="module" src="/mod.js"></script>
}

func ExampleNoscript() {
	core.Dump(script.Noscript()(text.Text("JavaScript required")))
	// Output: <noscript>JavaScript required</noscript>
}

func ExampleTemplate() {
	core.Dump(script.Template(attr.ID("item-tpl"))(text.Span()(text.Text("Item"))))
	// Output: <template id="item-tpl"><span>Item</span></template>
}

func ExampleCanvas() {
	core.Dump(script.Canvas(attr.ID("game"), attr.KV("width", "800"), attr.KV("height", "600")))
	// Output: <canvas id="game" width="800" height="600"></canvas>
}
