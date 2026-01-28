package output

import "github.com/protolambda/chord/core"

// For sets the for attribute for output elements.
func For(v string) core.Node { return core.Attribute("for", v) }

// Value sets the value attribute.
func Value(v string) core.Node { return core.Attribute("value", v) }

// Min sets the min attribute.
func Min(v string) core.Node { return core.Attribute("min", v) }

// Max sets the max attribute.
func Max(v string) core.Node { return core.Attribute("max", v) }

// Low sets the low attribute for meter elements.
func Low(v string) core.Node { return core.Attribute("low", v) }

// High sets the high attribute for meter elements.
func High(v string) core.Node { return core.Attribute("high", v) }

// Optimum sets the optimum attribute for meter elements.
func Optimum(v string) core.Node { return core.Attribute("optimum", v) }
