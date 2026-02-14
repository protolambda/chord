package on

import (
	"github.com/protolambda/chord/core/attr"
)

// PointerDown sets the onpointerdown event handler.
func PointerDown(v string) attr.Node { return attr.KV("onpointerdown", v) }

// PointerUp sets the onpointerup event handler.
func PointerUp(v string) attr.Node { return attr.KV("onpointerup", v) }

// PointerMove sets the onpointermove event handler.
func PointerMove(v string) attr.Node { return attr.KV("onpointermove", v) }

// PointerEnter sets the onpointerenter event handler.
func PointerEnter(v string) attr.Node { return attr.KV("onpointerenter", v) }

// PointerLeave sets the onpointerleave event handler.
func PointerLeave(v string) attr.Node { return attr.KV("onpointerleave", v) }

// PointerOver sets the onpointerover event handler.
func PointerOver(v string) attr.Node { return attr.KV("onpointerover", v) }

// PointerOut sets the onpointerout event handler.
func PointerOut(v string) attr.Node { return attr.KV("onpointerout", v) }

// PointerCancel sets the onpointercancel event handler.
func PointerCancel(v string) attr.Node { return attr.KV("onpointercancel", v) }

// GotPointerCapture sets the ongotpointercapture event handler.
func GotPointerCapture(v string) attr.Node { return attr.KV("ongotpointercapture", v) }

// LostPointerCapture sets the onlostpointercapture event handler.
func LostPointerCapture(v string) attr.Node { return attr.KV("onlostpointercapture", v) }
