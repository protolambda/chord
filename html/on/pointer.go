package on

import (
	"github.com/protolambda/chord/core/attr"
)

// PointerDown sets the onpointerdown event handler.
func PointerDown(v string) attr.Node { return attr.Name("onpointerdown").Value(v) }

// PointerUp sets the onpointerup event handler.
func PointerUp(v string) attr.Node { return attr.Name("onpointerup").Value(v) }

// PointerMove sets the onpointermove event handler.
func PointerMove(v string) attr.Node { return attr.Name("onpointermove").Value(v) }

// PointerEnter sets the onpointerenter event handler.
func PointerEnter(v string) attr.Node { return attr.Name("onpointerenter").Value(v) }

// PointerLeave sets the onpointerleave event handler.
func PointerLeave(v string) attr.Node { return attr.Name("onpointerleave").Value(v) }

// PointerOver sets the onpointerover event handler.
func PointerOver(v string) attr.Node { return attr.Name("onpointerover").Value(v) }

// PointerOut sets the onpointerout event handler.
func PointerOut(v string) attr.Node { return attr.Name("onpointerout").Value(v) }

// PointerCancel sets the onpointercancel event handler.
func PointerCancel(v string) attr.Node { return attr.Name("onpointercancel").Value(v) }

// GotPointerCapture sets the ongotpointercapture event handler.
func GotPointerCapture(v string) attr.Node { return attr.Name("ongotpointercapture").Value(v) }

// LostPointerCapture sets the onlostpointercapture event handler.
func LostPointerCapture(v string) attr.Node { return attr.Name("onlostpointercapture").Value(v) }
