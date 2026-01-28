package input

import "github.com/protolambda/chord/core"

// Checked sets the checked boolean attribute.
func Checked() core.Node { return core.BoolAttribute("checked") }

// Pattern sets the pattern attribute.
func Pattern(v string) core.Node { return core.Attribute("pattern", v) }

// Min sets the min attribute.
func Min(v string) core.Node { return core.Attribute("min", v) }

// Max sets the max attribute.
func Max(v string) core.Node { return core.Attribute("max", v) }

// Minlength sets the minlength attribute.
func Minlength(v string) core.Node { return core.Attribute("minlength", v) }

// Maxlength sets the maxlength attribute.
func Maxlength(v string) core.Node { return core.Attribute("maxlength", v) }

// Step sets the step attribute.
func Step(v string) core.Node { return core.Attribute("step", v) }

// Multiple sets the multiple boolean attribute.
func Multiple() core.Node { return core.BoolAttribute("multiple") }

// List sets the list attribute.
func List(v string) core.Node { return core.Attribute("list", v) }

// Accept sets the accept attribute.
func Accept(v string) core.Node { return core.Attribute("accept", v) }

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
