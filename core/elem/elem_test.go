package elem_test

import (
	"context"
	"testing"

	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/core/attr"
	"github.com/protolambda/chord/core/elem"
)

func ExampleName_New() {
	core.Dump(elem.Name("section").New()(elem.Raw("content")))
	// Output: <section>content</section>
}

func ExampleName_Void() {
	core.Dump(elem.Name("hr").Void())
	// Output: <hr/>
}

func TestParseName(t *testing.T) {
	for _, value := range []string{"div", "H1", "my-element", "svg:path", "x.y", "x_y", "dév", "emotion-😍"} {
		t.Run("valid "+value, func(t *testing.T) {
			got, ok := elem.ParseName(value)
			if !ok || string(got) != value {
				t.Fatalf("ParseName(%q) = %q, %t", value, got, ok)
			}
		})
	}

	for _, value := range []string{"", "1div", "-element", "élement", "two words", `x"y`, "x'y", "x=y", "x/y", "x<y", "x>y", "x\x00y", string([]byte{0xff})} {
		t.Run("invalid "+value, func(t *testing.T) {
			if got, ok := elem.ParseName(value); ok || got != "" {
				t.Fatalf("ParseName(%q) = %q, %t", value, got, ok)
			}
		})
	}
}

func TestTagConstructors(t *testing.T) {
	name, ok := elem.ParseName("custom-element")
	if !ok {
		t.Fatal("expected custom-element to be a valid tag")
	}
	node := name.New(attr.ID("content"))(elem.Raw("body"))
	obj, err := node.Eval(context.Background())
	if err != nil {
		t.Fatalf("eval element: %v", err)
	}
	if obj.Tag != "custom-element" || obj.Void {
		t.Fatalf("unexpected element: %+v", obj)
	}

	void, err := elem.Name("input").Void().Eval(context.Background())
	if err != nil {
		t.Fatalf("eval void element: %v", err)
	}
	if void.Tag != "input" || !void.Void {
		t.Fatalf("unexpected void element: %+v", void)
	}
}
