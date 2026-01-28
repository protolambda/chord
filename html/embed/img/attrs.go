package img

import "github.com/protolambda/chord/core"

// Src sets the src attribute.
func Src(v string) core.Node { return core.Attribute("src", v) }

// Alt sets the alt attribute.
func Alt(v string) core.Node { return core.Attribute("alt", v) }

// Width sets the width attribute.
func Width(v string) core.Node { return core.Attribute("width", v) }

// Height sets the height attribute.
func Height(v string) core.Node { return core.Attribute("height", v) }

// Srcset sets the srcset attribute.
func Srcset(v string) core.Node { return core.Attribute("srcset", v) }

// Sizes sets the sizes attribute.
func Sizes(v string) core.Node { return core.Attribute("sizes", v) }

// Loading sets the loading attribute.
func Loading(v string) core.Node { return core.Attribute("loading", v) }

// Decoding sets the decoding attribute.
func Decoding(v string) core.Node { return core.Attribute("decoding", v) }

// Fetchpriority sets the fetchpriority attribute.
func Fetchpriority(v string) core.Node { return core.Attribute("fetchpriority", v) }

// Ismap sets the ismap boolean attribute.
func Ismap() core.Node { return core.BoolAttribute("ismap") }

// Usemap sets the usemap attribute.
func Usemap(v string) core.Node { return core.Attribute("usemap", v) }
