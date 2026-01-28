// Package hx2 provides HTMX v2 attributes for building hypermedia-driven applications.
//
// HTMX v2 is a major update with breaking changes from v1. This package provides
// the v2 attribute API independently from the v1 package.
//
// # Key Changes from v1
//
//   - Event binding uses [On] with hx-on:event syntax (kebab-case event names)
//   - SSE/WebSocket moved to extensions: use [Ext]("sse") or [Ext]("ws")
//   - SSE uses [SSEConnect], [SSESwap], [SSEClose] instead of hx-sse
//   - WebSocket uses [WSConnect], [WSSend] instead of hx-ws
//   - DELETE requests use URL params by default (not form-encoded body)
//   - Cross-domain requests disabled by default
//   - IE support removed
//
// # Core Request Attributes
//
//   - [Get], [Post], [Put], [Patch], [Delete]: HTTP request triggers
//
// # Common Attributes
//
//   - [Target]: CSS selector for response target
//   - [Swap]: How to swap the response into the DOM
//   - [Trigger]: Event that triggers the request
//   - [Indicator]: Element to show during request
//   - [Vals]: Additional values to submit
//
// # Event Binding (v2 syntax)
//
//   - [On]: hx-on:event-name="handler" (replaces v1's hx-on="event: handler")
//
// # Extensions
//
//   - [Ext]: Enable HTMX extensions (hx-ext attribute)
//   - [SSEConnect], [SSESwap], [SSEClose]: SSE extension attributes
//   - [WSConnect], [WSSend]: WebSocket extension attributes
//
// Reference: https://htmx.org/reference/
// Migration guide: https://htmx.org/migration-guide-htmx-1/
//
// For HTMX v1, see the [hx1] package.
package hx2
