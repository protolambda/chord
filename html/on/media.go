package on

import (
	"github.com/protolambda/chord/core/attr"
)

// Play sets the onplay event handler.
func Play(v string) attr.Node { return attr.Name("onplay").Value(v) }

// Pause sets the onpause event handler.
func Pause(v string) attr.Node { return attr.Name("onpause").Value(v) }

// Ended sets the onended event handler.
func Ended(v string) attr.Node { return attr.Name("onended").Value(v) }

// VolumeChange sets the onvolumechange event handler.
func VolumeChange(v string) attr.Node { return attr.Name("onvolumechange").Value(v) }

// TimeUpdate sets the ontimeupdate event handler.
func TimeUpdate(v string) attr.Node { return attr.Name("ontimeupdate").Value(v) }

// Seeking sets the onseeking event handler.
func Seeking(v string) attr.Node { return attr.Name("onseeking").Value(v) }

// Seeked sets the onseeked event handler.
func Seeked(v string) attr.Node { return attr.Name("onseeked").Value(v) }

// LoadedData sets the onloadeddata event handler.
func LoadedData(v string) attr.Node { return attr.Name("onloadeddata").Value(v) }

// LoadedMetadata sets the onloadedmetadata event handler.
func LoadedMetadata(v string) attr.Node { return attr.Name("onloadedmetadata").Value(v) }

// CanPlay sets the oncanplay event handler.
func CanPlay(v string) attr.Node { return attr.Name("oncanplay").Value(v) }

// CanPlayThrough sets the oncanplaythrough event handler.
func CanPlayThrough(v string) attr.Node { return attr.Name("oncanplaythrough").Value(v) }

// Waiting sets the onwaiting event handler.
func Waiting(v string) attr.Node { return attr.Name("onwaiting").Value(v) }

// Playing sets the onplaying event handler.
func Playing(v string) attr.Node { return attr.Name("onplaying").Value(v) }

// DurationChange sets the ondurationchange event handler.
func DurationChange(v string) attr.Node { return attr.Name("ondurationchange").Value(v) }

// RateChange sets the onratechange event handler.
func RateChange(v string) attr.Node { return attr.Name("onratechange").Value(v) }

// Stalled sets the onstalled event handler.
func Stalled(v string) attr.Node { return attr.Name("onstalled").Value(v) }

// Suspend sets the onsuspend event handler.
func Suspend(v string) attr.Node { return attr.Name("onsuspend").Value(v) }

// Emptied sets the onemptied event handler.
func Emptied(v string) attr.Node { return attr.Name("onemptied").Value(v) }
