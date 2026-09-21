package inspect_test

import (
	"context"
	"errors"
	"slices"
	"strings"
	"testing"

	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/core/attr"
	"github.com/protolambda/chord/core/elem"
	"github.com/protolambda/chord/core/inspect"
)

func page(fn elem.Node) elem.Node {
	return elem.Name("html").New()(
		elem.Name("body").New(attr.Class("a"), attr.Class("b"))(
			elem.Name("main").New(attr.ID("content"))(
				elem.Name("h1").New()(elem.Text("Hello "), elem.Name("em").New()(elem.Text("world"))),
				elem.Comment("note"),
				elem.Raw("<hr>"),
				elem.Name("input").Void(attr.Name("type").Raw("text"), attr.Name("required").Bool()),
				fn,
			),
		),
	)
}

func build(t *testing.T, node elem.Node) *inspect.Document {
	t.Helper()
	doc, err := inspect.Build(context.Background(), node)
	if err != nil {
		t.Fatalf("build: %v", err)
	}
	return doc
}

func TestBuildStructure(t *testing.T) {
	doc := build(t, page(elem.Noop()))
	root := doc.Root()
	if root.Kind() != elem.KindFragment || root.Parent() != nil {
		t.Fatalf("unexpected root: kind %s", root.Kind())
	}

	var kinds []string
	for n := range root.Descendants() {
		kinds = append(kinds, n.Kind().String()+":"+n.Tag()+n.Data())
	}
	want := []string{
		"element:html", "element:body", "element:main",
		"element:h1", "text:Hello ", "element:em", "text:world",
		"comment:note", "raw:<hr>", "element:input",
	}
	if !slices.Equal(kinds, want) {
		t.Fatalf("descendants mismatch:\n got: %q\nwant: %q", kinds, want)
	}

	body := slices.Collect(root.Descendants())[1]
	class, ok := body.Attr("class")
	if !ok || class.Val != "a b" || class.Kind != attr.KindValue {
		t.Fatalf("unexpected class attribute: %+v (found %t)", class, ok)
	}
	input := slices.Collect(root.Descendants())[9]
	if !input.Void() {
		t.Fatal("expected void input")
	}
	if typ, _ := input.Attr("type"); typ.Kind != attr.KindRawValue || typ.Val != "text" {
		t.Fatalf("unexpected type attribute: %+v", typ)
	}
	if req, _ := input.Attr("required"); req.Kind != attr.KindBool {
		t.Fatalf("unexpected required attribute: %+v", req)
	}
	if _, ok := input.Attr("missing"); ok {
		t.Fatal("unexpected missing attribute")
	}
}

func TestNodeLocationsAndText(t *testing.T) {
	doc := build(t, page(elem.Noop()))
	nodes := slices.Collect(doc.Root().Descendants())
	tests := map[int]string{
		0: "html[0]",
		3: "html[0]/body[0]/main#content[0]/h1[0]",
		6: "html[0]/body[0]/main#content[0]/h1[0]/em[1]/[0]",
		7: "html[0]/body[0]/main#content[0]/[1]",
	}
	for i, want := range tests {
		if got := nodes[i].Location().String(); got != want {
			t.Errorf("node %d location: got %q, want %q", i, got, want)
		}
	}
	if got := doc.Root().Location().String(); got != "" {
		t.Errorf("root location: got %q, want empty", got)
	}
	if got, want := nodes[3].TextContent(), "Hello world"; got != want {
		t.Errorf("h1 text: got %q, want %q", got, want)
	}
	if got, want := nodes[2].TextContent(), "Hello world"; got != want {
		t.Errorf("main text ignores raw and comments: got %q, want %q", got, want)
	}
	if steps := nodes[3].Location().Steps(); len(steps) != 4 || steps[2].ID != "content" || steps[3].Tag != "h1" {
		t.Errorf("unexpected steps: %+v", steps)
	}
}

func TestBuildRunsApplicationCodeOnceAndReplays(t *testing.T) {
	calls := 0
	fn := elem.Fn(func(context.Context) (elem.Node, error) {
		calls++
		return elem.Name("p").New(attr.Data("n", "1"))(elem.Text("dynamic <x>")), nil
	})
	original := page(fn)
	doc := build(t, original)
	if calls != 1 {
		t.Fatalf("fn called %d times during build", calls)
	}

	for range 2 {
		if n := len(slices.Collect(doc.Root().Descendants())); n != 12 {
			t.Fatalf("expected 12 descendants on each iteration, got %d", n)
		}
	}

	var want, got strings.Builder
	if err := core.Render(context.Background(), page(fn), &want, core.WithIndent()); err != nil {
		t.Fatalf("render original: %v", err)
	}
	if err := core.Render(context.Background(), doc.Root(), &got, core.WithIndent()); err != nil {
		t.Fatalf("render snapshot: %v", err)
	}
	if got.String() != want.String() {
		t.Fatalf("replay mismatch:\n got: %s\nwant: %s", got.String(), want.String())
	}
	if calls != 2 {
		t.Fatalf("replay must not re-run application code: fn called %d times", calls)
	}
}

func TestBuildFailureReturnsNoDocument(t *testing.T) {
	errBoom := errors.New("boom")
	fn := elem.Fn(func(context.Context) (elem.Node, error) { return nil, errBoom })
	doc, err := inspect.Build(context.Background(), page(fn))
	if doc != nil {
		t.Fatal("expected no document")
	}
	if !errors.Is(err, errBoom) {
		t.Fatalf("expected cause in chain, got %v", err)
	}
	if want := "inspect: at html[0]/body[0]/main#content[0]/[4]: boom"; err.Error() != want {
		t.Fatalf("error:\n got: %s\nwant: %s", err, want)
	}
}
