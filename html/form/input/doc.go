// Package input provides the input element and its attributes.
//
// The [Input] element creates interactive form controls.
//
// Type attribute values (use [Type] or typed constructors):
//   - text, password, email, url, tel, search
//   - number, range
//   - date, time, datetime-local, month, week
//   - checkbox, radio
//   - file, color
//   - submit, reset, button
//   - hidden
//
// Common attributes:
//   - [Name], [Value]: Form submission
//   - [Placeholder]: Hint text
//   - [Required], [Disabled], [Readonly]: States
//   - [Min], [Max], [Step]: Numeric constraints
//   - [Minlength], [Maxlength]: Text length constraints
//   - [Pattern]: Regex validation
//   - [Autocomplete]: Browser autofill hints
//   - [Checked]: Initial checked state (checkbox/radio)
//   - [Multiple]: Allow multiple values (file/email)
//   - [Accept]: Accepted file types
//
// Reference: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input
package input
