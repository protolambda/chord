package ct

import (
	"fmt"
	"strings"
)

const (
	maxListedNodes  = 8
	maxOutlineLines = 30
	maxValueLength  = 48
	redactedValue   = "[redacted]"
)

// listNodes formats one line per node with its location and summary,
// bounded to maxListedNodes.
func listNodes(nodes []Node, redact func(tag, key string) bool) string {
	var b strings.Builder
	for i, n := range nodes {
		if i == maxListedNodes {
			fmt.Fprintf(&b, "  ... %d more\n", len(nodes)-i)
			break
		}
		fmt.Fprintf(&b, "  %s %s\n", locationOf(n), describeNode(n, redact))
	}
	return strings.TrimRight(b.String(), "\n")
}

// outline formats the subtrees of the nodes as an indented tree, bounded to
// maxOutlineLines in total.
func outline(nodes []Node, redact func(tag, key string) bool) string {
	var b strings.Builder
	lines := 0
	truncated := false
	for _, n := range nodes {
		if !outlineNode(&b, n, 1, &lines, redact) {
			truncated = true
			break
		}
	}
	if truncated {
		b.WriteString("  ...\n")
	}
	return strings.TrimRight(b.String(), "\n")
}

func outlineNode(b *strings.Builder, n Node, depth int, lines *int, redact func(tag, key string) bool) bool {
	if n.Kind() == KindDocument {
		// The document root is invisible: its children start at this depth.
		for c := range n.Children() {
			if !outlineNode(b, c, depth, lines, redact) {
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
	b.WriteString(describeNode(n, redact))
	b.WriteByte('\n')
	for c := range n.Children() {
		if !outlineNode(b, c, depth+1, lines, redact) {
			return false
		}
	}
	return true
}

// describeNode summarizes a node on one line, e.g. `form#profile.card [method="post"]`.
func describeNode(n Node, redact func(tag, key string) bool) string {
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
			val = redactValue(n, key, val, redact)
			if val == "" {
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
		return fmt.Sprintf("%q", truncate(collapseSpace(n.Data())))
	case KindRaw:
		return fmt.Sprintf("raw(%q)", truncate(collapseSpace(n.Data())))
	case KindComment:
		return fmt.Sprintf("<!-- %s -->", truncate(collapseSpace(n.Data())))
	case KindDoctype:
		return "<!DOCTYPE " + n.Data() + ">"
	default:
		return n.Kind().String()
	}
}

func locationOf(n Node) string {
	if loc := n.Location(); loc != "" {
		return loc
	}
	return "(root)"
}

func truncate(v string) string {
	if len(v) <= maxValueLength {
		return v
	}
	return v[:maxValueLength] + "…"
}

// secretWords are attribute name and control name fragments that indicate a
// value which must not appear in test output.
var secretWords = []string{"password", "secret", "token", "csrf", "authorization", "cookie"}

func looksSecret(s string) bool {
	s = strings.ToLower(s)
	for _, w := range secretWords {
		if strings.Contains(s, w) {
			return true
		}
	}
	return false
}

// redactValue hides attribute values that look like secrets: any attribute
// whose name looks secret, and the value of a control whose type is password
// or whose name looks secret. The optional caller predicate adds more.
func redactValue(n Node, key, val string, redact func(tag, key string) bool) string {
	if val == "" {
		return val
	}
	if redact != nil && redact(n.Tag(), key) {
		return redactedValue
	}
	if looksSecret(key) {
		return redactedValue
	}
	if key == "value" {
		if typ, _ := n.Attr("type"); strings.EqualFold(typ, "password") {
			return redactedValue
		}
		if name, _ := n.Attr("name"); looksSecret(name) {
			return redactedValue
		}
	}
	return val
}
