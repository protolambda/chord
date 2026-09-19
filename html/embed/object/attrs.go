package object

import (
	"github.com/protolambda/chord/core/attr"
)

// Data sets the data attribute.
func Data(v string) attr.Node { return attr.Name("data").Value(v) }

// Type sets the type attribute.
func Type(v string) attr.Node { return attr.Name("type").Value(v) }

// Name sets the name attribute.
func Name(v string) attr.Node { return attr.Name("name").Value(v) }

// Width sets the width attribute.
func Width(v string) attr.Node { return attr.Name("width").Value(v) }

// Height sets the height attribute.
func Height(v string) attr.Node { return attr.Name("height").Value(v) }
