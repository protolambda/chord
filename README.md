# chord

A Go DSL for generating HTML structures with a focus on composition and type safety.

## Installation

```bash
go get github.com/protolambda/chord@latest
```

## Overview

Chord provides a clean, composable API for building HTML documents in Go.
Structures are lazily evaluated, allowing reuse and context-aware rendering.

### Core Types

- `attr.Node`: A lazy-evaluated attribute (key-value, boolean, or bundle)
- `attr.Name`: A trusted, statically known attribute name
- `elem.Node`: A lazy-evaluated element (tag, text, raw HTML, or bundle)
- `elem.Name`: A trusted, statically known element tag name
- `elem.Scope`: A `func(...elem.Node) elem.Node`, a scope of sub-elements (`<div>`, etc.)

Constructors:
- `attr.KV(k, v)`: A validated runtime name with an HTML-escaped value
- `attr.Bool(k)`: A validated runtime name for a boolean attribute
- `attr.ParseName(v)`: Validate a runtime attribute name
- `attr.Name("key").Value(v)`: A trusted name with an HTML-escaped value
- `attr.Name("key").Raw(v)`: A trusted name and output-ready value (no escaping)
- `elem.ParseName(v)`: Validate a runtime element tag name
- `elem.Name("div").New(attrs...)`: A non-void HTML element (returns `Scope`)
- `elem.Name("input").Void(attrs...)`: A self-closing element (returns `elem.Node`)
- `elem.Raw(v)`: Raw HTML content (no escaping)
- `elem.Comment(v)`: An HTML comment
- `elem.Noop()`, `attr.Noop()`: Empty no-op

Core utils:
- `text.Text(v)`: HTML-escaped text content (returns `elem.Node`)
- `elem.If(bool, elem)`, `attr.If(bool, attr)`: Conditional content
- `elem.Fn(func(ctx) (elem, error))`, `attr.Fn(func(ctx) (attr, error))`: Dynamic content
- `core.Fallback(node, fallback func(ctx, err) elem)`: Element with recovery


Non-void elements use a two-step call: first attributes, then children:

```go
div.Div(attr.Class("outer"))(        // attributes
    text.P()(text.Text("content")),   // children
)
```

`Scope` also implements `elem.Node`, so elements without children don't need a trailing `()`:

```go
div.Div(attr.Class("empty"))  // renders as: <div class="empty"></div>
```

## Package Structure

```
chord/
├── core/             # Core rendering
│   ├── elem/         # Element core types and functions
│   └── attr/         # Attribute core types, functions, and global attributes
├── html/             # HTML elements and attributes
│   ├── aria/         # ARIA accessibility attributes
│   ├── on/           # DOM event handlers (onclick, onsubmit, etc.)
│   ├── meta/         # Document metadata (html, head, title, meta, link, style)
│   ├── section/      # Sectioning (body, article, nav, header, footer, h1-h6)
│   ├── text/         # Text semantics (span, a, em, strong, code, p, pre, etc.)
│   ├── group/        # Grouping content
│   │   ├── div/      # Div element
│   │   ├── list/     # Lists (ol, ul, li, menu)
│   │   ├── dl/       # Definition lists (dl, dt, dd)
│   │   └── figure/   # Figures (figure, figcaption)
│   ├── edit/         # Edit elements (ins, del)
│   ├── embed/        # Embedded content
│   │   ├── img/      # Images (img, picture, source)
│   │   ├── media/    # Media (video, audio, track)
│   │   ├── iframe/   # Inline frames
│   │   ├── object/   # Object/embed (legacy)
│   │   └── area/     # Image maps (map, area)
│   ├── table/        # Table elements
│   ├── form/         # Form elements
│   │   ├── input/    # Input element
│   │   ├── button/   # Button element
│   │   ├── select/   # Select, option, optgroup, datalist
│   │   ├── textarea/ # Textarea element
│   │   ├── label/    # Label element
│   │   └── output/   # Output, progress, meter
│   ├── interactive/  # Interactive elements (details, summary, dialog)
│   ├── script/       # Scripting (script, noscript, template, canvas)
│   └── webcomp/      # Web components (slot)
├── hx1/              # HTMX v1 attributes
├── hx2/              # HTMX v2 attributes
├── bs/               # Bootstrap 5.3 element components
├── ba/               # Bootstrap 5.3 attribute utilities
└── bi/               # Bootstrap Icons
```

## Example

```go
package main

import (
	"context"
	"log"
	"strings"

	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/core/attr"
	"github.com/protolambda/chord/core/elem"
	"github.com/protolambda/chord/html/group/div"
	"github.com/protolambda/chord/html/meta"
	"github.com/protolambda/chord/html/section"
	"github.com/protolambda/chord/html/text"
)

// Context key for authentication state
type ctxKey string

const isLoggedInKey ctxKey = "isLoggedIn"

func main() {
	// Build the page structure (can be reused with different contexts)
	page := meta.HTML()(
		meta.Head()(
			meta.Title()(text.Text("My Page")),
		),
		section.Body()(
			section.Header()(
				section.H1()(text.Text("Welcome")),
				// Use Fn to read from context and conditionally render
				elem.Fn(func(ctx context.Context) (elem.Node, error) {
					if loggedIn, _ := ctx.Value(isLoggedInKey).(bool); loggedIn {
						return div.Div(attr.Class("user-menu"))(text.Text("Logged in")), nil
					}
					return elem.Noop(), nil
				}),
			),
			section.Main()(
				text.P()(text.Text("Hello, world!")),
			),
		),
	)

	// Attach state to context and render
	ctx := context.WithValue(context.Background(), isLoggedInKey, true)
	var out strings.Builder
	if err := core.Render(ctx, page, &out, core.WithIndent()); err != nil {
		log.Fatal(err)
	}
	// out.String() contains the rendered HTML
}
```

## Extensions

### HTMX

```go
import "github.com/protolambda/chord/hx1" // HTMX v1
import "github.com/protolambda/chord/hx2" // HTMX v2

// Attributes go in the first call, children in the second
button.Button(hx1.Post("/api/submit"), hx1.Target("#result"))(
    text.Text("Submit"),
)
```

### Bootstrap 5.3

Bootstrap support is split into two packages:
- `bs`: Element components that create HTML elements (return `Scope`)
- `ba`: Attribute utilities that add classes to elements (return `Attrib`)

This separation prevents confusion: `bs.Row()` creates a `<div class="row">`,
while `ba.Row()` returns a class attribute you can apply to any element.

#### Elements (bs package)

```go
import "github.com/protolambda/chord/bs"

// Grid elements (attrs...)(children...)
bs.Container(attrs...)(children...)   // <div class="container">
bs.ContainerFluid(attrs...)(...)      // <div class="container-fluid">
bs.Row(attrs...)(children...)         // <div class="row">
bs.Col(attrs...)(children...)         // <div class="col">
bs.Col6(attrs...)(children...)        // <div class="col-6">
bs.ColMD(4, attrs...)(children...)    // <div class="col-md-4">

// Buttons
bs.Btn(attrs...)(children...)                 // <button class="btn">
bs.BtnPrimary(attrs...)(children...)          // <button class="btn btn-primary">
bs.BtnOutlineSecondary(attrs...)(children...) // <button class="btn btn-outline-secondary">

// Components
bs.Card{Header: ..., Body: ...}     // struct implementing elem.Node
bs.Modal{ID: "...", Title: ..., Body: ...}
bs.Dropdown{Toggle: ..., Items: ...}
bs.AlertDanger()(children...)
bs.BadgeSuccess()(children...)
bs.Nav()(children...), bs.Navbar(attrs...)(children...)
```

#### Attributes (ba package)

```go
import "github.com/protolambda/chord/ba"

// Spacing (margin/padding, values 0-5)
ba.M(3), ba.MT(4), ba.MB(2), ba.MX(3), ba.MY(2)
ba.P(3), ba.PT(4), ba.PB(2), ba.PX(3), ba.PY(2)

// Flexbox
ba.DFlex(), ba.DInlineFlex()
ba.FlexRow(), ba.FlexColumn(), ba.FlexWrap()
ba.JustifyContentCenter(), ba.JustifyContentBetween()
ba.AlignItemsCenter(), ba.AlignItemsStart()
ba.Gap(3)

// Colors
ba.BgPrimary(), ba.BgDark(), ba.BgLight()
ba.TextPrimary(), ba.TextMuted(), ba.TextWhite()

// Borders & Shadows
ba.Border(), ba.BorderPrimary(), ba.Border0()
ba.Rounded(), ba.RoundedCircle(), ba.RoundedPill()
ba.Shadow(), ba.ShadowSM(), ba.ShadowLG()

// Sizing
ba.W100(), ba.W50(), ba.H100(), ba.WAuto()

// Display
ba.DNone(), ba.DBlock(), ba.DInline()

// Text
ba.TextCenter(), ba.TextEnd()
ba.FWBold(), ba.FSItalic()

// Forms
ba.FormControl(), ba.FormSelect(), ba.FormCheckInput()
```

#### Combined Example

```go
bs.Container(ba.MT(4))(
    bs.Row()(
        bs.ColMD(6, ba.MB(3))(
            bs.Card{
                Header: text.Text("Users"),
                Body:   text.Text("42 active"),
            },
        ),
        bs.ColMD(6, ba.MB(3))(
            div.Div(ba.DFlex(), ba.JustifyContentBetween())(
                text.Span()(text.Text("Status")),
                bs.BadgeSuccess()(text.Text("Online")),
            ),
        ),
    ),
    bs.BtnPrimary(ba.MT(3))(text.Text("Refresh")),
)
```

### Bootstrap Icons

```go
import "github.com/protolambda/chord/bi"

bi.Check      // <i class="bi bi-check"></i>
bi.XLg        // <i class="bi bi-x-lg"></i>
bi.Github     // <i class="bi bi-github"></i>
```

## Design Philosophy

- **Composition over repetition**: Build reusable components easily
- **Type safety**: Native Go typing without code generation
- **Clean API**: Scoped packages prevent namespace bloat
- **Extensibility**: Create custom components with the same patterns
- **Zero dependencies**: Core library has no external dependencies

## License

MIT, see [`LICENSE`](./LICENSE) file.
