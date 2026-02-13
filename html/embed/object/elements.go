// Package object provides HTML object elements (object, param, embed).
package object

import (
	"github.com/protolambda/chord/core/attrib"
	"github.com/protolambda/chord/core/elem"
)

// Object creates an object element.
func Object(attrs ...attrib.Node) elem.Scope { return elem.New("object", attrs...) }

// Param creates a param element (void).
func Param(attrs ...attrib.Node) elem.Node { return elem.Void("param", attrs...) }

// Embed creates an embed element (void).
func Embed(attrs ...attrib.Node) elem.Node { return elem.Void("embed", attrs...) }
