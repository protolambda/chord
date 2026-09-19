package input

import (
	"github.com/protolambda/chord/core/attr"
)

// Checked sets the checked boolean attribute.
func Checked() attr.Node { return attr.Name("checked").Bool() }

// Pattern sets the pattern attribute.
func Pattern(v string) attr.Node { return attr.Name("pattern").Value(v) }

// Min sets the min attribute.
func Min(v string) attr.Node { return attr.Name("min").Value(v) }

// Max sets the max attribute.
func Max(v string) attr.Node { return attr.Name("max").Value(v) }

// Minlength sets the minlength attribute.
func Minlength(v string) attr.Node { return attr.Name("minlength").Value(v) }

// Maxlength sets the maxlength attribute.
func Maxlength(v string) attr.Node { return attr.Name("maxlength").Value(v) }

// Step sets the step attribute.
func Step(v string) attr.Node { return attr.Name("step").Value(v) }

// Multiple sets the multiple boolean attribute.
func Multiple() attr.Node { return attr.Name("multiple").Bool() }

// List sets the list attribute.
func List(v string) attr.Node { return attr.Name("list").Value(v) }

// Accept sets the accept attribute.
func Accept(v string) attr.Node { return attr.Name("accept").Value(v) }

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
