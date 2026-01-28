package textarea

import "github.com/protolambda/chord/core"

// Rows sets the rows attribute.
func Rows(v string) core.Node { return core.Attribute("rows", v) }

// Cols sets the cols attribute.
func Cols(v string) core.Node { return core.Attribute("cols", v) }

// Wrap sets the wrap attribute.
func Wrap(v string) core.Node { return core.Attribute("wrap", v) }

// Minlength sets the minlength attribute.
func Minlength(v string) core.Node { return core.Attribute("minlength", v) }

// Maxlength sets the maxlength attribute.
func Maxlength(v string) core.Node { return core.Attribute("maxlength", v) }
