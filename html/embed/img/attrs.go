package img

import (
	"github.com/protolambda/chord/core/attr"
)

// Src sets the src attribute.
func Src(v string) attr.Node { return attr.Name("src").Value(v) }

// Alt sets the alt attribute.
func Alt(v string) attr.Node { return attr.Name("alt").Value(v) }

// Width sets the width attribute.
func Width(v string) attr.Node { return attr.Name("width").Value(v) }

// Height sets the height attribute.
func Height(v string) attr.Node { return attr.Name("height").Value(v) }

// Srcset sets the srcset attribute.
func Srcset(v string) attr.Node { return attr.Name("srcset").Value(v) }

// Sizes sets the sizes attribute.
func Sizes(v string) attr.Node { return attr.Name("sizes").Value(v) }

// Loading sets the loading attribute.
func Loading(v string) attr.Node { return attr.Name("loading").Value(v) }

// Decoding sets the decoding attribute.
func Decoding(v string) attr.Node { return attr.Name("decoding").Value(v) }

// Fetchpriority sets the fetchpriority attribute.
func Fetchpriority(v string) attr.Node { return attr.Name("fetchpriority").Value(v) }

// Ismap sets the ismap boolean attribute.
func Ismap() attr.Node { return attr.Name("ismap").Bool() }

// Usemap sets the usemap attribute.
func Usemap(v string) attr.Node { return attr.Name("usemap").Value(v) }
