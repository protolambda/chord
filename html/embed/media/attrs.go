package media

import (
	"github.com/protolambda/chord/core/attrib"
)

// Src sets the src attribute.
func Src(v string) attrib.Node { return attrib.KV("src", v) }

// Controls sets the controls boolean attribute.
func Controls() attrib.Node { return attrib.Bool("controls") }

// Autoplay sets the autoplay boolean attribute.
func Autoplay() attrib.Node { return attrib.Bool("autoplay") }

// Loop sets the loop boolean attribute.
func Loop() attrib.Node { return attrib.Bool("loop") }

// Muted sets the muted boolean attribute.
func Muted() attrib.Node { return attrib.Bool("muted") }

// Preload sets the preload attribute.
func Preload(v string) attrib.Node { return attrib.KV("preload", v) }

// Poster sets the poster attribute.
func Poster(v string) attrib.Node { return attrib.KV("poster", v) }

// Playsinline sets the playsinline boolean attribute.
func Playsinline() attrib.Node { return attrib.Bool("playsinline") }

// Width sets the width attribute.
func Width(v string) attrib.Node { return attrib.KV("width", v) }

// Height sets the height attribute.
func Height(v string) attrib.Node { return attrib.KV("height", v) }
