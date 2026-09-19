// Package aria provides ARIA attributes for accessibility.
package aria

import (
	"github.com/protolambda/chord/core/attr"
)

// Role sets the role attribute.
func Role(v string) attr.Node { return attr.Name("role").Value(v) }

// Label sets the aria-label attribute.
func Label(v string) attr.Node { return attr.Name("aria-label").Value(v) }

// Labelledby sets the aria-labelledby attribute.
func Labelledby(v string) attr.Node { return attr.Name("aria-labelledby").Value(v) }

// Describedby sets the aria-describedby attribute.
func Describedby(v string) attr.Node { return attr.Name("aria-describedby").Value(v) }

// Hidden sets the aria-hidden attribute.
func Hidden(v string) attr.Node { return attr.Name("aria-hidden").Value(v) }

// Live sets the aria-live attribute.
func Live(v string) attr.Node { return attr.Name("aria-live").Value(v) }

// Atomic sets the aria-atomic attribute.
func Atomic(v string) attr.Node { return attr.Name("aria-atomic").Value(v) }

// Busy sets the aria-busy attribute.
func Busy(v string) attr.Node { return attr.Name("aria-busy").Value(v) }

// Controls sets the aria-controls attribute.
func Controls(v string) attr.Node { return attr.Name("aria-controls").Value(v) }

// Current sets the aria-current attribute.
func Current(v string) attr.Node { return attr.Name("aria-current").Value(v) }

// Disabled sets the aria-disabled attribute.
func Disabled(v string) attr.Node { return attr.Name("aria-disabled").Value(v) }

// Expanded sets the aria-expanded attribute.
func Expanded(v string) attr.Node { return attr.Name("aria-expanded").Value(v) }

// Haspopup sets the aria-haspopup attribute.
func Haspopup(v string) attr.Node { return attr.Name("aria-haspopup").Value(v) }

// Invalid sets the aria-invalid attribute.
func Invalid(v string) attr.Node { return attr.Name("aria-invalid").Value(v) }

// Modal sets the aria-modal attribute.
func Modal(v string) attr.Node { return attr.Name("aria-modal").Value(v) }

// Pressed sets the aria-pressed attribute.
func Pressed(v string) attr.Node { return attr.Name("aria-pressed").Value(v) }

// Selected sets the aria-selected attribute.
func Selected(v string) attr.Node { return attr.Name("aria-selected").Value(v) }

// Valuenow sets the aria-valuenow attribute.
func Valuenow(v string) attr.Node { return attr.Name("aria-valuenow").Value(v) }

// Valuemin sets the aria-valuemin attribute.
func Valuemin(v string) attr.Node { return attr.Name("aria-valuemin").Value(v) }

// Valuemax sets the aria-valuemax attribute.
func Valuemax(v string) attr.Node { return attr.Name("aria-valuemax").Value(v) }

// Valuetext sets the aria-valuetext attribute.
func Valuetext(v string) attr.Node { return attr.Name("aria-valuetext").Value(v) }
