package hx1

import (
	"github.com/protolambda/chord/core/attr"
)

// Boost sets the hx-boost attribute to convert links and forms to AJAX.
func Boost(v string) attr.Node { return attr.KV("hx-boost", v) }

// Preserve sets the hx-preserve boolean attribute to preserve element across swaps.
func Preserve() attr.Node { return attr.Bool("hx-preserve") }

// Confirm sets the hx-confirm attribute to show confirmation dialog before request.
func Confirm(v string) attr.Node { return attr.KV("hx-confirm", v) }

// Prompt sets the hx-prompt attribute to show prompt dialog before request.
func Prompt(v string) attr.Node { return attr.KV("hx-prompt", v) }

// Disable sets the hx-disable boolean attribute to disable HTMX processing.
func Disable() attr.Node { return attr.Bool("hx-disable") }

// DisabledElt sets the hx-disabled-elt attribute to disable elements during request.
func DisabledElt(v string) attr.Node { return attr.KV("hx-disabled-elt", v) }

// Disinherit sets the hx-disinherit attribute to prevent inheritance of attributes.
func Disinherit(v string) attr.Node { return attr.KV("hx-disinherit", v) }

// Encoding sets the hx-encoding attribute to set request encoding.
func Encoding(v string) attr.Node { return attr.KV("hx-encoding", v) }

// Ext sets the hx-ext attribute to enable HTMX extensions.
func Ext(v string) attr.Node { return attr.KV("hx-ext", v) }
