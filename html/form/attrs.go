package form

import (
	"github.com/protolambda/chord/core/attr"
)

// Action sets the action attribute.
func Action(v string) attr.Node { return attr.KV("action", v) }

// Target sets the target attribute.
func Target(v string) attr.Node { return attr.KV("target", v) }

// Novalidate sets the novalidate boolean attribute.
func Novalidate() attr.Node { return attr.Bool("novalidate") }

// Accept sets the accept attribute.
func Accept(v string) attr.Node { return attr.KV("accept", v) }

// Acceptcharset sets the accept-charset attribute.
func Acceptcharset(v string) attr.Node { return attr.KV("accept-charset", v) }
