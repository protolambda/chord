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
