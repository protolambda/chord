package object

import (
	"github.com/protolambda/chord/core/attr"
)

// Data sets the data attribute.
func Data(v string) attr.Node { return attr.KV("data", v) }

// Type sets the type attribute.
func Type(v string) attr.Node { return attr.KV("type", v) }

// Name sets the name attribute.
func Name(v string) attr.Node { return attr.KV("name", v) }

// Width sets the width attribute.
func Width(v string) attr.Node { return attr.KV("width", v) }

// Height sets the height attribute.
func Height(v string) attr.Node { return attr.KV("height", v) }
