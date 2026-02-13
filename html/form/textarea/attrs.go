package textarea

import (
	"github.com/protolambda/chord/core/attrib"
)

// Rows sets the rows attribute.
func Rows(v string) attrib.Node { return attrib.KV("rows", v) }

// Cols sets the cols attribute.
func Cols(v string) attrib.Node { return attrib.KV("cols", v) }

// Wrap sets the wrap attribute.
func Wrap(v string) attrib.Node { return attrib.KV("wrap", v) }

// Minlength sets the minlength attribute.
func Minlength(v string) attrib.Node { return attrib.KV("minlength", v) }

// Maxlength sets the maxlength attribute.
func Maxlength(v string) attrib.Node { return attrib.KV("maxlength", v) }
