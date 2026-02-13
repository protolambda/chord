package hx2

import (
	"github.com/protolambda/chord/core/attrib"
)

// Target sets the hx-target attribute to specify where to swap response content.
func Target(v string) attrib.Node { return attrib.KV("hx-target", v) }

// Swap sets the hx-swap attribute to control how response content is swapped.
func Swap(v string) attrib.Node { return attrib.KV("hx-swap", v) }

// SwapOOB sets the hx-swap-oob attribute for out-of-band swaps.
func SwapOOB(v string) attrib.Node { return attrib.KV("hx-swap-oob", v) }

// Select sets the hx-select attribute to select a subset of response content.
func Select(v string) attrib.Node { return attrib.KV("hx-select", v) }

// SelectOOB sets the hx-select-oob attribute for out-of-band selection.
func SelectOOB(v string) attrib.Node { return attrib.KV("hx-select-oob", v) }
