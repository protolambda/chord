// Package label provides the HTML label element and its attributes.
package label

import (
	"github.com/protolambda/chord/core/attr"
	"github.com/protolambda/chord/core/elem"
)

// Label creates a label element.
func Label(attrs ...attr.Node) elem.Scope { return elem.New("label", attrs...) }
