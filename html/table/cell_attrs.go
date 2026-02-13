package table

import (
	"github.com/protolambda/chord/core/attrib"
)

// Colspan sets the colspan attribute.
func Colspan(v string) attrib.Node { return attrib.KV("colspan", v) }

// Rowspan sets the rowspan attribute.
func Rowspan(v string) attrib.Node { return attrib.KV("rowspan", v) }

// Headers sets the headers attribute.
func Headers(v string) attrib.Node { return attrib.KV("headers", v) }

// Scope sets the scope attribute for th elements.
func Scope(v string) attrib.Node { return attrib.KV("scope", v) }

// Span sets the span attribute for col/colgroup elements.
func Span(v string) attrib.Node { return attrib.KV("span", v) }
