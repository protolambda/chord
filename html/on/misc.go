package on

import "github.com/protolambda/chord/core"

// Wheel sets the onwheel event handler.
func Wheel(v string) core.Node { return core.Attribute("onwheel", v) }

// BeforeInput sets the onbeforeinput event handler.
func BeforeInput(v string) core.Node { return core.Attribute("onbeforeinput", v) }

// FullscreenChange sets the onfullscreenchange event handler.
func FullscreenChange(v string) core.Node { return core.Attribute("onfullscreenchange", v) }

// FullscreenError sets the onfullscreenerror event handler.
func FullscreenError(v string) core.Node { return core.Attribute("onfullscreenerror", v) }
