// Package iframe provides the HTML iframe element.
package iframe

import "github.com/protolambda/chord/core"

// Iframe creates an iframe element.
func Iframe(opts ...core.Node) core.Node { return core.Element("iframe", opts...) }
