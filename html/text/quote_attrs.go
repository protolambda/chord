package text

import (
	"github.com/protolambda/chord/core/attr"
)

// QuoteCite sets the cite attribute for blockquote/q elements.
func QuoteCite(v string) attr.Node { return attr.KV("cite", v) }
