package attr

// Global Attributes

// Class sets the class attribute.
func Class(v string) Node { return KV("class", v) }

// ID sets the id attribute.
func ID(v string) Node { return KV("id", v) }

// Style sets the style attribute.
func Style(v string) Node { return KV("style", v) }

// Slot sets the slot attribute.
func Slot(v string) Node { return KV("slot", v) }

// Title sets the title attribute.
func Title(v string) Node { return KV("title", v) }

// Accesskey sets the accesskey attribute.
func Accesskey(v string) Node { return KV("accesskey", v) }

// Autocapitalize sets the autocapitalize attribute.
func Autocapitalize(v string) Node { return KV("autocapitalize", v) }

// Autofocus sets the autofocus boolean attribute.
func Autofocus() Node { return Bool("autofocus") }

// Contenteditable sets the contenteditable attribute.
func Contenteditable(v string) Node { return KV("contenteditable", v) }

// Dir sets the dir attribute.
func Dir(v string) Node { return KV("dir", v) }

// Draggable sets the draggable attribute.
func Draggable(v string) Node { return KV("draggable", v) }

// Enterkeyhint sets the enterkeyhint attribute.
func Enterkeyhint(v string) Node { return KV("enterkeyhint", v) }

// Hidden sets the hidden boolean attribute.
func Hidden() Node { return Bool("hidden") }

// Inert sets the inert boolean attribute.
func Inert() Node { return Bool("inert") }

// Inputmode sets the inputmode attribute.
func Inputmode(v string) Node { return KV("inputmode", v) }

// Is sets the is attribute for custom elements.
func Is(v string) Node { return KV("is", v) }

// Itemid sets the itemid attribute (microdata).
func Itemid(v string) Node { return KV("itemid", v) }

// Itemprop sets the itemprop attribute (microdata).
func Itemprop(v string) Node { return KV("itemprop", v) }

// Itemref sets the itemref attribute (microdata).
func Itemref(v string) Node { return KV("itemref", v) }

// Itemscope sets the itemscope boolean attribute (microdata).
func Itemscope() Node { return Bool("itemscope") }

// Itemtype sets the itemtype attribute (microdata).
func Itemtype(v string) Node { return KV("itemtype", v) }

// Lang sets the lang attribute.
func Lang(v string) Node { return KV("lang", v) }

// Nonce sets the nonce attribute.
func Nonce(v string) Node { return KV("nonce", v) }

// Popover sets the popover attribute.
func Popover(v string) Node { return KV("popover", v) }

// Spellcheck sets the spellcheck attribute.
func Spellcheck(v string) Node { return KV("spellcheck", v) }

// Tabindex sets the tabindex attribute.
func Tabindex(v string) Node { return KV("tabindex", v) }

// Translate sets the translate attribute.
func Translate(v string) Node { return KV("translate", v) }

// Data Attributes

// Data sets a data-* attribute.
func Data(name, value string) Node { return KV("data-"+name, value) }
