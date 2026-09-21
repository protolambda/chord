package ct_test

import (
	"context"
	"errors"
	"regexp"
	"strings"
	"testing"

	"github.com/protolambda/mustbe/assertion"

	"github.com/protolambda/chord/core/attr"
	"github.com/protolambda/chord/core/elem"
	"github.com/protolambda/chord/ct"
	"github.com/protolambda/chord/html/form"
	"github.com/protolambda/chord/html/form/button"
	"github.com/protolambda/chord/html/form/input"
	"github.com/protolambda/chord/html/form/label"
	"github.com/protolambda/chord/html/group/list"
	"github.com/protolambda/chord/html/meta"
	"github.com/protolambda/chord/html/section"
	"github.com/protolambda/chord/html/text"
)

// accountPage is the shared fixture: a page with navigation, a form, and an
// optional admin link.
func accountPage(admin bool) elem.Node {
	return meta.HTML()(
		meta.Head()(meta.Title()(text.Text("Account"))),
		section.Body()(
			section.H1()(text.Text("Account")),
			section.Nav()(
				list.UL()(
					list.LI()(text.A(text.Href("/"))(text.Text("Home"))),
					list.LI()(text.A(text.Href("/account"))(text.Text("Your "), text.EM()(text.Text("account")))),
					elem.If(admin, list.LI()(text.A(text.Href("/admin"))(text.Text("Admin")))),
				),
			),
			form.Form(attr.ID("profile"), form.Method(form.MethodPost), attr.Class("card shadow"))(
				label.Label(label.For("email"))(text.Text("Email")),
				input.Input(input.Type(input.Email), attr.ID("email"), attr.KV("name", "email")),
				input.Input(input.Type(input.Password), attr.KV("name", "password"), attr.KV("value", "hunter2")),
				input.Input(input.Type(input.Hidden), attr.KV("name", "csrf_token"), attr.KV("value", "abc123")),
				button.Button(button.Type(button.TypeSubmit), attr.Data("testid", "save"))(text.Text("Save")),
			),
		),
	)
}

func mustPass(t *testing.T, a assertion.Assertion) {
	t.Helper()
	if err := a.Check(t.Context()); err != nil {
		t.Fatalf("%s: unexpected failure: %v", a, err)
	}
}

func mustFail(t *testing.T, a assertion.Assertion, target error, contains ...string) error {
	t.Helper()
	err := a.Check(t.Context())
	if err == nil {
		t.Fatalf("%s: expected failure", a)
	}
	if !errors.Is(err, target) {
		t.Fatalf("%s: expected %v in chain, got: %v", a, target, err)
	}
	for _, c := range contains {
		if !strings.Contains(err.Error(), c) {
			t.Fatalf("%s: expected %q in error:\n%v", a, c, err)
		}
	}
	return err
}

func TestSubjectLoadsOnce(t *testing.T) {
	calls := 0
	page := ct.View(elem.Fn(func(context.Context) (elem.Node, error) {
		calls++
		return accountPage(true), nil
	}))

	mustPass(t, page)
	mustPass(t, page.Find(ct.Tag("h1")))
	mustPass(t, page.Find(ct.Tag("a"), ct.Attr("href", "/admin")))
	mustPass(t, page.Find(ct.Tag("table")).None())
	if _, err := page.Load(t.Context()); err != nil {
		t.Fatalf("load: %v", err)
	}
	if calls != 1 {
		t.Fatalf("view evaluated %d times, want 1", calls)
	}
}

func TestFailedLoadStaysFailedAndIsNeverAbsence(t *testing.T) {
	errDB := errors.New("database unavailable")
	page := ct.View(section.Body()(elem.Fn(func(context.Context) (elem.Node, error) {
		return nil, errDB
	})))

	err := mustFail(t, page, ct.ErrLoad, "chord view", "at body[0]/[0]")
	if !errors.Is(err, errDB) {
		t.Fatalf("expected cause in chain, got %v", err)
	}
	// Absence must not be concluded from a failed load.
	mustFail(t, page.Find(ct.Tag("a")).None(), ct.ErrLoad)
	mustFail(t, page.Find(ct.Tag("a")).Any(), ct.ErrLoad)
	mustFail(t, page.Find(ct.Tag("a")).Matches(ct.Tag("a")), ct.ErrLoad)
}

func TestSelectionCardinalities(t *testing.T) {
	page := ct.View(accountPage(false))
	links := page.Find(ct.Tag("a"))    // 2
	admin := page.Find(ct.ID("admin")) // 0
	form := page.Find(ct.Tag("form"))  // 1

	pass := []assertion.Assertion{
		form, links.Any(), links.Count(2), links.AtLeast(1), links.AtLeast(2), links.AtMost(2),
		admin.None(), admin.Count(0), admin.AtMost(0),
	}
	for _, a := range pass {
		mustPass(t, a)
	}

	fail := []struct {
		a    assertion.Assertion
		want string
	}{
		{links, "expected exactly one match of tag(\"a\") in chord view, found 2"},
		{admin, "expected exactly one match of id(\"admin\") in chord view, found 0"},
		{links.None(), "expected no match"},
		{admin.Any(), "expected at least one match"},
		{links.Count(3), "expected exactly 3 match"},
		{links.AtLeast(3), "expected at least 3 match"},
		{links.AtMost(1), "expected at most 1 match"},
	}
	for _, tc := range fail {
		mustFail(t, tc.a, ct.ErrCount, tc.want)
	}
}

func TestNestedScopes(t *testing.T) {
	page := ct.View(accountPage(true))

	form := page.Find(ct.Tag("form"), ct.ID("profile"))
	mustPass(t, form)
	mustPass(t, form.Find(ct.Tag("label")))
	mustPass(t, form.Find(ct.Tag("a")).None())

	// Descendants of multiple scope nodes are combined in document order.
	items := page.Find(ct.Tag("li"))
	mustPass(t, items.Count(3))
	mustPass(t, items.Find(ct.Tag("a")).Count(3))
	mustPass(t, items.Find(ct.Tag("a")).Texts("Home", "Your account", "Admin"))

	// Nested scope nodes do not produce duplicate matches.
	mustPass(t, page.Find(ct.Tag("body").Or(ct.Tag("nav"))).Find(ct.Tag("a")).Count(3))

	// A missing scope is reported as such.
	mustFail(t, page.Find(ct.Tag("table")).Find(ct.Tag("td")), ct.ErrCount,
		"found 0", `scope tag("table") matched 0`)
}

func TestMatches(t *testing.T) {
	page := ct.View(accountPage(false))
	email := page.Find(ct.ID("email"))

	mustPass(t, email.Matches(ct.Tag("input"), ct.Attr("type", "email"), ct.Attr("name", "email")))
	mustPass(t, email.Matches())
	mustFail(t, email.Matches(ct.Attr("type", "text")), ct.ErrMismatch,
		`expected input#email [type="email" name="email"] to match attr("type", "text")`)
	mustFail(t, page.Find(ct.Tag("input")).Matches(ct.Tag("input")), ct.ErrCount, "found 3")
}

func TestTexts(t *testing.T) {
	page := ct.View(accountPage(true))
	mustPass(t, page.Find(ct.Tag("li")).Texts("Home", "Your account", "Admin"))
	mustPass(t, page.Find(ct.Tag("table")).Texts())
	mustFail(t, page.Find(ct.Tag("li")).Texts("Home", "Account"), ct.ErrMismatch,
		`to be ["Home" "Account"], got ["Home" "Your account" "Admin"]`)
}

func TestPicks(t *testing.T) {
	page := ct.View(accountPage(true))
	links := page.Find(ct.Tag("a"))

	mustPass(t, links.First().Matches(ct.Attr("href", "/")))
	mustPass(t, links.Last().Matches(ct.Attr("href", "/admin")))
	mustPass(t, links.Nth(1).Matches(ct.Attr("href", "/account")))
	mustPass(t, links.Nth(-2).Matches(ct.Attr("href", "/account")))
	mustPass(t, links.Nth(3).None())
	mustPass(t, links.Nth(-4).None())
	mustPass(t, page.Find(ct.Tag("li")).Last().Find(ct.Tag("a")).Texts("Admin"))
	mustFail(t, links.Nth(3), ct.ErrCount, `expected exactly one match of nth(3) of tag("a") in chord view, found 0`)
	if got, want := links.First().String(), `exactlyOne(first of tag("a") in chord view)`; got != want {
		t.Fatalf("String(): got %s, want %s", got, want)
	}
}

func TestEachAndInOrder(t *testing.T) {
	page := ct.View(accountPage(true))
	links := page.Find(ct.Tag("a"))

	mustPass(t, links.Each(ct.HasAttr("href"), ct.AttrMatches("href", ct.Prefix("/"))))
	mustFail(t, links.Each(ct.Attr("href", "/")), ct.ErrMismatch,
		`expected every match of tag("a") in chord view to match attr("href", "/"), match 1 does not`)
	mustFail(t, page.Find(ct.Tag("table")).Each(ct.Tag("table")), ct.ErrCount, "expected at least one match")

	mustPass(t, links.InOrder(ct.Text("Home"), ct.TextContent("Your account"), ct.Attr("href", "/admin")))
	mustFail(t, links.InOrder(ct.Text("Home"), ct.Text("Admin"), ct.Attr("href", "/admin")), ct.ErrMismatch,
		`expected match 1 of tag("a") in chord view to match text("Admin")`)
	mustFail(t, links.InOrder(ct.Text("Home")), ct.ErrCount, "expected exactly 1 match", "found 3")

	mustPass(t, links.AttrValues("href", "/", "/account", "/admin"))
	mustFail(t, links.AttrValues("href", "/", "/account"), ct.ErrMismatch,
		`expected href values of tag("a") in chord view to be ["/" "/account"], got ["/" "/account" "/admin"]`)
	mustFail(t, links.AttrValues("id", "", "", ""), ct.ErrMismatch, `got ["(absent)" "(absent)" "(absent)"]`)
	mustPass(t, page.Find(ct.Tag("input")).AttrValues("name", "email", "password", "csrf_token"))

	if got, want := links.Each(ct.HasAttr("href")).String(), `each(tag("a") in chord view, hasAttr("href"))`; got != want {
		t.Fatalf("String(): got %s, want %s", got, want)
	}
	if got, want := links.AttrValues("href", "/").String(), `attrValues("href")(tag("a") in chord view, ["/"])`; got != want {
		t.Fatalf("String(): got %s, want %s", got, want)
	}
}

func TestQueries(t *testing.T) {
	page := ct.View(accountPage(true))
	errRef := errors.New("broken reference")
	broken := ct.Custom("broken()", func(ct.Node) (bool, error) { return false, errRef })

	cases := map[string]struct {
		query ct.Query
		count int
	}{
		"tag is case-insensitive":  {ct.Tag("FORM"), 1},
		"id":                       {ct.ID("profile"), 1},
		"class token":              {ct.Class("shadow"), 1},
		"class partial token":      {ct.Class("shad"), 0},
		"attr exact":               {ct.Attr("href", "/account"), 1},
		"attr matches":             {ct.AttrMatches("href", ct.Prefix("/a")), 2},
		"attr on raw enum value":   {ct.Attr("method", "post"), 1},
		"has attr":                 {ct.HasAttr("for"), 1},
		"test id":                  {ct.TestID("save"), 1},
		"own text":                 {ct.Text("Account"), 2}, // title and h1
		"own text excludes nested": {ct.Text("Your account"), 0},
		"own text of wrapper":      {ct.Text("Your"), 1},
		"text matches":             {ct.TextMatches(ct.Contains("count")), 3}, // title, h1, em
		"text content":             {ct.TextContent("Your account"), 2},
		"text content matches":     {ct.TextContentMatches(ct.Regexp(regexp.MustCompile(`^Your\s`))), 2},
		"has child":                {ct.HasChild(ct.Tag("a")), 3},
		"has descendant":           {ct.HasDescendant(ct.Tag("a")), 7},
		"and":                      {ct.Tag("a").And(ct.Attr("href", "/")), 1},
		"or":                       {ct.Tag("h1").Or(ct.Tag("label"), ct.Tag("nope")), 2},
		"not":                      {ct.Tag("input").And(ct.Attr("type", "hidden").Not()), 2},
		"and short-circuits":       {ct.Tag("nope").And(broken), 0},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			mustPass(t, page.Find(tc.query).Count(tc.count))
		})
	}

	t.Run("query errors propagate", func(t *testing.T) {
		for _, q := range []ct.Query{broken, broken.Not(), ct.Tag("h1").And(broken), ct.HasChild(broken), ct.HasDescendant(broken)} {
			err := mustFail(t, page.Find(q).None(), ct.ErrQuery, "broken()")
			if !errors.Is(err, errRef) {
				t.Fatalf("expected cause in chain, got %v", err)
			}
		}
	})

	t.Run("zero query is invalid", func(t *testing.T) {
		mustFail(t, page.Find(ct.Query{}).Any(), ct.ErrQuery, "zero query")
	})
}

func TestValueMatchers(t *testing.T) {
	cases := map[string]struct {
		m    ct.ValueMatch
		yes  string
		no   string
		desc string
	}{
		"exact":    {ct.Exact("a"), "a", "ab", `"a"`},
		"contains": {ct.Contains("b"), "abc", "ac", `contains("b")`},
		"prefix":   {ct.Prefix("ab"), "abc", "bab", `prefix("ab")`},
		"suffix":   {ct.Suffix("bc"), "abc", "bca", `suffix("bc")`},
		"regexp":   {ct.Regexp(regexp.MustCompile(`^\d+$`)), "123", "12a", `regexp("^\\d+$")`},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if !tc.m.Match(tc.yes) || tc.m.Match(tc.no) {
				t.Fatalf("matcher %s: yes=%t no=%t", tc.m, tc.m.Match(tc.yes), tc.m.Match(tc.no))
			}
			if tc.m.String() != tc.desc {
				t.Fatalf("desc: got %s, want %s", tc.m, tc.desc)
			}
		})
	}
	if (ct.ValueMatch{}).Match("") {
		t.Fatal("zero matcher must not match")
	}
}

func TestStrings(t *testing.T) {
	page := ct.View(accountPage(false))
	form := page.Find(ct.Tag("form"), ct.ID("profile"))
	email := form.Find(ct.Tag("input"))

	cases := []struct {
		a    assertion.Assertion
		want string
	}{
		{page, `loads(chord view)`},
		{form, `exactlyOne(and(tag("form"), id("profile")) in chord view)`},
		{email, `exactlyOne(tag("input") within and(tag("form"), id("profile")) in chord view)`},
		{email.None(), `none(tag("input") within and(tag("form"), id("profile")) in chord view)`},
		{email.Count(3), `count(3)(tag("input") within and(tag("form"), id("profile")) in chord view)`},
		{email.Matches(ct.ID("email")), `matches(tag("input") within and(tag("form"), id("profile")) in chord view, id("email"))`},
		{email.Texts("a"), `texts(tag("input") within and(tag("form"), id("profile")) in chord view, ["a"])`},
	}
	for _, tc := range cases {
		if got := tc.a.String(); got != tc.want {
			t.Errorf("String():\n got: %s\nwant: %s", got, tc.want)
		}
	}
}

func TestDiagnosticsRedactSecrets(t *testing.T) {
	page := ct.View(accountPage(false), ct.WithRedact(func(tag, key string) bool {
		return tag == "a" && key == "href"
	}))

	err := mustFail(t, page.Find(ct.Tag("input")), ct.ErrCount, "found 3", "[redacted]",
		`input [type="password" name="password" value="[redacted]"]`,
		`input [type="hidden" name="csrf_token" value="[redacted]"]`,
		`html[0]/body[1]/form#profile[2]/input[2]`)
	if strings.Contains(err.Error(), "hunter2") || strings.Contains(err.Error(), "abc123") {
		t.Fatalf("secret leaked in diagnostics:\n%v", err)
	}

	err = mustFail(t, page.Find(ct.Tag("nav")).Find(ct.Tag("table")), ct.ErrCount, "scope outline:", `a [href="[redacted]"]`)
	if strings.Contains(err.Error(), "/account") {
		t.Fatalf("custom redaction ignored:\n%v", err)
	}
}

func TestDiagnosticsOutlineScope(t *testing.T) {
	page := ct.View(accountPage(false))
	mustFail(t, page.Find(ct.Tag("form")).Find(ct.Tag("select")), ct.ErrCount,
		"scope outline:",
		"\n  form#profile.card.shadow [method=\"post\"]\n",
		"\n    label [for=\"email\"]\n",
		"\n      \"Email\"\n",
		"\n    button [type=\"submit\" data-testid=\"save\"]\n")
}
