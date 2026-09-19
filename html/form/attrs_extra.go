package form

import (
	"github.com/protolambda/chord/core/attr"
)

// FormMethod represents valid form method values.
type FormMethod string

const (
	MethodGet  FormMethod = "get"
	MethodPost FormMethod = "post"
)

// Method sets the method attribute with a typed FormMethod.
func Method(m FormMethod) attr.Node {
	switch m {
	case MethodGet, MethodPost:
		return attr.Name("method").Raw(string(m))
	default:
		return attr.Name("method").Value(string(m))
	}
}

// FormEnctype represents valid form enctype values.
type FormEnctype string

const (
	EnctypeURLEncoded FormEnctype = "application/x-www-form-urlencoded"
	EnctypeMultipart  FormEnctype = "multipart/form-data"
	EnctypeText       FormEnctype = "text/plain"
)

// Enctype sets the enctype attribute with a typed FormEnctype.
func Enctype(e FormEnctype) attr.Node {
	switch e {
	case EnctypeURLEncoded, EnctypeMultipart, EnctypeText:
		return attr.Name("enctype").Raw(string(e))
	default:
		return attr.Name("enctype").Value(string(e))
	}
}
