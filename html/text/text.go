package text

import (
	"html"

	"github.com/protolambda/chord/core"
)

// Text creates a node with text, no element wrapper. The text will be HTML-escaped.
func Text(v string) core.Node {
	return core.Raw(html.EscapeString(v))
}
