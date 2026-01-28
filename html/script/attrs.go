package script

import "github.com/protolambda/chord/core"

// Src sets the src attribute.
func Src(v string) core.Node { return core.Attribute("src", v) }

// Type sets the type attribute.
func Type(v string) core.Node { return core.Attribute("type", v) }

// Async sets the async boolean attribute.
func Async() core.Node { return core.BoolAttribute("async") }

// Defer sets the defer boolean attribute.
func Defer() core.Node { return core.BoolAttribute("defer") }

// Crossorigin sets the crossorigin attribute.
func Crossorigin(v string) core.Node { return core.Attribute("crossorigin", v) }

// Integrity sets the integrity attribute.
func Integrity(v string) core.Node { return core.Attribute("integrity", v) }

// Nomodule sets the nomodule boolean attribute.
func Nomodule() core.Node { return core.BoolAttribute("nomodule") }

// Referrerpolicy sets the referrerpolicy attribute.
func Referrerpolicy(v string) core.Node { return core.Attribute("referrerpolicy", v) }

// Blocking sets the blocking attribute.
func Blocking(v string) core.Node { return core.Attribute("blocking", v) }
