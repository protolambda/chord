// Package script provides HTML scripting elements.
package script

import (
	"github.com/protolambda/chord/core/attr"
	"github.com/protolambda/chord/core/elem"
)

// Script creates a script element. Text children are written literally, as
// browsers read script code; see [Inline].
func Script(attrs ...attr.Node) elem.Scope { return elem.Name("script").New(attrs...) }

// Inline creates a script element with inline code:
//
//	script.Inline(`document.body.dataset.ready = "1";`, script.Type("module"))
//
// It is the same as Script(attrs...)(elem.Text(js)). The code is written
// literally, not HTML-escaped, because browsers do not unescape script
// content. Rendering fails with [elem.ErrUnsafeText] when the code contains
// "</script" or "<!--" (ASCII case-insensitive), which would end the element
// early or change how it is parsed; write "<\/script" and "\x3C!--" in
// JavaScript strings instead. Inside an element whose content is text, such
// as noscript, its end tag is refused as well.
//
// The code runs with the page's privileges: never build it from untrusted
// input. Pass data in a data attribute, or as JSON from encoding/json, which
// escapes "<".
func Inline(js string, attrs ...attr.Node) elem.Node {
	return Script(attrs...)(elem.Text(js))
}

// Noscript creates a noscript element.
func Noscript(attrs ...attr.Node) elem.Scope { return elem.Name("noscript").New(attrs...) }

// Template creates a template element.
func Template(attrs ...attr.Node) elem.Scope { return elem.Name("template").New(attrs...) }

// Canvas creates a canvas element.
func Canvas(attrs ...attr.Node) elem.Scope { return elem.Name("canvas").New(attrs...) }
