package hx2

import "github.com/protolambda/chord/core"

// On sets the hx-on:* attribute for HTMX v2 event binding.
// This generates hx-on:event="handler" (e.g., hx-on:click="...").
//
// In v2, event names should be kebab-case: "htmx:before-request" not "htmx:beforeRequest".
//
// Note: This is distinct from the on.Click() function which generates
// standard DOM event attributes like onclick="handler".
// Use hx2.On for HTMX-specific event handling that integrates with
// HTMX's event system and lifecycle.
func On(event, handler string) core.Node { return core.Attribute("hx-on:"+event, handler) }

// Inherit sets the hx-inherit attribute to control attribute inheritance.
func Inherit(v string) core.Node { return core.Attribute("hx-inherit", v) }
