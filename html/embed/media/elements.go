// Package media provides HTML media elements (video, audio, track).
package media

import (
	"github.com/protolambda/chord/core/attr"
	"github.com/protolambda/chord/core/elem"
)

// Video creates a video element.
func Video(attrs ...attr.Node) elem.Scope { return elem.New("video", attrs...) }

// Audio creates an audio element.
func Audio(attrs ...attr.Node) elem.Scope { return elem.New("audio", attrs...) }

// Track creates a track element (void).
func Track(attrs ...attr.Node) elem.Node { return elem.Void("track", attrs...) }

// Source creates a source element (void), re-exported for media context.
func Source(attrs ...attr.Node) elem.Node { return elem.Void("source", attrs...) }
