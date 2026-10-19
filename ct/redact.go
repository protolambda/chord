package ct

import "strings"

// secretWords are fragments of attribute names, and of the names, ids and
// test ids of elements, that indicate a value which must not appear in test
// output. They are compared with lowercase names without the separators
// '-', '_', '.', ':' and spaces, so "api-key" is "apikey" and "seed_phrase"
// contains "seed".
var secretWords = []string{
	"password", "passwd", "passphrase", "secret", "token", "csrf", "xsrf",
	"authorization", "cookie", "credential", "apikey", "mnemonic", "seed", "private",
}

// secretAutocomplete are the autocomplete tokens of fields that hold
// credentials or payment card secrets.
var secretAutocomplete = map[string]bool{
	"current-password": true, "new-password": true, "one-time-code": true,
	"cc-number": true, "cc-csc": true,
}

// LooksSecret reports whether a name, such as the name of an attribute, an
// element or an HTTP header, looks like it holds a secret: whether it
// contains one of the words listed in the package documentation (see
// Redaction), ignoring case and the separators '-', '_', '.', ':' and
// spaces. Custom assertions use it to keep secret values out of failure
// messages; ct/cthttp redacts the values of such headers.
func LooksSecret(s string) bool {
	var b strings.Builder
	for _, r := range strings.ToLower(s) {
		switch r {
		case '-', '_', '.', ':', ' ':
		default:
			b.WriteRune(r)
		}
	}
	norm := b.String()
	for _, w := range secretWords {
		if strings.Contains(norm, w) {
			return true
		}
	}
	return false
}

// sensitive reports whether the element marks its content as secret: its
// type is password, its name, id or data-testid looks secret, its
// autocomplete is for a credential or card secret, or it matches a query of
// [WithRedactContent]. A query error counts as a match. o may be nil.
func (o *options) sensitive(n Node) bool {
	if n.Kind() != KindElement {
		return false
	}
	if typ, ok := n.Attr("type"); ok && strings.EqualFold(strings.TrimSpace(typ), "password") {
		return true
	}
	for _, key := range []string{"name", "id", "data-testid"} {
		if v, ok := n.Attr(key); ok && LooksSecret(v) {
			return true
		}
	}
	if ac, ok := n.Attr("autocomplete"); ok {
		for token := range strings.FieldsSeq(strings.ToLower(ac)) {
			if secretAutocomplete[token] {
				return true
			}
		}
	}
	if o != nil {
		for _, q := range o.redactContent {
			if ok, err := q.Match(n); ok || err != nil {
				return true
			}
		}
	}
	return false
}

// inSensitive reports whether n is, or is inside, a sensitive element.
func (o *options) inSensitive(n Node) bool {
	for cur := n; cur != nil; cur = cur.Parent() {
		if o.sensitive(cur) {
			return true
		}
	}
	return false
}

// redactValue hides attribute values that look like secrets: any attribute
// whose name looks secret, an attribute that the [WithRedact] predicate
// selects, and the value and content attributes of a sensitive element or
// of an element inside one. o may be nil.
func (o *options) redactValue(n Node, key, val string) string {
	if val == "" {
		return val
	}
	if o != nil && o.redact != nil && o.redact(n.Tag(), key) {
		return redactedValue
	}
	if LooksSecret(key) {
		return redactedValue
	}
	if (key == "value" || key == "content") && o.inSensitive(n) {
		return redactedValue
	}
	return val
}

// redactedInnerText is the inner text of n for diagnostics: the text of
// sensitive elements is replaced by a redaction marker.
func (o *options) redactedInnerText(n Node) string {
	if o.inSensitive(n) {
		if innerText(n) == "" {
			return ""
		}
		return redactedValue
	}
	w := textWriter{raw: !rendered(n), redact: o.sensitive}
	w.children(n)
	return collapseSpace(w.b.String())
}
