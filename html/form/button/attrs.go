package button

import (
	"github.com/protolambda/chord/core/attr"
)

// Popovertarget sets the popovertarget attribute.
func Popovertarget(v string) attr.Node { return attr.Name("popovertarget").Value(v) }

// Popovertargetaction sets the popovertargetaction attribute.
func Popovertargetaction(v string) attr.Node { return attr.Name("popovertargetaction").Value(v) }

// Formaction sets the formaction attribute.
func Formaction(v string) attr.Node { return attr.Name("formaction").Value(v) }

// Formenctype sets the formenctype attribute.
func Formenctype(v string) attr.Node { return attr.Name("formenctype").Value(v) }

// Formmethod sets the formmethod attribute.
func Formmethod(v string) attr.Node { return attr.Name("formmethod").Value(v) }

// Formnovalidate sets the formnovalidate boolean attribute.
func Formnovalidate() attr.Node { return attr.Name("formnovalidate").Bool() }

// Formtarget sets the formtarget attribute.
func Formtarget(v string) attr.Node { return attr.Name("formtarget").Value(v) }
