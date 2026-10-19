package ct

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

// ValueMatch is an immutable, descriptive string predicate, used where a
// query or assertion accepts more than exact equality. The zero ValueMatch
// matches nothing.
type ValueMatch struct {
	desc  string
	match func(v string) bool
}

func (m ValueMatch) String() string {
	return m.desc
}

// Match reports whether v satisfies the matcher.
func (m ValueMatch) Match(v string) bool {
	return m.match != nil && m.match(v)
}

// MatchFunc creates a matcher from a description and a predicate. The
// description names the matcher in assertion descriptions and failure
// messages, e.g. "jsonObject()". A nil predicate matches nothing.
func MatchFunc(desc string, match func(v string) bool) ValueMatch {
	return ValueMatch{desc: desc, match: match}
}

// And requires this matcher and all others to match.
func (m ValueMatch) And(others ...ValueMatch) ValueMatch {
	if len(others) == 0 {
		return m
	}
	all := append([]ValueMatch{m}, others...)
	return ValueMatch{
		desc: "and(" + describeMatches(all) + ")",
		match: func(v string) bool {
			for _, m := range all {
				if !m.Match(v) {
					return false
				}
			}
			return true
		},
	}
}

// Or requires this matcher or any of the others to match.
func (m ValueMatch) Or(others ...ValueMatch) ValueMatch {
	if len(others) == 0 {
		return m
	}
	all := append([]ValueMatch{m}, others...)
	return ValueMatch{
		desc: "or(" + describeMatches(all) + ")",
		match: func(v string) bool {
			for _, m := range all {
				if m.Match(v) {
					return true
				}
			}
			return false
		},
	}
}

// Not inverts the matcher. The inversion of the zero ValueMatch matches
// nothing too, so that a forgotten matcher does not turn into one that
// accepts every value.
func (m ValueMatch) Not() ValueMatch {
	inverted := ValueMatch{desc: "not(" + m.desc + ")"}
	if m.match != nil {
		inverted.match = func(v string) bool { return !m.match(v) }
	}
	return inverted
}

func describeMatches(matches []ValueMatch) string {
	parts := make([]string, len(matches))
	for i, m := range matches {
		parts[i] = m.String()
	}
	return strings.Join(parts, ", ")
}

// Exact matches the value exactly.
func Exact(value string) ValueMatch {
	return ValueMatch{
		desc:  fmt.Sprintf("%q", value),
		match: func(v string) bool { return v == value },
	}
}

// Contains matches values containing the substring.
func Contains(sub string) ValueMatch {
	return ValueMatch{
		desc:  fmt.Sprintf("contains(%q)", sub),
		match: func(v string) bool { return strings.Contains(v, sub) },
	}
}

// Prefix matches values starting with the prefix.
func Prefix(prefix string) ValueMatch {
	return ValueMatch{
		desc:  fmt.Sprintf("prefix(%q)", prefix),
		match: func(v string) bool { return strings.HasPrefix(v, prefix) },
	}
}

// Suffix matches values ending with the suffix.
func Suffix(suffix string) ValueMatch {
	return ValueMatch{
		desc:  fmt.Sprintf("suffix(%q)", suffix),
		match: func(v string) bool { return strings.HasSuffix(v, suffix) },
	}
}

// Regexp matches values containing a match of the expression.
func Regexp(re *regexp.Regexp) ValueMatch {
	return ValueMatch{
		desc:  fmt.Sprintf("regexp(%q)", re.String()),
		match: re.MatchString,
	}
}

// OneOf matches values equal to one of the given values. Without values it
// matches nothing.
func OneOf(values ...string) ValueMatch {
	values = slices.Clone(values)
	quoted := make([]string, len(values))
	for i, v := range values {
		quoted[i] = strconv.Quote(v)
	}
	return ValueMatch{
		desc:  "oneOf(" + strings.Join(quoted, ", ") + ")",
		match: func(v string) bool { return slices.Contains(values, v) },
	}
}

// Token matches a whitespace-separated list, such as the rel or
// autocomplete attribute, that contains the token, compared
// case-insensitively as HTML compares such keywords. Classes are
// case-sensitive: use the [Class] query for them.
func Token(token string) ValueMatch {
	return ValueMatch{
		desc: fmt.Sprintf("token(%q)", token),
		match: func(v string) bool {
			for f := range strings.FieldsSeq(v) {
				if strings.EqualFold(f, token) {
					return true
				}
			}
			return false
		},
	}
}

// CommaToken matches a comma-separated list, such as the Vary,
// Cache-Control or Connection header, with an item equal to the token,
// compared case-insensitively and without the whitespace around items.
// Items are compared whole: CommaToken("max-age=0") matches
// "no-cache, max-age=0", and CommaToken("max-age") does not.
func CommaToken(token string) ValueMatch {
	return ValueMatch{
		desc: fmt.Sprintf("commaToken(%q)", token),
		match: func(v string) bool {
			for item := range strings.SplitSeq(v, ",") {
				if strings.EqualFold(strings.TrimSpace(item), token) {
					return true
				}
			}
			return false
		},
	}
}
