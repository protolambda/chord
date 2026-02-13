package ba

import (
	"github.com/protolambda/chord/core/attrib"
	"github.com/protolambda/chord/html/attr"
)

// DNone creates a Bootstrap d-none class.
func DNone() attrib.Node { return attr.Class("d-none") }

// DInline creates a Bootstrap d-inline class.
func DInline() attrib.Node { return attr.Class("d-inline") }

// DInlineBlock creates a Bootstrap d-inline-block class.
func DInlineBlock() attrib.Node { return attr.Class("d-inline-block") }

// DBlock creates a Bootstrap d-block class.
func DBlock() attrib.Node { return attr.Class("d-block") }

// DGrid creates a Bootstrap d-grid class.
func DGrid() attrib.Node { return attr.Class("d-grid") }

// DFlex creates a Bootstrap d-flex class.
func DFlex() attrib.Node { return attr.Class("d-flex") }

// DInlineFlex creates a Bootstrap d-inline-flex class.
func DInlineFlex() attrib.Node { return attr.Class("d-inline-flex") }

// DTable creates a Bootstrap d-table class.
func DTable() attrib.Node { return attr.Class("d-table") }

// DTableRow creates a Bootstrap d-table-row class.
func DTableRow() attrib.Node { return attr.Class("d-table-row") }

// DTableCell creates a Bootstrap d-table-cell class.
func DTableCell() attrib.Node { return attr.Class("d-table-cell") }
