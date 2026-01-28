// Package script provides scripting and template elements.
//
// Elements:
//   - [Script]: JavaScript code or reference
//   - [Noscript]: Fallback for no JavaScript
//   - [Template]: HTML template (not rendered)
//   - [Canvas]: Graphics canvas
//
// Script attributes:
//   - [Src]: External script URL
//   - [Type]: Script type (module, text/javascript)
//   - [Async]: Load asynchronously
//   - [Defer]: Defer execution until DOM ready
//   - [Crossorigin]: CORS settings
//   - [Integrity]: Subresource integrity hash
//   - [Nomodule]: Skip for module-supporting browsers
//
// Canvas attributes:
//   - [Width], [Height]: Canvas dimensions
//
// Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element#scripting
package script
