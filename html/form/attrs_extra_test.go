package form_test

import (
	"context"
	"strings"
	"testing"

	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/core/attr"
	"github.com/protolambda/chord/core/elem"
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
			if got, want := obj.Kind, attr.KindValue; got != want {
				t.Fatalf("forged enum value must be a logical (escaped) value, got kind %s", got)
			}
			var out strings.Builder
			if err := core.Render(context.Background(), elem.Name("x").Void(node), &out); err != nil {
				t.Fatalf("render: %v", err)
			}
			if got, want := out.String(), map[string]string{
				"method":      `<x method="post&#34; onsubmit=&#34;attack()"/>`,
				"enctype":     `<x enctype="text/plain&#34; onsubmit=&#34;attack()"/>`,
				"button type": `<x type="button&#34; onclick=&#34;attack()"/>`,
				"input type":  `<x type="text&#34; onfocus=&#34;attack()"/>`,
			}[name]; got != want {
				t.Fatalf("rendered: got %s, want %s", got, want)
			}
		})
	}
}
