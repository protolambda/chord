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
	"github.com/protolambda/chord/html/group/list"
	"github.com/protolambda/chord/html/section"
	"github.com/protolambda/chord/html/table"
	"github.com/protolambda/chord/html/text"
)

// semanticPage exercises every supported label and role form.
func semanticPage() elem.Node {
	return section.Main()(
		section.H1()(text.Text("  Account\n  ")),
		section.H2(aria.Label("Details"))(text.Text("ignored")),
		div.Div(aria.Role("heading"), attr.KV("aria-level", "3"))(text.Text("Custom")),
		text.A(text.Href("/edit"))(text.Text("Edit "), text.EM()(text.Text("account"))),
		text.A()(text.Text("placeholder")),
		form.Form()(
			label.Label(label.For("email"))(text.Text("Email address")),
			input.Input(input.Type(input.Email), attr.ID("email")),
			label.Label()(text.Text("Remember me"), input.Input(input.Type(input.Checkbox), attr.ID("remember"))),
			input.Input(input.Type(input.Text), aria.Label("Nickname")),
			text.Span(attr.ID("pw-label"))(text.Text("Password")),
			input.Input(input.Type(input.Password), aria.Labelledby("pw-label")),
			input.Input(input.Type(input.Submit), attr.KV("value", "Save")),
			input.Input(input.Type(input.Reset)),
			button.Button(button.Type(button.TypeButton))(text.Span()(text.Text("Cancel"))),
			elem.Name("select").New(attr.KV("name", "role"))(elem.Name("option").New()(text.Text("Admin"))),
			elem.Name("select").New(attr.Name("multiple").Bool()),
		),
		elem.Name("img").Void(attr.KV("alt", "Logo")),
		elem.Name("img").Void(attr.KV("alt", "")),
		list.UL()(list.LI()(text.Text("Alpha"))),
		table.Table()(
			table.Caption()(text.Text("Users")),
			table.TR()(table.TH()(text.Text("Name")), table.TH(attr.KV("scope", "row"))(text.Text("Bob")), table.TD()(text.Text("Admin"))),
		),
		div.Div(aria.Role("alert"), attr.Title("Saved"))(),
	)
}

func TestLabelQueries(t *testing.T) {
	page := ct.View(semanticPage())

	mustPass(t, page.Find(ct.Label("Email address")).Matches(ct.ID("email")))
	mustPass(t, page.Find(ct.Label("Remember me")).Matches(ct.Attr("type", "checkbox")))
	mustPass(t, page.Find(ct.Label("Nickname")).Matches(ct.Attr("type", "text")))
	mustPass(t, page.Find(ct.Label("Password")).Matches(ct.Attr("type", "password")))
	mustPass(t, page.Find(ct.LabelMatches(ct.Prefix("Email"))).Matches(ct.Tag("input")))
	mustPass(t, page.Find(ct.Label("Email")).None())

	broken := ct.View(div.Div()(input.Input(aria.Labelledby("missing"))))
	mustFail(t, broken.Find(ct.Label("x")).None(), ct.ErrQuery, `missing id "missing"`)
}

func TestRoleQueries(t *testing.T) {
	page := ct.View(semanticPage())

	cases := map[string]struct {
		query ct.Query
		count int
	}{
		"link requires href":           {ct.Role("link"), 1},
		"link named from content":      {ct.Role("link", ct.Named("Edit account")), 1},
		"headings":                     {ct.Role("heading"), 3},
		"heading name collapses":       {ct.Role("heading", ct.Named("Account")), 1},
		"heading aria-label wins":      {ct.Role("heading", ct.Named("Details")), 1},
		"heading level from tag":       {ct.Role("heading", ct.Level(2)), 1},
		"heading level from aria":      {ct.Role("heading", ct.Level(3), ct.Named("Custom")), 1},
		"role is case-insensitive":     {ct.Role("HEADING"), 3},
		"buttons":                      {ct.Role("button"), 3},
		"button value name":            {ct.Role("button", ct.Named("Save")), 1},
		"button default reset name":    {ct.Role("button", ct.Named("Reset")), 1},
		"button content name":          {ct.Role("button", ct.NameMatches(ct.Prefix("Can"))), 1},
		"textboxes exclude password":   {ct.Role("textbox"), 2},
		"textbox named by label":       {ct.Role("textbox", ct.Named("Email address")), 1},
		"checkbox named by wrapper":    {ct.Role("checkbox", ct.Named("Remember me")), 1},
		"combobox":                     {ct.Role("combobox"), 1},
		"listbox when multiple":        {ct.Role("listbox"), 1},
		"option":                       {ct.Role("option", ct.Named("Admin")), 1},
		"img with alt":                 {ct.Role("img", ct.Named("Logo")), 1},
		"img with empty alt":           {ct.Role("presentation"), 1},
		"list and listitem":            {ct.Role("list").Or(ct.Role("listitem")), 2},
		"listitem has no content name": {ct.Role("listitem", ct.Named("Alpha")), 0},
		"table named by caption":       {ct.Role("table", ct.Named("Users")), 1},
		"column header":                {ct.Role("columnheader", ct.Named("Name")), 1},
		"row header":                   {ct.Role("rowheader", ct.Named("Bob")), 1},
		"cell":                         {ct.Role("cell", ct.Named("Admin")), 1},
		"explicit role":                {ct.Role("alert"), 1},
		"title as last resort":         {ct.Role("alert", ct.Named("Saved")), 1},
		"unsupported":                  {ct.Role("main"), 1},
		"no role for plain div":        {ct.Tag("div").And(ct.Role("generic")), 0},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			mustPass(t, page.Find(tc.query).Count(tc.count))
		})
	}

	if got := ct.Role("button", ct.Named("Save"), ct.Level(1)).String(); got != `role("button", named("Save"), level(1))` {
		t.Fatalf("unexpected description: %s", got)
	}
}

func TestAltQueries(t *testing.T) {
	page := ct.View(semanticPage())
	mustPass(t, page.Find(ct.Alt("Logo")).Matches(ct.Tag("img")))
	mustPass(t, page.Find(ct.Alt("")).Count(1))
	mustPass(t, page.Find(ct.AltMatches(ct.Contains("o"))).Count(1))
}
