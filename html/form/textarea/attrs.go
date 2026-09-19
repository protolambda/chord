package textarea

import (
	"github.com/protolambda/chord/core/attr"
)

// Rows sets the rows attribute.
func Rows(v string) attr.Node { return attr.Name("rows").Value(v) }

// Cols sets the cols attribute.
func Cols(v string) attr.Node { return attr.Name("cols").Value(v) }

// Wrap sets the wrap attribute.
func Wrap(v string) attr.Node { return attr.Name("wrap").Value(v) }

// Minlength sets the minlength attribute.
func Minlength(v string) attr.Node { return attr.Name("minlength").Value(v) }

// Maxlength sets the maxlength attribute.
func Maxlength(v string) attr.Node { return attr.Name("maxlength").Value(v) }
