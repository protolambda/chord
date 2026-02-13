package text

import (
	"github.com/protolambda/chord/core/attrib"
)

// Datetime sets the datetime attribute for time elements.
func Datetime(v string) attrib.Node { return attrib.KV("datetime", v) }
