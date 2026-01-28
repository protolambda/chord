// Package div provides general grouping elements like div, p, hr, pre, and blockquote.
package div

import "github.com/protolambda/chord/core"

// Div creates a div element.
func Div(opts ...core.Node) core.Node { return core.Element("div", opts...) }
