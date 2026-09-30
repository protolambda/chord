package ct

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

const (
	maxListedNodes  = 8
	maxOutlineLines = 30
	maxValueLength  = 48
	redactedValue   = "[redacted]"
)

// listNodes formats one line per node with its location and summary,
// bounded to maxListedNodes.
func listNodes(nodes []Node, o *options) string {
	var b strings.Builder
	for i, n := range nodes {
		if i == maxListedNodes {
			fmt.Fprintf(&b, "  ... %d more\n", len(nodes)-i)
			break
		}
		fmt.Fprintf(&b, "  %s %s\n", locationOf(n), describeNode(n, o))
	}
	return strings.TrimRight(b.String(), "\n")
}

// outline formats the subtrees of the nodes as an indented tree, bounded to
// maxOutlineLines in total.
func outline(nodes []Node, o *options) string {
	var b strings.Builder
	lines := 0
	truncated := false
	for _, n := range nodes {
		if !outlineNode(&b, n, 1, &lines, o) {
			truncated = true
			break
		}
	}
	if truncated {
		b.WriteString("  ...\n")
	}
	return strings.TrimRight(b.String(), "\n")
}

func outlineNode(b *strings.Builder, n Node, depth int, lines *int, o *options) bool {
	if n.Kind() == KindDocument {
		// The document root is invisible: its children start at this depth.
		for c := range n.Children() {
			if !outlineNode(b, c, depth, lines, o) {
				return false
			}
		}
		return true
	}
	if *lines == maxOutlineLines {
		return false
	}
	*lines++
	b.WriteString(strings.Repeat("  ", depth))
	b.WriteString(describeNode(n, o))
	b.WriteByte('\n')
	for c := range n.Children() {
		if !outlineNode(b, c, depth+1, lines, o) {
			return false
		}
	}
	return true
}

// describeNode summarizes a node on one line, e.g. `form#profile.card [method="post"]`.
func describeNode(n Node, o *options) string {
	switch n.Kind() {
	case KindElement:
		var b strings.Builder
		b.WriteString(n.Tag())
		if id, ok := n.Attr("id"); ok && id != "" {
			b.WriteString("#" + id)
		}
		if classes, ok := n.Attr("class"); ok {
			for c := range strings.FieldsSeq(classes) {
				b.WriteString("." + c)
			}
		}
		var attrs []string
		for key, val := range n.Attrs() {
			if key == "id" || key == "class" {
				continue
			}
			val = o.redactValue(n, key, val)
			if val == "" && boolAttr(n, key) {
				attrs = append(attrs, key)
			} else {
				attrs = append(attrs, fmt.Sprintf("%s=%q", key, truncate(val)))
			}
		}
		if len(attrs) > 0 {
			b.WriteString(" [" + strings.Join(attrs, " ") + "]")
		}
		return b.String()
	case KindText:
		if collapseSpace(n.Data()) != "" && o.inSensitive(n) {
			return redactedValue
		}
		return fmt.Sprintf("%q", truncate(collapseSpace(n.Data())))
	case KindRaw:
		if o.inSensitive(n) {
			return "raw(" + redactedValue + ")"
		}
		return fmt.Sprintf("raw(%q)", truncate(collapseSpace(n.Data())))
	case KindComment:
		if o.inSensitive(n) {
			return "<!-- " + redactedValue + " -->"
		}
		return fmt.Sprintf("<!-- %s -->", truncate(collapseSpace(n.Data())))
	case KindDoctype:
		return "<!DOCTYPE " + n.Data() + ">"
	default:
		return n.Kind().String()
	}
}

// boolAttr reports whether the node is known to have the attribute without
// a value.
func boolAttr(n Node, key string) bool {
	b, ok := n.(BoolAttrNode)
	return ok && b.BoolAttr(key)
}

func locationOf(n Node) string {
	if loc := n.Location(); loc != "" {
		return loc
	}
	return "(root)"
}

// truncate shortens v to at most maxValueLength bytes, at a character
// boundary, and marks the cut.
func truncate(v string) string {
	if len(v) <= maxValueLength {
		return v
	}
	cut := maxValueLength
	for cut > 0 && !utf8.RuneStart(v[cut]) {
		cut--
	}
	return v[:cut] + "…"
}
