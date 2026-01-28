package div

import (
	"github.com/protolambda/chord/core"
)

func Div(opts ...core.Node) core.Node {
	return core.Element("div", opts...)
}
