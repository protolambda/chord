package textarea

import (
	"github.com/protolambda/chord/core/attr"
)

// Rows sets the rows attribute.
func Rows(v string) attr.Node { return attr.KV("rows", v) }

// Cols sets the cols attribute.
func Cols(v string) attr.Node { return attr.KV("cols", v) }

// Wrap sets the wrap attribute.
func Wrap(v string) attr.Node { return attr.KV("wrap", v) }

// Minlength sets the minlength attribute.
func Minlength(v string) attr.Node { return attr.KV("minlength", v) }

// Maxlength sets the maxlength attribute.
func Maxlength(v string) attr.Node { return attr.KV("maxlength", v) }
