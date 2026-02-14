package output

import (
	"github.com/protolambda/chord/core/attr"
)

// For sets the for attribute for output elements.
func For(v string) attr.Node { return attr.KV("for", v) }

// Value sets the value attribute.
func Value(v string) attr.Node { return attr.KV("value", v) }

// Min sets the min attribute.
func Min(v string) attr.Node { return attr.KV("min", v) }

// Max sets the max attribute.
func Max(v string) attr.Node { return attr.KV("max", v) }

// Low sets the low attribute for meter elements.
func Low(v string) attr.Node { return attr.KV("low", v) }

// High sets the high attribute for meter elements.
func High(v string) attr.Node { return attr.KV("high", v) }

// Optimum sets the optimum attribute for meter elements.
func Optimum(v string) attr.Node { return attr.KV("optimum", v) }
