// Package table provides HTML table elements.
package table

import "github.com/protolambda/chord/core"

// Table creates a table element.
func Table(opts ...core.Node) core.Node { return core.Element("table", opts...) }

// Caption creates a caption element.
func Caption(opts ...core.Node) core.Node { return core.Element("caption", opts...) }

// Colgroup creates a colgroup element.
func Colgroup(opts ...core.Node) core.Node { return core.Element("colgroup", opts...) }

// Col creates a col element (void).
func Col(opts ...core.Node) core.Node { return core.VoidElement("col", opts...) }

// Tbody creates a tbody element.
func Tbody(opts ...core.Node) core.Node { return core.Element("tbody", opts...) }

// Thead creates a thead element.
func Thead(opts ...core.Node) core.Node { return core.Element("thead", opts...) }

// Tfoot creates a tfoot element.
func Tfoot(opts ...core.Node) core.Node { return core.Element("tfoot", opts...) }

// TR creates a tr element.
func TR(opts ...core.Node) core.Node { return core.Element("tr", opts...) }

// TD creates a td element.
func TD(opts ...core.Node) core.Node { return core.Element("td", opts...) }

// TH creates a th element.
func TH(opts ...core.Node) core.Node { return core.Element("th", opts...) }
