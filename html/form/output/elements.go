// Package output provides HTML output elements (output, progress, meter).
package output

import (
	"github.com/protolambda/chord/core/attrib"
	"github.com/protolambda/chord/core/elem"
)

// Output creates an output element.
func Output(attrs ...attrib.Node) elem.Scope { return elem.New("output", attrs...) }

// Progress creates a progress element.
func Progress(attrs ...attrib.Node) elem.Scope { return elem.New("progress", attrs...) }

// Meter creates a meter element.
func Meter(attrs ...attrib.Node) elem.Scope { return elem.New("meter", attrs...) }
