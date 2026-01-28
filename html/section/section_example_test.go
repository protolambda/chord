package section_test

import (
	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/html/section"
	"github.com/protolambda/chord/html/text"
)

func ExampleH1() {
	core.Dump(section.H1(text.Text("Main Title")))
	// Output: <h1>Main Title</h1>
}

func ExampleH2() {
	core.Dump(section.H2(text.Text("Section Title")))
	// Output: <h2>Section Title</h2>
}

func ExampleHeader() {
	core.Dump(section.Header(section.H1(text.Text("Site"))))
	// Output: <header><h1>Site</h1></header>
}

func ExampleFooter() {
	core.Dump(section.Footer(text.Text("Copyright 2024")))
	// Output: <footer>Copyright 2024</footer>
}

func ExampleNav() {
	core.Dump(section.Nav(text.A(text.Href("/"), text.Text("Home"))))
	// Output: <nav><a href="/">Home</a></nav>
}

func ExampleMain() {
	core.Dump(section.Main(text.P(text.Text("Content"))))
	// Output: <main><p>Content</p></main>
}

func ExampleArticle() {
	core.Dump(section.Article(section.H2(text.Text("Post")), text.P(text.Text("Body"))))
	// Output: <article><h2>Post</h2><p>Body</p></article>
}

func ExampleSection() {
	core.Dump(section.Section(section.H2(text.Text("About"))))
	// Output: <section><h2>About</h2></section>
}

func ExampleAside() {
	core.Dump(section.Aside(text.Text("Sidebar")))
	// Output: <aside>Sidebar</aside>
}

func ExampleAddress() {
	core.Dump(section.Address(text.Text("123 Main St")))
	// Output: <address>123 Main St</address>
}
