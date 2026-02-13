package ba

import (
	"github.com/protolambda/chord/core/attrib"
	"github.com/protolambda/chord/html/attr"
)

// Shadow creates a Bootstrap shadow class.
func Shadow() attrib.Node { return attr.Class("shadow") }

// ShadowSM creates a Bootstrap shadow-sm class.
func ShadowSM() attrib.Node { return attr.Class("shadow-sm") }

// ShadowLG creates a Bootstrap shadow-lg class.
func ShadowLG() attrib.Node { return attr.Class("shadow-lg") }

// ShadowNone creates a Bootstrap shadow-none class.
func ShadowNone() attrib.Node { return attr.Class("shadow-none") }
