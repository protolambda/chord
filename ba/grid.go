package ba

import (
	"fmt"

	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/html/attr"
)

// Container returns a "container" class attribute.
func Container() core.Node { return attr.Class("container") }

// ContainerFluid returns a "container-fluid" class attribute.
func ContainerFluid() core.Node { return attr.Class("container-fluid") }

// ContainerSM returns a "container-sm" class attribute.
func ContainerSM() core.Node { return attr.Class("container-sm") }

// ContainerMD returns a "container-md" class attribute.
func ContainerMD() core.Node { return attr.Class("container-md") }

// ContainerLG returns a "container-lg" class attribute.
func ContainerLG() core.Node { return attr.Class("container-lg") }

// ContainerXL returns a "container-xl" class attribute.
func ContainerXL() core.Node { return attr.Class("container-xl") }

// ContainerXXL returns a "container-xxl" class attribute.
func ContainerXXL() core.Node { return attr.Class("container-xxl") }

// Row returns a "row" class attribute.
func Row() core.Node { return attr.Class("row") }

// Col returns a "col" class attribute.
func Col() core.Node { return attr.Class("col") }

// Col1 returns a "col-1" class attribute.
func Col1() core.Node { return attr.Class("col-1") }

// Col2 returns a "col-2" class attribute.
func Col2() core.Node { return attr.Class("col-2") }

// Col3 returns a "col-3" class attribute.
func Col3() core.Node { return attr.Class("col-3") }

// Col4 returns a "col-4" class attribute.
func Col4() core.Node { return attr.Class("col-4") }

// Col5 returns a "col-5" class attribute.
func Col5() core.Node { return attr.Class("col-5") }

// Col6 returns a "col-6" class attribute.
func Col6() core.Node { return attr.Class("col-6") }

// Col7 returns a "col-7" class attribute.
func Col7() core.Node { return attr.Class("col-7") }

// Col8 returns a "col-8" class attribute.
func Col8() core.Node { return attr.Class("col-8") }

// Col9 returns a "col-9" class attribute.
func Col9() core.Node { return attr.Class("col-9") }

// Col10 returns a "col-10" class attribute.
func Col10() core.Node { return attr.Class("col-10") }

// Col11 returns a "col-11" class attribute.
func Col11() core.Node { return attr.Class("col-11") }

// Col12 returns a "col-12" class attribute.
func Col12() core.Node { return attr.Class("col-12") }

// ColSM returns a "col-sm-{n}" class attribute.
func ColSM(n int) core.Node { return attr.Class(fmt.Sprintf("col-sm-%d", n)) }

// ColMD returns a "col-md-{n}" class attribute.
func ColMD(n int) core.Node { return attr.Class(fmt.Sprintf("col-md-%d", n)) }

// ColLG returns a "col-lg-{n}" class attribute.
func ColLG(n int) core.Node { return attr.Class(fmt.Sprintf("col-lg-%d", n)) }

// ColXL returns a "col-xl-{n}" class attribute.
func ColXL(n int) core.Node { return attr.Class(fmt.Sprintf("col-xl-%d", n)) }

// ColXXL returns a "col-xxl-{n}" class attribute.
func ColXXL(n int) core.Node { return attr.Class(fmt.Sprintf("col-xxl-%d", n)) }

// ColAuto returns a "col-auto" class attribute.
func ColAuto() core.Node { return attr.Class("col-auto") }

// Offset returns an "offset-{n}" class attribute.
func Offset(n int) core.Node { return attr.Class(fmt.Sprintf("offset-%d", n)) }

// OffsetSM returns an "offset-sm-{n}" class attribute.
func OffsetSM(n int) core.Node { return attr.Class(fmt.Sprintf("offset-sm-%d", n)) }

// OffsetMD returns an "offset-md-{n}" class attribute.
func OffsetMD(n int) core.Node { return attr.Class(fmt.Sprintf("offset-md-%d", n)) }

// OffsetLG returns an "offset-lg-{n}" class attribute.
func OffsetLG(n int) core.Node { return attr.Class(fmt.Sprintf("offset-lg-%d", n)) }
