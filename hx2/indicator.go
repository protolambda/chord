package hx2

import (
	"github.com/protolambda/chord/core/attrib"
)

// Indicator sets the hx-indicator attribute to specify loading indicator element.
func Indicator(v string) attrib.Node { return attrib.KV("hx-indicator", v) }
