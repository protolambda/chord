// Package img provides HTML image elements.
package img

import "github.com/protolambda/chord/core"

// Img creates an img element (void).
func Img(opts ...core.Node) core.Node { return core.VoidElement("img", opts...) }

// Picture creates a picture element.
func Picture(opts ...core.Node) core.Node { return core.Element("picture", opts...) }

// Source creates a source element (void).
func Source(opts ...core.Node) core.Node { return core.VoidElement("source", opts...) }
