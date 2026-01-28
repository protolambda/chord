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

- `core.Node`: A lazy-evaluated piece of the document graph
- `core.Obj`: A static representation of evaluated content

Node types include:
- `core.Raw`: Raw document content (text, HTML)
- `core.Element`: An HTML element with children
- `core.VoidElement`: A self-closing element (e.g., `<br/>`, `<img/>`)
- `core.Attribute`: An attribute key-value pair
- `core.BoolAttribute`: A boolean attribute (no value)
- `core.Bundle`: A combination of multiple nodes
- `core.Noop`: An empty node

Nodes can be either attributes or elements, and bundles can mix both for easy composition.

## Package Structure

```
chord/
├── core/             # Core types (Node, Obj) and rendering
├── util/             # Utilities (If, Fn, Fallback)
├── html/             # HTML elements and attributes
│   ├── attr/         # Global attributes (class, id, style, data, etc.)
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
│   │   ├── select/   # Select, option, optgroup
│   │   ├── textarea/ # Textarea element
│   │   ├── label/    # Label element
│   │   └── output/   # Output, progress, meter
│   ├── interactive/  # Interactive elements (details, summary, dialog)
│   ├── script/       # Scripting (script, noscript, template, canvas)
│   └── webcomp/      # Web components (slot)
├── hx1/              # HTMX v1 attributes
├── hx2/              # HTMX v2 attributes
├── bs/               # Bootstrap 5.3 element components (buttons, cards, grids, etc.)
├── ba/               # Bootstrap 5.3 attribute utilities (spacing, colors, flexbox, etc.)
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
    "github.com/protolambda/chord/html/attr"
    "github.com/protolambda/chord/html/group/div"
    "github.com/protolambda/chord/html/meta"
    "github.com/protolambda/chord/html/section"
    "github.com/protolambda/chord/html/text"
    "github.com/protolambda/chord/util"
)

// Context key for authentication state
type ctxKey string
const isLoggedInKey ctxKey = "isLoggedIn"

func main() {
    // Build the page structure (can be reused with different contexts)
    page := meta.HTML(
        meta.Head(
            meta.Title(text.Text("My Page")),
        ),
        section.Body(
            section.Header(
                section.H1(text.Text("Welcome")),
                // Use Fn to read from context and conditionally render
                util.Fn(func(ctx context.Context) (core.Node, error) {
                    if loggedIn, _ := ctx.Value(isLoggedInKey).(bool); loggedIn {
                        return div.Div(attr.Class("user-menu"), text.Text("Logged in")), nil
                    }
                    return core.Noop(), nil
                }),
            ),
            section.Main(
                text.P(text.Text("Hello, world!")),
            ),
        ),
    )

    // Attach state to context and render
    ctx := context.WithValue(context.Background(), isLoggedInKey, true)
    var out strings.Builder
    if err := core.Render(ctx, page, &out); err != nil {
        log.Fatal(err)
    }
    // out.String() contains the rendered HTML
}
```

## Utilities

- `text.Text(v)`: HTML-escaped text content
- `util.If(cond, node)`: Conditional rendering
- `util.Fn(func)`: Dynamic content generation
- `util.Fallback(node, fallback)`: Error recovery

## Extensions

### HTMX

```go
import "github.com/protolambda/chord/hx1" // HTMX v1
import "github.com/protolambda/chord/hx2" // HTMX v2

button.Button(hx1.Post("/api/submit"), hx1.Target("#result"))
```

### Bootstrap 5.3

Bootstrap support is split into two packages:
- `bs`: Element components that create HTML elements
- `ba`: Attribute utilities that add classes to elements

This separation prevents confusion: `bs.Row()` creates a `<div class="row">`,
while `ba.Row()` returns a class attribute you can apply to any element.

#### Elements (bs package)

```go
import "github.com/protolambda/chord/bs"

// Grid elements
bs.Container(...)           // <div class="container">
bs.ContainerFluid(...)      // <div class="container-fluid">
bs.Row(...)                 // <div class="row">
bs.Col(...)                 // <div class="col">
bs.Col6(...)                // <div class="col-6">
bs.ColMD(4, ...)            // <div class="col-md-4">

// Buttons
bs.Btn(...)                 // <button class="btn">
bs.BtnPrimary(...)          // <button class="btn btn-primary">
bs.BtnOutlineSecondary(...) // <button class="btn btn-outline-secondary">

// Components
bs.Card{Header: ..., Body: ...}
bs.Alert(...), bs.AlertDanger(...)
bs.Badge(...), bs.BadgeSuccess(...)
bs.Nav(...), bs.Navbar(...)
bs.Modal{ID: "...", Title: ..., Body: ...}
bs.Dropdown{Toggle: ..., Items: ...}
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
bs.Container(
    ba.MT(4),
    bs.Row(
        bs.ColMD(6, ba.MB(3),
            bs.Card{
                Header: text.Text("Users"),
                Body:   text.Text("42 active"),
            },
        ),
        bs.ColMD(6, ba.MB(3),
            div.Div(ba.DFlex(), ba.JustifyContentBetween(),
                text.Span(text.Text("Status")),
                bs.BadgeSuccess(text.Text("Online")),
            ),
        ),
    ),
    bs.BtnPrimary(ba.MT(3), text.Text("Refresh")),
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
