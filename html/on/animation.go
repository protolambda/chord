package on

import (
	"github.com/protolambda/chord/core/attr"
)

// AnimationStart sets the onanimationstart event handler.
func AnimationStart(v string) attr.Node { return attr.Name("onanimationstart").Value(v) }

// AnimationEnd sets the onanimationend event handler.
func AnimationEnd(v string) attr.Node { return attr.Name("onanimationend").Value(v) }

// AnimationIteration sets the onanimationiteration event handler.
func AnimationIteration(v string) attr.Node { return attr.Name("onanimationiteration").Value(v) }

// AnimationCancel sets the onanimationcancel event handler.
func AnimationCancel(v string) attr.Node { return attr.Name("onanimationcancel").Value(v) }
