package attr

import (
	"github.com/protolambda/chord/core"
)

// Global Attributes

// Class sets the class attribute.
func Class(v string) core.Node { return core.Attribute("class", v) }

// ID sets the id attribute.
func ID(v string) core.Node { return core.Attribute("id", v) }

// Style sets the style attribute.
func Style(v string) core.Node { return core.Attribute("style", v) }

// Slot sets the slot attribute.
func Slot(v string) core.Node { return core.Attribute("slot", v) }

// Title sets the title attribute.
func Title(v string) core.Node { return core.Attribute("title", v) }

// Accesskey sets the accesskey attribute.
func Accesskey(v string) core.Node { return core.Attribute("accesskey", v) }

// Autocapitalize sets the autocapitalize attribute.
func Autocapitalize(v string) core.Node { return core.Attribute("autocapitalize", v) }

// Autofocus sets the autofocus boolean attribute.
func Autofocus() core.Node { return core.BoolAttribute("autofocus") }

// Contenteditable sets the contenteditable attribute.
func Contenteditable(v string) core.Node { return core.Attribute("contenteditable", v) }

// Dir sets the dir attribute.
func Dir(v string) core.Node { return core.Attribute("dir", v) }

// Draggable sets the draggable attribute.
func Draggable(v string) core.Node { return core.Attribute("draggable", v) }

// Enterkeyhint sets the enterkeyhint attribute.
func Enterkeyhint(v string) core.Node { return core.Attribute("enterkeyhint", v) }

// Hidden sets the hidden boolean attribute.
func Hidden() core.Node { return core.BoolAttribute("hidden") }

// Inert sets the inert boolean attribute.
func Inert() core.Node { return core.BoolAttribute("inert") }

// Inputmode sets the inputmode attribute.
func Inputmode(v string) core.Node { return core.Attribute("inputmode", v) }

// Is sets the is attribute for custom elements.
func Is(v string) core.Node { return core.Attribute("is", v) }

// Itemid sets the itemid attribute (microdata).
func Itemid(v string) core.Node { return core.Attribute("itemid", v) }

// Itemprop sets the itemprop attribute (microdata).
func Itemprop(v string) core.Node { return core.Attribute("itemprop", v) }

// Itemref sets the itemref attribute (microdata).
func Itemref(v string) core.Node { return core.Attribute("itemref", v) }

// Itemscope sets the itemscope boolean attribute (microdata).
func Itemscope() core.Node { return core.BoolAttribute("itemscope") }

// Itemtype sets the itemtype attribute (microdata).
func Itemtype(v string) core.Node { return core.Attribute("itemtype", v) }

// Lang sets the lang attribute.
func Lang(v string) core.Node { return core.Attribute("lang", v) }

// Nonce sets the nonce attribute.
func Nonce(v string) core.Node { return core.Attribute("nonce", v) }

// Popover sets the popover attribute.
func Popover(v string) core.Node { return core.Attribute("popover", v) }

// Spellcheck sets the spellcheck attribute.
func Spellcheck(v string) core.Node { return core.Attribute("spellcheck", v) }

// Tabindex sets the tabindex attribute.
func Tabindex(v string) core.Node { return core.Attribute("tabindex", v) }

// Translate sets the translate attribute.
func Translate(v string) core.Node { return core.Attribute("translate", v) }

// Data Attributes

// Data sets a data-* attribute.
func Data(name, value string) core.Node { return core.Attribute("data-"+name, value) }
