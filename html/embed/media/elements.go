// Package media provides HTML media elements (video, audio, track).
package media

import "github.com/protolambda/chord/core"

// Video creates a video element.
func Video(opts ...core.Node) core.Node { return core.Element("video", opts...) }

// Audio creates an audio element.
func Audio(opts ...core.Node) core.Node { return core.Element("audio", opts...) }

// Track creates a track element (void).
func Track(opts ...core.Node) core.Node { return core.VoidElement("track", opts...) }

// Source creates a source element (void), re-exported for media context.
func Source(opts ...core.Node) core.Node { return core.VoidElement("source", opts...) }
