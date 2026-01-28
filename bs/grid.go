package bs

import (
	"fmt"

	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/html/attr"
	"github.com/protolambda/chord/html/group/div"
)

// Container creates a Bootstrap container div element.
func Container(opts ...core.Node) core.Node {
	return div.Div(core.Compose(attr.Class("container"), opts...))
}

// ContainerFluid creates a Bootstrap container-fluid div element.
func ContainerFluid(opts ...core.Node) core.Node {
	return div.Div(core.Compose(attr.Class("container-fluid"), opts...))
}

// ContainerSM creates a Bootstrap container-sm div element.
func ContainerSM(opts ...core.Node) core.Node {
	return div.Div(core.Compose(attr.Class("container-sm"), opts...))
}

// ContainerMD creates a Bootstrap container-md div element.
func ContainerMD(opts ...core.Node) core.Node {
	return div.Div(core.Compose(attr.Class("container-md"), opts...))
}

// ContainerLG creates a Bootstrap container-lg div element.
func ContainerLG(opts ...core.Node) core.Node {
	return div.Div(core.Compose(attr.Class("container-lg"), opts...))
}

// ContainerXL creates a Bootstrap container-xl div element.
func ContainerXL(opts ...core.Node) core.Node {
	return div.Div(core.Compose(attr.Class("container-xl"), opts...))
}

// ContainerXXL creates a Bootstrap container-xxl div element.
func ContainerXXL(opts ...core.Node) core.Node {
	return div.Div(core.Compose(attr.Class("container-xxl"), opts...))
}

// Row creates a Bootstrap row div element.
func Row(opts ...core.Node) core.Node {
	return div.Div(core.Compose(attr.Class("row"), opts...))
}

// Col creates a Bootstrap col div element.
func Col(opts ...core.Node) core.Node {
	return div.Div(core.Compose(attr.Class("col"), opts...))
}

// Col1 creates a Bootstrap col-1 div element.
func Col1(opts ...core.Node) core.Node {
	return div.Div(core.Compose(attr.Class("col-1"), opts...))
}

// Col2 creates a Bootstrap col-2 div element.
func Col2(opts ...core.Node) core.Node {
	return div.Div(core.Compose(attr.Class("col-2"), opts...))
}

// Col3 creates a Bootstrap col-3 div element.
func Col3(opts ...core.Node) core.Node {
	return div.Div(core.Compose(attr.Class("col-3"), opts...))
}

// Col4 creates a Bootstrap col-4 div element.
func Col4(opts ...core.Node) core.Node {
	return div.Div(core.Compose(attr.Class("col-4"), opts...))
}

// Col5 creates a Bootstrap col-5 div element.
func Col5(opts ...core.Node) core.Node {
	return div.Div(core.Compose(attr.Class("col-5"), opts...))
}

// Col6 creates a Bootstrap col-6 div element.
func Col6(opts ...core.Node) core.Node {
	return div.Div(core.Compose(attr.Class("col-6"), opts...))
}

// Col7 creates a Bootstrap col-7 div element.
func Col7(opts ...core.Node) core.Node {
	return div.Div(core.Compose(attr.Class("col-7"), opts...))
}

// Col8 creates a Bootstrap col-8 div element.
func Col8(opts ...core.Node) core.Node {
	return div.Div(core.Compose(attr.Class("col-8"), opts...))
}

// Col9 creates a Bootstrap col-9 div element.
func Col9(opts ...core.Node) core.Node {
	return div.Div(core.Compose(attr.Class("col-9"), opts...))
}

// Col10 creates a Bootstrap col-10 div element.
func Col10(opts ...core.Node) core.Node {
	return div.Div(core.Compose(attr.Class("col-10"), opts...))
}

// Col11 creates a Bootstrap col-11 div element.
func Col11(opts ...core.Node) core.Node {
	return div.Div(core.Compose(attr.Class("col-11"), opts...))
}

// Col12 creates a Bootstrap col-12 div element.
func Col12(opts ...core.Node) core.Node {
	return div.Div(core.Compose(attr.Class("col-12"), opts...))
}

// ColSM creates a Bootstrap col-sm-{n} div element.
func ColSM(n int, opts ...core.Node) core.Node {
	return div.Div(core.Compose(attr.Class(fmt.Sprintf("col-sm-%d", n)), opts...))
}

// ColMD creates a Bootstrap col-md-{n} div element.
func ColMD(n int, opts ...core.Node) core.Node {
	return div.Div(core.Compose(attr.Class(fmt.Sprintf("col-md-%d", n)), opts...))
}

// ColLG creates a Bootstrap col-lg-{n} div element.
func ColLG(n int, opts ...core.Node) core.Node {
	return div.Div(core.Compose(attr.Class(fmt.Sprintf("col-lg-%d", n)), opts...))
}

// ColXL creates a Bootstrap col-xl-{n} div element.
func ColXL(n int, opts ...core.Node) core.Node {
	return div.Div(core.Compose(attr.Class(fmt.Sprintf("col-xl-%d", n)), opts...))
}

// ColXXL creates a Bootstrap col-xxl-{n} div element.
func ColXXL(n int, opts ...core.Node) core.Node {
	return div.Div(core.Compose(attr.Class(fmt.Sprintf("col-xxl-%d", n)), opts...))
}

// ColAuto creates a Bootstrap col-auto div element.
func ColAuto(opts ...core.Node) core.Node {
	return div.Div(core.Compose(attr.Class("col-auto"), opts...))
}
