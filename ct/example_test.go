package ct_test

import (
	"context"
	"fmt"

	"github.com/protolambda/mustbe"
	"github.com/protolambda/mustbe/assertion"

	"github.com/protolambda/chord/core/attr"
	"github.com/protolambda/chord/core/elem"
	"github.com/protolambda/chord/ct"
	"github.com/protolambda/chord/html/group/div"
	"github.com/protolambda/chord/html/group/list"
	"github.com/protolambda/chord/html/section"
	"github.com/protolambda/chord/html/text"
)

// printT is a minimal mustbe.T that prints failures instead of failing a test,
// so that examples can show what a failing assertion reports.
type printT struct{}

func (printT) Error(args ...any)        { fmt.Println(args...) }
func (printT) FailNow()                 {}
func (printT) Context() context.Context { return context.Background() }
func (printT) Helper()                  {}

// ProjectPage is an example view under test.
func ProjectPage(name string, canEdit bool) elem.Node {
	return section.Main()(
		section.H1()(text.Text(name)),
		list.UL(attr.ID("actions"))(
			list.LI()(text.A(text.Href("/projects/"+name))(text.Text("View"))),
			elem.If(canEdit, list.LI()(text.A(text.Href("/projects/"+name+"/edit"))(text.Text("Edit"))))),
	)
}

// A passing test reads as: build the subject, derive selections, assert.
func Example() {
	t := mustbe.WrapT(printT{}) // in a real test: mustbe.WrapT(t) or devtest.SerialT(t)
	page := ct.View(ProjectPage("chord", false))

	t.Must(page)
	t.Must(page.Find(ct.Tag("h1"), ct.Text("chord")))
	t.Must(page.Find(ct.Tag("a"), ct.Text("Edit")).None())

	actions := page.Find(ct.ID("actions"))
	t.Must(actions)
	t.Must(actions.Find(ct.Tag("a")).Texts("View"))
	t.Must(actions.Find(ct.Tag("a")).Matches(ct.Attr("href", "/projects/chord")))

	fmt.Println("all passed")
	// Output: all passed
}

// A failing cardinality assertion reports the expectation, the matches with
// their locations, and (when nothing matched) an outline of the scope.
func Example_failure() {
	t := mustbe.WrapT(printT{})
	page := ct.View(ProjectPage("chord", true))

	t.Must(page.Find(ct.Tag("a")))
	t.Must(page.Find(ct.ID("actions")).Find(ct.Tag("button")).Any())
	// Output:
	// assertion failed: unexpected number of matches: expected exactly one match of tag("a") in chord view, found 2
	// matches:
	//   main[0]/ul#actions[1]/li[0]/a[0] a [href="/projects/chord"]
	//   main[0]/ul#actions[1]/li[1]/a[0] a [href="/projects/chord/edit"]
	// assertion failed: unexpected number of matches: expected at least one match of tag("button") within id("actions") in chord view, found 0
	// scope id("actions") matched 1
	// scope outline:
	//   ul#actions
	//     li
	//       a [href="/projects/chord"]
	//         "View"
	//     li
	//       a [href="/projects/chord/edit"]
	//         "Edit"
}

// Assertions are values: a reusable contract can be applied to any subject.
func Example_contract() {
	editable := func(page *ct.Subject) assertion.Assertion {
		return page.Find(ct.Tag("a"), ct.Text("Edit")).Matches(ct.AttrMatches("href", ct.Suffix("/edit")))
	}

	t := mustbe.WrapT(printT{})
	t.Must(editable(ct.View(ProjectPage("chord", true))))
	t.Must(editable(ct.View(ProjectPage("chord", false))))
	// Output:
	// assertion failed: unexpected number of matches: expected exactly one match of and(tag("a"), text("Edit")) in chord view, found 0
	// scope outline:
	//   main
	//     h1
	//       "chord"
	//     ul#actions
	//       li
	//         a [href="/projects/chord"]
	//           "View"
}

// Sequence assertions check every match at once, in document order:
// Each requires a property of all matches, InOrder pairs the i-th match
// with the i-th query, and Texts and AttrValues compare extracted values.
func Example_sequences() {
	t := mustbe.WrapT(printT{})
	links := ct.View(ProjectPage("chord", true)).Find(ct.Role("link"))

	t.Must(links.Each(ct.AttrMatches("href", ct.Prefix("/projects/chord"))))
	t.Must(links.InOrder(ct.Text("View"), ct.Text("Edit")))
	t.Must(links.Texts("View", "Edit"))
	t.Must(links.AttrValues("href", "/projects/chord", "/projects/chord/edit"))

	// Failures name the offending match.
	t.Must(links.Each(ct.Attr("href", "/projects/chord")))
	t.Must(links.AttrValues("href", "/projects/chord"))
	// Output:
	// assertion failed: selected nodes do not match: expected every match of role("link") in chord view to match attr("href", "/projects/chord"), match 1 does not
	// matches:
	//   main[0]/ul#actions[1]/li[0]/a[0] a [href="/projects/chord"]
	//   main[0]/ul#actions[1]/li[1]/a[0] a [href="/projects/chord/edit"]
	// assertion failed: selected nodes do not match: expected href values of role("link") in chord view to be ["/projects/chord"], got ["/projects/chord" "/projects/chord/edit"]
	// matches:
	//   main[0]/ul#actions[1]/li[0]/a[0] a [href="/projects/chord"]
	//   main[0]/ul#actions[1]/li[1]/a[0] a [href="/projects/chord/edit"]
}

// First, Last, and Nth narrow a selection to one position, so that the
// exactly-one default and Matches apply to a single element. A position
// that does not exist matches nothing.
func Example_positions() {
	t := mustbe.WrapT(printT{})
	items := ct.View(ProjectPage("chord", true)).Find(ct.Role("listitem"))

	t.Must(items.First().Find(ct.Role("link")).Matches(ct.Text("View")))
	t.Must(items.Last().Find(ct.Role("link")).Matches(ct.Text("Edit")))
	t.Must(items.Nth(-1).Find(ct.Role("link")).Texts("Edit"))
	t.Must(items.Nth(2).None())

	t.Must(items.Nth(2))
	// Output:
	// assertion failed: unexpected number of matches: expected exactly one match of nth(2) of role("listitem") in chord view, found 0
	// scope role("listitem") matched 2
	// scope outline:
	//   li
	//     a [href="/projects/chord"]
	//       "View"
	//   li
	//     a [href="/projects/chord/edit"]
	//       "Edit"
}

// Texts read as users see the page: the text of block-level parts, such as
// the header and body of a card, is separated, while inline elements join.
// TextContent is the raw DOM textContent.
func ExampleInnerText() {
	t := mustbe.WrapT(printT{})
	page := ct.View(div.Div(attr.Class("card"))(
		div.Div(attr.Class("card-header"))(text.Text("Author")),
		div.Div(attr.Class("card-body"))(text.Text("Signed with "), text.Strong()(text.Text("dev")), text.Text("0")),
	))

	t.Must(page.Find(ct.Class("card")).Texts("Author Signed with dev0"))
	t.Must(page.Find(ct.InnerText("Author Signed with dev0")))
	t.Must(page.Find(ct.TextContent("AuthorSigned with dev0")))
	fmt.Println("all passed")
	// Output: all passed
}

// Attr and Text read a value of the single match, for values that a test
// uses further, such as the URL that a polling panel requests next.
func ExampleSelection_Attr() {
	page := ct.View(div.Div(attr.ID("panel"), attr.KV("hx-get", "/tasks/1/panel"), attr.KV("hx-trigger", "every 1s"))(
		text.Text("pending"),
	))
	panel := page.Find(ct.ID("panel"))

	next, polling, err := panel.Attr(context.Background(), "hx-get")
	fmt.Println(next, polling, err)
	status, err := panel.Text(context.Background())
	fmt.Println(status, err)
	_, polling, err = panel.Attr(context.Background(), "hx-swap")
	fmt.Println(polling, err)
	// Output:
	// /tasks/1/panel true <nil>
	// pending <nil>
	// false <nil>
}
