package chord

import (
	"context"
	"errors"

	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/core/attr"
	"github.com/protolambda/chord/core/elem"
	"github.com/protolambda/chord/html/form/button"
	"github.com/protolambda/chord/html/form/input"
	"github.com/protolambda/chord/html/form/label"
	"github.com/protolambda/chord/html/group/div"
	"github.com/protolambda/chord/html/group/list"
	"github.com/protolambda/chord/html/meta"
	"github.com/protolambda/chord/html/section"
	"github.com/protolambda/chord/html/table"
	"github.com/protolambda/chord/html/text"
)

// =============================================================================
// BASIC ELEMENTS
// =============================================================================

// The simplest case: an element with text content.
// text.Text() HTML-escapes the string for safe rendering.
func Example_basicElement() {
	core.Dump(div.Div()(text.Text("hello world")))
	// Output: <div>hello world</div>
}

// Elements can be nested arbitrarily deep.
func Example_nestedElements() {
	core.Dump(div.Div()(
		text.P()(text.Text("A paragraph")),
		text.Span()(text.Text("A span")),
	))
	// Output: <div><p>A paragraph</p><span>A span</span></div>
}

// Void elements like <input>, <br>, <img> self-close with />
func Example_voidElement() {
	core.Dump(input.Input(input.Type("text")))
	// Output: <input type="text"/>
}

// Raw HTML can be inserted directly when needed (use with caution - no escaping).
// Prefer text.Text() for user content to prevent XSS.
func Example_rawHTML() {
	core.Dump(div.Div()(elem.Raw("<strong>Bold</strong> and <em>italic</em>")))
	// Output: <div><strong>Bold</strong> and <em>italic</em></div>
}

// =============================================================================
// ATTRIBUTES
// =============================================================================

// Standard attributes: ID, Class, Style, etc.
func Example_basicAttributes() {
	core.Dump(div.Div(attr.ID("main"), attr.Class("container")))
	// Output: <div id="main" class="container"></div>
}

// Special behavior: multiple Class() calls are merged with spaces.
// This enables composable styling from different sources.
func Example_multipleClasses() {
	core.Dump(div.Div(attr.Class("foo"), attr.Class("bar"), attr.Class("baz")))
	// Output: <div class="foo bar baz"></div>
}

// Special behavior: multiple Style() calls are merged with semicolons.
func Example_multipleStyles() {
	core.Dump(div.Div(attr.Style("color:blue"), attr.Style("background:#222")))
	// Output: <div style="color:blue;background:#222"></div>
}

// Data attributes use the Data() helper.
func Example_dataAttributes() {
	core.Dump(div.Div(attr.Data("user-id", "123"), attr.Data("role", "admin")))
	// Output: <div data-user-id="123" data-role="admin"></div>
}

// Error case: duplicate non-mergeable attributes cause an error.
func Example_duplicateAttribute() {
	core.Dump(div.Div(attr.ID("first"), attr.ID("second")))
	// Output: ERROR: render: at div[0]: duplicate attribute "id"
}

// =============================================================================
// BUNDLES - Combining Nodes
// =============================================================================

// Bundle combines multiple nodes into one, useful for reusable attribute sets.
func Example_attributeBundle() {
	// Define a reusable style bundle
	cardStyle := attr.Bundle{attr.Class("card"), attr.Class("shadow"), attr.Style("padding:1rem")}
	core.Dump(div.Div(cardStyle)(text.Text("Card content")))
	// Output: <div class="card shadow" style="padding:1rem">Card content</div>
}

// Bundle also works for combining multiple elements.
func Example_elementBundle() {
	core.Dump(elem.Bundle{
		text.P()(text.Text("First")),
		text.P()(text.Text("Second")),
	})
	// Output: <p>First</p><p>Second</p>
}

// =============================================================================
// HTML DOCUMENT STRUCTURE
// =============================================================================

// Build a complete HTML document with proper structure.
func Example_fullDocument() {
	core.Dump(elem.Bundle{elem.Doctype(), meta.HTML()(
		meta.Head()(
			meta.Title()(text.Text("My Page")),
		),
		section.Body()(
			section.H1()(text.Text("Welcome")),
		),
	)})
	// Output: <!DOCTYPE html><html><head><title>My Page</title></head><body><h1>Welcome</h1></body></html>
}

// =============================================================================
// LISTS AND TABLES
// =============================================================================

func Example_list() {
	core.Dump(list.UL()(
		list.LI()(text.Text("Item 1")),
		list.LI()(text.Text("Item 2")),
	))
	// Output: <ul><li>Item 1</li><li>Item 2</li></ul>
}

func Example_table() {
	core.Dump(table.Table()(
		table.Thead()(table.TR()(table.TH()(text.Text("Name")), table.TH()(text.Text("Age")))),
		table.Tbody()(table.TR()(table.TD()(text.Text("Alice")), table.TD()(text.Text("30")))),
	))
	// Output: <table><thead><tr><th>Name</th><th>Age</th></tr></thead><tbody><tr><td>Alice</td><td>30</td></tr></tbody></table>
}

// =============================================================================
// FORMS
// =============================================================================

func Example_formInput() {
	core.Dump(input.Input(
		input.Type("email"),
		input.Name("email"),
		input.Placeholder("you@example.com"),
		input.Required(),
	))
	// Output: <input type="email" name="email" placeholder="you@example.com" required/>
}

func Example_checkbox() {
	core.Dump(input.Input(input.Type("checkbox"), input.Name("agree"), input.Checked()))
	// Output: <input type="checkbox" name="agree" checked/>
}

func Example_labeledInput() {
	core.Dump(elem.Bundle{
		label.Label(label.For("email"))(text.Text("Email:")),
		input.Input(attr.ID("email"), input.Type("email")),
	})
	// Output: <label for="email">Email:</label><input id="email" type="email"/>
}

// =============================================================================
// CUSTOM ATTRIBUTES
// =============================================================================

// Attribute vocabularies of libraries, such as htmx, are trusted names with
// escaped values. The packages of github.com/protolambda/webui (hx2, bs, ba,
// bi) are built this way.
func Example_customAttributes() {
	hxPost := attr.Name("hx-post")
	hxTarget := attr.Name("hx-target")
	core.Dump(button.Button(
		button.Type(button.TypeSubmit),
		hxPost.Value("/api/submit"),
		hxTarget.Value("#result"),
	)(text.Text("Submit")))
	// Output: <button type="submit" hx-post="/api/submit" hx-target="#result">Submit</button>
}

// A component is a plain Go function that composes elements, and accepts
// the attributes of the caller.
func Example_component() {
	alert := func(attrs ...attr.Node) elem.Scope {
		return div.Div(attr.Class("alert"), attr.Bundle(attrs))
	}
	core.Dump(alert(attr.Class("alert-danger"))(text.Text("Failed")))
	// Output: <div class="alert alert-danger">Failed</div>
}

// =============================================================================
// CONDITIONAL RENDERING - util.If
// =============================================================================

// If() includes the node only when the condition is true.
// When false, it returns a no-op node that renders nothing.
func Example_conditionalTrue() {
	isAdmin := true
	core.Dump(div.Div()(
		text.Text("Hello"),
		elem.If(isAdmin, text.Span(attr.Class("badge"))(text.Text("Admin"))),
	))
	// Output: <div>Hello<span class="badge">Admin</span></div>
}

func Example_conditionalFalse() {
	isAdmin := false
	core.Dump(div.Div()(
		text.Text("Hello"),
		elem.If(isAdmin, text.Span(attr.Class("badge"))(text.Text("Admin"))),
	))
	// Output: <div>Hello</div>
}

// =============================================================================
// DYNAMIC CONTENT - util.Fn
// =============================================================================

// Fn() generates content at render time using the context.
// This enables the same page structure to render differently based on state.
func Example_dynamicContent() {
	type ctxKey string
	const userKey ctxKey = "user"

	page := div.Div()(
		elem.Fn(func(ctx context.Context) (elem.Node, error) {
			if user, ok := ctx.Value(userKey).(string); ok {
				return text.Text("Hello, " + user), nil
			}
			return text.Text("Hello, Guest"), nil
		}),
	)

	// Render with user in context
	ctx := context.WithValue(context.Background(), userKey, "Alice")
	core.DumpCtx(ctx, page)
	// Output: <div>Hello, Alice</div>
}

// =============================================================================
// ERROR HANDLING - util.Fallback
// =============================================================================

// Fallback() catches render errors and substitutes alternative content.
// Useful for graceful degradation of page sections.
func Example_fallback() {
	unreliable := elem.Fn(func(ctx context.Context) (elem.Node, error) {
		return nil, errors.New("database unavailable")
	})

	core.Dump(core.Fallback(
		div.Div(attr.Class("content"))(unreliable),
		func(ctx context.Context, err error) elem.Node {
			return div.Div(attr.Class("error"))(text.Text("Failed to load content"))
		},
	))
	// Output: <div class="error">Failed to load content</div>
}

// =============================================================================
// PUTTING IT ALL TOGETHER
// =============================================================================

// A realistic example combining multiple features.
func Example_complete() {
	hxGet := attr.Name("hx-get")
	hxTarget := attr.Name("hx-target")
	core.Dump(div.Div(attr.Class("container"))(
		section.Header()(
			section.H1()(text.Text("Dashboard")),
		),
		div.Div(attr.ID("stats"))(
			text.P()(text.Text("Users: 42")),
		),
		div.Div(attr.Class("actions"))(
			button.Button(button.Type(button.TypeButton), hxGet.Value("/api/refresh"), hxTarget.Value("#stats"))(
				text.Text("Refresh"),
			),
		),
	), core.WithIndent())
	// Output:
	// <div class="container">
	//   <header>
	//     <h1>Dashboard</h1>
	//   </header>
	//   <div id="stats">
	//     <p>Users: 42</p>
	//   </div>
	//   <div class="actions"><button type="button" hx-get="/api/refresh" hx-target="#stats">Refresh</button></div>
	// </div>
}
