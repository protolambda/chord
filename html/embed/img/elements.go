// Package img provides HTML image elements.
package img

import (
	"github.com/protolambda/chord/core/attr"
	"github.com/protolambda/chord/core/elem"
)

// Img creates an img element (void).
func Img(attrs ...attr.Node) elem.Node { return elem.Void("img", attrs...) }

// Picture creates a picture element.
func Picture(attrs ...attr.Node) elem.Scope { return elem.New("picture", attrs...) }

// Source creates a source element (void).
func Source(attrs ...attr.Node) elem.Node { return elem.Void("source", attrs...) }
