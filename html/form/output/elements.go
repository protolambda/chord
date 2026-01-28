// Package output provides HTML output elements (output, progress, meter).
package output

import "github.com/protolambda/chord/core"

// Output creates an output element.
func Output(opts ...core.Node) core.Node { return core.Element("output", opts...) }

// Progress creates a progress element.
func Progress(opts ...core.Node) core.Node { return core.Element("progress", opts...) }

// Meter creates a meter element.
func Meter(opts ...core.Node) core.Node { return core.Element("meter", opts...) }
