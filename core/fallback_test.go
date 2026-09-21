package core_test

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/core/attr"
	"github.com/protolambda/chord/core/elem"
	"github.com/protolambda/chord/html/text"
)

// Fallback() catches evaluation errors and substitutes alternative content.
// Useful for graceful degradation of page sections.
func Example_fallback() {
	unreliable := elem.Fn(func(ctx context.Context) (elem.Node, error) {
		return nil, errors.New("database unavailable")
	})

	core.Dump(core.Fallback(
		elem.Name("div").New()(unreliable),
		func(ctx context.Context, err error) elem.Node {
			return elem.Name("div").New(attr.KV("class", "error"))(
				text.Text("Failed to load content"))
		},
	))
	// Output: <div class="error">Failed to load content</div>
}

func TestFallbackRendersPrimaryOnce(t *testing.T) {
	calls := 0
	primary := elem.Fn(func(context.Context) (elem.Node, error) {
		calls++
		return elem.Name("p").New()(elem.Text("primary <ok>")), nil
	})
	fallbackCalled := false
	node := core.Fallback(primary, func(context.Context, error) elem.Node {
		fallbackCalled = true
		return elem.Text("fallback")
	})

	if got, want := render(t, elem.Name("div").New()(node)), "<div><p>primary &lt;ok&gt;</p></div>"; got != want {
		t.Fatalf("rendered output mismatch:\n got: %s\nwant: %s", got, want)
	}
	if calls != 1 || fallbackCalled {
		t.Fatalf("primary called %d times, fallback called %t", calls, fallbackCalled)
	}
}

func TestFallbackReceivesLocatedError(t *testing.T) {
	errBoom := errors.New("boom")
	var seen error
	node := core.Fallback(
		elem.Name("section").New(attr.ID("news"))(
			elem.Text("ok"),
			elem.Fn(func(context.Context) (elem.Node, error) { return nil, errBoom }),
		),
		func(_ context.Context, err error) elem.Node {
			seen = err
			return elem.Text("fallback")
		},
	)

	if got, want := render(t, node), "fallback"; got != want {
		t.Fatalf("rendered output mismatch:\n got: %s\nwant: %s", got, want)
	}
	if !errors.Is(seen, errBoom) {
		t.Fatalf("expected cause in chain, got %v", seen)
	}
	if !strings.Contains(seen.Error(), "at section#news[0]/[1]") {
		t.Fatalf("expected location in error, got %v", seen)
	}
}

func TestFallbackIgnoresOutputFailures(t *testing.T) {
	fallbackCalled := false
	node := core.Fallback(elem.Name("div").New(), func(context.Context, error) elem.Node {
		fallbackCalled = true
		return elem.Text("fallback")
	})

	err := core.Render(context.Background(), node, failingStringWriter{})
	if !errors.Is(err, core.ErrRenderOutput) {
		t.Fatalf("expected ErrRenderOutput, got %v", err)
	}
	if fallbackCalled {
		t.Fatal("fallback must not run for output failures")
	}
}

func TestFallbackContentErrorsPropagate(t *testing.T) {
	errAlt := errors.New("alternative failed too")
	node := core.Fallback(
		elem.Fn(func(context.Context) (elem.Node, error) { return nil, errors.New("primary failed") }),
		func(context.Context, error) elem.Node {
			return elem.Name("div").New()(elem.Fn(func(context.Context) (elem.Node, error) { return nil, errAlt }))
		},
	)
	var out strings.Builder
	err := core.Render(context.Background(), node, &out)
	if !errors.Is(err, errAlt) {
		t.Fatalf("expected fallback content error, got %v", err)
	}
}
