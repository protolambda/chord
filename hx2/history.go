package hx2

import "github.com/protolambda/chord/core"

// PushURL sets the hx-push-url attribute to push a URL to browser history.
func PushURL(v string) core.Node { return core.Attribute("hx-push-url", v) }

// ReplaceURL sets the hx-replace-url attribute to replace browser history URL.
func ReplaceURL(v string) core.Node { return core.Attribute("hx-replace-url", v) }

// HistoryElt sets the hx-history-elt boolean attribute.
func HistoryElt() core.Node { return core.BoolAttribute("hx-history-elt") }
