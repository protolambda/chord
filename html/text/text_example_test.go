package text_test

import (
	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/html/attr"
	"github.com/protolambda/chord/html/text"
)

func ExampleText() {
	core.Dump(text.Text("Hello <world>"))
	// Output: Hello &lt;world&gt;
}

func ExampleP() {
	core.Dump(text.P(text.Text("Hello world")))
	// Output: <p>Hello world</p>
}

func ExampleSpan() {
	core.Dump(text.Span(attr.Class("highlight"), text.Text("important")))
	// Output: <span class="highlight">important</span>
}

func ExampleA() {
	core.Dump(text.A(text.Href("/page"), text.Text("Link")))
	// Output: <a href="/page">Link</a>
}

func ExampleStrong() {
	core.Dump(text.Strong(text.Text("Bold text")))
	// Output: <strong>Bold text</strong>
}

func ExampleEM() {
	core.Dump(text.EM(text.Text("Emphasized")))
	// Output: <em>Emphasized</em>
}

func ExampleCode() {
	core.Dump(text.Code(text.Text("fmt.Println()")))
	// Output: <code>fmt.Println()</code>
}

func ExamplePre() {
	core.Dump(text.Pre(text.Text("  indented\n  text")))
	// Output: <pre>  indented
	//   text</pre>
}

func ExampleBlockquote() {
	core.Dump(text.Blockquote(text.P(text.Text("A quote"))))
	// Output: <blockquote><p>A quote</p></blockquote>
}

func ExampleBR() {
	core.Dump(core.Bundle(text.Text("Line 1"), text.BR(), text.Text("Line 2")))
	// Output: Line 1<br/>Line 2
}

func ExampleHR() {
	core.Dump(core.Bundle(text.P(text.Text("Section 1")), text.HR(), text.P(text.Text("Section 2"))))
	// Output: <p>Section 1</p><hr/><p>Section 2</p>
}

func ExampleSmall() {
	core.Dump(text.Small(text.Text("fine print")))
	// Output: <small>fine print</small>
}

func ExampleMark() {
	core.Dump(text.Mark(text.Text("highlighted")))
	// Output: <mark>highlighted</mark>
}

func ExampleSub() {
	core.Dump(core.Bundle(text.Text("H"), text.Sub(text.Text("2")), text.Text("O")))
	// Output: H<sub>2</sub>O
}

func ExampleSup() {
	core.Dump(core.Bundle(text.Text("E=mc"), text.Sup(text.Text("2"))))
	// Output: E=mc<sup>2</sup>
}

func ExampleKbd() {
	core.Dump(text.Kbd(text.Text("Ctrl+C")))
	// Output: <kbd>Ctrl+C</kbd>
}

func ExampleAbbr() {
	core.Dump(text.Abbr(attr.Title("HyperText Markup Language"), text.Text("HTML")))
	// Output: <abbr title="HyperText Markup Language">HTML</abbr>
}

func ExampleTime() {
	core.Dump(text.Time(core.Attribute("datetime", "2024-01-15"), text.Text("January 15")))
	// Output: <time datetime="2024-01-15">January 15</time>
}

func ExampleQ() {
	core.Dump(text.Q(text.Text("inline quote")))
	// Output: <q>inline quote</q>
}

func ExampleCite() {
	core.Dump(text.Cite(text.Text("The Art of War")))
	// Output: <cite>The Art of War</cite>
}
