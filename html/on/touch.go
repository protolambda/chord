package on

import (
	"github.com/protolambda/chord/core/attr"
)

// TouchStart sets the ontouchstart event handler.
func TouchStart(v string) attr.Node { return attr.Name("ontouchstart").Value(v) }

// TouchMove sets the ontouchmove event handler.
func TouchMove(v string) attr.Node { return attr.Name("ontouchmove").Value(v) }

// TouchEnd sets the ontouchend event handler.
func TouchEnd(v string) attr.Node { return attr.Name("ontouchend").Value(v) }

// TouchCancel sets the ontouchcancel event handler.
func TouchCancel(v string) attr.Node { return attr.Name("ontouchcancel").Value(v) }
