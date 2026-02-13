package script

import (
	"github.com/protolambda/chord/core/attrib"
)

// Src sets the src attribute.
func Src(v string) attrib.Node { return attrib.KV("src", v) }

// Type sets the type attribute.
func Type(v string) attrib.Node { return attrib.KV("type", v) }

// Async sets the async boolean attribute.
func Async() attrib.Node { return attrib.Bool("async") }

// Defer sets the defer boolean attribute.
func Defer() attrib.Node { return attrib.Bool("defer") }

// Crossorigin sets the crossorigin attribute.
func Crossorigin(v string) attrib.Node { return attrib.KV("crossorigin", v) }

// Integrity sets the integrity attribute.
func Integrity(v string) attrib.Node { return attrib.KV("integrity", v) }

// Nomodule sets the nomodule boolean attribute.
func Nomodule() attrib.Node { return attrib.Bool("nomodule") }

// Referrerpolicy sets the referrerpolicy attribute.
func Referrerpolicy(v string) attrib.Node { return attrib.KV("referrerpolicy", v) }

// Blocking sets the blocking attribute.
func Blocking(v string) attrib.Node { return attrib.KV("blocking", v) }
