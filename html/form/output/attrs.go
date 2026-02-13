package output

import (
	"github.com/protolambda/chord/core/attrib"
)

// For sets the for attribute for output elements.
func For(v string) attrib.Node { return attrib.KV("for", v) }

// Value sets the value attribute.
func Value(v string) attrib.Node { return attrib.KV("value", v) }

// Min sets the min attribute.
func Min(v string) attrib.Node { return attrib.KV("min", v) }

// Max sets the max attribute.
func Max(v string) attrib.Node { return attrib.KV("max", v) }

// Low sets the low attribute for meter elements.
func Low(v string) attrib.Node { return attrib.KV("low", v) }

// High sets the high attribute for meter elements.
func High(v string) attrib.Node { return attrib.KV("high", v) }

// Optimum sets the optimum attribute for meter elements.
func Optimum(v string) attrib.Node { return attrib.KV("optimum", v) }
