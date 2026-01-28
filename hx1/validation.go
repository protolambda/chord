package hx1

import "github.com/protolambda/chord/core"

// Validate sets the hx-validate attribute to control form validation.
func Validate(v string) core.Node { return core.Attribute("hx-validate", v) }
