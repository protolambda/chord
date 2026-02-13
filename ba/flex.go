package ba

import (
	"fmt"

	"github.com/protolambda/chord/core/attrib"
	"github.com/protolambda/chord/html/attr"
)

// FlexRow creates a Bootstrap flex-row class.
func FlexRow() attrib.Node { return attr.Class("flex-row") }

// FlexColumn creates a Bootstrap flex-column class.
func FlexColumn() attrib.Node { return attr.Class("flex-column") }

// FlexRowReverse creates a Bootstrap flex-row-reverse class.
func FlexRowReverse() attrib.Node { return attr.Class("flex-row-reverse") }

// FlexColumnReverse creates a Bootstrap flex-column-reverse class.
func FlexColumnReverse() attrib.Node { return attr.Class("flex-column-reverse") }

// FlexWrap creates a Bootstrap flex-wrap class.
func FlexWrap() attrib.Node { return attr.Class("flex-wrap") }

// FlexNowrap creates a Bootstrap flex-nowrap class.
func FlexNowrap() attrib.Node { return attr.Class("flex-nowrap") }

// FlexWrapReverse creates a Bootstrap flex-wrap-reverse class.
func FlexWrapReverse() attrib.Node { return attr.Class("flex-wrap-reverse") }

// JustifyContentStart creates a Bootstrap justify-content-start class.
func JustifyContentStart() attrib.Node { return attr.Class("justify-content-start") }

// JustifyContentEnd creates a Bootstrap justify-content-end class.
func JustifyContentEnd() attrib.Node { return attr.Class("justify-content-end") }

// JustifyContentCenter creates a Bootstrap justify-content-center class.
func JustifyContentCenter() attrib.Node { return attr.Class("justify-content-center") }

// JustifyContentBetween creates a Bootstrap justify-content-between class.
func JustifyContentBetween() attrib.Node { return attr.Class("justify-content-between") }

// JustifyContentAround creates a Bootstrap justify-content-around class.
func JustifyContentAround() attrib.Node { return attr.Class("justify-content-around") }

// JustifyContentEvenly creates a Bootstrap justify-content-evenly class.
func JustifyContentEvenly() attrib.Node { return attr.Class("justify-content-evenly") }

// AlignItemsStart creates a Bootstrap align-items-start class.
func AlignItemsStart() attrib.Node { return attr.Class("align-items-start") }

// AlignItemsEnd creates a Bootstrap align-items-end class.
func AlignItemsEnd() attrib.Node { return attr.Class("align-items-end") }

// AlignItemsCenter creates a Bootstrap align-items-center class.
func AlignItemsCenter() attrib.Node { return attr.Class("align-items-center") }

// AlignItemsBaseline creates a Bootstrap align-items-baseline class.
func AlignItemsBaseline() attrib.Node { return attr.Class("align-items-baseline") }

// AlignItemsStretch creates a Bootstrap align-items-stretch class.
func AlignItemsStretch() attrib.Node { return attr.Class("align-items-stretch") }

// AlignSelfStart creates a Bootstrap align-self-start class.
func AlignSelfStart() attrib.Node { return attr.Class("align-self-start") }

// AlignSelfEnd creates a Bootstrap align-self-end class.
func AlignSelfEnd() attrib.Node { return attr.Class("align-self-end") }

// AlignSelfCenter creates a Bootstrap align-self-center class.
func AlignSelfCenter() attrib.Node { return attr.Class("align-self-center") }

// AlignSelfBaseline creates a Bootstrap align-self-baseline class.
func AlignSelfBaseline() attrib.Node { return attr.Class("align-self-baseline") }

// AlignSelfStretch creates a Bootstrap align-self-stretch class.
func AlignSelfStretch() attrib.Node { return attr.Class("align-self-stretch") }

// FlexFill creates a Bootstrap flex-fill class.
func FlexFill() attrib.Node { return attr.Class("flex-fill") }

// FlexGrow0 creates a Bootstrap flex-grow-0 class.
func FlexGrow0() attrib.Node { return attr.Class("flex-grow-0") }

// FlexGrow1 creates a Bootstrap flex-grow-1 class.
func FlexGrow1() attrib.Node { return attr.Class("flex-grow-1") }

// FlexShrink0 creates a Bootstrap flex-shrink-0 class.
func FlexShrink0() attrib.Node { return attr.Class("flex-shrink-0") }

// FlexShrink1 creates a Bootstrap flex-shrink-1 class.
func FlexShrink1() attrib.Node { return attr.Class("flex-shrink-1") }

// Gap creates a Bootstrap gap-{n} class.
func Gap(n int) attrib.Node { return attr.Class(fmt.Sprintf("gap-%d", n)) }
