package on

import (
	"github.com/protolambda/chord/core/attrib"
)

// Play sets the onplay event handler.
func Play(v string) attrib.Node { return attrib.KV("onplay", v) }

// Pause sets the onpause event handler.
func Pause(v string) attrib.Node { return attrib.KV("onpause", v) }

// Ended sets the onended event handler.
func Ended(v string) attrib.Node { return attrib.KV("onended", v) }

// VolumeChange sets the onvolumechange event handler.
func VolumeChange(v string) attrib.Node { return attrib.KV("onvolumechange", v) }

// TimeUpdate sets the ontimeupdate event handler.
func TimeUpdate(v string) attrib.Node { return attrib.KV("ontimeupdate", v) }

// Seeking sets the onseeking event handler.
func Seeking(v string) attrib.Node { return attrib.KV("onseeking", v) }

// Seeked sets the onseeked event handler.
func Seeked(v string) attrib.Node { return attrib.KV("onseeked", v) }

// LoadedData sets the onloadeddata event handler.
func LoadedData(v string) attrib.Node { return attrib.KV("onloadeddata", v) }

// LoadedMetadata sets the onloadedmetadata event handler.
func LoadedMetadata(v string) attrib.Node { return attrib.KV("onloadedmetadata", v) }

// CanPlay sets the oncanplay event handler.
func CanPlay(v string) attrib.Node { return attrib.KV("oncanplay", v) }

// CanPlayThrough sets the oncanplaythrough event handler.
func CanPlayThrough(v string) attrib.Node { return attrib.KV("oncanplaythrough", v) }

// Waiting sets the onwaiting event handler.
func Waiting(v string) attrib.Node { return attrib.KV("onwaiting", v) }

// Playing sets the onplaying event handler.
func Playing(v string) attrib.Node { return attrib.KV("onplaying", v) }

// DurationChange sets the ondurationchange event handler.
func DurationChange(v string) attrib.Node { return attrib.KV("ondurationchange", v) }

// RateChange sets the onratechange event handler.
func RateChange(v string) attrib.Node { return attrib.KV("onratechange", v) }

// Stalled sets the onstalled event handler.
func Stalled(v string) attrib.Node { return attrib.KV("onstalled", v) }

// Suspend sets the onsuspend event handler.
func Suspend(v string) attrib.Node { return attrib.KV("onsuspend", v) }

// Emptied sets the onemptied event handler.
func Emptied(v string) attrib.Node { return attrib.KV("onemptied", v) }
