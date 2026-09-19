package script

import (
	"github.com/protolambda/chord/core/attr"
)

// Src sets the src attribute.
func Src(v string) attr.Node { return attr.Name("src").Value(v) }

// Type sets the type attribute.
func Type(v string) attr.Node { return attr.Name("type").Value(v) }

// Async sets the async boolean attribute.
func Async() attr.Node { return attr.Name("async").Bool() }

// Defer sets the defer boolean attribute.
func Defer() attr.Node { return attr.Name("defer").Bool() }

// Crossorigin sets the crossorigin attribute.
func Crossorigin(v string) attr.Node { return attr.Name("crossorigin").Value(v) }

// Integrity sets the integrity attribute.
func Integrity(v string) attr.Node { return attr.Name("integrity").Value(v) }

// Nomodule sets the nomodule boolean attribute.
func Nomodule() attr.Node { return attr.Name("nomodule").Bool() }

// Referrerpolicy sets the referrerpolicy attribute.
func Referrerpolicy(v string) attr.Node { return attr.Name("referrerpolicy").Value(v) }

// Blocking sets the blocking attribute.
func Blocking(v string) attr.Node { return attr.Name("blocking").Value(v) }
