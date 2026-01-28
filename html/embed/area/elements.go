// Package area provides HTML map and area elements.
package area

import "github.com/protolambda/chord/core"

// Map creates a map element.
func Map(opts ...core.Node) core.Node { return core.Element("map", opts...) }

// Area creates an area element (void).
func Area(opts ...core.Node) core.Node { return core.VoidElement("area", opts...) }
