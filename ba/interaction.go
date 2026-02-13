package ba

import (
	"github.com/protolambda/chord/core/attrib"
	"github.com/protolambda/chord/html/attr"
)

// UserSelectAll creates a Bootstrap user-select-all class.
func UserSelectAll() attrib.Node { return attr.Class("user-select-all") }

// UserSelectAuto creates a Bootstrap user-select-auto class.
func UserSelectAuto() attrib.Node { return attr.Class("user-select-auto") }

// UserSelectNone creates a Bootstrap user-select-none class.
func UserSelectNone() attrib.Node { return attr.Class("user-select-none") }

// PENone creates a Bootstrap pe-none (pointer-events) class.
func PENone() attrib.Node { return attr.Class("pe-none") }

// PEAuto creates a Bootstrap pe-auto (pointer-events) class.
func PEAuto() attrib.Node { return attr.Class("pe-auto") }
