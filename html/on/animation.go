package on

import "github.com/protolambda/chord/core"

// AnimationStart sets the onanimationstart event handler.
func AnimationStart(v string) core.Node { return core.Attribute("onanimationstart", v) }

// AnimationEnd sets the onanimationend event handler.
func AnimationEnd(v string) core.Node { return core.Attribute("onanimationend", v) }

// AnimationIteration sets the onanimationiteration event handler.
func AnimationIteration(v string) core.Node { return core.Attribute("onanimationiteration", v) }

// AnimationCancel sets the onanimationcancel event handler.
func AnimationCancel(v string) core.Node { return core.Attribute("onanimationcancel", v) }
