package core

type renderConfig struct {
	Indent bool
}

// Option applies a configuration to the renderer.
type Option func(*renderConfig)

// WithIndent puts block-level elements on their own lines, indented with two
// spaces per depth level.
//
// Line breaks only go where browsers render no whitespace by default: between
// two block-level elements (such as div, p, li, and table parts, and
// everything in head), and between a block-level element and the start or
// end of its parent. Inline content (text, links, buttons, raw content,
// comments) stays on one line, and nothing is added inside pre, textarea,
// script, style, and title. The page therefore looks the same, but its DOM
// gains whitespace-only text nodes.
//
// The rules follow the default display of each element. Content whose CSS
// changes that (block elements shown inline or inline-block, or white-space:
// pre) may show the added whitespace; render it without indentation.
func WithIndent() Option {
	return func(cfg *renderConfig) {
		cfg.Indent = true
	}
}
