package chord

import (
	"context"
	"errors"

	"github.com/protolambda/chord/ba"
	"github.com/protolambda/chord/bi"
	"github.com/protolambda/chord/bs"
	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/html/attr"
	"github.com/protolambda/chord/html/form/button"
	"github.com/protolambda/chord/html/form/input"
	"github.com/protolambda/chord/html/form/label"
	"github.com/protolambda/chord/html/group/div"
	"github.com/protolambda/chord/html/group/list"
	"github.com/protolambda/chord/html/meta"
	"github.com/protolambda/chord/html/section"
	"github.com/protolambda/chord/html/table"
	"github.com/protolambda/chord/html/text"
	"github.com/protolambda/chord/hx2"
	"github.com/protolambda/chord/util"
)

// =============================================================================
// BASIC ELEMENTS
// =============================================================================

// The simplest case: an element with text content.
// text.Text() HTML-escapes the string for safe rendering.
func Example_basicElement() {
	core.Dump(div.Div(text.Text("hello world")))
	// Output: <div>hello world</div>
}

// Elements can be nested arbitrarily deep.
func Example_nestedElements() {
	core.Dump(div.Div(
		text.P(text.Text("A paragraph")),
		text.Span(text.Text("A span")),
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
	core.Dump(div.Div(core.Raw("<strong>Bold</strong> and <em>italic</em>")))
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
	// Output: ERROR: failed to render: duplicate attribute "id"
}

// =============================================================================
// BUNDLES - Combining Nodes
// =============================================================================

// Bundle combines multiple nodes into one, useful for reusable attribute sets.
func Example_attributeBundle() {
	// Define a reusable style bundle
	cardStyle := core.Bundle(attr.Class("card"), attr.Class("shadow"), attr.Style("padding:1rem"))
	core.Dump(div.Div(cardStyle, text.Text("Card content")))
	// Output: <div class="card shadow" style="padding:1rem">Card content</div>
}

// Bundle also works for combining multiple elements.
func Example_elementBundle() {
	core.Dump(core.Bundle(
		text.P(text.Text("First")),
		text.P(text.Text("Second")),
	))
	// Output: <p>First</p><p>Second</p>
}

// =============================================================================
// HTML DOCUMENT STRUCTURE
// =============================================================================

// Build a complete HTML document with proper structure.
func Example_fullDocument() {
	core.Dump(meta.HTML(
		meta.Head(
			meta.Title(text.Text("My Page")),
		),
		section.Body(
			section.H1(text.Text("Welcome")),
		),
	))
	// Output: <html><head><title>My Page</title></head><body><h1>Welcome</h1></body></html>
}

// =============================================================================
// LISTS AND TABLES
// =============================================================================

func Example_list() {
	core.Dump(list.UL(
		list.LI(text.Text("Item 1")),
		list.LI(text.Text("Item 2")),
	))
	// Output: <ul><li>Item 1</li><li>Item 2</li></ul>
}

func Example_table() {
	core.Dump(table.Table(
		table.Thead(table.TR(table.TH(text.Text("Name")), table.TH(text.Text("Age")))),
		table.Tbody(table.TR(table.TD(text.Text("Alice")), table.TD(text.Text("30")))),
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
	core.Dump(core.Bundle(
		label.Label(label.For("email"), text.Text("Email:")),
		input.Input(attr.ID("email"), input.Type("email")),
	))
	// Output: <label for="email">Email:</label><input id="email" type="email"/>
}

// =============================================================================
// BOOTSTRAP 5.3 - Elements and Attributes
// =============================================================================

// bs.Container(), bs.Row(), bs.Col*() are ELEMENT constructors that create divs.
func Example_bootstrapGrid() {
	core.Dump(bs.Container(
		bs.Row(
			bs.ColMD(6, text.Text("Left")),
			bs.ColMD(6, text.Text("Right")),
		),
	), core.WithIndent())
	// Output:
	// <div class="container">
	//   <div class="row">
	//     <div class="col-md-6">
	//       Left
	//     </div>
	//     <div class="col-md-6">
	//       Right
	//     </div>
	//   </div>
	// </div>
}

// Bootstrap utility classes (from ba package) can be combined freely.
func Example_bootstrapUtilities() {
	core.Dump(div.Div(ba.DFlex(), ba.JustifyContentBetween(), ba.AlignItemsCenter(), ba.P(3)))
	// Output: <div class="d-flex justify-content-between align-items-center p-3"></div>
}

// Bootstrap button components ARE elements (they create the button tag).
func Example_bootstrapButton() {
	core.Dump(bs.BtnPrimary(text.Text("Click me")))
	// Output: <button class="btn btn-primary">Click me</button>
}

// =============================================================================
// BOOTSTRAP ICONS
// =============================================================================

// Icons are typed constants that render as <i class="bi bi-{name}"></i>.
// They implement core.Node so can be used directly in element trees.
func Example_bootstrapIcon() {
	core.Dump(button.Button(bi.Check, text.Text(" Save")))
	// Output: <button><i class="bi bi-check"></i> Save</button>
}

// =============================================================================
// HTMX v2
// =============================================================================

// HTMX attributes enable hypermedia-driven interactions.
func Example_htmxBasic() {
	core.Dump(button.Button(
		hx2.Post("/api/submit"),
		hx2.Target("#result"),
		hx2.Swap("innerHTML"),
		text.Text("Submit"),
	))
	// Output: <button hx-post="/api/submit" hx-target="#result" hx-swap="innerHTML">Submit</button>
}

// hx2.On() uses v2 syntax: hx-on:event="handler"
func Example_htmxEvents() {
	core.Dump(div.Div(hx2.On("htmx:after-swap", "console.log('swapped')")))
	// Output: <div hx-on:htmx:after-swap="console.log('swapped')"></div>
}

// SSE in v2 requires the extension: hx-ext="sse"
func Example_htmxSSE() {
	core.Dump(div.Div(hx2.Ext("sse"), hx2.SSEConnect("/events"), hx2.SSESwap("message")))
	// Output: <div hx-ext="sse" sse-connect="/events" sse-swap="message"></div>
}

// =============================================================================
// CONDITIONAL RENDERING - util.If
// =============================================================================

// If() includes the node only when the condition is true.
// When false, it returns a no-op node that renders nothing.
func Example_conditionalTrue() {
	isAdmin := true
	core.Dump(div.Div(
		text.Text("Hello"),
		util.If(isAdmin, text.Span(attr.Class("badge"), text.Text("Admin"))),
	))
	// Output: <div>Hello<span class="badge">Admin</span></div>
}

func Example_conditionalFalse() {
	isAdmin := false
	core.Dump(div.Div(
		text.Text("Hello"),
		util.If(isAdmin, text.Span(attr.Class("badge"), text.Text("Admin"))),
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

	page := div.Div(
		util.Fn(func(ctx context.Context) (core.Node, error) {
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
	unreliable := util.Fn(func(ctx context.Context) (core.Node, error) {
		return nil, errors.New("database unavailable")
	})

	core.Dump(util.Fallback(
		div.Div(attr.Class("content"), unreliable),
		func(ctx context.Context, err error) core.Node {
			return div.Div(attr.Class("error"), text.Text("Failed to load content"))
		},
	))
	// Output: <div class="error">Failed to load content</div>
}

// =============================================================================
// PUTTING IT ALL TOGETHER
// =============================================================================

// A realistic example combining multiple features.
func Example_complete() {
	core.Dump(bs.Container(
		ba.MT(4),
		section.Header(
			section.H1(attr.Class("mb-3"), text.Text("Dashboard")),
		),
		bs.Row(
			bs.ColMD(4,
				bs.Card{Body: text.Text("Users: 42")},
			),
		),
		bs.BtnPrimary(
			ba.MT(3),
			hx2.Get("/api/refresh"), hx2.Target("#stats"),
			bi.ArrowClockwise, text.Text(" Refresh"),
		),
	), core.WithIndent())
	// Output:
	// <div class="container mt-4">
	//   <header>
	//     <h1 class="mb-3">
	//       Dashboard
	//     </h1>
	//   </header>
	//   <div class="row">
	//     <div class="col-md-4">
	//       <div class="card">
	//         <div class="card-body">
	//           Users: 42
	//         </div>
	//       </div>
	//     </div>
	//   </div>
	//   <button class="btn btn-primary mt-3" hx-get="/api/refresh" hx-target="#stats">
	//     <i class="bi bi-arrow-clockwise"></i>
	//      Refresh
	//   </button>
	// </div>
	//
}
