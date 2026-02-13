package on

import (
	"github.com/protolambda/chord/core/attrib"
)

// Submit sets the onsubmit event handler.
func Submit(v string) attrib.Node { return attrib.KV("onsubmit", v) }

// Reset sets the onreset event handler.
func Reset(v string) attrib.Node { return attrib.KV("onreset", v) }

// Change sets the onchange event handler.
func Change(v string) attrib.Node { return attrib.KV("onchange", v) }

// Input sets the oninput event handler.
func Input(v string) attrib.Node { return attrib.KV("oninput", v) }

// Invalid sets the oninvalid event handler.
func Invalid(v string) attrib.Node { return attrib.KV("oninvalid", v) }

// Select sets the onselect event handler.
func Select(v string) attrib.Node { return attrib.KV("onselect", v) }
