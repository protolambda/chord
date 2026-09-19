package iframe

import (
	"github.com/protolambda/chord/core/attr"
)

// Src sets the src attribute.
func Src(v string) attr.Node { return attr.Name("src").Value(v) }

// Srcdoc sets the srcdoc attribute.
func Srcdoc(v string) attr.Node { return attr.Name("srcdoc").Value(v) }

// Name sets the name attribute.
func Name(v string) attr.Node { return attr.Name("name").Value(v) }

// Width sets the width attribute.
func Width(v string) attr.Node { return attr.Name("width").Value(v) }

// Height sets the height attribute.
func Height(v string) attr.Node { return attr.Name("height").Value(v) }

// Loading sets the loading attribute.
func Loading(v string) attr.Node { return attr.Name("loading").Value(v) }

// Sandbox sets the sandbox attribute.
func Sandbox(v string) attr.Node { return attr.Name("sandbox").Value(v) }

// Allow sets the allow attribute.
func Allow(v string) attr.Node { return attr.Name("allow").Value(v) }

// Allowfullscreen sets the allowfullscreen boolean attribute.
func Allowfullscreen() attr.Node { return attr.Name("allowfullscreen").Bool() }

// Referrerpolicy sets the referrerpolicy attribute.
func Referrerpolicy(v string) attr.Node { return attr.Name("referrerpolicy").Value(v) }
