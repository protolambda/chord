package input

import (
	"github.com/protolambda/chord/core/attr"
)

// Checked sets the checked boolean attribute.
func Checked() attr.Node { return attr.Bool("checked") }

// Pattern sets the pattern attribute.
func Pattern(v string) attr.Node { return attr.KV("pattern", v) }

// Min sets the min attribute.
func Min(v string) attr.Node { return attr.KV("min", v) }

// Max sets the max attribute.
func Max(v string) attr.Node { return attr.KV("max", v) }

// Minlength sets the minlength attribute.
func Minlength(v string) attr.Node { return attr.KV("minlength", v) }

// Maxlength sets the maxlength attribute.
func Maxlength(v string) attr.Node { return attr.KV("maxlength", v) }

// Step sets the step attribute.
func Step(v string) attr.Node { return attr.KV("step", v) }

// Multiple sets the multiple boolean attribute.
func Multiple() attr.Node { return attr.Bool("multiple") }

// List sets the list attribute.
func List(v string) attr.Node { return attr.KV("list", v) }

// Accept sets the accept attribute.
func Accept(v string) attr.Node { return attr.KV("accept", v) }

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
