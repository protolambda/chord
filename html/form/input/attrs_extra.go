package input

import (
	"github.com/protolambda/chord/core/attr"
)

// InputType represents valid input type values.
type InputType string

const (
	Text          InputType = "text"
	Password      InputType = "password"
	Email         InputType = "email"
	Number        InputType = "number"
	Tel           InputType = "tel"
	URL           InputType = "url"
	Search        InputType = "search"
	Date          InputType = "date"
	Time          InputType = "time"
	DatetimeLocal InputType = "datetime-local"
	Month         InputType = "month"
	Week          InputType = "week"
	Color         InputType = "color"
	File          InputType = "file"
	Checkbox      InputType = "checkbox"
	Radio         InputType = "radio"
	Range         InputType = "range"
	Hidden        InputType = "hidden"
	Submit        InputType = "submit"
	Reset         InputType = "reset"
	Button        InputType = "button"
	Image         InputType = "image"
)

// Type sets the type attribute with a typed InputType.
func Type(t InputType) attr.Node { return attr.KV("type", string(t)) }
