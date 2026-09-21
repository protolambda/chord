package ct

import (
	"fmt"
	"regexp"
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
