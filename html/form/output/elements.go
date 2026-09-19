// Package output provides HTML output elements (output, progress, meter).
package output

import (
	"github.com/protolambda/chord/core/attr"
	"github.com/protolambda/chord/core/elem"
)

// Output creates an output element.
func Output(attrs ...attr.Node) elem.Scope { return elem.Name("output").New(attrs...) }

// Progress creates a progress element.
func Progress(attrs ...attr.Node) elem.Scope { return elem.Name("progress").New(attrs...) }

// Meter creates a meter element.
func Meter(attrs ...attr.Node) elem.Scope { return elem.Name("meter").New(attrs...) }
