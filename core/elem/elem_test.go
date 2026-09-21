package elem_test

import (
	"context"
	"errors"
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
	if obj.Kind != elem.KindElement || obj.Tag != "custom-element" || obj.Void {
		t.Fatalf("unexpected element: %+v", obj)
	}

	void, err := elem.Name("input").Void().Eval(context.Background())
	if err != nil {
		t.Fatalf("eval void element: %v", err)
	}
	if void.Kind != elem.KindElement || void.Tag != "input" || !void.Void {
		t.Fatalf("unexpected void element: %+v", void)
	}
}

func TestContentConstructors(t *testing.T) {
	tests := map[string]struct {
		node elem.Node
		kind elem.Kind
		data string
	}{
		"text":    {elem.Text("a < b"), elem.KindText, "a < b"},
		"raw":     {elem.Raw("<b>x</b>"), elem.KindRaw, "<b>x</b>"},
		"comment": {elem.Comment("note"), elem.KindComment, "note"},
		"noop":    {elem.Noop(), elem.KindNoop, ""},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			obj, err := tc.node.Eval(context.Background())
			if err != nil {
				t.Fatalf("eval: %v", err)
			}
			if obj.Kind != tc.kind || obj.Data != tc.data {
				t.Fatalf("got kind %s data %q, want kind %s data %q", obj.Kind, obj.Data, tc.kind, tc.data)
			}
		})
	}
}

func TestObjValidate(t *testing.T) {
	valid := map[string]struct {
		obj  elem.Obj
		kind elem.Kind
	}{
		"noop":            {elem.Obj{}, elem.KindNoop},
		"legacy element":  {elem.Obj{Tag: "div"}, elem.KindElement},
		"legacy fragment": {elem.Obj{Children: func(func(elem.Node) bool) {}}, elem.KindFragment},
		"void":            {elem.Obj{Kind: elem.KindElement, Tag: "br", Void: true}, elem.KindElement},
		"text":            {elem.Obj{Kind: elem.KindText, Data: "x"}, elem.KindText},
		"empty raw":       {elem.Obj{Kind: elem.KindRaw}, elem.KindRaw},
	}
	for name, tc := range valid {
		t.Run("valid "+name, func(t *testing.T) {
			kind, err := tc.obj.Validate()
			if err != nil {
				t.Fatalf("expected valid object, got %v", err)
			}
			if kind != tc.kind {
				t.Fatalf("kind: got %s, want %s", kind, tc.kind)
			}
		})
	}

	invalid := map[string]elem.Obj{
		"noop with data":        {Data: "x"},
		"fragment with tag":     {Kind: elem.KindFragment, Tag: "div"},
		"element without tag":   {Kind: elem.KindElement},
		"element with data":     {Kind: elem.KindElement, Tag: "div", Data: "x"},
		"void with children":    {Kind: elem.KindElement, Tag: "br", Void: true, Children: func(func(elem.Node) bool) {}},
		"text with tag":         {Kind: elem.KindText, Tag: "div", Data: "x"},
		"comment with children": {Kind: elem.KindComment, Children: func(func(elem.Node) bool) {}},
		"unknown kind":          {Kind: elem.Kind(42)},
	}
	for name, obj := range invalid {
		t.Run("invalid "+name, func(t *testing.T) {
			if _, err := obj.Validate(); !errors.Is(err, elem.ErrInvalidObj) {
				t.Fatalf("expected ErrInvalidObj, got %v", err)
			}
		})
	}
}
