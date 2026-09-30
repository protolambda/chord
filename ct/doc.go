// Package ct provides test subjects, queries, and assertions for Chord views.
//
// Every test operation is an ordinary [assertion.Assertion] value, so it
// works with mustbe.Must, mustbe.WrapT, and devtest.T without this package
// depending on any test runner. A test builds a subject, derives selections
// with queries, and asserts the resulting objects:
//
//	t := mustbe.WrapT(gt)
//	page := ct.View(AccountPage(account, viewer))
//
//	t.Must(page)                                                  // evaluation succeeds
//	t.Must(page.Find(ct.Role("heading", ct.Named("Account"))))   // exactly one match
//	t.Must(page.Find(ct.Role("link", ct.Named("Admin"))).None()) // absent
//
//	form := page.Find(ct.Tag("form"), ct.ID("profile"))           // every query must match
//	t.Must(form)
//	t.Must(form.Find(ct.Label("Email")).Matches(ct.Tag("input"), ct.Attr("name", "email")))
//	t.Must(page.Find(ct.Role("listitem")).Texts("Alpha", "Beta"))
//
// # Subjects and selections
//
// A [Subject] wraps a [Source] such as a Chord view ([View]), a parsed HTML
// document (package ct/html), or an HTTP response (package ct/http). It loads
// at most once, with the context of the first check, and every selection
// derived from it shares that result. Checking a subject asserts that it
// loads. A subject belongs to one test or subtest; do not share it across
// tests with different contexts.
//
// A [Selection] binds queries to a scope. Checking a selection asserts that
// exactly one element matches. [Selection.None], [Selection.Any],
// [Selection.Count], [Selection.AtLeast], and [Selection.AtMost] express
// other cardinalities. [Selection.Matches] checks the single match,
// [Selection.Each] checks every match, and [Selection.InOrder],
// [Selection.Texts], and [Selection.AttrValues] check the matched sequence in
// document order. [Selection.First], [Selection.Last], and [Selection.Nth]
// narrow a selection to one position. [Selection.Find] searches the
// descendants of every matched element, so assert the parent selection too
// when it must be unique.
//
// When a test needs a value from the page, such as the version in a hidden
// input or the URL that a panel polls, [Selection.Attr] and
// [Selection.Text] return it from the single match, and [Selection.Nodes]
// returns all matches for custom assertions.
//
// Negative assertions such as None load the subject first, and report an
// evaluation failure as a failure: "the query could not run" is never
// evidence of absence. Prefer them over a generic inversion of a selection.
//
// # Queries
//
// Queries are immutable descriptive values. Prefer, in order: [Role] with
// [Named], [Label], [Text], [InnerText] and [Alt], protocol attributes such
// as href or name through [Attr], [TestID], and finally [Tag], [ID], and
// [Class] as implementation-level escape hatches. Combine them with
// [Query.And], [Query.Or], [Query.Not], [HasChild], and [HasDescendant].
//
// Text queries operate on logical text. [InnerText], [Selection.Texts] and
// accessible names separate the text of block-level boxes as a browser's
// innerText does, so a card with a header "Author" and a body "Signed"
// reads "Author Signed", and they leave out the content of script, style,
// template and head elements; the layout is approximated from the default
// display of HTML elements, display styles, and Bootstrap classes.
// [TextContent] is the raw DOM textContent, which joins that text. Text
// queries do not model visibility, hidden attributes, or script behavior.
// [Role] and [Named] implement a documented subset of implicit roles and
// accessible names, not a browser accessibility tree.
//
// # Diagnostics
//
// Failures describe the expectation, the scope, the number of matches with
// their locations, and an outline of the scope. Elements show as
// tag#id.class [name="value"]; a bare attribute name is a boolean attribute
// of a view (see [BoolAttrNode]), while parsed pages give boolean attributes
// the empty value, as the DOM does.
//
// # Redaction
//
// Failure messages hide values that look secret, and show [redacted]
// instead:
//   - the value of an attribute whose name looks secret;
//   - inside a sensitive element (the element and its descendants): text,
//     raw content, comments, and the value and content attributes. This
//     also applies to the texts and values that [Selection.Texts] and
//     [Selection.AttrValues] report when they fail.
//
// A name looks secret when it contains one of password, passwd,
// passphrase, secret, token, csrf, xsrf, authorization, cookie,
// credential, apikey, mnemonic, seed, or private, ignoring case and the
// separators '-', '_', '.', ':' and spaces. An element is sensitive when
// its type is password, when its name, id, or data-testid looks secret, or
// when its autocomplete is current-password, new-password, one-time-code,
// cc-number, or cc-csc. [WithRedact] hides more attributes and
// [WithRedactContent] marks more elements as sensitive.
//
// The heuristics can hide more than needed: a form with the id
// "password-reset" hides the texts inside it. They cannot recognize a
// secret by its content, so keep real secrets out of test fixtures. Values
// that a test reads ([Selection.Attr], [Selection.Text], [Selection.Nodes])
// and the expected values that it states are not redacted.
package ct
