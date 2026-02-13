package on

import (
	"github.com/protolambda/chord/core/attrib"
)

// AnimationStart sets the onanimationstart event handler.
func AnimationStart(v string) attrib.Node { return attrib.KV("onanimationstart", v) }

// AnimationEnd sets the onanimationend event handler.
func AnimationEnd(v string) attrib.Node { return attrib.KV("onanimationend", v) }

// AnimationIteration sets the onanimationiteration event handler.
func AnimationIteration(v string) attrib.Node { return attrib.KV("onanimationiteration", v) }

// AnimationCancel sets the onanimationcancel event handler.
func AnimationCancel(v string) attrib.Node { return attrib.KV("onanimationcancel", v) }
