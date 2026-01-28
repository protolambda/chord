package table

import "github.com/protolambda/chord/core"

// Colspan sets the colspan attribute.
func Colspan(v string) core.Node { return core.Attribute("colspan", v) }

// Rowspan sets the rowspan attribute.
func Rowspan(v string) core.Node { return core.Attribute("rowspan", v) }

// Headers sets the headers attribute.
func Headers(v string) core.Node { return core.Attribute("headers", v) }

// Scope sets the scope attribute for th elements.
func Scope(v string) core.Node { return core.Attribute("scope", v) }

// Span sets the span attribute for col/colgroup elements.
func Span(v string) core.Node { return core.Attribute("span", v) }
