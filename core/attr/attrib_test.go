package attr_test

import (
	"context"
	"errors"
	"testing"

	"github.com/protolambda/chord/core/attr"
)

func TestKVValidatesNameAndKeepsLogicalValue(t *testing.T) {
	obj, err := attr.KV("data-user", `one" two<&`).Eval(context.Background())
	if err != nil {
		t.Fatalf("eval: %v", err)
	}
	if got, want := obj.Key, "data-user"; got != want {
		t.Fatalf("key: got %q, want %q", got, want)
	}
	if got, want := obj.Kind, attr.KindValue; got != want {
		t.Fatalf("kind: got %s, want %s", got, want)
	}
	if got, want := obj.Val, `one" two<&`; got != want {
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

	logical, err := name.Value(`one" two`).Eval(context.Background())
	if err != nil {
		t.Fatalf("logical eval: %v", err)
	}
	if logical.Kind != attr.KindValue || logical.Val != `one" two` {
		t.Fatalf("unexpected logical attribute: %+v", logical)
	}

	raw, err := name.Raw(`one&#34; two`).Eval(context.Background())
	if err != nil {
		t.Fatalf("raw eval: %v", err)
	}
	if raw.Kind != attr.KindRawValue || raw.Val != `one&#34; two` {
		t.Fatalf("unexpected raw attribute: %+v", raw)
	}

	boolean, err := attr.Name("disabled").Bool().Eval(context.Background())
	if err != nil {
		t.Fatalf("boolean eval: %v", err)
	}
	if boolean.Kind != attr.KindBool || boolean.Key != "disabled" {
		t.Fatalf("unexpected boolean attribute: %+v", boolean)
	}
}

func TestObjValidate(t *testing.T) {
	valid := map[string]attr.Obj{
		"noop":          {},
		"legacy bundle": {Sub: func(func(attr.Node) bool) {}},
		"value":         {Kind: attr.KindValue, Key: "id", Val: "x"},
		"raw":           {Kind: attr.KindRawValue, Key: "id", Val: "x"},
		"bool":          {Kind: attr.KindBool, Key: "disabled"},
	}
	for name, obj := range valid {
		t.Run("valid "+name, func(t *testing.T) {
			if _, err := obj.Validate(); err != nil {
				t.Fatalf("expected valid object, got %v", err)
			}
		})
	}

	invalid := map[string]attr.Obj{
		"key without kind":  {Key: "id", Val: "x"},
		"bundle with key":   {Kind: attr.KindBundle, Key: "id"},
		"value without key": {Kind: attr.KindValue, Val: "x"},
		"bool with value":   {Kind: attr.KindBool, Key: "disabled", Val: "x"},
		"unknown kind":      {Kind: attr.Kind(42), Key: "x"},
	}
	for name, obj := range invalid {
		t.Run("invalid "+name, func(t *testing.T) {
			if _, err := obj.Validate(); !errors.Is(err, attr.ErrInvalidObj) {
				t.Fatalf("expected ErrInvalidObj, got %v", err)
			}
		})
	}
}
