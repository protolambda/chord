package attr

// Global Attributes

// Class sets the class attribute.
func Class(v string) Node { return Name("class").Value(v) }

// ID sets the id attribute.
func ID(v string) Node { return Name("id").Value(v) }

// Style sets the style attribute.
func Style(v string) Node { return Name("style").Value(v) }

// Slot sets the slot attribute.
func Slot(v string) Node { return Name("slot").Value(v) }

// Title sets the title attribute.
func Title(v string) Node { return Name("title").Value(v) }

// Accesskey sets the accesskey attribute.
func Accesskey(v string) Node { return Name("accesskey").Value(v) }

// Autocapitalize sets the autocapitalize attribute.
func Autocapitalize(v string) Node { return Name("autocapitalize").Value(v) }

// Autofocus sets the autofocus boolean attribute.
func Autofocus() Node { return Name("autofocus").Bool() }

// Contenteditable sets the contenteditable attribute.
func Contenteditable(v string) Node { return Name("contenteditable").Value(v) }

// Dir sets the dir attribute.
func Dir(v string) Node { return Name("dir").Value(v) }

// Draggable sets the draggable attribute.
func Draggable(v string) Node { return Name("draggable").Value(v) }

// Enterkeyhint sets the enterkeyhint attribute.
func Enterkeyhint(v string) Node { return Name("enterkeyhint").Value(v) }

// Hidden sets the hidden boolean attribute.
func Hidden() Node { return Name("hidden").Bool() }

// Inert sets the inert boolean attribute.
func Inert() Node { return Name("inert").Bool() }

// Inputmode sets the inputmode attribute.
func Inputmode(v string) Node { return Name("inputmode").Value(v) }

// Is sets the is attribute for custom elements.
func Is(v string) Node { return Name("is").Value(v) }

// Itemid sets the itemid attribute (microdata).
func Itemid(v string) Node { return Name("itemid").Value(v) }

// Itemprop sets the itemprop attribute (microdata).
func Itemprop(v string) Node { return Name("itemprop").Value(v) }

// Itemref sets the itemref attribute (microdata).
func Itemref(v string) Node { return Name("itemref").Value(v) }

// Itemscope sets the itemscope boolean attribute (microdata).
func Itemscope() Node { return Name("itemscope").Bool() }

// Itemtype sets the itemtype attribute (microdata).
func Itemtype(v string) Node { return Name("itemtype").Value(v) }

// Lang sets the lang attribute.
func Lang(v string) Node { return Name("lang").Value(v) }

// Nonce sets the nonce attribute.
func Nonce(v string) Node { return Name("nonce").Value(v) }

// Popover sets the popover attribute.
func Popover(v string) Node { return Name("popover").Value(v) }

// Spellcheck sets the spellcheck attribute.
func Spellcheck(v string) Node { return Name("spellcheck").Value(v) }

// Tabindex sets the tabindex attribute.
func Tabindex(v string) Node { return Name("tabindex").Value(v) }

// Translate sets the translate attribute.
func Translate(v string) Node { return Name("translate").Value(v) }

// Data Attributes

// Data sets a data-* attribute.
func Data(name, value string) Node { return KV("data-"+name, value) }
