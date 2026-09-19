package media

import (
	"github.com/protolambda/chord/core/attr"
)

// Src sets the src attribute.
func Src(v string) attr.Node { return attr.Name("src").Value(v) }

// Controls sets the controls boolean attribute.
func Controls() attr.Node { return attr.Name("controls").Bool() }

// Autoplay sets the autoplay boolean attribute.
func Autoplay() attr.Node { return attr.Name("autoplay").Bool() }

// Loop sets the loop boolean attribute.
func Loop() attr.Node { return attr.Name("loop").Bool() }

// Muted sets the muted boolean attribute.
func Muted() attr.Node { return attr.Name("muted").Bool() }

// Preload sets the preload attribute.
func Preload(v string) attr.Node { return attr.Name("preload").Value(v) }

// Poster sets the poster attribute.
func Poster(v string) attr.Node { return attr.Name("poster").Value(v) }

// Playsinline sets the playsinline boolean attribute.
func Playsinline() attr.Node { return attr.Name("playsinline").Bool() }

// Width sets the width attribute.
func Width(v string) attr.Node { return attr.Name("width").Value(v) }

// Height sets the height attribute.
func Height(v string) attr.Node { return attr.Name("height").Value(v) }
