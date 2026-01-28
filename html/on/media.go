package on

import "github.com/protolambda/chord/core"

// Play sets the onplay event handler.
func Play(v string) core.Node { return core.Attribute("onplay", v) }

// Pause sets the onpause event handler.
func Pause(v string) core.Node { return core.Attribute("onpause", v) }

// Ended sets the onended event handler.
func Ended(v string) core.Node { return core.Attribute("onended", v) }

// VolumeChange sets the onvolumechange event handler.
func VolumeChange(v string) core.Node { return core.Attribute("onvolumechange", v) }

// TimeUpdate sets the ontimeupdate event handler.
func TimeUpdate(v string) core.Node { return core.Attribute("ontimeupdate", v) }

// Seeking sets the onseeking event handler.
func Seeking(v string) core.Node { return core.Attribute("onseeking", v) }

// Seeked sets the onseeked event handler.
func Seeked(v string) core.Node { return core.Attribute("onseeked", v) }

// LoadedData sets the onloadeddata event handler.
func LoadedData(v string) core.Node { return core.Attribute("onloadeddata", v) }

// LoadedMetadata sets the onloadedmetadata event handler.
func LoadedMetadata(v string) core.Node { return core.Attribute("onloadedmetadata", v) }

// CanPlay sets the oncanplay event handler.
func CanPlay(v string) core.Node { return core.Attribute("oncanplay", v) }

// CanPlayThrough sets the oncanplaythrough event handler.
func CanPlayThrough(v string) core.Node { return core.Attribute("oncanplaythrough", v) }

// Waiting sets the onwaiting event handler.
func Waiting(v string) core.Node { return core.Attribute("onwaiting", v) }

// Playing sets the onplaying event handler.
func Playing(v string) core.Node { return core.Attribute("onplaying", v) }

// DurationChange sets the ondurationchange event handler.
func DurationChange(v string) core.Node { return core.Attribute("ondurationchange", v) }

// RateChange sets the onratechange event handler.
func RateChange(v string) core.Node { return core.Attribute("onratechange", v) }

// Stalled sets the onstalled event handler.
func Stalled(v string) core.Node { return core.Attribute("onstalled", v) }

// Suspend sets the onsuspend event handler.
func Suspend(v string) core.Node { return core.Attribute("onsuspend", v) }

// Emptied sets the onemptied event handler.
func Emptied(v string) core.Node { return core.Attribute("onemptied", v) }
