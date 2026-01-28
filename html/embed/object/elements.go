// Package object provides HTML object elements (object, param, embed).
package object

import "github.com/protolambda/chord/core"

// Object creates an object element.
func Object(opts ...core.Node) core.Node { return core.Element("object", opts...) }

// Param creates a param element (void).
func Param(opts ...core.Node) core.Node { return core.VoidElement("param", opts...) }

// Embed creates an embed element (void).
func Embed(opts ...core.Node) core.Node { return core.VoidElement("embed", opts...) }
