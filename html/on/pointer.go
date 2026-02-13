package on

import (
	"github.com/protolambda/chord/core/attrib"
)

// PointerDown sets the onpointerdown event handler.
func PointerDown(v string) attrib.Node { return attrib.KV("onpointerdown", v) }

// PointerUp sets the onpointerup event handler.
func PointerUp(v string) attrib.Node { return attrib.KV("onpointerup", v) }

// PointerMove sets the onpointermove event handler.
func PointerMove(v string) attrib.Node { return attrib.KV("onpointermove", v) }

// PointerEnter sets the onpointerenter event handler.
func PointerEnter(v string) attrib.Node { return attrib.KV("onpointerenter", v) }

// PointerLeave sets the onpointerleave event handler.
func PointerLeave(v string) attrib.Node { return attrib.KV("onpointerleave", v) }

// PointerOver sets the onpointerover event handler.
func PointerOver(v string) attrib.Node { return attrib.KV("onpointerover", v) }

// PointerOut sets the onpointerout event handler.
func PointerOut(v string) attrib.Node { return attrib.KV("onpointerout", v) }

// PointerCancel sets the onpointercancel event handler.
func PointerCancel(v string) attrib.Node { return attrib.KV("onpointercancel", v) }

// GotPointerCapture sets the ongotpointercapture event handler.
func GotPointerCapture(v string) attrib.Node { return attrib.KV("ongotpointercapture", v) }

// LostPointerCapture sets the onlostpointercapture event handler.
func LostPointerCapture(v string) attrib.Node { return attrib.KV("onlostpointercapture", v) }
