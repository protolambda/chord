// Package core provides core rendering functionality.
//
// Use [Render] to evaluate a node graph once and write it to an
// [io.StringWriter], or [Dump] to write to stdout for debugging. Text and
// logical attribute values are HTML-escaped by the renderer; raw content is
// written verbatim. Render errors wrap their cause and carry the location of
// the failing node, such as "at html[0]/body[1]/form#login[0]/[2]".
//
// [Fallback] evaluates a subtree transactionally and substitutes alternative
// content when evaluation fails. The [github.com/protolambda/chord/core/inspect]
// package evaluates a node graph into a read-only snapshot through the same
// normalized walk that Render uses.
package core
