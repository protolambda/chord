// Package selectel provides HTML select elements (select, optgroup, option, datalist).
package selectel

import (
	"github.com/protolambda/chord/core/attrib"
	"github.com/protolambda/chord/core/elem"
)

// Select creates a select element.
func Select(attrs ...attrib.Node) elem.Scope { return elem.New("select", attrs...) }

// Optgroup creates an optgroup element.
func Optgroup(attrs ...attrib.Node) elem.Scope { return elem.New("optgroup", attrs...) }

// Option creates an option element.
func Option(attrs ...attrib.Node) elem.Scope { return elem.New("option", attrs...) }

// Datalist creates a datalist element.
func Datalist(attrs ...attrib.Node) elem.Scope { return elem.New("datalist", attrs...) }
