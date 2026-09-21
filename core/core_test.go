package core_test

import (
	"context"
	"errors"
	"io"
	"slices"
	"strings"
	"testing"

	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/core/attr"
	"github.com/protolambda/chord/core/elem"
)

func render(t testing.TB, node elem.Node) string {
	t.Helper()

	var out strings.Builder
	if err := core.Render(context.Background(), node, &out); err != nil {
		t.Fatalf("render: %v", err)
	}
	return out.String()
}

func TestRenderMergesClassAndStyleAttributes(t *testing.T) {
	node := elem.Name("div").New(
		attr.Class(`one"`),
		attr.Style(`color:"red"`),
		attr.Class("two&"),
		attr.Style("display:block"),
	)

	if got, want := render(t, node), `<div class="one&#34; two&amp;" style="color:&#34;red&#34;;display:block"></div>`; got != want {
		t.Fatalf("rendered output mismatch:\n got: %s\nwant: %s", got, want)
	}
}

func TestRenderEscapesAttributeValues(t *testing.T) {
	node := elem.Name("div").New(attr.ID(`profile" autofocus onfocus="attack()`))

	if got, want := render(t, node), `<div id="profile&#34; autofocus onfocus=&#34;attack()"></div>`; got != want {
		t.Fatalf("rendered output mismatch:\n got: %s\nwant: %s", got, want)
	}
}

func TestRenderPreservesTrustedRawAttributeValue(t *testing.T) {
	node := elem.Name("div").New(attr.Name("data-safe").Raw("one&amp;two"))

	if got, want := render(t, node), `<div data-safe="one&amp;two"></div>`; got != want {
		t.Fatalf("rendered output mismatch:\n got: %s\nwant: %s", got, want)
	}
}

func TestRenderRejectsInvalidDynamicAttributeName(t *testing.T) {
	var out strings.Builder
	err := core.Render(context.Background(), elem.Name("div").New(attr.KV(`id" onclick`, "attack()")), &out)
	if !errors.Is(err, attr.ErrInvalidName) {
		t.Fatalf("expected ErrInvalidName, got %v", err)
	}
}

func TestRenderRejectsDuplicateAttribute(t *testing.T) {
	var out strings.Builder
	err := core.Render(context.Background(), elem.Name("div").New(attr.ID("one"), attr.ID("two")), &out)
	if err == nil || !strings.Contains(err.Error(), `duplicate attribute "id"`) {
		t.Fatalf("expected duplicate id error, got %v", err)
	}
}

func TestRenderPreservesEvaluationErrors(t *testing.T) {
	errEval := errors.New("evaluation failed")
	tests := map[string]elem.Node{
		"root": elem.Fn(func(context.Context) (elem.Node, error) {
			return nil, errEval
		}),
		"child": elem.Name("div").New()(elem.Fn(func(context.Context) (elem.Node, error) {
			return nil, errEval
		})),
		"attribute": elem.Name("div").New(attr.Fn(func(context.Context) (attr.Node, error) {
			return nil, errEval
		})),
	}

	for name, node := range tests {
		t.Run(name, func(t *testing.T) {
			var out strings.Builder
			err := core.Render(context.Background(), node, &out)
			if !errors.Is(err, errEval) {
				t.Fatalf("expected evaluation error in chain, got %v", err)
			}
		})
	}
}

func TestRenderFlattensBundlesSequencesAndNoops(t *testing.T) {
	attrs := attr.Bundle{
		attr.Noop(),
		attr.Seq(slices.Values([]attr.Node{attr.Data("one", "1"), attr.Data("two", "2")})),
	}
	children := elem.Bundle{
		elem.Noop(),
		elem.Raw("a"),
		elem.Seq(slices.Values([]elem.Node{elem.Raw("b"), elem.Noop(), elem.Raw("c")})),
	}
	node := elem.Bundle{elem.Noop(), elem.Name("div").New(attrs)(children), elem.Noop()}

	if got, want := render(t, node), `<div data-one="1" data-two="2">abc</div>`; got != want {
		t.Fatalf("rendered output mismatch:\n got: %s\nwant: %s", got, want)
	}
}

func TestRenderDistinguishesVoidAndEmptyElements(t *testing.T) {
	node := elem.Bundle{elem.Name("input").Void(), elem.Name("div").New()}

	if got, want := render(t, node), `<input/><div></div>`; got != want {
		t.Fatalf("rendered output mismatch:\n got: %s\nwant: %s", got, want)
	}
}

var errWriterUnavailable = errors.New("writer unavailable")

type failingStringWriter struct{}

func (failingStringWriter) WriteString(string) (int, error) {
	return 0, errWriterUnavailable
}

func TestRenderReturnsWriterError(t *testing.T) {
	err := core.Render(context.Background(), elem.Name("div").New(), failingStringWriter{})
	if !errors.Is(err, core.ErrRenderOutput) {
		t.Fatalf("expected ErrRenderOutput, got %v", err)
	}
	if !errors.Is(err, errWriterUnavailable) {
		t.Fatalf("expected underlying writer error, got %v", err)
	}
}

type shortStringWriter struct{}

func (shortStringWriter) WriteString(value string) (int, error) {
	if value == "" {
		return 0, nil
	}
	return len(value) - 1, nil
}

func TestRenderReturnsShortWrite(t *testing.T) {
	err := core.Render(context.Background(), elem.Name("div").New(), shortStringWriter{})
	if !errors.Is(err, core.ErrRenderOutput) {
		t.Fatalf("expected ErrRenderOutput, got %v", err)
	}
	if !errors.Is(err, io.ErrShortWrite) {
		t.Fatalf("expected io.ErrShortWrite, got %v", err)
	}
}

func TestRenderEscapesTextAndCommentsAtOutput(t *testing.T) {
	node := elem.Name("div").New()(
		elem.Text(`<script>alert("x")</script> & more`),
		elem.Comment("ends --> early?"),
		elem.Raw("<b>trusted</b>"),
		elem.Raw(""),
	)

	want := `<div>&lt;script&gt;alert(&#34;x&#34;)&lt;/script&gt; &amp; more<!-- ends --&gt; early? --><b>trusted</b></div>`
	if got := render(t, node); got != want {
		t.Fatalf("rendered output mismatch:\n got: %s\nwant: %s", got, want)
	}
}

func TestRenderMergesMixedRawAndLogicalClasses(t *testing.T) {
	node := elem.Name("div").New(
		attr.Name("class").Raw("btn &amp;"),
		attr.Class("user<class>"),
		attr.Style("a:1"),
		attr.Name("style").Raw("b:&quot;2&quot;"),
	)

	want := `<div class="btn &amp; user&lt;class&gt;" style="a:1;b:&quot;2&quot;"></div>`
	if got := render(t, node); got != want {
		t.Fatalf("rendered output mismatch:\n got: %s\nwant: %s", got, want)
	}
}

func TestRenderRejectsInvalidObjectLiterals(t *testing.T) {
	tests := map[string]elem.Node{
		"element":   elem.Obj{Kind: elem.KindText, Tag: "div"},
		"attribute": elem.Name("div").New(attr.Obj{Key: "id", Val: "x"}),
	}
	for name, node := range tests {
		t.Run(name, func(t *testing.T) {
			var out strings.Builder
			err := core.Render(context.Background(), node, &out)
			if !errors.Is(err, elem.ErrInvalidObj) && !errors.Is(err, attr.ErrInvalidObj) {
				t.Fatalf("expected invalid object error, got %v", err)
			}
		})
	}
}

func TestRenderAcceptsLegacyLiterals(t *testing.T) {
	node := elem.Obj{Children: slices.Values([]elem.Node{
		elem.Obj{Tag: "p", Children: slices.Values([]elem.Node{elem.Text("legacy")})},
		elem.Obj{Tag: "hr", Void: true},
	})}

	if got, want := render(t, node), `<p>legacy</p><hr/>`; got != want {
		t.Fatalf("rendered output mismatch:\n got: %s\nwant: %s", got, want)
	}
}

func TestRenderStopsOnCanceledContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	var out strings.Builder
	err := core.Render(ctx, elem.Name("div").New()(elem.Text("never")), &out)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("expected context.Canceled, got %v", err)
	}
	if out.Len() != 0 {
		t.Fatalf("expected no output, got %q", out.String())
	}
}

func TestRenderIndentKeepsPreformattedContent(t *testing.T) {
	node := elem.Name("div").New()(
		elem.Name("pre").New()(
			elem.Text("line 1\n  line 2"),
			elem.Name("b").New()(elem.Text("bold")),
		),
		elem.Name("p").New()(elem.Text("after")),
	)
	var out strings.Builder
	if err := core.Render(context.Background(), node, &out, core.WithIndent()); err != nil {
		t.Fatalf("render: %v", err)
	}
	want := "<div>\n  <pre>line 1\n  line 2<b>bold</b></pre>\n  <p>\n    after\n  </p>\n</div>\n"
	if got := out.String(); got != want {
		t.Fatalf("rendered output mismatch:\n got: %q\nwant: %q", got, want)
	}
}

func TestRenderLocatesFailures(t *testing.T) {
	errEval := errors.New("evaluation failed")
	node := elem.Name("main").New(attr.ID("page"))(
		elem.Name("ul").New()(
			elem.Name("li").New()(elem.Text("ok")),
			elem.Name("li").New()(elem.Fn(func(context.Context) (elem.Node, error) { return nil, errEval })),
		),
	)
	var out strings.Builder
	err := core.Render(context.Background(), node, &out)
	if !errors.Is(err, errEval) {
		t.Fatalf("expected evaluation error in chain, got %v", err)
	}
	if want := "render: at main#page[0]/ul[0]/li[1]/[0]: evaluation failed"; err.Error() != want {
		t.Fatalf("error message:\n got: %s\nwant: %s", err, want)
	}
}
