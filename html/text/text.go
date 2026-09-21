package text

import (
	"github.com/protolambda/chord/core/elem"
)

// Text creates a node with text, no element wrapper. The text will be HTML-escaped.
// It is the same as [elem.Text].
func Text(v string) elem.Node {
	return elem.Text(v)
}
