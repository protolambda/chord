package on

import (
	"github.com/protolambda/chord/core/attr"
)

// Wheel sets the onwheel event handler.
func Wheel(v string) attr.Node { return attr.Name("onwheel").Value(v) }

// BeforeInput sets the onbeforeinput event handler.
func BeforeInput(v string) attr.Node { return attr.Name("onbeforeinput").Value(v) }

// FullscreenChange sets the onfullscreenchange event handler.
func FullscreenChange(v string) attr.Node { return attr.Name("onfullscreenchange").Value(v) }

// FullscreenError sets the onfullscreenerror event handler.
func FullscreenError(v string) attr.Node { return attr.Name("onfullscreenerror").Value(v) }
