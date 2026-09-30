// Package elem provides types and utils for node elements.
//
// The core abstractions are:
//   - [Node]: A lazy-evaluated element.
//   - [Obj]: A static representation of evaluated content, tagged with a [Kind].
//   - [Name]: A trusted, statically known element tag name.
//
// Element utils include:
//   - [ParseName]: Validates a runtime element tag name.
//   - [Name.New]: Produces an element, returning a Scope to complete
//     with child elements (e.g. <div>scope goes here</div>).
//   - [Name.Void]: A self-closing element (e.g., <br/>, <img/>).
//   - [Text]: Logical text content, HTML-escaped when rendered.
//   - [Raw]: Raw document content (trusted HTML, written verbatim).
//   - [Comment]: An HTML comment.
//   - [Doctype]: The HTML document type declaration, <!DOCTYPE html>.
//   - [Scope]: A function to call to fill in child elements.
//   - [Bundle]: A slice of multiple elements.
//   - [Seq]: A sequence of multiple elements.
//   - [Cons]: A head element followed by a tail of elements.
//   - [Noop]: An empty element that produces no output.
//   - [If]: A conditional element.
//   - [IfElse]: A choice between elements.
//   - [Fn]: A dynamically produced element.
//
// A nil [Node] is not "no content": rendering it fails with [ErrNilNode],
// located like any other evaluation error, and so does an [Fn] that returns a
// nil node without an error. Use [Noop] or [If] for optional content.
package elem
