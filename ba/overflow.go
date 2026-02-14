package ba

import (
	"github.com/protolambda/chord/core/attr"
)

// OverflowAuto creates a Bootstrap overflow-auto class.
func OverflowAuto() attr.Node { return attr.Class("overflow-auto") }

// OverflowHidden creates a Bootstrap overflow-hidden class.
func OverflowHidden() attr.Node { return attr.Class("overflow-hidden") }

// OverflowVisible creates a Bootstrap overflow-visible class.
func OverflowVisible() attr.Node { return attr.Class("overflow-visible") }

// OverflowScroll creates a Bootstrap overflow-scroll class.
func OverflowScroll() attr.Node { return attr.Class("overflow-scroll") }
