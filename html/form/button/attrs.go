package button

import "github.com/protolambda/chord/core"

// Popovertarget sets the popovertarget attribute.
func Popovertarget(v string) core.Node { return core.Attribute("popovertarget", v) }

// Popovertargetaction sets the popovertargetaction attribute.
func Popovertargetaction(v string) core.Node { return core.Attribute("popovertargetaction", v) }

// Formaction sets the formaction attribute.
func Formaction(v string) core.Node { return core.Attribute("formaction", v) }

// Formenctype sets the formenctype attribute.
func Formenctype(v string) core.Node { return core.Attribute("formenctype", v) }

// Formmethod sets the formmethod attribute.
func Formmethod(v string) core.Node { return core.Attribute("formmethod", v) }

// Formnovalidate sets the formnovalidate boolean attribute.
func Formnovalidate() core.Node { return core.BoolAttribute("formnovalidate") }

// Formtarget sets the formtarget attribute.
func Formtarget(v string) core.Node { return core.Attribute("formtarget", v) }
