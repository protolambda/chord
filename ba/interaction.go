package ba

import (
	"github.com/protolambda/chord/core/attr"
)

// UserSelectAll creates a Bootstrap user-select-all class.
func UserSelectAll() attr.Node { return attr.Class("user-select-all") }

// UserSelectAuto creates a Bootstrap user-select-auto class.
func UserSelectAuto() attr.Node { return attr.Class("user-select-auto") }

// UserSelectNone creates a Bootstrap user-select-none class.
func UserSelectNone() attr.Node { return attr.Class("user-select-none") }

// PENone creates a Bootstrap pe-none (pointer-events) class.
func PENone() attr.Node { return attr.Class("pe-none") }

// PEAuto creates a Bootstrap pe-auto (pointer-events) class.
func PEAuto() attr.Node { return attr.Class("pe-auto") }
