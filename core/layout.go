package core

// blockLevel reports whether an element with tag, inside an element with
// parentTag, is laid out as a block by default. Browsers render no
// whitespace between two such elements, nor between one of them and the
// start or end of its parent, so indentation can go there without changing
// the page. Everything in head is included, since head is not rendered.
func blockLevel(tag, parentTag string) bool {
	return blockTags[lowerASCII(tag)] || lowerASCII(parentTag) == "head"
}

// blockTags are the elements that the HTML rendering rules display as block,
// list-item, or table parts, plus the document elements. Elements that are
// not displayed (such as script and template) are excluded: two of them
// between inline content would add a visible space.
var blockTags = map[string]bool{
	"address": true, "article": true, "aside": true, "blockquote": true, "body": true,
	"caption": true, "col": true, "colgroup": true, "dd": true, "details": true,
	"div": true, "dl": true, "dt": true, "fieldset": true, "figcaption": true,
	"figure": true, "footer": true, "form": true, "h1": true, "h2": true,
	"h3": true, "h4": true, "h5": true, "h6": true, "head": true,
	"header": true, "hgroup": true, "hr": true, "html": true, "legend": true,
	"li": true, "main": true, "menu": true, "nav": true, "ol": true,
	"optgroup": true, "option": true, "p": true, "pre": true, "search": true,
	"section": true, "summary": true, "table": true, "tbody": true, "td": true,
	"tfoot": true, "th": true, "thead": true, "tr": true, "ul": true,
}

// preformatted reports whether whitespace inside an element with tag is
// part of its content, so that indentation must not be added inside it:
// preformatted text, form values, script and style code, and the title.
func preformatted(tag string) bool {
	switch lowerASCII(tag) {
	case "pre", "listing", "plaintext", "xmp", "textarea", "script", "style", "title":
		return true
	}
	return false
}

// dropsLeadingNewline reports whether the HTML parser drops a newline that
// directly follows the start tag of an element with tag.
func dropsLeadingNewline(tag string) bool {
	switch lowerASCII(tag) {
	case "pre", "listing", "textarea":
		return true
	}
	return false
}

// lowerASCII lowercases ASCII letters only, as HTML does for tag names.
func lowerASCII(s string) string {
	for i := 0; i < len(s); i++ {
		if 'A' <= s[i] && s[i] <= 'Z' {
			b := []byte(s)
			for j := i; j < len(b); j++ {
				if 'A' <= b[j] && b[j] <= 'Z' {
					b[j] += 'a' - 'A'
				}
			}
			return string(b)
		}
	}
	return s
}
