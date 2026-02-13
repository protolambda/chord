package form

import (
	"github.com/protolambda/chord/core/attrib"
)

// FormMethod represents valid form method values.
type FormMethod string

const (
	MethodGet  FormMethod = "get"
	MethodPost FormMethod = "post"
)

// Method sets the method attribute with a typed FormMethod.
func Method(m FormMethod) attrib.Node { return attrib.KV("method", string(m)) }

// FormEnctype represents valid form enctype values.
type FormEnctype string

const (
	EnctypeURLEncoded FormEnctype = "application/x-www-form-urlencoded"
	EnctypeMultipart  FormEnctype = "multipart/form-data"
	EnctypeText       FormEnctype = "text/plain"
)

// Enctype sets the enctype attribute with a typed FormEnctype.
func Enctype(e FormEnctype) attrib.Node { return attrib.KV("enctype", string(e)) }
