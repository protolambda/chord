package hx1

import "github.com/protolambda/chord/core"

// SSE sets the hx-sse attribute for Server-Sent Events connections.
func SSE(v string) core.Node { return core.Attribute("hx-sse", v) }

// WS sets the hx-ws attribute for WebSocket connections.
func WS(v string) core.Node { return core.Attribute("hx-ws", v) }
