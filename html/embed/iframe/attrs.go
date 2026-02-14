package iframe

import (
	"github.com/protolambda/chord/core/attr"
)

// Src sets the src attribute.
func Src(v string) attr.Node { return attr.KV("src", v) }

// Srcdoc sets the srcdoc attribute.
func Srcdoc(v string) attr.Node { return attr.KV("srcdoc", v) }

// Name sets the name attribute.
func Name(v string) attr.Node { return attr.KV("name", v) }

// Width sets the width attribute.
func Width(v string) attr.Node { return attr.KV("width", v) }

// Height sets the height attribute.
func Height(v string) attr.Node { return attr.KV("height", v) }

// Loading sets the loading attribute.
func Loading(v string) attr.Node { return attr.KV("loading", v) }

// Sandbox sets the sandbox attribute.
func Sandbox(v string) attr.Node { return attr.KV("sandbox", v) }

// Allow sets the allow attribute.
func Allow(v string) attr.Node { return attr.KV("allow", v) }

// Allowfullscreen sets the allowfullscreen boolean attribute.
func Allowfullscreen() attr.Node { return attr.Bool("allowfullscreen") }

// Referrerpolicy sets the referrerpolicy attribute.
func Referrerpolicy(v string) attr.Node { return attr.KV("referrerpolicy", v) }
