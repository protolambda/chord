package text

import "github.com/protolambda/chord/core"

// QuoteCite sets the cite attribute for blockquote/q elements.
func QuoteCite(v string) core.Node { return core.Attribute("cite", v) }
