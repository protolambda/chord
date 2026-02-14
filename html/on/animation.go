package on

import (
	"github.com/protolambda/chord/core/attr"
)

// AnimationStart sets the onanimationstart event handler.
func AnimationStart(v string) attr.Node { return attr.KV("onanimationstart", v) }

// AnimationEnd sets the onanimationend event handler.
func AnimationEnd(v string) attr.Node { return attr.KV("onanimationend", v) }

// AnimationIteration sets the onanimationiteration event handler.
func AnimationIteration(v string) attr.Node { return attr.KV("onanimationiteration", v) }

// AnimationCancel sets the onanimationcancel event handler.
func AnimationCancel(v string) attr.Node { return attr.KV("onanimationcancel", v) }
