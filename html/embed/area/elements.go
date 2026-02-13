// Package area provides HTML map and area elements.
package area

import (
	"github.com/protolambda/chord/core/attrib"
	"github.com/protolambda/chord/core/elem"
)

// Map creates a map element.
func Map(attrs ...attrib.Node) elem.Scope { return elem.New("map", attrs...) }

// Area creates an area element (void).
func Area(attrs ...attrib.Node) elem.Node { return elem.Void("area", attrs...) }
