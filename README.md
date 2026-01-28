# chord

*note: implementation is a work in progress*:
- [x] core functionality
- [ ] common elements and attributes
- [ ] full event attributes coverage
- [ ] extension packages for boostrap, htmx, etc.

This is a Go DSL for generating HTML structures.

The rendering is evaluated, so structures can be re-used,
and adapt their output to context changes.

General API:
- `core.Obj`: static representation of a piece of the document graph
- `core.Node`: yet-to-be-evaluated piece of the document graph, this can be:
  - `core.Raw`: raw document content (text, raw html, etc.)
  - `core.Element`: an element, with optional sub elements (lazy eval)
  - `core.VoidElement`: an element that self-closes, e.g. `<br/>`
  - `core.Attribute`: an attribute, to be applied to the scope
  - `core.BoolAttribute`: attribute without value
  - `core.Bundle`: Combination of nodes
  - `core.Noop`: No-op node (empty output)
- Important: `Node` is either attribute or element: bundles can mix these for easy combinations.
- `core.Render`: eval a node graph and write it to output
- `<package>`: scoped packages for attributes specific to an HTML tag or other group scope
- core: all common element types are available in the main package, for easy dot-import.
- utils:
  - `Text`: html-escaped string content
  - `Fn`: dynamic content
  - `If`: conditional sub-structures
  - `Fallback`: scoped render, catch error, fallback

## Example

See [Example test](./chord_example_test.go).

## Design

The design is focused on composition to handle repetition, hence the "chord" name.
Stringing together elements and attributes is important to be easy.

The eval approach takes some inspiration from react and other similar HTML DSLs.

This library is designed to:
- provide a clean scoped API (no bloat)
- render with reasonable performance
- prevent code-gen for end-users
- offer native Go typing
- be extensible (custom components!)
- be composable (mix component attribute bundles)
- not require dependencies

## License

MIT, see [`LICENSE`](./LICENSE) file.
