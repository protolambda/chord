package table

import (
	"github.com/protolambda/chord/core/attr"
)

// Colspan sets the colspan attribute.
func Colspan(v string) attr.Node { return attr.Name("colspan").Value(v) }

// Rowspan sets the rowspan attribute.
func Rowspan(v string) attr.Node { return attr.Name("rowspan").Value(v) }

// Headers sets the headers attribute.
func Headers(v string) attr.Node { return attr.Name("headers").Value(v) }

// Scope sets the scope attribute for th elements.
func Scope(v string) attr.Node { return attr.Name("scope").Value(v) }

// Span sets the span attribute for col/colgroup elements.
func Span(v string) attr.Node { return attr.Name("span").Value(v) }
