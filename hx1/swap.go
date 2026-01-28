package hx1

import "github.com/protolambda/chord/core"

// Target sets the hx-target attribute to specify where to swap response content.
func Target(v string) core.Node { return core.Attribute("hx-target", v) }

// Swap sets the hx-swap attribute to control how response content is swapped.
func Swap(v string) core.Node { return core.Attribute("hx-swap", v) }

// SwapOOB sets the hx-swap-oob attribute for out-of-band swaps.
func SwapOOB(v string) core.Node { return core.Attribute("hx-swap-oob", v) }

// Select sets the hx-select attribute to select a subset of response content.
func Select(v string) core.Node { return core.Attribute("hx-select", v) }

// SelectOOB sets the hx-select-oob attribute for out-of-band selection.
func SelectOOB(v string) core.Node { return core.Attribute("hx-select-oob", v) }
