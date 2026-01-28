package text

import "github.com/protolambda/chord/core"

// Datetime sets the datetime attribute for time elements.
func Datetime(v string) core.Node { return core.Attribute("datetime", v) }
