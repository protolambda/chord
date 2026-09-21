package ct

import (
	"fmt"
	"strconv"
	"strings"
)

// Label matches elements associated with a label whose text equals the value.
//
// The supported label forms are: an aria-label attribute, the elements
// referenced by aria-labelledby, label elements whose for attribute matches
// the element id, and an enclosing label element. Any of them matching is
// enough. A reference to a missing id is a query error, not a non-match.
func Label(value string) Query {
	return Query{
		desc:  fmt.Sprintf("label(%q)", value),
		match: labelMatcher(Exact(value)),
	}
}

// LabelMatches is [Label] with a value matcher.
func LabelMatches(value ValueMatch) Query {
	return Query{
		desc:  fmt.Sprintf("label(%s)", value),
		match: labelMatcher(value),
	}
}

func labelMatcher(value ValueMatch) func(Node) (bool, error) {
	return func(n Node) (bool, error) {
		labels, err := labelTexts(n)
		if err != nil {
			return false, err
		}
		for _, l := range labels {
			if value.Match(l.text) {
				return true, nil
			}
		}
		return false, nil
	}
}

// Alt matches elements with exactly the given alt attribute.
func Alt(value string) Query {
	return Query{
		desc:  fmt.Sprintf("alt(%q)", value),
		match: func(n Node) (bool, error) { return attrEquals(n, "alt", value), nil },
	}
}

// AltMatches matches elements whose alt attribute satisfies the matcher.
func AltMatches(value ValueMatch) Query {
	return Query{
		desc: fmt.Sprintf("alt(%s)", value),
		match: func(n Node) (bool, error) {
			v, ok := n.Attr("alt")
			return ok && value.Match(v), nil
		},
	}
}

// RoleOption refines a [Role] query.
type RoleOption struct {
	desc  string
	name  *ValueMatch
	level *int
}

func (o RoleOption) String() string {
	return o.desc
}

// Named requires the accessible name to equal the value.
func Named(name string) RoleOption {
	m := Exact(name)
	return RoleOption{desc: fmt.Sprintf("named(%q)", name), name: &m}
}

// NameMatches requires the accessible name to satisfy the matcher.
func NameMatches(name ValueMatch) RoleOption {
	return RoleOption{desc: fmt.Sprintf("named(%s)", name), name: &name}
}

// Level requires a heading level: the digit of h1 to h6, or aria-level.
func Level(level int) RoleOption {
	return RoleOption{desc: fmt.Sprintf("level(%d)", level), level: &level}
}

// Role matches elements with the given role, optionally refined by
// accessible name and heading level. Role names are case-insensitive.
//
// An explicit role attribute (its first token) wins. Otherwise a documented
// subset of implicit HTML roles applies, see [ImplicitRole]. The accessible
// name is computed by [AccessibleName]. Unsupported elements have no role
// and never match; use structural queries for them.
func Role(role string, options ...RoleOption) Query {
	role = strings.ToLower(role)
	parts := []string{fmt.Sprintf("%q", role)}
	for _, o := range options {
		parts = append(parts, o.desc)
	}
	return Query{
		desc: "role(" + strings.Join(parts, ", ") + ")",
		match: func(n Node) (bool, error) {
			if roleOf(n) != role {
				return false, nil
			}
			for _, o := range options {
				if o.name != nil {
					name, err := AccessibleName(n)
					if err != nil {
						return false, err
					}
					if !o.name.Match(name) {
						return false, nil
					}
				}
				if o.level != nil && headingLevel(n) != *o.level {
					return false, nil
				}
			}
			return true, nil
		},
	}
}

// roleOf returns the explicit role, or the implicit role.
func roleOf(n Node) string {
	if explicit, ok := n.Attr("role"); ok {
		if tokens := strings.Fields(explicit); len(tokens) > 0 {
			return strings.ToLower(tokens[0])
		}
	}
	return ImplicitRole(n)
}

// ImplicitRole returns the implicit ARIA role of an element for the
// supported subset of HTML, or "" when the element has no supported role.
//
// Supported: a/area with href (link), button and button-like inputs
// (button), checkbox, radio, range (slider), number (spinbutton), search
// (searchbox), text/email/tel/url inputs and textarea (textbox), select
// (combobox, or listbox when multiple or sized), option, h1-h6 (heading),
// nav, main, header (banner), footer (contentinfo), aside (complementary),
// article, section (region), form, search, ul/ol/menu (list), li
// (listitem), table, thead/tbody/tfoot (rowgroup), tr (row), td (cell), th
// (columnheader, or rowheader with scope=row), img (img, or presentation
// with empty alt), dialog, details/fieldset/optgroup (group), hr
// (separator), progress (progressbar), meter, output (status), p
// (paragraph), figure, caption, blockquote, code, em (emphasis), strong,
// time, mark, del (deletion), ins (insertion), sub/sup (subscript,
// superscript), dfn (term), and datalist (listbox).
func ImplicitRole(n Node) string {
	if n.Kind() != KindElement {
		return ""
	}
	switch tag := n.Tag(); tag {
	case "a", "area":
		if _, ok := n.Attr("href"); ok {
			return "link"
		}
		return ""
	case "input":
		return inputRole(n)
	case "textarea":
		return "textbox"
	case "select":
		if _, multiple := n.Attr("multiple"); multiple {
			return "listbox"
		}
		if size, ok := n.Attr("size"); ok {
			if v, err := strconv.Atoi(strings.TrimSpace(size)); err == nil && v > 1 {
				return "listbox"
			}
		}
		return "combobox"
	case "h1", "h2", "h3", "h4", "h5", "h6":
		return "heading"
	case "th":
		if scope, _ := n.Attr("scope"); strings.EqualFold(scope, "row") {
			return "rowheader"
		}
		return "columnheader"
	case "img":
		if alt, ok := n.Attr("alt"); ok && alt == "" {
			return "presentation"
		}
		return "img"
	default:
		return implicitRoles[tag]
	}
}

var implicitRoles = map[string]string{
	"button": "button", "option": "option", "nav": "navigation", "main": "main",
	"header": "banner", "footer": "contentinfo", "aside": "complementary",
	"article": "article", "section": "region", "form": "form", "search": "search",
	"ul": "list", "ol": "list", "menu": "list", "li": "listitem",
	"table": "table", "thead": "rowgroup", "tbody": "rowgroup", "tfoot": "rowgroup",
	"tr": "row", "td": "cell", "dialog": "dialog", "details": "group",
	"fieldset": "group", "optgroup": "group", "hr": "separator",
	"progress": "progressbar", "meter": "meter", "output": "status",
	"p": "paragraph", "figure": "figure", "caption": "caption",
	"blockquote": "blockquote", "code": "code", "em": "emphasis", "strong": "strong",
	"time": "time", "mark": "mark", "del": "deletion", "ins": "insertion",
	"sub": "subscript", "sup": "superscript", "dfn": "term", "datalist": "listbox",
}

func inputType(n Node) string {
	typ, ok := n.Attr("type")
	if !ok || strings.TrimSpace(typ) == "" {
		return "text"
	}
	return strings.ToLower(strings.TrimSpace(typ))
}

func inputRole(n Node) string {
	switch inputType(n) {
	case "button", "submit", "reset", "image":
		return "button"
	case "checkbox":
		return "checkbox"
	case "radio":
		return "radio"
	case "range":
		return "slider"
	case "number":
		return "spinbutton"
	case "search":
		return "searchbox"
	case "text", "email", "tel", "url":
		if _, ok := n.Attr("list"); ok {
			return "combobox"
		}
		return "textbox"
	default:
		return ""
	}
}

func headingLevel(n Node) int {
	if v, ok := n.Attr("aria-level"); ok {
		if level, err := strconv.Atoi(strings.TrimSpace(v)); err == nil {
			return level
		}
	}
	if tag := n.Tag(); len(tag) == 2 && tag[0] == 'h' && tag[1] >= '1' && tag[1] <= '6' {
		return int(tag[1] - '0')
	}
	return 0
}

// nameFromContentRoles are the roles whose accessible name may come from
// their content when nothing more specific applies.
var nameFromContentRoles = map[string]bool{
	"button": true, "cell": true, "checkbox": true, "columnheader": true,
	"gridcell": true, "heading": true, "link": true, "menuitem": true,
	"menuitemcheckbox": true, "menuitemradio": true, "option": true,
	"radio": true, "row": true, "rowheader": true, "sectionhead": true,
	"switch": true, "tab": true, "tooltip": true, "treeitem": true,
}

// AccessibleName computes the accessible name of an element for a
// documented subset of the accessible name computation, in this order:
// aria-labelledby (the texts of the referenced elements), aria-label,
// native labels (associated label elements for form controls, the value of
// button-like inputs, alt of image inputs and images, the caption of a
// table, the legend of a fieldset, the figcaption of a figure), the text
// content for roles that allow naming from content, and finally the title
// attribute. Whitespace is collapsed. A reference to a missing id is an
// error.
func AccessibleName(n Node) (string, error) {
	if ids, ok := n.Attr("aria-labelledby"); ok {
		var parts []string
		for id := range strings.FieldsSeq(ids) {
			ref := elementByID(root(n), id)
			if ref == nil {
				return "", fmt.Errorf("aria-labelledby references missing id %q", id)
			}
			if label, ok := ref.Attr("aria-label"); ok && collapseSpace(label) != "" {
				parts = append(parts, collapseSpace(label))
			} else {
				parts = append(parts, contentText(ref))
			}
		}
		if name := collapseSpace(strings.Join(parts, " ")); name != "" {
			return name, nil
		}
	}
	if label, ok := n.Attr("aria-label"); ok && collapseSpace(label) != "" {
		return collapseSpace(label), nil
	}
	if name, err := nativeName(n); err != nil || name != "" {
		return name, err
	}
	if nameFromContentRoles[roleOf(n)] {
		if name := contentText(n); name != "" {
			return name, nil
		}
	}
	if title, ok := n.Attr("title"); ok {
		return collapseSpace(title), nil
	}
	return "", nil
}

func nativeName(n Node) (string, error) {
	switch tag := n.Tag(); tag {
	case "input":
		switch inputType(n) {
		case "button", "submit", "reset":
			if v, ok := n.Attr("value"); ok && collapseSpace(v) != "" {
				return collapseSpace(v), nil
			}
			switch inputType(n) {
			case "submit":
				return "Submit", nil
			case "reset":
				return "Reset", nil
			}
			return "", nil
		case "image":
			alt, _ := n.Attr("alt")
			return collapseSpace(alt), nil
		case "hidden":
			return "", nil
		}
		return labelName(n)
	case "textarea", "select", "meter", "output", "progress":
		return labelName(n)
	case "button":
		return labelName(n)
	case "img", "area":
		alt, _ := n.Attr("alt")
		return collapseSpace(alt), nil
	case "table":
		return childText(n, "caption"), nil
	case "fieldset":
		return childText(n, "legend"), nil
	case "figure":
		return childText(n, "figcaption"), nil
	}
	return "", nil
}

// labelName joins the texts of associated label elements.
func labelName(n Node) (string, error) {
	labels, err := labelTexts(n)
	if err != nil {
		return "", err
	}
	// aria-label and aria-labelledby were already considered by the caller;
	// keep only label elements here.
	var parts []string
	for _, l := range labels {
		if l.element {
			parts = append(parts, l.text)
		}
	}
	return collapseSpace(strings.Join(parts, " ")), nil
}

// labelText is one associated label of an element.
type labelText struct {
	text    string
	element bool // from a label element rather than an aria attribute
}

// labelTexts returns all label texts associated with n, in the order:
// aria-label, aria-labelledby, label[for], enclosing label.
func labelTexts(n Node) ([]labelText, error) {
	var out []labelText
	if v, ok := n.Attr("aria-label"); ok {
		out = append(out, labelText{text: collapseSpace(v)})
	}
	if ids, ok := n.Attr("aria-labelledby"); ok {
		var parts []string
		for id := range strings.FieldsSeq(ids) {
			ref := elementByID(root(n), id)
			if ref == nil {
				return nil, fmt.Errorf("aria-labelledby references missing id %q", id)
			}
			parts = append(parts, contentText(ref))
		}
		out = append(out, labelText{text: collapseSpace(strings.Join(parts, " "))})
	}
	if id, ok := n.Attr("id"); ok && id != "" {
		for d := range descendants(root(n)) {
			if d.Kind() == KindElement && d.Tag() == "label" && attrEquals(d, "for", id) {
				out = append(out, labelText{text: contentText(d), element: true})
			}
		}
	}
	for p := n.Parent(); p != nil; p = p.Parent() {
		if p.Kind() == KindElement && p.Tag() == "label" {
			if _, explicit := p.Attr("for"); !explicit {
				out = append(out, labelText{text: contentText(p), element: true})
			}
			break
		}
	}
	return out, nil
}

// contentText is the text of a subtree for naming purposes: text nodes, and
// the alt of images, with whitespace collapsed.
func contentText(n Node) string {
	var b strings.Builder
	appendContentText(&b, n)
	return collapseSpace(b.String())
}

func appendContentText(b *strings.Builder, n Node) {
	switch n.Kind() {
	case KindText:
		b.WriteString(n.Data())
		return
	case KindElement:
		if n.Tag() == "img" {
			alt, _ := n.Attr("alt")
			b.WriteString(" " + alt + " ")
			return
		}
	}
	for c := range n.Children() {
		appendContentText(b, c)
	}
}

func childText(n Node, tag string) string {
	for c := range n.Children() {
		if c.Kind() == KindElement && c.Tag() == tag {
			return contentText(c)
		}
	}
	return ""
}

// elementByID returns the first element with the id in the tree, or nil.
func elementByID(root Node, id string) Node {
	for d := range descendants(root) {
		if d.Kind() == KindElement && attrEquals(d, "id", id) {
			return d
		}
	}
	return nil
}
