package ba

import (
	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/html/attr"
)

// W25 creates a Bootstrap w-25 class.
func W25() core.Node { return attr.Class("w-25") }

// W50 creates a Bootstrap w-50 class.
func W50() core.Node { return attr.Class("w-50") }

// W75 creates a Bootstrap w-75 class.
func W75() core.Node { return attr.Class("w-75") }

// W100 creates a Bootstrap w-100 class.
func W100() core.Node { return attr.Class("w-100") }

// WAuto creates a Bootstrap w-auto class.
func WAuto() core.Node { return attr.Class("w-auto") }

// H25 creates a Bootstrap h-25 class.
func H25() core.Node { return attr.Class("h-25") }

// H50 creates a Bootstrap h-50 class.
func H50() core.Node { return attr.Class("h-50") }

// H75 creates a Bootstrap h-75 class.
func H75() core.Node { return attr.Class("h-75") }

// H100 creates a Bootstrap h-100 class.
func H100() core.Node { return attr.Class("h-100") }

// HAuto creates a Bootstrap h-auto class.
func HAuto() core.Node { return attr.Class("h-auto") }

// MW100 creates a Bootstrap mw-100 class.
func MW100() core.Node { return attr.Class("mw-100") }

// MH100 creates a Bootstrap mh-100 class.
func MH100() core.Node { return attr.Class("mh-100") }

// VWMin100 creates a Bootstrap min-vw-100 class.
func VWMin100() core.Node { return attr.Class("min-vw-100") }

// VHMin100 creates a Bootstrap min-vh-100 class.
func VHMin100() core.Node { return attr.Class("min-vh-100") }
