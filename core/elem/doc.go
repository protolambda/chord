// Package elem provides types and utils for node elements.
//
// The core abstractions are:
//   - [Node]: A lazy-evaluated element.
//   - [Obj]: A static representation of evaluated content.
//
// Element utils include:
//   - [New]: Produces an element, returns a Scope to complete
//     with child elements (e.g. <div>scope goes here</div>).
//   - [Void]: A self-closing element (e.g., <br/>, <img/>).
//   - [Raw]: Raw document content (text, HTML).
//   - [Scope]: A function to call to fill in child elements.
//   - [Bundle]: A slice of multiple elements.
//   - [Seq]: A sequence of multiple elements.
//   - [Cons]: A head element followed by a tail of elements.
//   - [Noop]: An empty element that produces no output.
//   - [If]: A conditional element.
//   - [IfElse]: A choice between elements.
//   - [Fn]: A dynamically produced element.
package elem
