// Package interactive provides HTML interactive elements.
package interactive

import (
	"github.com/protolambda/chord/core/attrib"
	"github.com/protolambda/chord/core/elem"
)

// Details creates a details element.
func Details(attrs ...attrib.Node) elem.Scope { return elem.New("details", attrs...) }

// Summary creates a summary element.
func Summary(attrs ...attrib.Node) elem.Scope { return elem.New("summary", attrs...) }

// Dialog creates a dialog element.
func Dialog(attrs ...attrib.Node) elem.Scope { return elem.New("dialog", attrs...) }
