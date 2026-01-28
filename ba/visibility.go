package ba

import (
	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/html/attr"
)

// Visible creates a Bootstrap visible class.
func Visible() core.Node { return attr.Class("visible") }

// Invisible creates a Bootstrap invisible class.
func Invisible() core.Node { return attr.Class("invisible") }
