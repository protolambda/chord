package on

import (
	"github.com/protolambda/chord/core/attr"
)

// Wheel sets the onwheel event handler.
func Wheel(v string) attr.Node { return attr.KV("onwheel", v) }

// BeforeInput sets the onbeforeinput event handler.
func BeforeInput(v string) attr.Node { return attr.KV("onbeforeinput", v) }

// FullscreenChange sets the onfullscreenchange event handler.
func FullscreenChange(v string) attr.Node { return attr.KV("onfullscreenchange", v) }

// FullscreenError sets the onfullscreenerror event handler.
func FullscreenError(v string) attr.Node { return attr.KV("onfullscreenerror", v) }
