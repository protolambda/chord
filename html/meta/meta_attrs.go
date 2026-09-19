package meta

import (
	"github.com/protolambda/chord/core/attr"
)

// Charset sets the charset attribute.
func Charset(v string) attr.Node { return attr.Name("charset").Value(v) }

// Content sets the content attribute.
func Content(v string) attr.Node { return attr.Name("content").Value(v) }

// HttpEquiv sets the http-equiv attribute.
func HttpEquiv(v string) attr.Node { return attr.Name("http-equiv").Value(v) }

// Name sets the name attribute.
func Name(v string) attr.Node { return attr.Name("name").Value(v) }
