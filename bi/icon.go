// Package bi provides bootstrap-icons
package bi

import (
	"context"

	"github.com/protolambda/chord/core"
)

// Icon is a bootstrap icon enum value, and can be used directly as element.
// The icon renders as an `<i "bi bi-icon-name-here"></i>`, corresponding to the icon.
// This assumes the icons CSS is available.
type Icon string

var _ core.Node = Icon("")

func (i Icon) Eval(ctx context.Context) (core.Obj, error) {
	return core.Element("i", core.Attribute("class", "bi bi-"+string(i))).Eval(ctx)
}

// Icons can be found at:
// https://icons.getbootstrap.com/assets/font/bootstrap-icons.css
//
// To get the list from the CSS excerpt:
// Replace 1:
// \.bi-(.+)::before.+
// \u$1 Icon = "$1"
//
// Replace 2 (repeat, convert spaces/snake case):
// (.+)[ -](.+) Icon = (.+)
// $1\u$2 Icon = $3
//
// Replace 3 (fix num prefix):
// (^\d)
// Num$1
