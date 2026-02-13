package hx1

import (
	"github.com/protolambda/chord/core/attrib"
)

// Boost sets the hx-boost attribute to convert links and forms to AJAX.
func Boost(v string) attrib.Node { return attrib.KV("hx-boost", v) }

// Preserve sets the hx-preserve boolean attribute to preserve element across swaps.
func Preserve() attrib.Node { return attrib.Bool("hx-preserve") }

// Confirm sets the hx-confirm attribute to show confirmation dialog before request.
func Confirm(v string) attrib.Node { return attrib.KV("hx-confirm", v) }

// Prompt sets the hx-prompt attribute to show prompt dialog before request.
func Prompt(v string) attrib.Node { return attrib.KV("hx-prompt", v) }

// Disable sets the hx-disable boolean attribute to disable HTMX processing.
func Disable() attrib.Node { return attrib.Bool("hx-disable") }

// DisabledElt sets the hx-disabled-elt attribute to disable elements during request.
func DisabledElt(v string) attrib.Node { return attrib.KV("hx-disabled-elt", v) }

// Disinherit sets the hx-disinherit attribute to prevent inheritance of attributes.
func Disinherit(v string) attrib.Node { return attrib.KV("hx-disinherit", v) }

// Encoding sets the hx-encoding attribute to set request encoding.
func Encoding(v string) attrib.Node { return attrib.KV("hx-encoding", v) }

// Ext sets the hx-ext attribute to enable HTMX extensions.
func Ext(v string) attrib.Node { return attrib.KV("hx-ext", v) }
