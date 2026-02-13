package on

import (
	"github.com/protolambda/chord/core/attrib"
)

// Wheel sets the onwheel event handler.
func Wheel(v string) attrib.Node { return attrib.KV("onwheel", v) }

// BeforeInput sets the onbeforeinput event handler.
func BeforeInput(v string) attrib.Node { return attrib.KV("onbeforeinput", v) }

// FullscreenChange sets the onfullscreenchange event handler.
func FullscreenChange(v string) attrib.Node { return attrib.KV("onfullscreenchange", v) }

// FullscreenError sets the onfullscreenerror event handler.
func FullscreenError(v string) attrib.Node { return attrib.KV("onfullscreenerror", v) }
