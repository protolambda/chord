// Package area provides HTML map and area elements.
package area

import (
	"github.com/protolambda/chord/core/attr"
	"github.com/protolambda/chord/core/elem"
)

// Map creates a map element.
func Map(attrs ...attr.Node) elem.Scope { return elem.New("map", attrs...) }

// Area creates an area element (void).
func Area(attrs ...attr.Node) elem.Node { return elem.Void("area", attrs...) }
