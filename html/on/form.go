package on

import (
	"github.com/protolambda/chord/core/attr"
)

// Submit sets the onsubmit event handler.
func Submit(v string) attr.Node { return attr.KV("onsubmit", v) }

// Reset sets the onreset event handler.
func Reset(v string) attr.Node { return attr.KV("onreset", v) }

// Change sets the onchange event handler.
func Change(v string) attr.Node { return attr.KV("onchange", v) }

// Input sets the oninput event handler.
func Input(v string) attr.Node { return attr.KV("oninput", v) }

// Invalid sets the oninvalid event handler.
func Invalid(v string) attr.Node { return attr.KV("oninvalid", v) }

// Select sets the onselect event handler.
func Select(v string) attr.Node { return attr.KV("onselect", v) }
