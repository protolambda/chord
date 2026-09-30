package core_test

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"testing"
	"unicode/utf8"

	"golang.org/x/net/html"

	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/core/elem"
)

// literalContainers are the elements whose content the parser reads as
// text, up to their end tag: rawText ones as written, rcdata ones with
// character references decoded.
var literalContainers = []struct {
	tag    string
	rcdata bool
}{
	{tag: ""}, // no container
	{tag: "noscript"}, {tag: "style"}, {tag: "script"}, {tag: "xmp"}, {tag: "iframe"},
	{tag: "noembed"}, {tag: "noframes"}, {tag: "textarea", rcdata: true}, {tag: "title", rcdata: true},
}

// FuzzLiteralTextRoundTrip checks that text written literally into a script
// or style element, possibly inside an element whose content is text,
// either is refused, or parses back as exactly that text, with the elements
// after it intact.
func FuzzLiteralTextRoundTrip(f *testing.F) {
	for _, seed := range []string{
		"", "a < b && c > d", `"</script>"`, "</SCRIPT ", "<!-- x -->", "<script>", "<!--<script>",
		"</style>", "</scrip", "-->", "<\\/script>", "\x3C!--", "</sty", "</stylex",
		"</noscript><img src=x onerror=alert(1)>", "</textarea>", "</TITLE>", "</xmp", "</iframe>",
		"</noembed>", "</noframes>", "&amp;&lt", "</select>",
	} {
		for i := range literalContainers {
			f.Add(seed, false, uint8(i))
			f.Add(seed, true, uint8(i))
		}
	}
	f.Fuzz(func(t *testing.T, text string, style bool, container uint8) {
		// The HTML parser normalizes these, so they cannot round-trip.
		if !utf8.ValidString(text) || strings.ContainsAny(text, "\r\x00") {
			return
		}
		tag := "script"
		if style {
			tag = "style"
		}
		outer := literalContainers[int(container)%len(literalContainers)]
		if outer.tag == tag {
			return
		}
		var content elem.Node = elem.Name(tag).New()(elem.Text(text))
		if outer.tag != "" {
			content = elem.Name(outer.tag).New()(content)
		}
		page := elem.Name("body").New()(content, elem.Name("p").New()(elem.Text("after")))
		var out strings.Builder
		if err := core.Render(context.Background(), page, &out); err != nil {
			if !errors.Is(err, elem.ErrUnsafeText) {
				t.Fatalf("unexpected error: %v", err)
			}
			return
		}
		want := tag + "(" + quoteText(text) + ")"
		if outer.tag != "" {
			inner := "<" + tag + ">" + text + "</" + tag + ">"
			if outer.rcdata {
				inner = html.UnescapeString(inner)
			}
			want = outer.tag + "(" + quoteText(inner) + ")"
		}
		want += ` p("after")`
		if got := parsedBody(t, out.String(), true); got != want {
			t.Fatalf("round trip of %q:\n got: %s\nwant: %s", out.String(), got, want)
		}
		if outer.tag == "noscript" {
			// Without scripting, the content of noscript is markup.
			want := "noscript(" + tag + "(" + quoteText(text) + `)) p("after")`
			if got := parsedBody(t, out.String(), false); got != want {
				t.Fatalf("round trip of %q without scripting:\n got: %s\nwant: %s", out.String(), got, want)
			}
		}
	})
}

// quoteText returns the outline of a text child: quoted, or nothing when
// empty, since the parser creates no empty text nodes.
func quoteText(s string) string {
	if s == "" {
		return ""
	}
	return strconv.Quote(s)
}

// parsedBody parses a page and outlines the content of its body as
// tag(children) and quoted text.
func parsedBody(t *testing.T, page string, scripting bool) string {
	t.Helper()
	doc, err := html.ParseWithOptions(strings.NewReader(page), html.ParseOptionEnableScripting(scripting))
	if err != nil {
		t.Fatalf("parse %q: %v", page, err)
	}
	for n := range doc.Descendants() {
		if n.Type == html.ElementNode && n.Data == "body" {
			return outlineChildren(n)
		}
	}
	t.Fatalf("no body in %q", page)
	return ""
}

func outlineChildren(n *html.Node) string {
	var parts []string
	for c := range n.ChildNodes() {
		switch c.Type {
		case html.TextNode:
			parts = append(parts, strconv.Quote(c.Data))
		case html.ElementNode:
			parts = append(parts, c.Data+"("+outlineChildren(c)+")")
		default:
			parts = append(parts, fmt.Sprintf("node%d(%q)", c.Type, c.Data))
		}
	}
	return strings.Join(parts, " ")
}

// Literal text in a script or style element inside an element whose content
// is text must not end the outer element either, and after a frameset, or
// for a style inside select, text is escaped.
func TestRenderLiteralTextCannotEndEnclosingTextElement(t *testing.T) {
	n := func(tag string) elem.Scope { return elem.Name(tag).New() }
	const img = "<img src=x onerror=alert(1)>"
	refused := map[string]elem.Node{
		"style in noscript":   n("noscript")(n("style")(elem.Text("</noscript>" + img))),
		"script in style":     n("style")(n("script")(elem.Text("</style>" + img))),
		"script in textarea":  n("textarea")(n("script")(elem.Text("</textarea>" + img))),
		"script in title":     n("title")(n("script")(elem.Text("</TITLE>" + img))),
		"deeper container":    n("iframe")(n("div")(n("style")(elem.Text("</iframe>" + img)))),
		"split over nodes":    n("noscript")(n("style")(elem.Text("</nos"), elem.Text("cript>"+img))),
		"comment in script":   n("script")(n("style")(elem.Text("<!--"))),
		"through a fallback":  n("noscript")(core.Fallback(n("style")(elem.Text("</noscript>"+img)), nil)),
		"both containers end": n("xmp")(n("noembed")(n("style")(elem.Text("</xmp>" + img)))),
	}
	for name, node := range refused {
		t.Run(name, func(t *testing.T) {
			var out strings.Builder
			err := core.Render(context.Background(), n("body")(node), &out)
			if !errors.Is(err, elem.ErrUnsafeText) {
				t.Fatalf("expected ErrUnsafeText, got %v with output %q", err, out.String())
			}
		})
	}
	accepted := map[string]struct {
		node elem.Node
		want string
	}{
		"css in noscript": {
			n("noscript")(n("style")(elem.Text(`a > b { content: "x"; }`))),
			`<noscript><style>a > b { content: "x"; }</style></noscript>`,
		},
		"style in select": {
			n("select")(n("option")(n("style")(elem.Text("</select>" + img)))),
			`<select><option><style>&lt;/select&gt;&lt;img src=x onerror=alert(1)&gt;</style></option></select>`,
		},
		"script in select": {
			n("select")(n("script")(elem.Text("a < b"))),
			`<select><script>a < b</script></select>`,
		},
		"after a frameset": {
			elem.Bundle{n("frameset")(n("script")(elem.Text("<frame>"))), n("script")(elem.Text("<frame>"))},
			`<frameset><script>&lt;frame&gt;</script></frameset><script>&lt;frame&gt;</script>`,
		},
	}
	for name, tc := range accepted {
		t.Run(name, func(t *testing.T) {
			if got := render(t, tc.node); got != tc.want {
				t.Fatalf("rendered output mismatch:\n got: %s\nwant: %s", got, tc.want)
			}
		})
	}
}

func TestRenderKeepsLeadingNewlineOfPreformattedText(t *testing.T) {
	tests := map[string]struct {
		children []elem.Node
		want     string // the text content of the element, as parsed
		markup   bool   // whether the case has child markup (text in a textarea)
	}{
		"newline":            {children: []elem.Node{elem.Text("\nfirst line\n"), elem.Text("\n")}, want: "\nfirst line\n\n"},
		"crlf":               {children: []elem.Node{elem.Text("\r\nfirst")}, want: "\nfirst"},
		"cr":                 {children: []elem.Node{elem.Text("\rfirst")}, want: "\nfirst"},
		"after empty text":   {children: []elem.Node{elem.Text(""), elem.Text("\nfirst")}, want: "\nfirst"},
		"after empty raw":    {children: []elem.Node{elem.Raw(""), elem.Noop(), elem.Text("\r\nfirst")}, want: "\nfirst"},
		"after text":         {children: []elem.Node{elem.Text("a"), elem.Text("\nb")}, want: "a\nb"},
		"after a comment":    {children: []elem.Node{elem.Comment("c"), elem.Text("\nb")}, want: "\nb", markup: true},
		"after an element":   {children: []elem.Node{elem.Name("b").New(), elem.Text("\nb")}, want: "\nb", markup: true},
		"inside an element":  {children: []elem.Node{elem.Name("b").New()(elem.Text("\nb"))}, want: "\nb", markup: true},
		"no leading newline": {children: []elem.Node{elem.Text("x\r\ny")}, want: "x\ny"},
	}
	for _, tag := range []string{"pre", "textarea", "LISTING"} {
		for name, tc := range tests {
			if tc.markup && tag == "textarea" {
				continue
			}
			for _, indent := range []bool{false, true} {
				t.Run(fmt.Sprintf("%s %s indent=%t", tag, name, indent), func(t *testing.T) {
					node := elem.Name("body").New()(elem.Name(tag).New()(tc.children...))
					var opts []core.Option
					if indent {
						opts = append(opts, core.WithIndent())
					}
					var out strings.Builder
					if err := core.Render(context.Background(), node, &out, opts...); err != nil {
						t.Fatalf("render: %v", err)
					}
					doc, err := html.Parse(strings.NewReader(out.String()))
					if err != nil {
						t.Fatalf("parse: %v", err)
					}
					for n := range doc.Descendants() {
						if n.Type == html.ElementNode && strings.EqualFold(n.Data, tag) {
							var got strings.Builder
							for d := range n.Descendants() {
								if d.Type == html.TextNode {
									got.WriteString(d.Data)
								}
							}
							if got.String() != tc.want {
								t.Fatalf("text of %q parses as %q, want %q", out.String(), got.String(), tc.want)
							}
							return
						}
					}
					t.Fatalf("no %s in %q", tag, out.String())
				})
			}
		}
	}
}
