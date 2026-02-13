package text

import (
	"html"

	"github.com/protolambda/chord/core/elem"
)

// Text creates a node with text, no element wrapper. The text will be HTML-escaped.
func Text(v string) elem.Node {
	return elem.Raw(html.EscapeString(v))
}
