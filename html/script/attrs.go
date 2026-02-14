package script

import (
	"github.com/protolambda/chord/core/attr"
)

// Src sets the src attribute.
func Src(v string) attr.Node { return attr.KV("src", v) }

// Type sets the type attribute.
func Type(v string) attr.Node { return attr.KV("type", v) }

// Async sets the async boolean attribute.
func Async() attr.Node { return attr.Bool("async") }

// Defer sets the defer boolean attribute.
func Defer() attr.Node { return attr.Bool("defer") }

// Crossorigin sets the crossorigin attribute.
func Crossorigin(v string) attr.Node { return attr.KV("crossorigin", v) }

// Integrity sets the integrity attribute.
func Integrity(v string) attr.Node { return attr.KV("integrity", v) }

// Nomodule sets the nomodule boolean attribute.
func Nomodule() attr.Node { return attr.Bool("nomodule") }

// Referrerpolicy sets the referrerpolicy attribute.
func Referrerpolicy(v string) attr.Node { return attr.KV("referrerpolicy", v) }

// Blocking sets the blocking attribute.
func Blocking(v string) attr.Node { return attr.KV("blocking", v) }
