// Package dl provides HTML definition list elements (dl, dt, dd).
package dl

import "github.com/protolambda/chord/core"

// DL creates a dl element.
func DL(opts ...core.Node) core.Node { return core.Element("dl", opts...) }

// DT creates a dt element.
func DT(opts ...core.Node) core.Node { return core.Element("dt", opts...) }

// DD creates a dd element.
func DD(opts ...core.Node) core.Node { return core.Element("dd", opts...) }
