package img

import (
	"github.com/protolambda/chord/core/attrib"
)

// Src sets the src attribute.
func Src(v string) attrib.Node { return attrib.KV("src", v) }

// Alt sets the alt attribute.
func Alt(v string) attrib.Node { return attrib.KV("alt", v) }

// Width sets the width attribute.
func Width(v string) attrib.Node { return attrib.KV("width", v) }

// Height sets the height attribute.
func Height(v string) attrib.Node { return attrib.KV("height", v) }

// Srcset sets the srcset attribute.
func Srcset(v string) attrib.Node { return attrib.KV("srcset", v) }

// Sizes sets the sizes attribute.
func Sizes(v string) attrib.Node { return attrib.KV("sizes", v) }

// Loading sets the loading attribute.
func Loading(v string) attrib.Node { return attrib.KV("loading", v) }

// Decoding sets the decoding attribute.
func Decoding(v string) attrib.Node { return attrib.KV("decoding", v) }

// Fetchpriority sets the fetchpriority attribute.
func Fetchpriority(v string) attrib.Node { return attrib.KV("fetchpriority", v) }

// Ismap sets the ismap boolean attribute.
func Ismap() attrib.Node { return attrib.Bool("ismap") }

// Usemap sets the usemap attribute.
func Usemap(v string) attrib.Node { return attrib.KV("usemap", v) }
