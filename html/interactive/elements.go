// Package interactive provides HTML interactive elements.
package interactive

import (
	"github.com/protolambda/chord/core/attr"
	"github.com/protolambda/chord/core/elem"
)

// Details creates a details element.
func Details(attrs ...attr.Node) elem.Scope { return elem.Name("details").New(attrs...) }

// Summary creates a summary element.
func Summary(attrs ...attr.Node) elem.Scope { return elem.Name("summary").New(attrs...) }

// Dialog creates a dialog element.
func Dialog(attrs ...attr.Node) elem.Scope { return elem.Name("dialog").New(attrs...) }
