package ct

import (
	"fmt"
	"strings"
)

// Text matches elements whose own text equals the value. Own text is the
// concatenation of the element's direct text children, with whitespace
// collapsed, so that a wrapper does not match the text of a nested element.
// Text does not model visibility.
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

// TextContent matches elements whose whole subtree text equals the value,
// with whitespace collapsed. Raw content and comments do not count as text.
func TextContent(value string) Query {
	return Query{
		desc:  fmt.Sprintf("textContent(%q)", value),
		match: func(n Node) (bool, error) { return textContent(n) == value, nil },
	}
}

// TextContentMatches matches elements whose subtree text satisfies the matcher.
func TextContentMatches(value ValueMatch) Query {
	return Query{
		desc:  fmt.Sprintf("textContent(%s)", value),
		match: func(n Node) (bool, error) { return value.Match(textContent(n)), nil },
	}
}

// ownText returns the collapsed text of the direct text children of n.
func ownText(n Node) string {
	var b strings.Builder
	for c := range n.Children() {
		if c.Kind() == KindText {
			b.WriteString(c.Data())
		}
	}
	return collapseSpace(b.String())
}

// textContent returns the collapsed text of all text nodes in the subtree of n.
func textContent(n Node) string {
	var b strings.Builder
	appendText(&b, n)
	return collapseSpace(b.String())
}

func appendText(b *strings.Builder, n Node) {
	if n.Kind() == KindText {
		b.WriteString(n.Data())
	}
	for c := range n.Children() {
		appendText(b, c)
	}
}

// collapseSpace trims the string and replaces every run of whitespace with
// one space.
func collapseSpace(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
