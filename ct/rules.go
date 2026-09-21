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
	desc  string
	check func(root Node) []string
}

func (r Rule) String() string {
	return r.desc
}

// Violations runs the rule against a document root and returns one
// description per violation. It exists for custom assertions; prefer
// [Subject.Valid].
func (r Rule) Violations(root Node) []string {
	return r.check(root)
}

// CustomRule creates a rule from a description and a check that returns one
// description per violation.
func CustomRule(desc string, check func(root Node) []string) Rule {
	return Rule{desc: desc, check: check}
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
		for _, v := range r.check(doc.Root()) {
			fmt.Fprintf(&b, "\n  %s: %s", r.desc, v)
		}
	}
	if b.Len() == 0 {
		return nil
	}
	return fmt.Errorf("%w in %s:%s", ErrRule, a.subject.source, b.String())
}

// elements iterates all element nodes of the tree in document order.
func elements(root Node) func(yield func(Node) bool) {
	return func(yield func(Node) bool) {
		for d := range descendants(root) {
			if d.Kind() == KindElement && !yield(d) {
				return
			}
		}
	}
}

// UniqueIDs requires every non-empty id to occur once.
func UniqueIDs() Rule {
	return Rule{desc: "uniqueIDs()", check: func(root Node) []string {
		seen := make(map[string]Node)
		var out []string
		for e := range elements(root) {
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
	return Rule{desc: "labelReferences()", check: func(root Node) []string {
		var out []string
		for e := range elements(root) {
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
				out = append(out, fmt.Sprintf("label at %s references non-labelable %s", locationOf(e), describeNode(target, nil)))
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
	return Rule{desc: "ariaReferences()", check: func(root Node) []string {
		var out []string
		for e := range elements(root) {
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
	return Rule{desc: "localTargets()", check: func(root Node) []string {
		var out []string
		for e := range elements(root) {
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
	return Rule{desc: "imagesHaveAlt()", check: func(root Node) []string {
		var out []string
		for e := range elements(root) {
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
	return Rule{desc: "buttonsHaveType()", check: func(root Node) []string {
		var out []string
		for e := range elements(root) {
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

// NoInlineHandlers forbids inline event handler attributes (on*).
func NoInlineHandlers() Rule {
	return Rule{desc: "noInlineHandlers()", check: func(root Node) []string {
		var out []string
		for e := range elements(root) {
			for key := range e.Attrs() {
				if len(key) > 2 && strings.HasPrefix(strings.ToLower(key), "on") {
					out = append(out, fmt.Sprintf("inline handler %s at %s", key, locationOf(e)))
				}
			}
		}
		return out
	}}
}

// NoRaw forbids raw HTML nodes. Only direct views contain them; parsed
// documents never do.
func NoRaw() Rule {
	return Rule{desc: "noRaw()", check: func(root Node) []string {
		var out []string
		for d := range descendants(root) {
			if d.Kind() == KindRaw {
				out = append(out, fmt.Sprintf("raw content at %s: %s", locationOf(d), describeNode(d, nil)))
			}
		}
		return out
	}}
}
