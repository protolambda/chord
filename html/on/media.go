package on

import (
	"github.com/protolambda/chord/core/attr"
)

// Play sets the onplay event handler.
func Play(v string) attr.Node { return attr.KV("onplay", v) }

// Pause sets the onpause event handler.
func Pause(v string) attr.Node { return attr.KV("onpause", v) }

// Ended sets the onended event handler.
func Ended(v string) attr.Node { return attr.KV("onended", v) }

// VolumeChange sets the onvolumechange event handler.
func VolumeChange(v string) attr.Node { return attr.KV("onvolumechange", v) }

// TimeUpdate sets the ontimeupdate event handler.
func TimeUpdate(v string) attr.Node { return attr.KV("ontimeupdate", v) }

// Seeking sets the onseeking event handler.
func Seeking(v string) attr.Node { return attr.KV("onseeking", v) }

// Seeked sets the onseeked event handler.
func Seeked(v string) attr.Node { return attr.KV("onseeked", v) }

// LoadedData sets the onloadeddata event handler.
func LoadedData(v string) attr.Node { return attr.KV("onloadeddata", v) }

// LoadedMetadata sets the onloadedmetadata event handler.
func LoadedMetadata(v string) attr.Node { return attr.KV("onloadedmetadata", v) }

// CanPlay sets the oncanplay event handler.
func CanPlay(v string) attr.Node { return attr.KV("oncanplay", v) }

// CanPlayThrough sets the oncanplaythrough event handler.
func CanPlayThrough(v string) attr.Node { return attr.KV("oncanplaythrough", v) }

// Waiting sets the onwaiting event handler.
func Waiting(v string) attr.Node { return attr.KV("onwaiting", v) }

// Playing sets the onplaying event handler.
func Playing(v string) attr.Node { return attr.KV("onplaying", v) }

// DurationChange sets the ondurationchange event handler.
func DurationChange(v string) attr.Node { return attr.KV("ondurationchange", v) }

// RateChange sets the onratechange event handler.
func RateChange(v string) attr.Node { return attr.KV("onratechange", v) }

// Stalled sets the onstalled event handler.
func Stalled(v string) attr.Node { return attr.KV("onstalled", v) }

// Suspend sets the onsuspend event handler.
func Suspend(v string) attr.Node { return attr.KV("onsuspend", v) }

// Emptied sets the onemptied event handler.
func Emptied(v string) attr.Node { return attr.KV("onemptied", v) }
