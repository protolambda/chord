package cthtml_test

import (
	"context"
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/protolambda/mustbe/assertion"

	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/core/attr"
	"github.com/protolambda/chord/core/elem"
	"github.com/protolambda/chord/ct"
	"github.com/protolambda/chord/ct/cthtml"
	"github.com/protolambda/chord/html/form"
	"github.com/protolambda/chord/html/form/button"
	"github.com/protolambda/chord/html/form/input"
	"github.com/protolambda/chord/html/form/label"
	"github.com/protolambda/chord/html/group/list"
	"github.com/protolambda/chord/html/meta"
	"github.com/protolambda/chord/html/section"
	"github.com/protolambda/chord/html/table"
	"github.com/protolambda/chord/html/text"
)

func mustPass(t *testing.T, a assertion.Assertion) {
	t.Helper()
	if err := a.Check(t.Context()); err != nil {
		t.Fatalf("%s: unexpected failure: %v", a, err)
	}
}

func mustFail(t *testing.T, a assertion.Assertion, target error, contains ...string) {
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
}

func TestPageParsesDocument(t *testing.T) {
	page := cthtml.PageString(`<!DOCTYPE html><title>Hi &amp; bye</title><p class="a  b" hidden>one<br>two<!-- c --></p>`)

	mustPass(t, page)
	// The parser inserts html, head, and body, and decodes entities.
	mustPass(t, page.Find(ct.Tag("html")).Find(ct.Tag("head")).Find(ct.Tag("title"), ct.Text("Hi & bye")))
	mustPass(t, page.Find(ct.Tag("body")).Find(ct.Tag("p"), ct.Class("b"), ct.HasAttr("hidden")))
	mustPass(t, page.Find(ct.Tag("p")).Texts("onetwo"))
	mustPass(t, page.Find(ct.Tag("p"), ct.Text("onetwo")))
	mustFail(t, page.Find(ct.Tag("p")).Find(ct.Tag("span")), ct.ErrCount,
		"scope outline:", `p.a.b [hidden]`, `"one"`, `br`, `<!-- c -->`)
	mustFail(t, page.Find(ct.Tag("p").Or(ct.Tag("br"))), ct.ErrCount,
		"found 2", "html[1]/body[1]/p[0] p.a.b [hidden]", "html[1]/body[1]/p[0]/br[1] br")

	doc, err := page.Load(t.Context())
	if err != nil {
		t.Fatalf("load: %v", err)
	}
	var kinds []ct.NodeKind
	for c := range doc.Root().Children() {
		kinds = append(kinds, c.Kind())
	}
	if len(kinds) != 2 || kinds[0] != ct.KindDoctype || kinds[1] != ct.KindElement {
		t.Fatalf("unexpected root children kinds: %v", kinds)
	}
	if _, ok := doc.(cthtml.Document); !ok {
		t.Fatal("expected access to the parsed html root")
	}
}

func TestFragmentUsesContext(t *testing.T) {
	rows := cthtml.FragmentString("tbody", `<tr><td>a</td><td>b</td></tr><tr><td>c</td></tr>`)
	mustPass(t, rows.Find(ct.Tag("tr")).Count(2))
	mustPass(t, rows.Find(ct.Tag("html")).None())
	mustPass(t, rows.Find(ct.Tag("td")).Texts("a", "b", "c"))
	mustFail(t, rows.Find(ct.Tag("th")), ct.ErrCount, "html fragment in tbody")

	// Without a suitable context, table content is dropped by the parser.
	dropped := cthtml.FragmentString("div", `<tr><td>a</td></tr>`)
	mustPass(t, dropped.Find(ct.Tag("td")).None())
}

type countingReader struct {
	io.Reader
	reads int
}

func (c *countingReader) Read(p []byte) (int, error) {
	c.reads++
	return c.Reader.Read(p)
}

func TestReaderIsConsumedOnceOnFirstCheck(t *testing.T) {
	r := &countingReader{Reader: strings.NewReader("<p>x</p>")}
	page := cthtml.Page(r)
	if r.reads != 0 {
		t.Fatal("reader consumed before the first check")
	}
	mustPass(t, page.Find(ct.Tag("p")))
	reads := r.reads
	mustPass(t, page.Find(ct.Tag("p")).Any())
	mustPass(t, page)
	if r.reads != reads || reads == 0 {
		t.Fatalf("reader read %d times before and %d after", reads, r.reads)
	}
}

func TestBytesAreCopied(t *testing.T) {
	data := []byte("<p>before</p>")
	page := cthtml.PageBytes(data)
	copy(data, "<p>after!</p>")
	mustPass(t, page.Find(ct.Tag("p"), ct.Text("before")))
}

func TestFailedReadIsALoadFailure(t *testing.T) {
	errRead := errors.New("connection reset")
	page := cthtml.Page(failingReader{err: errRead})
	err := page.Check(context.Background())
	if !errors.Is(err, ct.ErrLoad) || !errors.Is(err, errRead) {
		t.Fatalf("expected ErrLoad wrapping the read error, got %v", err)
	}
	mustFail(t, page.Find(ct.Tag("p")).None(), ct.ErrLoad)
}

type failingReader struct{ err error }

func (f failingReader) Read([]byte) (int, error) { return 0, f.err }

// accountPage is rendered and parsed, so that the same queries can be
// checked against a direct view and the parsed output.
func accountPage() elem.Node {
	return meta.HTML()(
		meta.Head()(meta.Title()(text.Text("Account & more"))),
		section.Body()(
			section.H1()(text.Text("Account")),
			section.Nav()(list.UL()(
				list.LI()(text.A(text.Href("/"))(text.Text("Home"))),
				list.LI()(text.A(text.Href("/account"))(text.Text("Your "), text.EM()(text.Text("account")))),
			)),
			form.Form(attr.ID("profile"), form.Method(form.MethodPost), attr.Class("card shadow"))(
				label.Label(label.For("email"))(text.Text("Email address")),
				input.Input(input.Type(input.Email), attr.ID("email"), attr.KV("name", "email"), attr.KV("value", `a "quoted" <value>`)),
				label.Label()(text.Text("Remember me"), input.Input(input.Type(input.Checkbox), attr.Name("checked").Bool())),
				button.Button(button.Type(button.TypeSubmit))(text.Text("Save")),
			),
			table.Table()(table.Caption()(text.Text("Users")), table.TR()(table.TH()(text.Text("Name")), table.TD()(text.Text("Bob")))),
			elem.Comment("footer"),
		),
	)
}

func TestParsedAndDirectViewsAgree(t *testing.T) {
	var out strings.Builder
	if err := core.Render(context.Background(), accountPage(), &out, core.WithIndent()); err != nil {
		t.Fatalf("render: %v", err)
	}
	subjects := map[string]*ct.Subject{
		"direct": ct.View(accountPage()),
		"parsed": cthtml.PageString(out.String()),
	}

	checks := map[string]func(page *ct.Subject) assertion.Assertion{
		"title text": func(p *ct.Subject) assertion.Assertion { return p.Find(ct.Tag("title"), ct.Text("Account & more")) },
		"heading": func(p *ct.Subject) assertion.Assertion {
			return p.Find(ct.Role("heading", ct.Named("Account"), ct.Level(1)))
		},
		"links":     func(p *ct.Subject) assertion.Assertion { return p.Find(ct.Role("link")).Texts("Home", "Your account") },
		"link name": func(p *ct.Subject) assertion.Assertion { return p.Find(ct.Role("link", ct.Named("Your account"))) },
		"form": func(p *ct.Subject) assertion.Assertion {
			return p.Find(ct.Tag("form"), ct.ID("profile"), ct.Class("shadow"), ct.Attr("method", "post"))
		},
		"escaped value": func(p *ct.Subject) assertion.Assertion {
			return p.Find(ct.ID("email")).Matches(ct.Attr("value", `a "quoted" <value>`))
		},
		"label for": func(p *ct.Subject) assertion.Assertion {
			return p.Find(ct.Label("Email address")).Matches(ct.Tag("input"))
		},
		"label wrapper": func(p *ct.Subject) assertion.Assertion {
			return p.Find(ct.Label("Remember me")).Matches(ct.HasAttr("checked"), ct.Attr("checked", ""))
		},
		"button": func(p *ct.Subject) assertion.Assertion { return p.Find(ct.Role("button", ct.Named("Save"))) },
		"table": func(p *ct.Subject) assertion.Assertion {
			return p.Find(ct.Role("table", ct.Named("Users"))).Find(ct.Role("cell")).Texts("Bob")
		},
		"has child": func(p *ct.Subject) assertion.Assertion {
			return p.Find(ct.Tag("li"), ct.HasChild(ct.Tag("a"))).Count(2)
		},
		"no admin": func(p *ct.Subject) assertion.Assertion { return p.Find(ct.Role("link", ct.Named("Admin"))).None() },
		"rules": func(p *ct.Subject) assertion.Assertion {
			return p.Valid(ct.UniqueIDs(), ct.LabelReferences(), ct.ButtonsHaveType(), ct.NoRaw())
		},
		"comment excluded": func(p *ct.Subject) assertion.Assertion {
			return p.Find(ct.Tag("body")).Find(ct.TextContentMatches(ct.Contains("footer"))).None()
		},
	}
	for name, check := range checks {
		for source, page := range subjects {
			t.Run(name+"/"+source, func(t *testing.T) {
				mustPass(t, check(page))
			})
		}
	}
}
