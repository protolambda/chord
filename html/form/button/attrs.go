package button

import (
	"github.com/protolambda/chord/core/attr"
)

// Popovertarget sets the popovertarget attribute.
func Popovertarget(v string) attr.Node { return attr.KV("popovertarget", v) }

// Popovertargetaction sets the popovertargetaction attribute.
func Popovertargetaction(v string) attr.Node { return attr.KV("popovertargetaction", v) }

// Formaction sets the formaction attribute.
func Formaction(v string) attr.Node { return attr.KV("formaction", v) }

// Formenctype sets the formenctype attribute.
func Formenctype(v string) attr.Node { return attr.KV("formenctype", v) }

// Formmethod sets the formmethod attribute.
func Formmethod(v string) attr.Node { return attr.KV("formmethod", v) }

// Formnovalidate sets the formnovalidate boolean attribute.
func Formnovalidate() attr.Node { return attr.Bool("formnovalidate") }

// Formtarget sets the formtarget attribute.
func Formtarget(v string) attr.Node { return attr.KV("formtarget", v) }
