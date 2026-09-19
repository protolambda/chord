package form

import (
	"github.com/protolambda/chord/core/attr"
)

// Action sets the action attribute.
func Action(v string) attr.Node { return attr.Name("action").Value(v) }

// Target sets the target attribute.
func Target(v string) attr.Node { return attr.Name("target").Value(v) }

// Novalidate sets the novalidate boolean attribute.
func Novalidate() attr.Node { return attr.Name("novalidate").Bool() }

// Accept sets the accept attribute.
func Accept(v string) attr.Node { return attr.Name("accept").Value(v) }

// Acceptcharset sets the accept-charset attribute.
func Acceptcharset(v string) attr.Node { return attr.Name("accept-charset").Value(v) }
