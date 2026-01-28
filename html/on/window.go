package on

import "github.com/protolambda/chord/core"

// Resize sets the onresize event handler.
func Resize(v string) core.Node { return core.Attribute("onresize", v) }

// Scroll sets the onscroll event handler.
func Scroll(v string) core.Node { return core.Attribute("onscroll", v) }

// ScrollEnd sets the onscrollend event handler.
func ScrollEnd(v string) core.Node { return core.Attribute("onscrollend", v) }
