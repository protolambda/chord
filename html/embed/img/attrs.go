package img

import (
	"github.com/protolambda/chord/core/attr"
)

// Src sets the src attribute.
func Src(v string) attr.Node { return attr.KV("src", v) }

// Alt sets the alt attribute.
func Alt(v string) attr.Node { return attr.KV("alt", v) }

// Width sets the width attribute.
func Width(v string) attr.Node { return attr.KV("width", v) }

// Height sets the height attribute.
func Height(v string) attr.Node { return attr.KV("height", v) }

// Srcset sets the srcset attribute.
func Srcset(v string) attr.Node { return attr.KV("srcset", v) }

// Sizes sets the sizes attribute.
func Sizes(v string) attr.Node { return attr.KV("sizes", v) }

// Loading sets the loading attribute.
func Loading(v string) attr.Node { return attr.KV("loading", v) }

// Decoding sets the decoding attribute.
func Decoding(v string) attr.Node { return attr.KV("decoding", v) }

// Fetchpriority sets the fetchpriority attribute.
func Fetchpriority(v string) attr.Node { return attr.KV("fetchpriority", v) }

// Ismap sets the ismap boolean attribute.
func Ismap() attr.Node { return attr.Bool("ismap") }

// Usemap sets the usemap attribute.
func Usemap(v string) attr.Node { return attr.KV("usemap", v) }
