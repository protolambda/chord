// Package attrib provides types and utils for node attributes.
//
// The core abstractions are:
//   - [Node]: A lazy-evaluated attribute.
//   - [Obj]: A static representation of evaluated content.
//
// Attribute utils include:
//   - [KV]: An attribute key-value pair.
//   - [Bool]: A boolean attribute (no value).
//   - [Bundle]: A slice of multiple attributes.
//   - [Seq]: A sequence of multiple attributes.
//   - [Cons]: A head attribute followed by a tail of attributes.
//   - [Noop]: An empty attribute that produces no output.
//   - [If]: A conditional attribute.
//   - [IfElse]: A choice between attributes.
//   - [Fn]: A dynamically produced attribute.
package attrib
