package on

import "github.com/protolambda/chord/core"

// Submit sets the onsubmit event handler.
func Submit(v string) core.Node { return core.Attribute("onsubmit", v) }

// Reset sets the onreset event handler.
func Reset(v string) core.Node { return core.Attribute("onreset", v) }

// Change sets the onchange event handler.
func Change(v string) core.Node { return core.Attribute("onchange", v) }

// Input sets the oninput event handler.
func Input(v string) core.Node { return core.Attribute("oninput", v) }

// Invalid sets the oninvalid event handler.
func Invalid(v string) core.Node { return core.Attribute("oninvalid", v) }

// Select sets the onselect event handler.
func Select(v string) core.Node { return core.Attribute("onselect", v) }
