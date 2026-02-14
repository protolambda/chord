package button

import (
	"github.com/protolambda/chord/core/attr"
)

// ButtonType represents valid button type values.
type ButtonType string

const (
	TypeSubmit ButtonType = "submit"
	TypeReset  ButtonType = "reset"
	TypeButton ButtonType = "button"
)

// Type sets the type attribute with a typed ButtonType.
func Type(t ButtonType) attr.Node { return attr.KV("type", string(t)) }
