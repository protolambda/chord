package meta_test

import (
	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/html/meta"
	"github.com/protolambda/chord/html/text"
)

func ExampleHTML() {
	core.Dump(meta.HTML(core.Attribute("lang", "en")))
	// Output: <html lang="en"></html>
}

func ExampleHead() {
	core.Dump(meta.Head(meta.Title(text.Text("Page"))))
	// Output: <head><title>Page</title></head>
}

func ExampleTitle() {
	core.Dump(meta.Title(text.Text("My Page")))
	// Output: <title>My Page</title>
}

func ExampleMeta() {
	core.Dump(meta.Meta(meta.Charset("utf-8")))
	// Output: <meta charset="utf-8"/>
}

func ExampleMeta_viewport() {
	core.Dump(meta.Meta(meta.Name("viewport"), meta.Content("width=device-width")))
	// Output: <meta name="viewport" content="width=device-width"/>
}

func ExampleLink() {
	core.Dump(meta.Link(meta.Rel("stylesheet"), meta.Href("/style.css")))
	// Output: <link rel="stylesheet" href="/style.css"/>
}

func ExampleBase() {
	core.Dump(meta.Base(meta.Href("https://example.com/")))
	// Output: <base href="https://example.com/"/>
}

func ExampleStyle() {
	core.Dump(meta.Style(core.Raw("body { margin: 0; }")))
	// Output: <style>body { margin: 0; }</style>
}
