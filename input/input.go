package input

import (
	"github.com/protolambda/chord/core"
)

func Input(opts ...core.Node) core.Node {
	return core.VoidElement("input", opts...)
}

func Size(v string) core.Node {
	return core.Attribute("size", v)
}

func Type(v string) core.Node {
	return core.Attribute("type", v)
}

func Checked() core.Node {
	return core.BoolAttribute("checked")
}
