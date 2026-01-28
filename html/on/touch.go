package on

import "github.com/protolambda/chord/core"

// TouchStart sets the ontouchstart event handler.
func TouchStart(v string) core.Node { return core.Attribute("ontouchstart", v) }

// TouchMove sets the ontouchmove event handler.
func TouchMove(v string) core.Node { return core.Attribute("ontouchmove", v) }

// TouchEnd sets the ontouchend event handler.
func TouchEnd(v string) core.Node { return core.Attribute("ontouchend", v) }

// TouchCancel sets the ontouchcancel event handler.
func TouchCancel(v string) core.Node { return core.Attribute("ontouchcancel", v) }
