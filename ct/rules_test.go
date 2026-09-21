package ct_test

import (
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
		"no raw": {
			ct.NoRaw(),
			div.Div()(text.Text("<safe>")),
			div.Div()(elem.Raw("<script>")),
			`raw content at div[0]/[0]: raw("<script>")`,
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			mustPass(t, ct.View(tc.valid).Valid(tc.rule))
			mustFail(t, ct.View(tc.invalid).Valid(tc.rule), ct.ErrRule, tc.rule.String()+": "+tc.violation)
		})
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
