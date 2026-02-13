// Package table provides HTML table elements.
package table

import (
	"github.com/protolambda/chord/core/attrib"
	"github.com/protolambda/chord/core/elem"
)

// Table creates a table element.
func Table(attrs ...attrib.Node) elem.Scope { return elem.New("table", attrs...) }

// Caption creates a caption element.
func Caption(attrs ...attrib.Node) elem.Scope { return elem.New("caption", attrs...) }

// Colgroup creates a colgroup element.
func Colgroup(attrs ...attrib.Node) elem.Scope { return elem.New("colgroup", attrs...) }

// Col creates a col element (void).
func Col(attrs ...attrib.Node) elem.Node { return elem.Void("col", attrs...) }

// Tbody creates a tbody element.
func Tbody(attrs ...attrib.Node) elem.Scope { return elem.New("tbody", attrs...) }

// Thead creates a thead element.
func Thead(attrs ...attrib.Node) elem.Scope { return elem.New("thead", attrs...) }

// Tfoot creates a tfoot element.
func Tfoot(attrs ...attrib.Node) elem.Scope { return elem.New("tfoot", attrs...) }

// TR creates a tr element.
func TR(attrs ...attrib.Node) elem.Scope { return elem.New("tr", attrs...) }

// TD creates a td element.
func TD(attrs ...attrib.Node) elem.Scope { return elem.New("td", attrs...) }

// TH creates a th element.
func TH(attrs ...attrib.Node) elem.Scope { return elem.New("th", attrs...) }
