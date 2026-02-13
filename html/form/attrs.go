package form

import (
	"github.com/protolambda/chord/core/attrib"
)

// Action sets the action attribute.
func Action(v string) attrib.Node { return attrib.KV("action", v) }

// Target sets the target attribute.
func Target(v string) attrib.Node { return attrib.KV("target", v) }

// Novalidate sets the novalidate boolean attribute.
func Novalidate() attrib.Node { return attrib.Bool("novalidate") }

// Accept sets the accept attribute.
func Accept(v string) attrib.Node { return attrib.KV("accept", v) }

// Acceptcharset sets the accept-charset attribute.
func Acceptcharset(v string) attrib.Node { return attrib.KV("accept-charset", v) }
