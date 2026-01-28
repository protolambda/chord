// Package core provides the fundamental types and functions for building HTML documents.
//
// The core abstractions are:
//   - [Node]: A lazy-evaluated piece of the document graph
//   - [Obj]: A static representation of evaluated content
//
// Node types include:
//   - [Raw]: Raw document content (text, HTML)
//   - [Element]: An HTML element with optional children
//   - [VoidElement]: A self-closing element (e.g., <br/>, <img/>)
//   - [Attribute]: An attribute key-value pair
//   - [BoolAttribute]: A boolean attribute (no value)
//   - [Bundle]: A combination of multiple nodes
//   - [Noop]: An empty node that produces no output
//
// Nodes can be either attributes or elements, and bundles can mix both
// for easy composition of reusable components.
//
// Use [Render] to evaluate a node graph and write it to an [io.Writer],
// or [Dump] to write to stdout for debugging.
package core
