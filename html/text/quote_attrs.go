package text

import (
	"github.com/protolambda/chord/core/attrib"
)

// QuoteCite sets the cite attribute for blockquote/q elements.
func QuoteCite(v string) attrib.Node { return attrib.KV("cite", v) }
