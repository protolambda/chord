// Package attr provides types and utils for node attributes.
//
// The core abstractions are:
//   - [Node]: A lazy-evaluated attribute.
//   - [Obj]: A static representation of evaluated content.
//   - [Name]: A trusted, statically known attribute name with escaped and raw value constructors.
//
// Attribute utils include:
//   - [ParseName]: Validates a runtime attribute name.
//   - [KV]: A validated runtime name and HTML-escaped value.
//   - [Bool]: A validated runtime name for a boolean attribute (no value).
//   - [Bundle]: A slice of multiple attributes.
//   - [Seq]: A sequence of multiple attributes.
//   - [Cons]: A head attribute followed by a tail of attributes.
//   - [Noop]: An empty attribute that produces no output.
//   - [If]: A conditional attribute.
//   - [IfElse]: A choice between attributes.
//   - [Fn]: A dynamically produced attribute.
//
// Values from callers should use [KV], [Name.Value], or the attribute helpers in
// this package. Chord internals may use [Name.Raw] for values already known to be
// safe; Raw does not escape its value.
//
// This package also provides global HTML attributes that can be applied to any element.
//
// Common attributes:
//   - [Class]: CSS class names
//   - [ID]: Unique element identifier
//   - [Style]: Inline CSS styles
//   - [Title]: Advisory title
//   - [Data]: Custom data attributes (data-*)
//   - [Lang]: Language of the element
//   - [Dir]: Text directionality
//   - [Hidden]: Hide the element
//   - [Tabindex]: Tab order
//   - [Contenteditable]: Make content editable
//   - [Draggable]: Enable drag and drop
//   - [Spellcheck]: Enable spell checking
//
// Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes
package attr
