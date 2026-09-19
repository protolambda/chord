package attr_test

import (
	"context"
	"errors"
	"testing"

	"github.com/protolambda/chord/core/attr"
)

func TestKVValidatesNameAndEscapesValue(t *testing.T) {
	obj, err := attr.KV("data-user", `one" two<&`).Eval(context.Background())
	if err != nil {
		t.Fatalf("eval: %v", err)
	}
	if got, want := obj.Key, "data-user"; got != want {
		t.Fatalf("key: got %q, want %q", got, want)
	}
	if got, want := obj.Val, "one&#34; two&lt;&amp;"; got != want {
		t.Fatalf("value: got %q, want %q", got, want)
	}
}

func TestKVRejectsInvalidNames(t *testing.T) {
	for _, name := range []string{"", "two words", `x"y`, "x'y", "x=y", "x/y", "x<y", "x>y", "x\x00y", string([]byte{0xff})} {
		t.Run(name, func(t *testing.T) {
			_, err := attr.KV(name, "value").Eval(context.Background())
			if !errors.Is(err, attr.ErrInvalidName) {
				t.Fatalf("expected ErrInvalidName, got %v", err)
			}
		})
	}
}

func TestBoolRejectsInvalidName(t *testing.T) {
	_, err := attr.Bool("two words").Eval(context.Background())
	if !errors.Is(err, attr.ErrInvalidName) {
		t.Fatalf("expected ErrInvalidName, got %v", err)
	}
}

func TestKVAcceptsExtensionNames(t *testing.T) {
	for _, name := range []string{"data-user", "aria-label", "hx-on:click", "@click", ":class"} {
		t.Run(name, func(t *testing.T) {
			if _, err := attr.KV(name, "value").Eval(context.Background()); err != nil {
				t.Fatalf("eval: %v", err)
			}
		})
	}
}

func TestParseName(t *testing.T) {
	for _, value := range []string{"class", "data-user", "aria-label", "hx-on:click", "@click", ":class"} {
		t.Run("valid "+value, func(t *testing.T) {
			got, ok := attr.ParseName(value)
			if !ok || string(got) != value {
				t.Fatalf("ParseName(%q) = %q, %t", value, got, ok)
			}
		})
	}

	for _, value := range []string{"", "two words", `x"y`, "x'y", "x=y", "x/y", "x<y", "x>y", "x\x00y", string([]byte{0xff})} {
		t.Run("invalid "+value, func(t *testing.T) {
			if got, ok := attr.ParseName(value); ok || got != "" {
				t.Fatalf("ParseName(%q) = %q, %t", value, got, ok)
			}
		})
	}
}

func TestNameOperations(t *testing.T) {
	name := attr.Name("title")

	escaped, err := name.Value(`one" two`).Eval(context.Background())
	if err != nil {
		t.Fatalf("escaped eval: %v", err)
	}
	if got, want := escaped.Val, "one&#34; two"; got != want {
		t.Fatalf("escaped value: got %q, want %q", got, want)
	}

	raw, err := name.Raw(`one&#34; two`).Eval(context.Background())
	if err != nil {
		t.Fatalf("raw eval: %v", err)
	}
	if got, want := raw.Val, `one&#34; two`; got != want {
		t.Fatalf("raw value: got %q, want %q", got, want)
	}

	boolean, err := attr.Name("disabled").Bool().Eval(context.Background())
	if err != nil {
		t.Fatalf("boolean eval: %v", err)
	}
	if !boolean.Bool || boolean.Key != "disabled" {
		t.Fatalf("unexpected boolean attribute: %+v", boolean)
	}
}
