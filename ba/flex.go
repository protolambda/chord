package ba

import (
	"fmt"

	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/html/attr"
)

// FlexRow creates a Bootstrap flex-row class.
func FlexRow() core.Node { return attr.Class("flex-row") }

// FlexColumn creates a Bootstrap flex-column class.
func FlexColumn() core.Node { return attr.Class("flex-column") }

// FlexRowReverse creates a Bootstrap flex-row-reverse class.
func FlexRowReverse() core.Node { return attr.Class("flex-row-reverse") }

// FlexColumnReverse creates a Bootstrap flex-column-reverse class.
func FlexColumnReverse() core.Node { return attr.Class("flex-column-reverse") }

// FlexWrap creates a Bootstrap flex-wrap class.
func FlexWrap() core.Node { return attr.Class("flex-wrap") }

// FlexNowrap creates a Bootstrap flex-nowrap class.
func FlexNowrap() core.Node { return attr.Class("flex-nowrap") }

// FlexWrapReverse creates a Bootstrap flex-wrap-reverse class.
func FlexWrapReverse() core.Node { return attr.Class("flex-wrap-reverse") }

// JustifyContentStart creates a Bootstrap justify-content-start class.
func JustifyContentStart() core.Node { return attr.Class("justify-content-start") }

// JustifyContentEnd creates a Bootstrap justify-content-end class.
func JustifyContentEnd() core.Node { return attr.Class("justify-content-end") }

// JustifyContentCenter creates a Bootstrap justify-content-center class.
func JustifyContentCenter() core.Node { return attr.Class("justify-content-center") }

// JustifyContentBetween creates a Bootstrap justify-content-between class.
func JustifyContentBetween() core.Node { return attr.Class("justify-content-between") }

// JustifyContentAround creates a Bootstrap justify-content-around class.
func JustifyContentAround() core.Node { return attr.Class("justify-content-around") }

// JustifyContentEvenly creates a Bootstrap justify-content-evenly class.
func JustifyContentEvenly() core.Node { return attr.Class("justify-content-evenly") }

// AlignItemsStart creates a Bootstrap align-items-start class.
func AlignItemsStart() core.Node { return attr.Class("align-items-start") }

// AlignItemsEnd creates a Bootstrap align-items-end class.
func AlignItemsEnd() core.Node { return attr.Class("align-items-end") }

// AlignItemsCenter creates a Bootstrap align-items-center class.
func AlignItemsCenter() core.Node { return attr.Class("align-items-center") }

// AlignItemsBaseline creates a Bootstrap align-items-baseline class.
func AlignItemsBaseline() core.Node { return attr.Class("align-items-baseline") }

// AlignItemsStretch creates a Bootstrap align-items-stretch class.
func AlignItemsStretch() core.Node { return attr.Class("align-items-stretch") }

// AlignSelfStart creates a Bootstrap align-self-start class.
func AlignSelfStart() core.Node { return attr.Class("align-self-start") }

// AlignSelfEnd creates a Bootstrap align-self-end class.
func AlignSelfEnd() core.Node { return attr.Class("align-self-end") }

// AlignSelfCenter creates a Bootstrap align-self-center class.
func AlignSelfCenter() core.Node { return attr.Class("align-self-center") }

// AlignSelfBaseline creates a Bootstrap align-self-baseline class.
func AlignSelfBaseline() core.Node { return attr.Class("align-self-baseline") }

// AlignSelfStretch creates a Bootstrap align-self-stretch class.
func AlignSelfStretch() core.Node { return attr.Class("align-self-stretch") }

// FlexFill creates a Bootstrap flex-fill class.
func FlexFill() core.Node { return attr.Class("flex-fill") }

// FlexGrow0 creates a Bootstrap flex-grow-0 class.
func FlexGrow0() core.Node { return attr.Class("flex-grow-0") }

// FlexGrow1 creates a Bootstrap flex-grow-1 class.
func FlexGrow1() core.Node { return attr.Class("flex-grow-1") }

// FlexShrink0 creates a Bootstrap flex-shrink-0 class.
func FlexShrink0() core.Node { return attr.Class("flex-shrink-0") }

// FlexShrink1 creates a Bootstrap flex-shrink-1 class.
func FlexShrink1() core.Node { return attr.Class("flex-shrink-1") }

// Gap creates a Bootstrap gap-{n} class.
func Gap(n int) core.Node { return attr.Class(fmt.Sprintf("gap-%d", n)) }
