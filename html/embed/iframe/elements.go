// Package iframe provides the HTML iframe element.
package iframe

import (
	"github.com/protolambda/chord/core/attrib"
	"github.com/protolambda/chord/core/elem"
)

// Iframe creates an iframe element.
func Iframe(attrs ...attrib.Node) elem.Scope { return elem.New("iframe", attrs...) }
