// Package label provides the HTML label element and its attributes.
package label

import (
	"github.com/protolambda/chord/core/attrib"
	"github.com/protolambda/chord/core/elem"
)

// Label creates a label element.
func Label(attrs ...attrib.Node) elem.Scope { return elem.New("label", attrs...) }
