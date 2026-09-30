package ct

import "strings"

// display is the part of the CSS display of an element that matters for
// text: whether its box is block-level, so that its text is separated from
// the text around it, and whether it is a flex or grid container, whose
// children are laid out as blocks (blockified).
type display struct {
	block     bool
	container bool
}

var (
	displayInline     = display{}
	displayBlock      = display{block: true}
	displayFlex       = display{block: true, container: true}
	displayInlineFlex = display{container: true}
)

// blockBox reports whether the element n is laid out as a block-level box,
// so that its text starts and ends a line in the inner text.
func blockBox(n Node) bool {
	if p := n.Parent(); p != nil && p.Kind() == KindElement && displayOf(p).container {
		return true
	}
	return displayOf(n).block
}

// displayOf approximates the display of an element, in order of
// precedence: a Bootstrap display utility class (these are !important),
// the display declared in the style attribute, a Bootstrap helper or
// component class, and finally the default display of the tag in the HTML
// rendering rules. Only unprefixed utilities count, as at the smallest
// viewport. Displays that hide an element (none) or remove its box
// (contents) are ignored: text queries do not model visibility.
func displayOf(n Node) display {
	var component, utility *display
	if classes, ok := n.Attr("class"); ok {
		for c := range strings.FieldsSeq(classes) {
			if d, ok := utilityDisplays[c]; ok {
				utility = &d
			} else if d, ok := componentDisplays[c]; ok && component == nil {
				component = &d
			}
		}
	}
	if utility != nil {
		return *utility
	}
	if style, ok := n.Attr("style"); ok {
		if d, ok := styleDisplay(style); ok {
			return d
		}
	}
	if component != nil {
		return *component
	}
	if blockTags[n.Tag()] {
		return displayBlock
	}
	return displayInline
}

// styleDisplay returns the display declared last in an inline style.
func styleDisplay(style string) (display, bool) {
	var out display
	found := false
	for decl := range strings.SplitSeq(style, ";") {
		prop, value, ok := strings.Cut(decl, ":")
		if !ok || !strings.EqualFold(strings.TrimSpace(prop), "display") {
			continue
		}
		value = strings.ToLower(value)
		value = strings.TrimSpace(strings.TrimSuffix(strings.TrimSpace(value), "!important"))
		if d, ok := parseDisplay(value); ok {
			out, found = d, true
		}
	}
	return out, found
}

// parseDisplay reads a CSS display value, in the one- or two-keyword form.
func parseDisplay(value string) (display, bool) {
	keywords := strings.Fields(value)
	if len(keywords) == 0 {
		return display{}, false
	}
	var d display
	d.block = true
	for _, k := range keywords {
		switch k {
		case "inline", "ruby":
			d.block = false
		case "inline-block", "inline-table", "inline-list-item":
			d.block = false
		case "inline-flex", "inline-grid":
			d.block, d.container = false, true
		case "flex", "grid":
			d.container = true
		case "block", "flow", "flow-root", "list-item", "table", "run-in",
			"table-row-group", "table-header-group", "table-footer-group", "table-row",
			"table-cell", "table-column-group", "table-column", "table-caption":
		default:
			// none, contents, global keywords, and unknown values.
			return display{}, false
		}
	}
	return d, true
}

// blockTags are the elements that the HTML rendering rules display as
// block, list-item, or table parts (including cells and rows, which the
// inner text separates too), plus the document elements.
var blockTags = map[string]bool{
	"address": true, "article": true, "aside": true, "blockquote": true, "body": true,
	"caption": true, "center": true, "col": true, "colgroup": true, "dd": true,
	"details": true, "dialog": true, "dir": true, "div": true, "dl": true,
	"dt": true, "fieldset": true, "figcaption": true, "figure": true, "footer": true,
	"form": true, "h1": true, "h2": true, "h3": true, "h4": true,
	"h5": true, "h6": true, "head": true, "header": true, "hgroup": true,
	"hr": true, "html": true, "legend": true, "li": true, "listing": true,
	"main": true, "menu": true, "nav": true, "ol": true, "optgroup": true,
	"option": true, "p": true, "plaintext": true, "pre": true, "search": true,
	"section": true, "summary": true, "table": true, "tbody": true, "td": true,
	"tfoot": true, "th": true, "thead": true, "tr": true, "ul": true,
	"xmp": true,
}

// utilityDisplays are the unprefixed Bootstrap 5.3 display utilities (see
// package ba). d-none is left out: text queries do not model visibility.
var utilityDisplays = map[string]display{
	"d-inline": displayInline, "d-inline-block": displayInline, "d-block": displayBlock,
	"d-grid": displayFlex, "d-inline-grid": displayInlineFlex, "d-table": displayBlock,
	"d-table-row": displayBlock, "d-table-cell": displayBlock, "d-flex": displayFlex,
	"d-inline-flex": displayInlineFlex,
}

// componentDisplays are the Bootstrap 5.3 helper and component classes that
// make an element a flex container, or that display a link or span as a
// block: the ones whose content is commonly inline elements.
var componentDisplays = map[string]display{
	"hstack": displayFlex, "vstack": displayFlex,
	"btn-group": displayInlineFlex, "btn-group-vertical": displayInlineFlex,
	"btn-toolbar": displayFlex, "input-group": displayFlex, "nav": displayFlex,
	"navbar": displayFlex, "navbar-nav": displayFlex, "pagination": displayFlex,
	"breadcrumb": displayFlex, "list-group": displayFlex, "card": displayFlex,
	"modal-header": displayFlex, "modal-footer": displayFlex, "toast-header": displayFlex,
	"offcanvas-header": displayFlex, "row": displayFlex,
	"nav-link": displayBlock, "dropdown-item": displayBlock, "dropdown-header": displayBlock,
	"dropdown-item-text": displayBlock, "list-group-item": displayBlock,
	"invalid-feedback": displayBlock, "valid-feedback": displayBlock,
}
