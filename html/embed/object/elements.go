// Package object provides HTML object elements (object, param, embed).
package object

import (
	"github.com/protolambda/chord/core/attr"
	"github.com/protolambda/chord/core/elem"
)

// Object creates an object element.
func Object(attrs ...attr.Node) elem.Scope { return elem.Name("object").New(attrs...) }

// Param creates a param element (void).
func Param(attrs ...attr.Node) elem.Node { return elem.Name("param").Void(attrs...) }

// Embed creates an embed element (void).
func Embed(attrs ...attr.Node) elem.Node { return elem.Name("embed").Void(attrs...) }
