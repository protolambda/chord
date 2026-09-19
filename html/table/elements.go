// Package table provides HTML table elements.
package table

import (
	"github.com/protolambda/chord/core/attr"
	"github.com/protolambda/chord/core/elem"
)

// Table creates a table element.
func Table(attrs ...attr.Node) elem.Scope { return elem.Name("table").New(attrs...) }

// Caption creates a caption element.
func Caption(attrs ...attr.Node) elem.Scope { return elem.Name("caption").New(attrs...) }

// Colgroup creates a colgroup element.
func Colgroup(attrs ...attr.Node) elem.Scope { return elem.Name("colgroup").New(attrs...) }

// Col creates a col element (void).
func Col(attrs ...attr.Node) elem.Node { return elem.Name("col").Void(attrs...) }

// Tbody creates a tbody element.
func Tbody(attrs ...attr.Node) elem.Scope { return elem.Name("tbody").New(attrs...) }

// Thead creates a thead element.
func Thead(attrs ...attr.Node) elem.Scope { return elem.Name("thead").New(attrs...) }

// Tfoot creates a tfoot element.
func Tfoot(attrs ...attr.Node) elem.Scope { return elem.Name("tfoot").New(attrs...) }

// TR creates a tr element.
func TR(attrs ...attr.Node) elem.Scope { return elem.Name("tr").New(attrs...) }

// TD creates a td element.
func TD(attrs ...attr.Node) elem.Scope { return elem.Name("td").New(attrs...) }

// TH creates a th element.
func TH(attrs ...attr.Node) elem.Scope { return elem.Name("th").New(attrs...) }
