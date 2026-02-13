package input

import (
	"github.com/protolambda/chord/core/attrib"
)

// Checked sets the checked boolean attribute.
func Checked() attrib.Node { return attrib.Bool("checked") }

// Pattern sets the pattern attribute.
func Pattern(v string) attrib.Node { return attrib.KV("pattern", v) }

// Min sets the min attribute.
func Min(v string) attrib.Node { return attrib.KV("min", v) }

// Max sets the max attribute.
func Max(v string) attrib.Node { return attrib.KV("max", v) }

// Minlength sets the minlength attribute.
func Minlength(v string) attrib.Node { return attrib.KV("minlength", v) }

// Maxlength sets the maxlength attribute.
func Maxlength(v string) attrib.Node { return attrib.KV("maxlength", v) }

// Step sets the step attribute.
func Step(v string) attrib.Node { return attrib.KV("step", v) }

// Multiple sets the multiple boolean attribute.
func Multiple() attrib.Node { return attrib.Bool("multiple") }

// List sets the list attribute.
func List(v string) attrib.Node { return attrib.KV("list", v) }

// Accept sets the accept attribute.
func Accept(v string) attrib.Node { return attrib.KV("accept", v) }

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
