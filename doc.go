// Package chord provides a composable, lazily evaluated HTML DSL for Go.
//
// Chord deliberately uses small semantic packages. The package name at a call
// site identifies the relevant HTML element or extension, and allows related
// packages to expose natural names without collisions. This package is the
// starting index; application code normally imports the packages below.
//
// # Package tree
//
// Paths in this tree are relative to github.com/protolambda/chord.
//
//	chord/
//	├── core/             # Core rendering
//	│   ├── elem/         # Element core types and functions
//	│   ├── attr/         # Attribute core types, functions, and global attributes
//	│   └── inspect/      # Evaluated snapshots for inspection
//	├── ct/               # Testing: subjects, queries, and mustbe assertions
//	│   ├── cthtml/       # Parsed HTML page and fragment subjects
//	│   └── cthttp/       # Handler execution and response subjects
//	├── html/             # HTML elements and attributes
//	│   ├── aria/         # ARIA accessibility attributes
//	│   ├── on/           # DOM event handlers (onclick, onsubmit, etc.)
//	│   ├── meta/         # Document metadata (html, head, title, meta, link, style)
//	│   ├── section/      # Sectioning (body, article, nav, header, footer, h1-h6)
//	│   ├── text/         # Text semantics (span, a, em, strong, code, p, pre, etc.)
//	│   ├── group/        # Grouping content
//	│   │   ├── div/      # Div element
//	│   │   ├── list/     # Lists (ol, ul, li, menu)
//	│   │   ├── dl/       # Definition lists (dl, dt, dd)
//	│   │   └── figure/   # Figures (figure, figcaption)
//	│   ├── edit/         # Edit elements (ins, del)
//	│   ├── embed/        # Embedded content
//	│   │   ├── img/      # Images (img, picture, source)
//	│   │   ├── media/    # Media (video, audio, track)
//	│   │   ├── iframe/   # Inline frames
//	│   │   ├── object/   # Object/embed (legacy)
//	│   │   └── area/     # Image maps (map, area)
//	│   ├── table/        # Table elements
//	│   ├── form/         # Form elements
//	│   │   ├── input/    # Input element
//	│   │   ├── button/   # Button element
//	│   │   ├── select/   # Select, option, optgroup, datalist
//	│   │   ├── textarea/ # Textarea element
//	│   │   ├── label/    # Label element
//	│   │   └── output/   # Output, progress, meter
//	│   ├── interactive/  # Interactive elements (details, summary, dialog)
//	│   ├── script/       # Scripting (script, noscript, template, canvas)
//	│   └── webcomp/      # Web components (slot)
//	├── hx1/              # HTMX v1 attributes
//	├── hx2/              # HTMX v2 attributes
//	├── bs/               # Bootstrap 5.3 element components
//	├── ba/               # Bootstrap 5.3 attribute utilities
//	└── bi/               # Bootstrap Icons
//
// The [github.com/protolambda/chord/core/elem] and
// [github.com/protolambda/chord/core/attr] packages define the common node
// interfaces composed by all other packages. Start with
// [github.com/protolambda/chord/core.Render] to render a completed element tree,
// and with [github.com/protolambda/chord/ct.View] to test one.
package chord
