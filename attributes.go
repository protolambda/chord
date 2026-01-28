package chord

import (
	"github.com/protolambda/chord/core"
)

func Class(v string) core.Node {
	return core.Attribute("class", v)
}

func ID(v string) core.Node {
	return core.Attribute("id", v)
}

func Style(v string) core.Node {
	return core.Attribute("style", v)
}

func Slot(v string) core.Node {
	return core.Attribute("slot", v)
}

func Title(v string) core.Node {
	return core.Attribute("title", v)
}
