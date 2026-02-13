package on

import (
	"github.com/protolambda/chord/core/attrib"
)

// Resize sets the onresize event handler.
func Resize(v string) attrib.Node { return attrib.KV("onresize", v) }

// Scroll sets the onscroll event handler.
func Scroll(v string) attrib.Node { return attrib.KV("onscroll", v) }

// ScrollEnd sets the onscrollend event handler.
func ScrollEnd(v string) attrib.Node { return attrib.KV("onscrollend", v) }
