// Package interactive provides HTML interactive elements.
package interactive

import "github.com/protolambda/chord/core"

// Details creates a details element.
func Details(opts ...core.Node) core.Node { return core.Element("details", opts...) }

// Summary creates a summary element.
func Summary(opts ...core.Node) core.Node { return core.Element("summary", opts...) }

// Dialog creates a dialog element.
func Dialog(opts ...core.Node) core.Node { return core.Element("dialog", opts...) }
