// Package aria provides ARIA attributes for accessibility.
package aria

import (
	"github.com/protolambda/chord/core/attrib"
)

// Role sets the role attribute.
func Role(v string) attrib.Node { return attrib.KV("role", v) }

// Label sets the aria-label attribute.
func Label(v string) attrib.Node { return attrib.KV("aria-label", v) }

// Labelledby sets the aria-labelledby attribute.
func Labelledby(v string) attrib.Node { return attrib.KV("aria-labelledby", v) }

// Describedby sets the aria-describedby attribute.
func Describedby(v string) attrib.Node { return attrib.KV("aria-describedby", v) }

// Hidden sets the aria-hidden attribute.
func Hidden(v string) attrib.Node { return attrib.KV("aria-hidden", v) }

// Live sets the aria-live attribute.
func Live(v string) attrib.Node { return attrib.KV("aria-live", v) }

// Atomic sets the aria-atomic attribute.
func Atomic(v string) attrib.Node { return attrib.KV("aria-atomic", v) }

// Busy sets the aria-busy attribute.
func Busy(v string) attrib.Node { return attrib.KV("aria-busy", v) }

// Controls sets the aria-controls attribute.
func Controls(v string) attrib.Node { return attrib.KV("aria-controls", v) }

// Current sets the aria-current attribute.
func Current(v string) attrib.Node { return attrib.KV("aria-current", v) }

// Disabled sets the aria-disabled attribute.
func Disabled(v string) attrib.Node { return attrib.KV("aria-disabled", v) }

// Expanded sets the aria-expanded attribute.
func Expanded(v string) attrib.Node { return attrib.KV("aria-expanded", v) }

// Haspopup sets the aria-haspopup attribute.
func Haspopup(v string) attrib.Node { return attrib.KV("aria-haspopup", v) }

// Invalid sets the aria-invalid attribute.
func Invalid(v string) attrib.Node { return attrib.KV("aria-invalid", v) }

// Modal sets the aria-modal attribute.
func Modal(v string) attrib.Node { return attrib.KV("aria-modal", v) }

// Pressed sets the aria-pressed attribute.
func Pressed(v string) attrib.Node { return attrib.KV("aria-pressed", v) }

// Selected sets the aria-selected attribute.
func Selected(v string) attrib.Node { return attrib.KV("aria-selected", v) }

// Valuenow sets the aria-valuenow attribute.
func Valuenow(v string) attrib.Node { return attrib.KV("aria-valuenow", v) }

// Valuemin sets the aria-valuemin attribute.
func Valuemin(v string) attrib.Node { return attrib.KV("aria-valuemin", v) }

// Valuemax sets the aria-valuemax attribute.
func Valuemax(v string) attrib.Node { return attrib.KV("aria-valuemax", v) }

// Valuetext sets the aria-valuetext attribute.
func Valuetext(v string) attrib.Node { return attrib.KV("aria-valuetext", v) }
