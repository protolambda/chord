package ba

import (
	"github.com/protolambda/chord/core/attrib"
	"github.com/protolambda/chord/html/attr"
)

// OverflowAuto creates a Bootstrap overflow-auto class.
func OverflowAuto() attrib.Node { return attr.Class("overflow-auto") }

// OverflowHidden creates a Bootstrap overflow-hidden class.
func OverflowHidden() attrib.Node { return attr.Class("overflow-hidden") }

// OverflowVisible creates a Bootstrap overflow-visible class.
func OverflowVisible() attrib.Node { return attr.Class("overflow-visible") }

// OverflowScroll creates a Bootstrap overflow-scroll class.
func OverflowScroll() attrib.Node { return attr.Class("overflow-scroll") }
