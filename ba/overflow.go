package ba

import (
	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/html/attr"
)

// OverflowAuto creates a Bootstrap overflow-auto class.
func OverflowAuto() core.Node { return attr.Class("overflow-auto") }

// OverflowHidden creates a Bootstrap overflow-hidden class.
func OverflowHidden() core.Node { return attr.Class("overflow-hidden") }

// OverflowVisible creates a Bootstrap overflow-visible class.
func OverflowVisible() core.Node { return attr.Class("overflow-visible") }

// OverflowScroll creates a Bootstrap overflow-scroll class.
func OverflowScroll() core.Node { return attr.Class("overflow-scroll") }
