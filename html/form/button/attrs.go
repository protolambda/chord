package button

import (
	"github.com/protolambda/chord/core/attrib"
)

// Popovertarget sets the popovertarget attribute.
func Popovertarget(v string) attrib.Node { return attrib.KV("popovertarget", v) }

// Popovertargetaction sets the popovertargetaction attribute.
func Popovertargetaction(v string) attrib.Node { return attrib.KV("popovertargetaction", v) }

// Formaction sets the formaction attribute.
func Formaction(v string) attrib.Node { return attrib.KV("formaction", v) }

// Formenctype sets the formenctype attribute.
func Formenctype(v string) attrib.Node { return attrib.KV("formenctype", v) }

// Formmethod sets the formmethod attribute.
func Formmethod(v string) attrib.Node { return attrib.KV("formmethod", v) }

// Formnovalidate sets the formnovalidate boolean attribute.
func Formnovalidate() attrib.Node { return attrib.Bool("formnovalidate") }

// Formtarget sets the formtarget attribute.
func Formtarget(v string) attrib.Node { return attrib.KV("formtarget", v) }
