package hx1

import (
	"github.com/protolambda/chord/core/attrib"
)

// SSE sets the hx-sse attribute for Server-Sent Events connections.
func SSE(v string) attrib.Node { return attrib.KV("hx-sse", v) }

// WS sets the hx-ws attribute for WebSocket connections.
func WS(v string) attrib.Node { return attrib.KV("hx-ws", v) }
