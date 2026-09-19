package on

import (
	"github.com/protolambda/chord/core/attr"
)

// Submit sets the onsubmit event handler.
func Submit(v string) attr.Node { return attr.Name("onsubmit").Value(v) }

// Reset sets the onreset event handler.
func Reset(v string) attr.Node { return attr.Name("onreset").Value(v) }

// Change sets the onchange event handler.
func Change(v string) attr.Node { return attr.Name("onchange").Value(v) }

// Input sets the oninput event handler.
func Input(v string) attr.Node { return attr.Name("oninput").Value(v) }

// Invalid sets the oninvalid event handler.
func Invalid(v string) attr.Node { return attr.Name("oninvalid").Value(v) }

// Select sets the onselect event handler.
func Select(v string) attr.Node { return attr.Name("onselect").Value(v) }
