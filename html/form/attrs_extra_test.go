package form_test

import (
	"context"
	"testing"

	"github.com/protolambda/chord/core/attr"
	"github.com/protolambda/chord/html/form"
	"github.com/protolambda/chord/html/form/button"
	"github.com/protolambda/chord/html/form/input"
)

func TestForgedEnumValuesAreEscaped(t *testing.T) {
	tests := map[string]attr.Node{
		"method":      form.Method(`post" onsubmit="attack()`),
		"enctype":     form.Enctype(`text/plain" onsubmit="attack()`),
		"button type": button.Type(`button" onclick="attack()`),
		"input type":  input.Type(`text" onfocus="attack()`),
	}

	for name, node := range tests {
		t.Run(name, func(t *testing.T) {
			obj, err := node.Eval(context.Background())
			if err != nil {
				t.Fatalf("eval: %v", err)
			}
			if got, want := obj.Val, map[string]string{
				"method":      "post&#34; onsubmit=&#34;attack()",
				"enctype":     "text/plain&#34; onsubmit=&#34;attack()",
				"button type": "button&#34; onclick=&#34;attack()",
				"input type":  "text&#34; onfocus=&#34;attack()",
			}[name]; got != want {
				t.Fatalf("value: got %q, want %q", got, want)
			}
		})
	}
}
