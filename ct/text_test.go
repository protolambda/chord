package ct_test

import (
	"testing"

	"github.com/protolambda/chord/core/attr"
	"github.com/protolambda/chord/core/elem"
	"github.com/protolambda/chord/ct"
	"github.com/protolambda/chord/html/group/div"
	"github.com/protolambda/chord/html/section"
	"github.com/protolambda/chord/html/table"
	"github.com/protolambda/chord/html/text"
)

// card is a Bootstrap card: block-level parts, and a flex footer of links.
func card() elem.Node {
	return div.Div(attr.Class("card"))(
		div.Div(attr.Class("card-header"))(section.H3(attr.Class("h6"))(text.Text("Author"))),
		div.Div(attr.Class("card-body"))(text.Text("Signed with "), text.Code()(text.Text("dev0"))),
		div.Div(attr.Class("card-footer d-flex gap-2"))(
			text.A(text.Href("/a"), attr.Class("btn"))(text.Text("One")),
			text.A(text.Href("/b"), attr.Class("btn"))(text.Text("Two")),
		),
	)
}

func TestTextsSeparateBlocks(t *testing.T) {
	page := ct.View(card())
	mustPass(t, page.Find(ct.Class("card")).Texts("Author Signed with dev0 One Two"))
	mustPass(t, page.Find(ct.InnerText("Author Signed with dev0 One Two")))
	mustPass(t, page.Find(ct.InnerTextMatches(ct.Prefix("Author Signed"))).Count(1))
	// TextContent is the raw DOM textContent: the text nodes, joined.
	mustPass(t, page.Find(ct.TextContent("AuthorSigned with dev0OneTwo")))
	mustPass(t, page.Find(ct.TextContentMatches(ct.Prefix("AuthorSigned"))).Count(1))
}

func TestInnerTextLayout(t *testing.T) {
	cases := map[string]struct {
		node elem.Node
		want string
	}{
		"inline elements join": {
			text.P()(text.Text("a"), text.B()(text.Text("b")), text.Text("c"), text.Span()(text.Text("5")), text.Small()(text.Text("ETH"))),
			"abc5ETH",
		},
		"block elements separate": {
			div.Div()(text.Text("a"), div.Div()(text.Text("b")), text.Text("c"), text.P()(text.Text("d"))),
			"a b c d",
		},
		"br separates": {
			text.P()(text.Text("one"), text.BR(), text.Text("two")),
			"one two",
		},
		"table cells separate": {
			table.Table()(table.TR()(table.TD()(text.Text("a")), table.TD()(text.Text("b"))), table.TR()(table.TD()(text.Text("c")))),
			"a b c",
		},
		"list items separate": {
			elem.Name("ul").New()(elem.Name("li").New()(text.Text("x")), elem.Name("li").New()(text.Text("y"))),
			"x y",
		},
		"style display block": {
			div.Div()(text.Span(attr.Style("color: red; DISPLAY: Block !important"))(text.Text("a")), text.Text("b")),
			"a b",
		},
		"style display inline": {
			div.Div()(div.Div(attr.Style("display:inline"))(text.Text("a")), div.Div(attr.Style("display:inline-block"))(text.Text("b"))),
			"ab",
		},
		"style flex container": {
			div.Div(attr.Style("display: flex"))(text.Span()(text.Text("a")), text.Span()(text.Text("b"))),
			"a b",
		},
		"style inline-grid container": {
			// The items are blocks, also in an inline-level container.
			div.Div()(text.Text("x"), text.Span(attr.Style("display: inline-grid"))(text.Span()(text.Text("a")), text.Span()(text.Text("b"))), text.Text("y")),
			"x a b y",
		},
		"style two-value display": {
			div.Div()(text.Span(attr.Style("display: block flex"))(text.Span()(text.Text("a")), text.Span()(text.Text("b"))), text.Text("c")),
			"a b c",
		},
		"unknown style display keeps the default": {
			div.Div()(text.Span(attr.Style("display: var(--d)"))(text.Text("a")), text.Text("b")),
			"ab",
		},
		"flex item text runs": {
			div.Div(attr.Class("d-flex"))(text.Text("Hello"), text.Span()(text.Text("x"))),
			"Hello x",
		},
		"bootstrap display utilities": {
			div.Div()(div.Div(attr.Class("d-inline"))(text.Text("a")), div.Div(attr.Class("d-inline-block"))(text.Text("b")), text.Span(attr.Class("d-block"))(text.Text("c"))),
			"ab c",
		},
		"bootstrap utilities win over inline style": {
			div.Div()(text.Span(attr.Class("d-block"), attr.Style("display:inline"))(text.Text("a")), text.Text("b")),
			"a b",
		},
		"bootstrap responsive utilities are ignored": {
			div.Div()(text.Span(attr.Class("d-md-block"))(text.Text("a")), text.Text("b")),
			"ab",
		},
		"bootstrap flex components": {
			div.Div()(
				div.Div(attr.Class("btn-group"))(text.A(attr.Class("btn"))(text.Text("a")), text.A(attr.Class("btn"))(text.Text("b"))),
				div.Div(attr.Class("hstack"))(text.Span()(text.Text("c")), text.Span()(text.Text("d"))),
			),
			"a b c d",
		},
		"bootstrap block components": {
			div.Div()(text.A(attr.Class("dropdown-item"))(text.Text("a")), text.A(attr.Class("dropdown-item"))(text.Text("b"))),
			"a b",
		},
		"flex items are blockified": {
			div.Div(attr.Class("d-flex"))(text.Span(attr.Class("d-inline"))(text.Text("a")), text.Span()(text.Text("b"))),
			"a b",
		},
		"inline flex container is inline": {
			div.Div()(text.Text("x"), text.Span(attr.Class("d-inline-flex"))(text.Text("a")), text.Text("y")),
			"xay",
		},
		"inline flex container items are blocks": {
			div.Div()(text.Text("x"), text.Span(attr.Class("d-inline-flex"))(text.Span()(text.Text("a")), text.Span()(text.Text("b"))), text.Text("y")),
			"x a b y",
		},
		"raw and comments are not text": {
			div.Div()(text.Text("a"), elem.Raw("<b>raw</b>"), elem.Comment("c"), text.Text("b")),
			"ab",
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			page := ct.View(section.Main()(tc.node))
			mustPass(t, page.Find(ct.Tag("main")).Texts(tc.want))
			mustPass(t, page.Find(ct.Tag("main"), ct.InnerText(tc.want)))
		})
	}
}

func TestOwnTextSeparatesBlocks(t *testing.T) {
	page := ct.View(div.Div(attr.ID("x"))(
		text.Text("a"), div.Div()(text.Text("b")), text.Text("c"), text.BR(), text.Text("d"), text.EM()(text.Text("e")), text.Text("f"),
	))
	// The text of the inline em is not own text, and does not separate.
	mustPass(t, page.Find(ct.ID("x"), ct.Text("a c df")))
	mustPass(t, page.Find(ct.ID("x"), ct.TextMatches(ct.Prefix("a c"))))
}

func TestNamesSeparateBlocks(t *testing.T) {
	page := ct.View(section.Main()(
		text.A(text.Href("/a"))(div.Div()(text.Text("Author")), div.Div()(text.Text("Signed"))),
		text.A(text.Href("/b"))(text.Text("Your "), text.EM()(text.Text("acc")), text.Text("ount")),
	))
	mustPass(t, page.Find(ct.Role("link", ct.Named("Author Signed"))))
	mustPass(t, page.Find(ct.Role("link", ct.Named("Your account"))))
}

func TestInnerTextLeavesOutUnrenderedContent(t *testing.T) {
	page := ct.View(elem.Name("html").New()(
		elem.Name("head").New()(elem.Name("title").New()(text.Text("Account")), elem.Name("style").New()(text.Text("p{}"))),
		section.Body()(
			text.P()(text.Text("Hello")),
			elem.Name("script").New()(text.Text("var x = 1;")),
			elem.Name("template").New()(text.P()(text.Text("row"))),
			elem.Name("noscript").New()(text.Text("Enable JavaScript")),
			text.P()(text.Text("World")),
		),
	))
	mustPass(t, page.Find(ct.Tag("html")).Texts("Hello World"))
	mustPass(t, page.Find(ct.Tag("body"), ct.InnerText("Hello World")))
	// The text of such an element itself is its text content.
	mustPass(t, page.Find(ct.Tag("head")).Texts("Accountp{}"))
	mustPass(t, page.Find(ct.Tag("title")).Texts("Account"))
	mustPass(t, page.Find(ct.Tag("script")).Texts("var x = 1;"))
	mustPass(t, page.Find(ct.Tag("template")).Find(ct.Tag("p")).Texts("row"))
	// TextContent keeps everything.
	mustPass(t, page.Find(ct.Tag("body"), ct.TextContent("Hellovar x = 1;rowEnable JavaScriptWorld")))
}
