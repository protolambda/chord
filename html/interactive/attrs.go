package interactive

import "github.com/protolambda/chord/core"

// Open sets the open boolean attribute.
func Open() core.Node { return core.BoolAttribute("open") }
