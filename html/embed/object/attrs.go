package object

import (
	"github.com/protolambda/chord/core/attrib"
)

// Data sets the data attribute.
func Data(v string) attrib.Node { return attrib.KV("data", v) }

// Type sets the type attribute.
func Type(v string) attrib.Node { return attrib.KV("type", v) }

// Name sets the name attribute.
func Name(v string) attrib.Node { return attrib.KV("name", v) }

// Width sets the width attribute.
func Width(v string) attrib.Node { return attrib.KV("width", v) }

// Height sets the height attribute.
func Height(v string) attrib.Node { return attrib.KV("height", v) }
