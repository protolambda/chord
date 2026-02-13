package attr

import (
	"github.com/protolambda/chord/core/attrib"
)

// Global Attributes

// Class sets the class attribute.
func Class(v string) attrib.Node { return attrib.KV("class", v) }

// ID sets the id attribute.
func ID(v string) attrib.Node { return attrib.KV("id", v) }

// Style sets the style attribute.
func Style(v string) attrib.Node { return attrib.KV("style", v) }

// Slot sets the slot attribute.
func Slot(v string) attrib.Node { return attrib.KV("slot", v) }

// Title sets the title attribute.
func Title(v string) attrib.Node { return attrib.KV("title", v) }

// Accesskey sets the accesskey attribute.
func Accesskey(v string) attrib.Node { return attrib.KV("accesskey", v) }

// Autocapitalize sets the autocapitalize attribute.
func Autocapitalize(v string) attrib.Node { return attrib.KV("autocapitalize", v) }

// Autofocus sets the autofocus boolean attribute.
func Autofocus() attrib.Node { return attrib.Bool("autofocus") }

// Contenteditable sets the contenteditable attribute.
func Contenteditable(v string) attrib.Node { return attrib.KV("contenteditable", v) }

// Dir sets the dir attribute.
func Dir(v string) attrib.Node { return attrib.KV("dir", v) }

// Draggable sets the draggable attribute.
func Draggable(v string) attrib.Node { return attrib.KV("draggable", v) }

// Enterkeyhint sets the enterkeyhint attribute.
func Enterkeyhint(v string) attrib.Node { return attrib.KV("enterkeyhint", v) }

// Hidden sets the hidden boolean attribute.
func Hidden() attrib.Node { return attrib.Bool("hidden") }

// Inert sets the inert boolean attribute.
func Inert() attrib.Node { return attrib.Bool("inert") }

// Inputmode sets the inputmode attribute.
func Inputmode(v string) attrib.Node { return attrib.KV("inputmode", v) }

// Is sets the is attribute for custom elements.
func Is(v string) attrib.Node { return attrib.KV("is", v) }

// Itemid sets the itemid attribute (microdata).
func Itemid(v string) attrib.Node { return attrib.KV("itemid", v) }

// Itemprop sets the itemprop attribute (microdata).
func Itemprop(v string) attrib.Node { return attrib.KV("itemprop", v) }

// Itemref sets the itemref attribute (microdata).
func Itemref(v string) attrib.Node { return attrib.KV("itemref", v) }

// Itemscope sets the itemscope boolean attribute (microdata).
func Itemscope() attrib.Node { return attrib.Bool("itemscope") }

// Itemtype sets the itemtype attribute (microdata).
func Itemtype(v string) attrib.Node { return attrib.KV("itemtype", v) }

// Lang sets the lang attribute.
func Lang(v string) attrib.Node { return attrib.KV("lang", v) }

// Nonce sets the nonce attribute.
func Nonce(v string) attrib.Node { return attrib.KV("nonce", v) }

// Popover sets the popover attribute.
func Popover(v string) attrib.Node { return attrib.KV("popover", v) }

// Spellcheck sets the spellcheck attribute.
func Spellcheck(v string) attrib.Node { return attrib.KV("spellcheck", v) }

// Tabindex sets the tabindex attribute.
func Tabindex(v string) attrib.Node { return attrib.KV("tabindex", v) }

// Translate sets the translate attribute.
func Translate(v string) attrib.Node { return attrib.KV("translate", v) }

// Data Attributes

// Data sets a data-* attribute.
func Data(name, value string) attrib.Node { return attrib.KV("data-"+name, value) }
