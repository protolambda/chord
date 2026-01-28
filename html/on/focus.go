package on

import "github.com/protolambda/chord/core"

// Focus sets the onfocus event handler.
func Focus(v string) core.Node { return core.Attribute("onfocus", v) }

// Blur sets the onblur event handler.
func Blur(v string) core.Node { return core.Attribute("onblur", v) }

// FocusIn sets the onfocusin event handler.
func FocusIn(v string) core.Node { return core.Attribute("onfocusin", v) }

// FocusOut sets the onfocusout event handler.
func FocusOut(v string) core.Node { return core.Attribute("onfocusout", v) }
