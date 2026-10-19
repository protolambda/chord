package ct

import (
	"context"
	"fmt"
	"strings"

	"github.com/protolambda/mustbe/assertion"
)

// Rule is a document-wide structural check, such as unique ids or resolving
// references. Rules are not full HTML or accessibility validation.
type Rule struct {
	desc string
	// check describes the violations; o holds the redaction options of the
	// subject for node descriptions, and may be nil.
	check func(root Node, o *options) []string
}

func (r Rule) String() string {
	return r.desc
}

// Violations runs the rule against a document root and returns one
// description per violation. It exists for custom assertions; prefer
// [Subject.Valid], which also applies the redaction options of the subject
// to the descriptions (Violations only applies the built-in redaction).
func (r Rule) Violations(root Node) []string {
	return r.check(root, nil)
}

// CustomRule creates a rule from a description and a check that returns one
// description per violation. Keep secret values out of the descriptions:
// they appear in failure messages as written. [DescribedRule] helps with
// that.
func CustomRule(desc string, check func(root Node) []string) Rule {
	return Rule{desc: desc, check: func(root Node, _ *options) []string { return check(root) }}
}

// DescribedRule creates a rule like [CustomRule], whose check formats nodes
// and values for its descriptions with a [Describer]: [Subject.Valid]
// passes one with the redaction options of the subject, and
// [Rule.Violations] the zero Describer, which applies the built-in
// redaction only.
func DescribedRule(desc string, check func(root Node, d Describer) []string) Rule {
	return Rule{desc: desc, check: func(root Node, o *options) []string { return check(root, Describer{opts: o}) }}
}

// Valid asserts that the document satisfies every rule.
func (s *Subject) Valid(rules ...Rule) assertion.Assertion {
	return validAssertion{subject: s, rules: rules}
}

type validAssertion struct {
	subject *Subject
	rules   []Rule
}

var _ assertion.Assertion = validAssertion{}

func (a validAssertion) String() string {
	parts := make([]string, 0, len(a.rules)+1)
	parts = append(parts, a.subject.source.String())
	for _, r := range a.rules {
		parts = append(parts, r.desc)
	}
	return "valid(" + strings.Join(parts, ", ") + ")"
}

func (a validAssertion) Check(ctx context.Context) error {
	doc, err := a.subject.Load(ctx)
	if err != nil {
		return err
	}
	var b strings.Builder
	for _, r := range a.rules {
		for _, v := range r.check(doc.Root(), &a.subject.opts) {
			fmt.Fprintf(&b, "\n  %s: %s", r.desc, v)
		}
	}
	if b.Len() == 0 {
		return nil
	}
	return fmt.Errorf("%w in %s:%s", ErrRule, a.subject.source, b.String())
}

// UniqueIDs requires every non-empty id to occur once.
func UniqueIDs() Rule {
	return Rule{desc: "uniqueIDs()", check: func(root Node, _ *options) []string {
		seen := make(map[string]Node)
		var out []string
		for e := range Elements(root) {
			id, ok := e.Attr("id")
			if !ok || id == "" {
				continue
			}
			if first, dup := seen[id]; dup {
				out = append(out, fmt.Sprintf("duplicate id %q at %s and %s", id, locationOf(first), locationOf(e)))
				continue
			}
			seen[id] = e
		}
		return out
	}}
}

// labelable reports whether an element can be labeled by a label element.
func labelable(n Node) bool {
	switch n.Tag() {
	case "button", "meter", "output", "progress", "select", "textarea":
		return true
	case "input":
		return inputType(n) != "hidden"
	}
	return false
}

// LabelReferences requires every label with a for attribute to reference a
// labelable element.
func LabelReferences() Rule {
	return Rule{desc: "labelReferences()", check: func(root Node, o *options) []string {
		var out []string
		for e := range Elements(root) {
			if e.Tag() != "label" {
				continue
			}
			id, ok := e.Attr("for")
			if !ok {
				continue
			}
			target := elementByID(root, id)
			switch {
			case target == nil:
				out = append(out, fmt.Sprintf("label at %s references missing id %q", locationOf(e), id))
			case !labelable(target):
				out = append(out, fmt.Sprintf("label at %s references non-labelable %s", locationOf(e), describeNode(target, o)))
			}
		}
		return out
	}}
}

// ariaReferenceAttrs are the ARIA attributes whose value is an id list.
var ariaReferenceAttrs = []string{
	"aria-labelledby", "aria-describedby", "aria-controls", "aria-owns",
	"aria-activedescendant", "aria-details", "aria-errormessage", "aria-flowto",
}

// ARIAReferences requires every id in an ARIA id-list attribute to resolve.
func ARIAReferences() Rule {
	return Rule{desc: "ariaReferences()", check: func(root Node, _ *options) []string {
		var out []string
		for e := range Elements(root) {
			for _, key := range ariaReferenceAttrs {
				ids, ok := e.Attr(key)
				if !ok {
					continue
				}
				for id := range strings.FieldsSeq(ids) {
					if elementByID(root, id) == nil {
						out = append(out, fmt.Sprintf("%s at %s references missing id %q", key, locationOf(e), id))
					}
				}
			}
		}
		return out
	}}
}

// LocalTargets requires every link to a local fragment ("#id") to resolve.
// The bare "#" and "#top" targets are always valid.
func LocalTargets() Rule {
	return Rule{desc: "localTargets()", check: func(root Node, _ *options) []string {
		var out []string
		for e := range Elements(root) {
			if e.Tag() != "a" && e.Tag() != "area" {
				continue
			}
			href, ok := e.Attr("href")
			if !ok || !strings.HasPrefix(href, "#") {
				continue
			}
			id := href[1:]
			if id == "" || id == "top" {
				continue
			}
			if elementByID(root, id) == nil {
				out = append(out, fmt.Sprintf("link at %s targets missing id %q", locationOf(e), id))
			}
		}
		return out
	}}
}

// ImagesHaveAlt requires every img to have an alt attribute, possibly empty.
func ImagesHaveAlt() Rule {
	return Rule{desc: "imagesHaveAlt()", check: func(root Node, _ *options) []string {
		var out []string
		for e := range Elements(root) {
			if e.Tag() != "img" {
				continue
			}
			if _, ok := e.Attr("alt"); !ok {
				out = append(out, fmt.Sprintf("image at %s has no alt attribute", locationOf(e)))
			}
		}
		return out
	}}
}

// ButtonsHaveType requires every button inside a form to state its type,
// so that a button does not submit the form by accident.
func ButtonsHaveType() Rule {
	return Rule{desc: "buttonsHaveType()", check: func(root Node, _ *options) []string {
		var out []string
		for e := range Elements(root) {
			if e.Tag() != "button" {
				continue
			}
			if _, ok := e.Attr("type"); ok {
				continue
			}
			for p := e.Parent(); p != nil; p = p.Parent() {
				if p.Kind() == KindElement && p.Tag() == "form" {
					out = append(out, fmt.Sprintf("button at %s inside a form has no type", locationOf(e)))
					break
				}
			}
		}
		return out
	}}
}

// NoInlineHandlers forbids inline event handler attributes (on*). It does
// not report the inline handlers of htmx, such as hx-on:click, whose names
// do not start with "on": see [NoHxOn].
func NoInlineHandlers() Rule {
	return Rule{desc: "noInlineHandlers()", check: func(root Node, _ *options) []string {
		var out []string
		for e := range Elements(root) {
			for key := range e.Attrs() {
				if len(key) > 2 && strings.HasPrefix(strings.ToLower(key), "on") {
					out = append(out, fmt.Sprintf("inline handler %s at %s", key, locationOf(e)))
				}
			}
		}
		return out
	}}
}

// NoHxOn forbids the inline handler attributes of htmx, which htmx runs as
// code, as browsers run on* handlers: hx-on:* and hx-on-* (such as
// hx-on:click and hx-on::after-request), the older hx-on attribute of
// htmx 1, and each with a data- prefix. Names are compared
// case-insensitively.
func NoHxOn() Rule {
	return Rule{desc: "noHxOn()", check: func(root Node, _ *options) []string {
		var out []string
		for e := range Elements(root) {
			for key := range e.Attrs() {
				if hxOn(key) {
					out = append(out, fmt.Sprintf("htmx handler %s at %s", key, locationOf(e)))
				}
			}
		}
		return out
	}}
}

// hxOn reports whether an attribute name is an htmx inline handler.
func hxOn(key string) bool {
	name := strings.TrimPrefix(strings.ToLower(key), "data-")
	return name == "hx-on" || strings.HasPrefix(name, "hx-on:") || strings.HasPrefix(name, "hx-on-")
}

// NoRaw forbids raw HTML nodes. Only direct views contain them; parsed
// documents never do.
func NoRaw() Rule {
	return Rule{desc: "noRaw()", check: func(root Node, o *options) []string {
		var out []string
		for d := range Descendants(root) {
			if d.Kind() == KindRaw {
				out = append(out, fmt.Sprintf("raw content at %s: %s", locationOf(d), describeNode(d, o)))
			}
		}
		return out
	}}
}
