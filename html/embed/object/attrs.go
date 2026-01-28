package object

import "github.com/protolambda/chord/core"

// Data sets the data attribute.
func Data(v string) core.Node { return core.Attribute("data", v) }

// Type sets the type attribute.
func Type(v string) core.Node { return core.Attribute("type", v) }

// Name sets the name attribute.
func Name(v string) core.Node { return core.Attribute("name", v) }

// Width sets the width attribute.
func Width(v string) core.Node { return core.Attribute("width", v) }

// Height sets the height attribute.
func Height(v string) core.Node { return core.Attribute("height", v) }
