package on

import (
	"github.com/protolambda/chord/core/attrib"
)

// TouchStart sets the ontouchstart event handler.
func TouchStart(v string) attrib.Node { return attrib.KV("ontouchstart", v) }

// TouchMove sets the ontouchmove event handler.
func TouchMove(v string) attrib.Node { return attrib.KV("ontouchmove", v) }

// TouchEnd sets the ontouchend event handler.
func TouchEnd(v string) attrib.Node { return attrib.KV("ontouchend", v) }

// TouchCancel sets the ontouchcancel event handler.
func TouchCancel(v string) attrib.Node { return attrib.KV("ontouchcancel", v) }
