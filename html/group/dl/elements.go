// Package dl provides HTML definition list elements (dl, dt, dd).
package dl

import (
	"github.com/protolambda/chord/core/attr"
	"github.com/protolambda/chord/core/elem"
)

// DL creates a dl element.
func DL(attrs ...attr.Node) elem.Scope { return elem.New("dl", attrs...) }

// DT creates a dt element.
func DT(attrs ...attr.Node) elem.Scope { return elem.New("dt", attrs...) }

// DD creates a dd element.
func DD(attrs ...attr.Node) elem.Scope { return elem.New("dd", attrs...) }
