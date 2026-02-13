package ba

import (
	"github.com/protolambda/chord/core/attrib"
	"github.com/protolambda/chord/html/attr"
)

// Visible creates a Bootstrap visible class.
func Visible() attrib.Node { return attr.Class("visible") }

// Invisible creates a Bootstrap invisible class.
func Invisible() attrib.Node { return attr.Class("invisible") }
