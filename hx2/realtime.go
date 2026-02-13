package hx2

import (
	"github.com/protolambda/chord/core/attrib"
)

// SSE Extension Attributes
//
// In HTMX v2, SSE support was moved to an extension. You must include hx-ext="sse"
// on a parent element to use these attributes.
// Reference: https://htmx.org/extensions/sse/

// SSEConnect sets the sse-connect attribute to establish an SSE connection.
// Use with hx-ext="sse" on the element or a parent.
func SSEConnect(url string) attrib.Node { return attrib.KV("sse-connect", url) }

// SSESwap sets the sse-swap attribute to specify which SSE message triggers a swap.
// Use with hx-ext="sse" on the element or a parent.
func SSESwap(messageName string) attrib.Node { return attrib.KV("sse-swap", messageName) }

// SSEClose sets the sse-close attribute to close the connection on a specific message.
// Use with hx-ext="sse" on the element or a parent.
func SSEClose(messageName string) attrib.Node { return attrib.KV("sse-close", messageName) }

// WebSocket Extension Attributes
//
// In HTMX v2, WebSocket support was moved to an extension. You must include hx-ext="ws"
// on a parent element to use these attributes.
// Reference: https://htmx.org/extensions/ws/

// WSConnect sets the ws-connect attribute to establish a WebSocket connection.
// Use with hx-ext="ws" on the element or a parent.
func WSConnect(url string) attrib.Node { return attrib.KV("ws-connect", url) }

// WSSend sets the ws-send attribute to mark an element that sends data over WebSocket.
// Use with hx-ext="ws" on the element or a parent. Typically used on forms.
func WSSend() attrib.Node { return attrib.Bool("ws-send") }
