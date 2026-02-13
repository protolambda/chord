package meta

import (
	"github.com/protolambda/chord/core/attrib"
)

// Charset sets the charset attribute.
func Charset(v string) attrib.Node { return attrib.KV("charset", v) }

// Content sets the content attribute.
func Content(v string) attrib.Node { return attrib.KV("content", v) }

// HttpEquiv sets the http-equiv attribute.
func HttpEquiv(v string) attrib.Node { return attrib.KV("http-equiv", v) }

// Name sets the name attribute.
func Name(v string) attrib.Node { return attrib.KV("name", v) }
