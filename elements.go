package chord

import (
	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/div"
	"github.com/protolambda/chord/input"
)

func Input(opts ...core.Node) core.Node {
	return input.Input(opts...)
}

func Div(opts ...core.Node) core.Node {
	return div.Div(opts...)
}
