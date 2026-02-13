package webcomp_test

import (
	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/core/attrib"
	"github.com/protolambda/chord/html/text"
	"github.com/protolambda/chord/html/webcomp"
)

func ExampleSlot() {
	core.Dump(webcomp.Slot(attrib.KV("name", "header")))
	// Output: <slot name="header"></slot>
}

func ExampleSlot_default() {
	core.Dump(webcomp.Slot()(text.Text("Default content")))
	// Output: <slot>Default content</slot>
}
