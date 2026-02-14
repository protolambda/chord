package hx2

import (
	"github.com/protolambda/chord/core/attr"
)

// Validate sets the hx-validate attribute to control form validation.
func Validate(v string) attr.Node { return attr.KV("hx-validate", v) }
