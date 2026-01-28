package on

import "github.com/protolambda/chord/core"

// PointerDown sets the onpointerdown event handler.
func PointerDown(v string) core.Node { return core.Attribute("onpointerdown", v) }

// PointerUp sets the onpointerup event handler.
func PointerUp(v string) core.Node { return core.Attribute("onpointerup", v) }

// PointerMove sets the onpointermove event handler.
func PointerMove(v string) core.Node { return core.Attribute("onpointermove", v) }

// PointerEnter sets the onpointerenter event handler.
func PointerEnter(v string) core.Node { return core.Attribute("onpointerenter", v) }

// PointerLeave sets the onpointerleave event handler.
func PointerLeave(v string) core.Node { return core.Attribute("onpointerleave", v) }

// PointerOver sets the onpointerover event handler.
func PointerOver(v string) core.Node { return core.Attribute("onpointerover", v) }

// PointerOut sets the onpointerout event handler.
func PointerOut(v string) core.Node { return core.Attribute("onpointerout", v) }

// PointerCancel sets the onpointercancel event handler.
func PointerCancel(v string) core.Node { return core.Attribute("onpointercancel", v) }

// GotPointerCapture sets the ongotpointercapture event handler.
func GotPointerCapture(v string) core.Node { return core.Attribute("ongotpointercapture", v) }

// LostPointerCapture sets the onlostpointercapture event handler.
func LostPointerCapture(v string) core.Node { return core.Attribute("onlostpointercapture", v) }
