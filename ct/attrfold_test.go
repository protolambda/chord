package ct_test

import (
	"testing"

	"github.com/protolambda/chord/core/attr"
	"github.com/protolambda/chord/ct"
	"github.com/protolambda/chord/ct/cthtml"
	"github.com/protolambda/chord/html/group/div"
)

// loadRoot loads s and returns its root.
func loadRoot(t *testing.T, s *ct.Subject) ct.Node {
	t.Helper()
	doc, err := s.Load(t.Context())
	if err != nil {
		t.Fatal(err)
	}
	return doc.Root()
}

// firstElement returns the first element below root with the tag.
func firstElement(t *testing.T, root ct.Node, tag string) ct.Node {
	t.Helper()
	for e := range ct.Elements(root) {
		if e.Tag() == tag {
			return e
		}
	}
	t.Fatalf("no %s element", tag)
	return nil
}

type foldCase struct {
	name  string
	value string
	ok    bool
}

func checkFold(t *testing.T, n ct.Node, cases []foldCase) {
	t.Helper()
	for _, c := range cases {
		v, ok := n.AttrFold(c.name)
		if v != c.value || ok != c.ok {
			t.Errorf("AttrFold(%q) = %q, %v; want %q, %v", c.name, v, ok, c.value, c.ok)
		}
	}
}

func TestAttrFoldView(t *testing.T) {
	root := loadRoot(t, ct.View(div.Div(
		attr.KV("Data-Mode", "dark"),
		attr.KV("HX-Get", "/tasks"),
		attr.KV("title", "a\r\nb"),
		attr.Bool("Hidden"),
		attr.KV("ID", "upper"),
		attr.KV("id", "lower"),
	)()))
	e := firstElement(t, root, "div")
	checkFold(t, e, []foldCase{
		{"data-mode", "dark", true},
		{"DATA-MODE", "dark", true},
		{"hx-get", "/tasks", true},
		{"title", "a\nb", true}, // the value as Attr gives it
		{"TITLE", "a\nb", true},
		{"hidden", "", true},
		{"id", "lower", true}, // an exact match wins
		{"ID", "upper", true},
		{"Id", "upper", true}, // else the first in source order
		{"hx-post", "", false},
		{"data-mod", "", false},
	})
}

func TestAttrFoldParsed(t *testing.T) {
	root := loadRoot(t, cthtml.PageString(`<!doctype html><html><body>
<div DATA-Mode="dark" hx-get="/tasks"></div>
<svg viewBox="0 0 16 16"><path d="M0 0"/></svg>
</body></html>`))
	d := firstElement(t, root, "div")
	checkFold(t, d, []foldCase{
		{"data-mode", "dark", true}, // the parser lower-cases names
		{"Data-Mode", "dark", true},
		{"HX-GET", "/tasks", true},
		{"hx-post", "", false},
	})
	svg := firstElement(t, root, "svg")
	if _, ok := svg.Attr("viewbox"); ok {
		t.Fatal("the parser no longer re-cases viewBox; this test needs another attribute")
	}
	checkFold(t, svg, []foldCase{
		{"viewBox", "0 0 16 16", true},
		{"viewbox", "0 0 16 16", true},
	})
}

func TestAttrFoldDoesNotAllocate(t *testing.T) {
	view := firstElement(t, loadRoot(t, ct.View(div.Div(
		attr.KV("Class", "a"), attr.KV("id", "b"), attr.KV("data-x", "c"),
	)())), "div")
	parsed := firstElement(t, loadRoot(t, cthtml.FragmentString("body",
		`<div class="a" id="b" data-x="c"></div>`)), "div")
	for _, n := range []ct.Node{view, parsed} {
		allocs := testing.AllocsPerRun(100, func() {
			n.AttrFold("hx-get") // a miss, the common case
			n.AttrFold("class")
			n.AttrFold("ID")
		})
		if allocs != 0 {
			t.Errorf("%T: %v allocations per run", n, allocs)
		}
	}
}
