package ct_test

import (
	"strings"
	"testing"

	"github.com/protolambda/chord/core/attr"
	"github.com/protolambda/chord/core/elem"
	"github.com/protolambda/chord/ct"
	"github.com/protolambda/chord/html/form"
	"github.com/protolambda/chord/html/form/input"
	"github.com/protolambda/chord/html/form/label"
	"github.com/protolambda/chord/html/form/textarea"
	"github.com/protolambda/chord/html/group/div"
	"github.com/protolambda/chord/html/meta"
	"github.com/protolambda/chord/html/script"
	"github.com/protolambda/chord/html/section"
	"github.com/protolambda/chord/html/text"
)

// secrets are the values of secretsPage that must never appear in diagnostics.
var secrets = []string{
	"abandon-ability", "seed-words", "0xprivate", "open-sesame", "key-123", "otp-987",
	"csrf-meta", "backup-words", "backup-input", "recovery-code", "card-4242",
}

func secretsPage() elem.Node {
	return section.Main()(
		meta.Meta(attr.KV("name", "csrf-token"), attr.KV("content", "csrf-meta")),
		form.Form(attr.ID("import"))(
			textarea.Textarea(attr.KV("name", "mnemonic"))(text.Text("abandon-ability")),
			input.Input(input.Type(input.Text), attr.KV("name", "seed_phrase"), attr.KV("value", "seed-words")),
			input.Input(input.Type(input.Text), attr.KV("name", "privateKey"), attr.KV("value", "0xprivate")),
			input.Input(input.Type(input.Text), attr.KV("name", "passphrase"), attr.KV("value", "open-sesame")),
			input.Input(input.Type(input.Text), attr.KV("name", "api-key"), attr.KV("value", "key-123")),
			input.Input(input.Type(input.Text), attr.KV("name", "code"), attr.KV("autocomplete", "one-time-code"), attr.KV("value", "otp-987")),
			input.Input(input.Type(input.Text), attr.KV("name", "card"), attr.KV("autocomplete", "cc-number"), attr.KV("value", "card-4242")),
			text.P()(text.Text("Keep this safe")),
		),
		div.Div(attr.ID("backup"), attr.Data("testid", "seed"))(
			text.Text("Your words: "),
			text.Code()(text.Text("backup-words")),
			input.Input(input.Type(input.Text), attr.KV("value", "backup-input")),
			elem.Comment("backup-words"),
		),
		text.Pre(attr.Class("codes"))(text.Text("recovery-code")),
	)
}

func noSecrets(t *testing.T, err error) {
	t.Helper()
	for _, s := range secrets {
		if strings.Contains(err.Error(), s) {
			t.Fatalf("secret %q leaked in diagnostics:\n%v", s, err)
		}
	}
}

func TestRedactSecretValues(t *testing.T) {
	page := ct.View(secretsPage(), ct.WithRedactContent(ct.Class("codes")))

	err := mustFail(t, page.Find(ct.Tag("main")).Find(ct.Tag("select")), ct.ErrCount,
		`meta [name="csrf-token" content="[redacted]"]`,
		`textarea [name="mnemonic"]`,
		`input [type="text" name="seed_phrase" value="[redacted]"]`,
		`input [type="text" name="privateKey" value="[redacted]"]`,
		`input [type="text" name="passphrase" value="[redacted]"]`,
		`input [type="text" name="api-key" value="[redacted]"]`,
		`input [type="text" name="code" autocomplete="one-time-code" value="[redacted]"]`,
		`input [type="text" name="card" autocomplete="cc-number" value="[redacted]"]`,
		`"Keep this safe"`,
	)
	noSecrets(t, err)
}

func TestRedactSensitiveText(t *testing.T) {
	page := ct.View(secretsPage(), ct.WithRedactContent(ct.Class("codes")))

	// Outlines: text inside a sensitive element, and the values in it.
	err := mustFail(t, page.Find(ct.ID("import")).Find(ct.Tag("select")), ct.ErrCount,
		"\n      [redacted]\n", `"Keep this safe"`)
	noSecrets(t, err)
	err = mustFail(t, page.Find(ct.ID("backup")).Find(ct.Tag("select")), ct.ErrCount,
		"\n    [redacted]\n", `input [type="text" value="[redacted]"]`, "<!-- [redacted] -->")
	noSecrets(t, err)
	err = mustFail(t, page.Find(ct.Tag("pre")).Find(ct.Tag("select")), ct.ErrCount, "pre.codes\n    [redacted]")
	noSecrets(t, err)

	// Texts and AttrValues report what they got, redacted.
	err = mustFail(t, page.Find(ct.Tag("main")).Texts("x"), ct.ErrMismatch, `got ["[redacted] Keep this safe [redacted] [redacted]"]`)
	noSecrets(t, err)
	err = mustFail(t, page.Find(ct.ID("backup")).Texts("x"), ct.ErrMismatch, `got ["[redacted]"]`)
	noSecrets(t, err)
	err = mustFail(t, page.Find(ct.Tag("input")).AttrValues("value", "x"), ct.ErrMismatch,
		`got ["[redacted]" "[redacted]" "[redacted]" "[redacted]" "[redacted]" "[redacted]" "[redacted]"]`)
	noSecrets(t, err)

	// Values that tests read are not redacted.
	words, err := page.Find(ct.ID("backup")).Find(ct.Tag("code")).Text(t.Context())
	if err != nil || words != "backup-words" {
		t.Fatalf("text: %q %v", words, err)
	}
}

// Rule violations that describe nodes apply the options of the subject too.
func TestRedactRuleViolations(t *testing.T) {
	page := ct.View(section.Main()(
		div.Div(attr.Class("codes"))(elem.Raw("recovery-code")),
		label.Label(attr.KV("for", "target"))(text.Text("Codes")),
		div.Div(attr.ID("target"), attr.Class("codes"), attr.Data("otp", "otp-987"), attr.KV("content", "card-4242")),
	),
		ct.WithRedactContent(ct.Class("codes")),
		ct.WithRedact(func(tag, key string) bool { return tag == "div" && key == "data-otp" }),
	)
	err := mustFail(t, page.Valid(ct.NoRaw(), ct.LabelReferences()), ct.ErrRule,
		"raw([redacted])", `data-otp="[redacted]"`, `content="[redacted]"`)
	noSecrets(t, err)
}

func TestRedactContentWithoutOption(t *testing.T) {
	page := ct.View(secretsPage())
	// Without the option, only the built-in markers apply.
	err := mustFail(t, page.Find(ct.Tag("pre")).Find(ct.Tag("select")), ct.ErrCount, `"recovery-code"`)
	if strings.Contains(err.Error(), "abandon-ability") {
		t.Fatalf("secret leaked:\n%v", err)
	}
}

func TestRedactInsideUnrenderedElements(t *testing.T) {
	page := ct.View(section.Main()(
		elem.Name("template").New()(text.Text("Words: "), textarea.Textarea(attr.KV("name", "mnemonic"))(text.Text("abandon-ability"))),
	))
	err := mustFail(t, page.Find(ct.Tag("template")).Texts("x"), ct.ErrMismatch, `got ["Words: [redacted]"]`)
	noSecrets(t, err)
}

// The redacted texts of a failure are the compared inner texts with the
// sensitive parts replaced: sensitive content that the inner text leaves
// out (in a script, template or noscript) adds no marker.
func TestRedactedTextsFollowInnerText(t *testing.T) {
	page := ct.View(div.Div(attr.ID("box"))(
		script.Inline(`window.cfg = {"a": "session-secret"};`, attr.ID("session-token"), script.Type("application/json")),
		text.P()(text.Text("Visible")),
		div.Div(attr.ID("token-holder"))(elem.Name("template").New()(text.Text("template-secret"))),
		script.Noscript(attr.KV("name", "csrf"))(text.Text("noscript-secret")),
	))
	box := page.Find(ct.ID("box"))
	got, err := box.Text(t.Context())
	if err != nil || got != "Visible" {
		t.Fatalf("text: %q %v", got, err)
	}
	for _, sel := range []*ct.Selection{box, page.Find(ct.ID("token-holder"))} {
		err := mustFail(t, sel.Texts("Other"), ct.ErrMismatch)
		if strings.Contains(err.Error(), "[redacted]") {
			t.Fatalf("redaction marker for text that is not compared:\n%v", err)
		}
		for _, s := range []string{"session-secret", "template-secret", "noscript-secret"} {
			if strings.Contains(err.Error(), s) {
				t.Fatalf("secret %q leaked in diagnostics:\n%v", s, err)
			}
		}
	}
	mustFail(t, box.Texts("Other"), ct.ErrMismatch, `got ["Visible"]`)
	mustFail(t, page.Find(ct.ID("token-holder")).Texts("Other"), ct.ErrMismatch, `got [""]`)
	// Texts of an unrendered element compare its text content, redacted.
	err = mustFail(t, page.Find(ct.ID("session-token")).Texts("Other"), ct.ErrMismatch, `got ["[redacted]"]`)
	if strings.Contains(err.Error(), "session-secret") {
		t.Fatalf("secret leaked in diagnostics:\n%v", err)
	}
}
