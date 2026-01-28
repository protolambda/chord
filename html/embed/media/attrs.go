package media

import "github.com/protolambda/chord/core"

// Src sets the src attribute.
func Src(v string) core.Node { return core.Attribute("src", v) }

// Controls sets the controls boolean attribute.
func Controls() core.Node { return core.BoolAttribute("controls") }

// Autoplay sets the autoplay boolean attribute.
func Autoplay() core.Node { return core.BoolAttribute("autoplay") }

// Loop sets the loop boolean attribute.
func Loop() core.Node { return core.BoolAttribute("loop") }

// Muted sets the muted boolean attribute.
func Muted() core.Node { return core.BoolAttribute("muted") }

// Preload sets the preload attribute.
func Preload(v string) core.Node { return core.Attribute("preload", v) }

// Poster sets the poster attribute.
func Poster(v string) core.Node { return core.Attribute("poster", v) }

// Playsinline sets the playsinline boolean attribute.
func Playsinline() core.Node { return core.BoolAttribute("playsinline") }

// Width sets the width attribute.
func Width(v string) core.Node { return core.Attribute("width", v) }

// Height sets the height attribute.
func Height(v string) core.Node { return core.Attribute("height", v) }
