// Package text provides text-level semantic elements and the Text helper.
//
// The [Text] function creates HTML-escaped text content.
//
// Inline elements:
//   - [Span]: Generic inline container
//   - [A]: Hyperlinks (with [Href], [Target], [Rel] attributes)
//   - [Em], [Strong]: Emphasis and importance
//   - [Small]: Side comments
//   - [S]: Strikethrough
//   - [Cite]: Citation
//   - [Q]: Inline quotation
//   - [Code], [Kbd], [Samp], [Var]: Code and technical text
//   - [Sub], [Sup]: Subscript and superscript
//   - [Mark]: Highlighted text
//   - [Time]: Date/time (with [Datetime] attribute)
//   - [Abbr]: Abbreviation
//   - [Dfn]: Definition term
//   - [B], [I], [U]: Bold, italic, underline (stylistic)
//   - [Bdi], [Bdo]: Bidirectional text
//   - [Ruby], [Rt], [Rp]: Ruby annotations
//   - [Data]: Machine-readable data
//   - [Wbr]: Word break opportunity
//
// Block elements in this package:
//   - [P]: Paragraph
//   - [Hr]: Horizontal rule
//   - [Pre]: Preformatted text
//   - [Blockquote]: Block quotation
//   - [Br]: Line break
//
// Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element#inline_text_semantics
package text
