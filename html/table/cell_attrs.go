package table

import (
	"github.com/protolambda/chord/core/attr"
)

// Colspan sets the colspan attribute.
func Colspan(v string) attr.Node { return attr.KV("colspan", v) }

// Rowspan sets the rowspan attribute.
func Rowspan(v string) attr.Node { return attr.KV("rowspan", v) }

// Headers sets the headers attribute.
func Headers(v string) attr.Node { return attr.KV("headers", v) }

// Scope sets the scope attribute for th elements.
func Scope(v string) attr.Node { return attr.KV("scope", v) }

// Span sets the span attribute for col/colgroup elements.
func Span(v string) attr.Node { return attr.KV("span", v) }
