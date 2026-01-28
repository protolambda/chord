// Package aria provides ARIA attributes for accessibility.
package aria

import "github.com/protolambda/chord/core"

// Role sets the role attribute.
func Role(v string) core.Node { return core.Attribute("role", v) }

// Label sets the aria-label attribute.
func Label(v string) core.Node { return core.Attribute("aria-label", v) }

// Labelledby sets the aria-labelledby attribute.
func Labelledby(v string) core.Node { return core.Attribute("aria-labelledby", v) }

// Describedby sets the aria-describedby attribute.
func Describedby(v string) core.Node { return core.Attribute("aria-describedby", v) }

// Hidden sets the aria-hidden attribute.
func Hidden(v string) core.Node { return core.Attribute("aria-hidden", v) }

// Live sets the aria-live attribute.
func Live(v string) core.Node { return core.Attribute("aria-live", v) }

// Atomic sets the aria-atomic attribute.
func Atomic(v string) core.Node { return core.Attribute("aria-atomic", v) }

// Busy sets the aria-busy attribute.
func Busy(v string) core.Node { return core.Attribute("aria-busy", v) }

// Controls sets the aria-controls attribute.
func Controls(v string) core.Node { return core.Attribute("aria-controls", v) }

// Current sets the aria-current attribute.
func Current(v string) core.Node { return core.Attribute("aria-current", v) }

// Disabled sets the aria-disabled attribute.
func Disabled(v string) core.Node { return core.Attribute("aria-disabled", v) }

// Expanded sets the aria-expanded attribute.
func Expanded(v string) core.Node { return core.Attribute("aria-expanded", v) }

// Haspopup sets the aria-haspopup attribute.
func Haspopup(v string) core.Node { return core.Attribute("aria-haspopup", v) }

// Invalid sets the aria-invalid attribute.
func Invalid(v string) core.Node { return core.Attribute("aria-invalid", v) }

// Modal sets the aria-modal attribute.
func Modal(v string) core.Node { return core.Attribute("aria-modal", v) }

// Pressed sets the aria-pressed attribute.
func Pressed(v string) core.Node { return core.Attribute("aria-pressed", v) }

// Selected sets the aria-selected attribute.
func Selected(v string) core.Node { return core.Attribute("aria-selected", v) }

// Valuenow sets the aria-valuenow attribute.
func Valuenow(v string) core.Node { return core.Attribute("aria-valuenow", v) }

// Valuemin sets the aria-valuemin attribute.
func Valuemin(v string) core.Node { return core.Attribute("aria-valuemin", v) }

// Valuemax sets the aria-valuemax attribute.
func Valuemax(v string) core.Node { return core.Attribute("aria-valuemax", v) }

// Valuetext sets the aria-valuetext attribute.
func Valuetext(v string) core.Node { return core.Attribute("aria-valuetext", v) }
