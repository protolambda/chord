package on

import (
	"github.com/protolambda/chord/core/attr"
)

// Resize sets the onresize event handler.
func Resize(v string) attr.Node { return attr.Name("onresize").Value(v) }

// Scroll sets the onscroll event handler.
func Scroll(v string) attr.Node { return attr.Name("onscroll").Value(v) }

// ScrollEnd sets the onscrollend event handler.
func ScrollEnd(v string) attr.Node { return attr.Name("onscrollend").Value(v) }
