package hx2

import (
	"github.com/protolambda/chord/core/attr"
)

// PushURL sets the hx-push-url attribute to push a URL to browser history.
func PushURL(v string) attr.Node { return attr.KV("hx-push-url", v) }

// ReplaceURL sets the hx-replace-url attribute to replace browser history URL.
func ReplaceURL(v string) attr.Node { return attr.KV("hx-replace-url", v) }

// HistoryElt sets the hx-history-elt boolean attribute.
func HistoryElt() attr.Node { return attr.Bool("hx-history-elt") }
