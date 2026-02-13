package hx1

import (
	"github.com/protolambda/chord/core/attrib"
)

// PushURL sets the hx-push-url attribute to push a URL to browser history.
func PushURL(v string) attrib.Node { return attrib.KV("hx-push-url", v) }

// ReplaceURL sets the hx-replace-url attribute to replace browser history URL.
func ReplaceURL(v string) attrib.Node { return attrib.KV("hx-replace-url", v) }

// HistoryElt sets the hx-history-elt boolean attribute.
func HistoryElt() attrib.Node { return attrib.Bool("hx-history-elt") }
