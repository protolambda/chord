package meta

import "github.com/protolambda/chord/core"

// Charset sets the charset attribute.
func Charset(v string) core.Node { return core.Attribute("charset", v) }

// Content sets the content attribute.
func Content(v string) core.Node { return core.Attribute("content", v) }

// HttpEquiv sets the http-equiv attribute.
func HttpEquiv(v string) core.Node { return core.Attribute("http-equiv", v) }

// Name sets the name attribute.
func Name(v string) core.Node { return core.Attribute("name", v) }
