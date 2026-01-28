package ba

import (
	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/html/attr"
)

// PositionStatic creates a Bootstrap position-static class.
func PositionStatic() core.Node { return attr.Class("position-static") }

// PositionRelative creates a Bootstrap position-relative class.
func PositionRelative() core.Node { return attr.Class("position-relative") }

// PositionAbsolute creates a Bootstrap position-absolute class.
func PositionAbsolute() core.Node { return attr.Class("position-absolute") }

// PositionFixed creates a Bootstrap position-fixed class.
func PositionFixed() core.Node { return attr.Class("position-fixed") }

// PositionSticky creates a Bootstrap position-sticky class.
func PositionSticky() core.Node { return attr.Class("position-sticky") }

// Top0 creates a Bootstrap top-0 class.
func Top0() core.Node { return attr.Class("top-0") }

// Top50 creates a Bootstrap top-50 class.
func Top50() core.Node { return attr.Class("top-50") }

// Top100 creates a Bootstrap top-100 class.
func Top100() core.Node { return attr.Class("top-100") }

// Bottom0 creates a Bootstrap bottom-0 class.
func Bottom0() core.Node { return attr.Class("bottom-0") }

// Bottom50 creates a Bootstrap bottom-50 class.
func Bottom50() core.Node { return attr.Class("bottom-50") }

// Bottom100 creates a Bootstrap bottom-100 class.
func Bottom100() core.Node { return attr.Class("bottom-100") }

// Start0 creates a Bootstrap start-0 class.
func Start0() core.Node { return attr.Class("start-0") }

// Start50 creates a Bootstrap start-50 class.
func Start50() core.Node { return attr.Class("start-50") }

// Start100 creates a Bootstrap start-100 class.
func Start100() core.Node { return attr.Class("start-100") }

// End0 creates a Bootstrap end-0 class.
func End0() core.Node { return attr.Class("end-0") }

// End50 creates a Bootstrap end-50 class.
func End50() core.Node { return attr.Class("end-50") }

// End100 creates a Bootstrap end-100 class.
func End100() core.Node { return attr.Class("end-100") }

// TranslateMiddle creates a Bootstrap translate-middle class.
func TranslateMiddle() core.Node { return attr.Class("translate-middle") }

// TranslateMiddleX creates a Bootstrap translate-middle-x class.
func TranslateMiddleX() core.Node { return attr.Class("translate-middle-x") }

// TranslateMiddleY creates a Bootstrap translate-middle-y class.
func TranslateMiddleY() core.Node { return attr.Class("translate-middle-y") }
