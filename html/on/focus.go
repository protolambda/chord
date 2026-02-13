package on

import (
	"github.com/protolambda/chord/core/attrib"
)

// Focus sets the onfocus event handler.
func Focus(v string) attrib.Node { return attrib.KV("onfocus", v) }

// Blur sets the onblur event handler.
func Blur(v string) attrib.Node { return attrib.KV("onblur", v) }

// FocusIn sets the onfocusin event handler.
func FocusIn(v string) attrib.Node { return attrib.KV("onfocusin", v) }

// FocusOut sets the onfocusout event handler.
func FocusOut(v string) attrib.Node { return attrib.KV("onfocusout", v) }
