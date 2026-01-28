// Package selectel provides HTML select elements (select, optgroup, option, datalist).
package selectel

import "github.com/protolambda/chord/core"

// Select creates a select element.
func Select(opts ...core.Node) core.Node { return core.Element("select", opts...) }

// Optgroup creates an optgroup element.
func Optgroup(opts ...core.Node) core.Node { return core.Element("optgroup", opts...) }

// Option creates an option element.
func Option(opts ...core.Node) core.Node { return core.Element("option", opts...) }

// Datalist creates a datalist element.
func Datalist(opts ...core.Node) core.Node { return core.Element("datalist", opts...) }
