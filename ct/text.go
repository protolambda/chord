package ct

import (
	"fmt"
	"strings"
)

// Text matches elements whose own text equals the value. Own text is the
// text of the element's direct text children, so that a wrapper does not
// match the text of a nested element. A br or block-level child in between
// separates the texts, as in [InnerText]; whitespace is collapsed. Text does
// not model visibility.
func Text(value string) Query {
	return Query{
		desc:  fmt.Sprintf("text(%q)", value),
		match: func(n Node) (bool, error) { return ownText(n) == value, nil },
	}
}

// TextMatches matches elements whose own text satisfies the matcher.
func TextMatches(value ValueMatch) Query {
	return Query{
		desc:  fmt.Sprintf("text(%s)", value),
		match: func(n Node) (bool, error) { return value.Match(ownText(n)), nil },
	}
}

// InnerText matches elements whose subtree text, laid out as users read
// it, equals the value. Like a browser's innerText, it separates the text
// of block-level boxes and of br elements, and whitespace is collapsed, so
// a card with a header "Author" and a body "Signed" reads "Author Signed",
// while inline elements join: "5<small>ETH</small>" reads "5ETH".
//
// Without CSS, the layout is approximated: the default display of the HTML
// elements; a display in the style attribute; the unprefixed Bootstrap 5.3
// display utilities (d-block, d-flex, ...), stacks (hstack, vstack), and the
// Bootstrap components that are flex containers (card, nav, navbar,
// btn-group, input-group, list-group, pagination, breadcrumb, row, and the
// modal, toast and offcanvas headers) or that show links as blocks
// (nav-link, dropdown-item, list-group-item). The children of flex and grid
// containers are separated, as their items are blocks.
//
// The content of head, title, script, style, template and noscript
// elements is never rendered and left out; for such an element itself,
// the inner text is its text content, as in a browser. Otherwise inner text
// does not model visibility: hidden elements, display: none and d-none
// count. Raw content and comments do not count as text.
func InnerText(value string) Query {
	return Query{
		desc:  fmt.Sprintf("innerText(%q)", value),
		match: func(n Node) (bool, error) { return innerText(n) == value, nil },
	}
}

// InnerTextMatches matches elements whose inner text satisfies the
// matcher. See [InnerText].
func InnerTextMatches(value ValueMatch) Query {
	return Query{
		desc:  fmt.Sprintf("innerText(%s)", value),
		match: func(n Node) (bool, error) { return value.Match(innerText(n)), nil },
	}
}

// TextContent matches elements whose whole subtree text equals the value,
// with whitespace collapsed. It is the raw DOM textContent: text nodes are
// joined without separation, so a header "Author" followed by a body
// "Signed" reads "AuthorSigned". Prefer [InnerText], which reads as users
// see the page. Raw content and comments do not count as text.
func TextContent(value string) Query {
	return Query{
		desc:  fmt.Sprintf("textContent(%q)", value),
		match: func(n Node) (bool, error) { return textContent(n) == value, nil },
	}
}

// TextContentMatches matches elements whose subtree text satisfies the
// matcher. See [TextContent].
func TextContentMatches(value ValueMatch) Query {
	return Query{
		desc:  fmt.Sprintf("textContent(%s)", value),
		match: func(n Node) (bool, error) { return value.Match(textContent(n)), nil },
	}
}

// ownText returns the collapsed text of the direct text children of n,
// separated where a br or block-level child element sits between them.
func ownText(n Node) string {
	var b strings.Builder
	for c := range n.Children() {
		switch c.Kind() {
		case KindText:
			b.WriteString(c.Data())
		case KindElement:
			if c.Tag() == "br" || blockBox(c) {
				b.WriteByte('\n')
			}
		}
	}
	return collapseSpace(b.String())
}

// textContent returns the collapsed text of all text nodes in the subtree of n.
func textContent(n Node) string {
	w := textWriter{raw: true}
	w.children(n)
	return collapseSpace(w.b.String())
}

// innerText returns the collapsed text of the subtree of n, laid out: see
// [InnerText].
func innerText(n Node) string {
	if !rendered(n) {
		return textContent(n)
	}
	w := textWriter{}
	w.children(n)
	return collapseSpace(w.b.String())
}

// unrenderedTags are the elements whose content is never rendered: the
// document head, scripts, styles, templates, and noscript (as in a browser
// with scripting). The inner text leaves them out.
var unrenderedTags = map[string]bool{
	"head": true, "title": true, "script": true, "style": true, "template": true, "noscript": true,
}

// rendered reports whether neither n nor an ancestor is an unrendered element.
func rendered(n Node) bool {
	for cur := n; cur != nil; cur = cur.Parent() {
		if cur.Kind() == KindElement && unrenderedTags[cur.Tag()] {
			return false
		}
	}
	return true
}

// textWriter lays out the text of a subtree: text nodes, with line breaks
// around block-level boxes and for br elements, leaving out unrendered
// elements. With raw set, it writes the text nodes only, as textContent.
// With alt set, images contribute their alt text, as in accessible names.
// Elements for which redact reports true are replaced by a redaction marker
// when they contribute text, so that the result is the text without
// redaction with the sensitive parts replaced.
type textWriter struct {
	b      strings.Builder
	raw    bool
	alt    bool
	redact func(n Node) bool
}

func (w *textWriter) children(n Node) {
	for c := range n.Children() {
		w.node(c)
	}
}

func (w *textWriter) node(n Node) {
	switch n.Kind() {
	case KindText:
		w.b.WriteString(n.Data())
	case KindElement:
		if !w.raw && unrenderedTags[n.Tag()] {
			return
		}
		if w.redact != nil && w.redact(n) {
			plain := textWriter{raw: w.raw, alt: w.alt}
			plain.node(n)
			if collapseSpace(plain.b.String()) != "" {
				w.b.WriteString(" " + redactedValue + " ")
			}
			return
		}
		if w.raw {
			w.children(n)
			return
		}
		switch n.Tag() {
		case "br":
			w.b.WriteByte('\n')
			return
		case "img":
			if w.alt {
				alt, _ := n.Attr("alt")
				w.b.WriteString(" " + alt + " ")
				return
			}
		}
		block := blockBox(n)
		if block {
			w.b.WriteByte('\n')
		}
		w.children(n)
		if block {
			w.b.WriteByte('\n')
		}
	}
}

// collapseSpace trims the string and replaces every run of whitespace with
// one space.
func collapseSpace(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
