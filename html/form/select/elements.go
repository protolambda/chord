// Package selectel provides HTML select elements (select, optgroup, option, datalist).
package selectel

import (
	"github.com/protolambda/chord/core/attr"
	"github.com/protolambda/chord/core/elem"
)

// Select creates a select element.
func Select(attrs ...attr.Node) elem.Scope { return elem.New("select", attrs...) }

// Optgroup creates an optgroup element.
func Optgroup(attrs ...attr.Node) elem.Scope { return elem.New("optgroup", attrs...) }

// Option creates an option element.
func Option(attrs ...attr.Node) elem.Scope { return elem.New("option", attrs...) }

// Datalist creates a datalist element.
func Datalist(attrs ...attr.Node) elem.Scope { return elem.New("datalist", attrs...) }
