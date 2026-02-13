package iframe

import (
	"github.com/protolambda/chord/core/attrib"
)

// Src sets the src attribute.
func Src(v string) attrib.Node { return attrib.KV("src", v) }

// Srcdoc sets the srcdoc attribute.
func Srcdoc(v string) attrib.Node { return attrib.KV("srcdoc", v) }

// Name sets the name attribute.
func Name(v string) attrib.Node { return attrib.KV("name", v) }

// Width sets the width attribute.
func Width(v string) attrib.Node { return attrib.KV("width", v) }

// Height sets the height attribute.
func Height(v string) attrib.Node { return attrib.KV("height", v) }

// Loading sets the loading attribute.
func Loading(v string) attrib.Node { return attrib.KV("loading", v) }

// Sandbox sets the sandbox attribute.
func Sandbox(v string) attrib.Node { return attrib.KV("sandbox", v) }

// Allow sets the allow attribute.
func Allow(v string) attrib.Node { return attrib.KV("allow", v) }

// Allowfullscreen sets the allowfullscreen boolean attribute.
func Allowfullscreen() attrib.Node { return attrib.Bool("allowfullscreen") }

// Referrerpolicy sets the referrerpolicy attribute.
func Referrerpolicy(v string) attrib.Node { return attrib.KV("referrerpolicy", v) }
