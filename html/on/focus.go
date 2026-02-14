package on

import (
	"github.com/protolambda/chord/core/attr"
)

// Focus sets the onfocus event handler.
func Focus(v string) attr.Node { return attr.KV("onfocus", v) }

// Blur sets the onblur event handler.
func Blur(v string) attr.Node { return attr.KV("onblur", v) }

// FocusIn sets the onfocusin event handler.
func FocusIn(v string) attr.Node { return attr.KV("onfocusin", v) }

// FocusOut sets the onfocusout event handler.
func FocusOut(v string) attr.Node { return attr.KV("onfocusout", v) }
