package ct_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/protolambda/chord/core/attr"
	"github.com/protolambda/chord/core/elem"
	"github.com/protolambda/chord/ct"
	"github.com/protolambda/chord/html/aria"
	"github.com/protolambda/chord/html/form"
	"github.com/protolambda/chord/html/form/button"
	"github.com/protolambda/chord/html/form/input"
	"github.com/protolambda/chord/html/form/label"
	"github.com/protolambda/chord/html/group/div"
	"github.com/protolambda/chord/html/text"
)

func TestRules(t *testing.T) {
	cases := map[string]struct {
		rule      ct.Rule
		valid     elem.Node
		invalid   elem.Node
		violation string
	}{
		"unique ids": {
			ct.UniqueIDs(),
			div.Div()(div.Div(attr.ID("a")), div.Div(attr.ID("b")), div.Div(attr.ID(""))),
			div.Div()(div.Div(attr.ID("a")), div.Div(attr.ID("a"))),
			`duplicate id "a" at div[0]/div#a[0] and div[0]/div#a[1]`,
		},
		"label references": {
			ct.LabelReferences(),
			form.Form()(label.Label(label.For("e"))(text.Text("Email")), input.Input(attr.ID("e"))),
			form.Form()(label.Label(label.For("e"))(text.Text("Email")), div.Div(attr.ID("e"))),
			`label at form[0]/label[0] references non-labelable div#e`,
		},
		"label references missing": {
			ct.LabelReferences(),
			form.Form()(label.Label()(text.Text("Email"), input.Input())),
			form.Form()(label.Label(label.For("e"))(text.Text("Email"))),
			`label at form[0]/label[0] references missing id "e"`,
		},
		"aria references": {
			ct.ARIAReferences(),
			div.Div()(div.Div(attr.ID("h")), div.Div(aria.Labelledby("h"), aria.Describedby("h"))),
			div.Div()(div.Div(aria.Describedby("h x"))),
			`aria-describedby at div[0]/div[0] references missing id "x"`,
		},
		"local targets": {
			ct.LocalTargets(),
			div.Div()(div.Div(attr.ID("s")), text.A(text.Href("#s")), text.A(text.Href("#")), text.A(text.Href("#top")), text.A(text.Href("/x#y"))),
			div.Div()(text.A(text.Href("#missing"))),
			`link at div[0]/a[0] targets missing id "missing"`,
		},
		"images have alt": {
			ct.ImagesHaveAlt(),
			div.Div()(elem.Name("img").Void(attr.KV("alt", ""))),
			div.Div()(elem.Name("img").Void(attr.KV("src", "x.png"))),
			`image at div[0]/img[0] has no alt attribute`,
		},
		"buttons have type": {
			ct.ButtonsHaveType(),
			div.Div()(form.Form()(button.Button(button.Type(button.TypeSubmit))), button.Button()),
			form.Form()(div.Div()(button.Button()(text.Text("Go")))),
			`button at form[0]/div[0]/button[0] inside a form has no type`,
		},
		"no inline handlers": {
			ct.NoInlineHandlers(),
			div.Div(attr.KV("data-on", "x"), attr.KV("on", "x")),
			div.Div(attr.KV("onclick", "attack()")),
			`inline handler onclick at div[0]`,
		},
		"no hx on": {
			ct.NoHxOn(),
			div.Div(attr.KV("hx-get", "/x"), attr.KV("data-hx-target", "#t"), attr.KV("hx-online", "x"), attr.KV("onx", "x")),
			div.Div(attr.KV("hx-on:click", "attack()")),
			`htmx handler hx-on:click at div[0]`,
		},
		"no raw": {
			ct.NoRaw(),
			div.Div()(text.Text("<safe>")),
			div.Div()(elem.Raw("<script>")),
			`raw content at div[0]/[0]: raw("<script>")`,
		},
		"no raw doctype": {
			ct.NoRaw(),
			elem.Bundle{elem.Doctype(), elem.Name("html").New()},
			elem.Bundle{elem.Raw("<!DOCTYPE html>"), elem.Name("html").New()},
			`raw content at [0]: raw("<!DOCTYPE html>")`,
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			mustPass(t, ct.View(tc.valid).Valid(tc.rule))
			mustFail(t, ct.View(tc.invalid).Valid(tc.rule), ct.ErrRule, tc.rule.String()+": "+tc.violation)
		})
	}
}

func TestNoHxOnForms(t *testing.T) {
	page := ct.View(div.Div(
		attr.KV("hx-on::after-request", "a()"),
		attr.KV("hx-on--after-request", "b()"),
		attr.KV("HX-On:Click", "c()"),
		attr.KV("data-hx-on:submit", "d()"),
		attr.KV("hx-on", "click: e()"),
	))
	mustFail(t, page.Valid(ct.NoHxOn()), ct.ErrRule,
		"hx-on::after-request at", "hx-on--after-request at", "HX-On:Click at", "data-hx-on:submit at", "htmx handler hx-on at")
	// NoInlineHandlers keeps its meaning: hx-on is not an on* attribute.
	mustPass(t, page.Valid(ct.NoInlineHandlers()))
}

// A described rule formats nodes with the redaction of the subject when it
// runs through Valid, and with the built-in redaction through Violations.
func TestDescribedRule(t *testing.T) {
	view := div.Div()(
		div.Div(attr.Class("codes"), attr.KV("content", "card-4242"), attr.Data("token", "tok-1"))(text.Text("recovery-code")),
	)
	rule := ct.DescribedRule("describeCodes()", func(root ct.Node, d ct.Describer) []string {
		var out []string
		for e := range ct.Elements(root) {
			if ok, _ := ct.Class("codes").Match(e); !ok {
				continue
			}
			content, _ := d.Attr(e, "content")
			token, _ := d.Attr(e, "data-token")
			_, missing := d.Attr(e, "missing")
			out = append(out, fmt.Sprintf("%s at %s: content=%s token=%s text=%s missing=%t",
				d.Node(e), d.Location(e), content, token, d.Text(e), missing))
		}
		return out
	})

	page := ct.View(view, ct.WithRedactContent(ct.Class("codes")))
	err := mustFail(t, page.Valid(rule), ct.ErrRule,
		`describeCodes(): div.codes [content="[redacted]" data-token="[redacted]"] at div[0]/div[0]: content=[redacted] token=[redacted] text=[redacted] missing=false`)
	for _, secret := range []string{"card-4242", "tok-1", "recovery-code"} {
		if strings.Contains(err.Error(), secret) {
			t.Fatalf("secret %q leaked:\n%v", secret, err)
		}
	}

	doc, err := page.Load(t.Context())
	if err != nil {
		t.Fatal(err)
	}
	got := rule.Violations(doc.Root())
	want := `div.codes [content="card-4242" data-token="[redacted]"] at div[0]/div[0]: content=card-4242 token=[redacted] text=recovery-code missing=false`
	if len(got) != 1 || got[0] != want {
		t.Fatalf("Violations:\n got: %q\nwant: %q", got, want)
	}

	d := page.Describer()
	codes, err := page.Find(ct.Class("codes")).Nodes(t.Context())
	if err != nil || len(codes) != 1 {
		t.Fatalf("nodes: %v %v", codes, err)
	}
	if got := d.Text(codes[0]); got != "[redacted]" {
		t.Fatalf("Subject.Describer().Text: got %q", got)
	}
	if got := (ct.Describer{}).Location(doc.Root()); got != "(root)" {
		t.Fatalf("Location of the root: got %q", got)
	}
}

func TestTreeWalking(t *testing.T) {
	page := ct.View(div.Div(attr.ID("a"))(
		text.P()(text.Text("one"), text.Span(attr.ID("s"))(text.Text("two"))),
		text.P()(text.Text("three")),
	))
	doc, err := page.Load(t.Context())
	if err != nil {
		t.Fatal(err)
	}
	var kinds, tags []string
	for n := range ct.Descendants(doc.Root()) {
		kinds = append(kinds, n.Kind().String())
	}
	for e := range ct.Elements(doc.Root()) {
		tags = append(tags, e.Tag())
	}
	if got, want := strings.Join(kinds, " "), "element element text element text element text"; got != want {
		t.Fatalf("Descendants: got %s, want %s", got, want)
	}
	if got, want := strings.Join(tags, " "), "div p span p"; got != want {
		t.Fatalf("Elements: got %s, want %s", got, want)
	}
	// Iterators stop early and restart.
	for range ct.Elements(doc.Root()) {
		break
	}
	n := 0
	for range ct.Elements(doc.Root()) {
		n++
	}
	if n != 4 {
		t.Fatalf("Elements restarted: got %d elements", n)
	}

	span, err := page.Find(ct.ID("s")).Nodes(t.Context())
	if err != nil {
		t.Fatal(err)
	}
	if got := ct.Closest(span[0], func(n ct.Node) bool { return n.Tag() == "span" }); got != span[0] {
		t.Fatalf("Closest includes the node itself: got %v", got)
	}
	if got := ct.Closest(span[0], func(n ct.Node) bool { return n.Tag() == "div" }); got == nil || got.Tag() != "div" {
		t.Fatalf("Closest div: got %v", got)
	}
	if got := ct.Closest(span[0], func(n ct.Node) bool { return n.Tag() == "form" }); got != nil {
		t.Fatalf("Closest form: got %v", got)
	}
	if ct.Closest(nil, func(ct.Node) bool { return true }) != nil {
		t.Fatal("Closest(nil) must be nil")
	}
	for range ct.Descendants(nil) {
		t.Fatal("Descendants(nil) must yield nothing")
	}
	for range ct.Elements(nil) {
		t.Fatal("Elements(nil) must yield nothing")
	}
}

func TestValidCombinesRules(t *testing.T) {
	page := ct.View(div.Div()(div.Div(attr.ID("a")), div.Div(attr.ID("a")), elem.Name("img").Void()))
	a := page.Valid(ct.UniqueIDs(), ct.ImagesHaveAlt(), ct.NoRaw())
	if got, want := a.String(), "valid(chord view, uniqueIDs(), imagesHaveAlt(), noRaw())"; got != want {
		t.Fatalf("String(): got %s, want %s", got, want)
	}
	mustFail(t, a, ct.ErrRule, "uniqueIDs(): duplicate id", "imagesHaveAlt(): image at")

	custom := ct.CustomRule("noDivs()", func(root ct.Node) []string {
		return []string{"divs everywhere"}
	})
	mustFail(t, page.Valid(custom), ct.ErrRule, "noDivs(): divs everywhere")
	mustPass(t, page.Valid())
}
