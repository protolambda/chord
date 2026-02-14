package on

import (
	"github.com/protolambda/chord/core/attr"
)

// TouchStart sets the ontouchstart event handler.
func TouchStart(v string) attr.Node { return attr.KV("ontouchstart", v) }

// TouchMove sets the ontouchmove event handler.
func TouchMove(v string) attr.Node { return attr.KV("ontouchmove", v) }

// TouchEnd sets the ontouchend event handler.
func TouchEnd(v string) attr.Node { return attr.KV("ontouchend", v) }

// TouchCancel sets the ontouchcancel event handler.
func TouchCancel(v string) attr.Node { return attr.KV("ontouchcancel", v) }
