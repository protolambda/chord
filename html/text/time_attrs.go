package text

import (
	"github.com/protolambda/chord/core/attr"
)

// Datetime sets the datetime attribute for time elements.
func Datetime(v string) attr.Node { return attr.KV("datetime", v) }
