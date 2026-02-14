package media

import (
	"github.com/protolambda/chord/core/attr"
)

// Src sets the src attribute.
func Src(v string) attr.Node { return attr.KV("src", v) }

// Controls sets the controls boolean attribute.
func Controls() attr.Node { return attr.Bool("controls") }

// Autoplay sets the autoplay boolean attribute.
func Autoplay() attr.Node { return attr.Bool("autoplay") }

// Loop sets the loop boolean attribute.
func Loop() attr.Node { return attr.Bool("loop") }

// Muted sets the muted boolean attribute.
func Muted() attr.Node { return attr.Bool("muted") }

// Preload sets the preload attribute.
func Preload(v string) attr.Node { return attr.KV("preload", v) }

// Poster sets the poster attribute.
func Poster(v string) attr.Node { return attr.KV("poster", v) }

// Playsinline sets the playsinline boolean attribute.
func Playsinline() attr.Node { return attr.Bool("playsinline") }

// Width sets the width attribute.
func Width(v string) attr.Node { return attr.KV("width", v) }

// Height sets the height attribute.
func Height(v string) attr.Node { return attr.KV("height", v) }
