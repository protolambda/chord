package hx1

import (
	"github.com/protolambda/chord/core/attrib"
)

// Validate sets the hx-validate attribute to control form validation.
func Validate(v string) attrib.Node { return attrib.KV("hx-validate", v) }
