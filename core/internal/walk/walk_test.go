package walk_test

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"strings"
	"testing"

	"github.com/protolambda/chord/core/attr"
	"github.com/protolambda/chord/core/elem"
	"github.com/protolambda/chord/core/internal/walk"
)

// recorder records events as compact strings, and can fail on demand.
type recorder struct {
	events []string
	failOn string // event prefix that triggers failErr
	err    error
}

var errReceiver = errors.New("receiver failed")

func (r *recorder) record(event string) error {
	if r.failOn != "" && strings.HasPrefix(event, r.failOn) {
		return errReceiver
	}
	r.events = append(r.events, event)
	return nil
}

func (r *recorder) Open(el walk.Element) error {
	var b strings.Builder
	b.WriteString("open " + el.Tag)
	if el.Void {
		b.WriteString(" void")
	}
	for _, a := range el.Attrs {
		switch a.Kind {
		case attr.KindBool:
			fmt.Fprintf(&b, " %s", a.Key)
		case attr.KindValue:
			fmt.Fprintf(&b, " %s=%q", a.Key, a.Val)
		case attr.KindRawValue:
			fmt.Fprintf(&b, " %s=raw%q", a.Key, a.Val)
		}
	}
	return r.record(b.String())
}

func (r *recorder) Text(v string) error    { return r.record("text " + v) }
func (r *recorder) Raw(v string) error     { return r.record("raw " + v) }
func (r *recorder) Comment(v string) error { return r.record("comment " + v) }
func (r *recorder) Close() error           { return r.record("close") }

func events(t *testing.T, node elem.Node) []string {
	t.Helper()
	var r recorder
	if err := walk.Walk(context.Background(), node, &r); err != nil {
		t.Fatalf("walk: %v", err)
	}
	return r.events
}

func expectEvents(t *testing.T, got, want []string) {
	t.Helper()
	if !slices.Equal(got, want) {
		t.Fatalf("events mismatch:\n got: %q\nwant: %q", got, want)
	}
}

func TestWalkFlattensAndBalances(t *testing.T) {
	node := elem.Bundle{
		elem.Noop(),
		elem.Name("div").New(attr.ID("a"), attr.Noop(), attr.Bundle{attr.Class("x"), attr.Name("hidden").Bool()})(
			elem.Seq(slices.Values([]elem.Node{elem.Text("hi"), elem.Noop()})),
			elem.Name("br").Void(),
			elem.Comment("c"),
			elem.Raw("<b>"),
			elem.Name("span").New(),
		),
	}
	expectEvents(t, events(t, node), []string{
		`open div id="a" class="x" hidden`,
		`text hi`,
		`open br void`,
		`close`,
		`comment c`,
		`raw <b>`,
		`open span`,
		`close`,
		`close`,
	})
}

func TestWalkMergesAttributesByKind(t *testing.T) {
	node := elem.Name("i").New(
		attr.Class("a"), attr.Class("b"),
		attr.Name("style").Raw("x:1"), attr.Style("y:<2>"),
	)
	expectEvents(t, events(t, node), []string{
		`open i class="a b" style=raw"x:1;y:&lt;2&gt;"`,
		`close`,
	})
}

func TestWalkEvaluatesSequencesOnce(t *testing.T) {
	calls := 0
	items := elem.Seq(func(yield func(elem.Node) bool) {
		calls++
		yield(elem.Text("once"))
	})
	fn := 0
	dynamic := elem.Fn(func(context.Context) (elem.Node, error) {
		fn++
		return elem.Text("fn"), nil
	})
	node := elem.Name("ul").New()(items, dynamic)
	expectEvents(t, events(t, node), []string{`open ul`, `text once`, `text fn`, `close`})
	if calls != 1 || fn != 1 {
		t.Fatalf("sequence iterated %d times, fn called %d times", calls, fn)
	}
}

func TestWalkLocatesErrors(t *testing.T) {
	errBoom := errors.New("boom")
	failing := elem.Fn(func(context.Context) (elem.Node, error) { return nil, errBoom })

	tests := map[string]struct {
		node elem.Node
		loc  string
	}{
		"root": {failing, "at [0]"},
		"second root child": {
			elem.Bundle{elem.Text("a"), elem.Name("p").New(), failing},
			"at [2]",
		},
		"nested child": {
			elem.Name("div").New(attr.ID("outer"))(
				elem.Name("section").New()(elem.Text("x"), failing),
			),
			"at div#outer[0]/section[0]/[1]",
		},
		"attribute": {
			elem.Name("div").New()(
				elem.Name("input").Void(attr.Fn(func(context.Context) (attr.Node, error) { return nil, errBoom })),
			),
			"at div[0]/input[0]: attribute",
		},
		"duplicate attribute": {
			elem.Name("form").New(attr.ID("login"), attr.ID("again")),
			`at form[0]: duplicate attribute "id"`,
		},
		"unsafe id is omitted": {
			elem.Name("div").New(attr.ID(`x"y`))(failing),
			"at div[0]/[0]",
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			var r recorder
			err := walk.Walk(context.Background(), tc.node, &r)
			if err == nil {
				t.Fatal("expected error")
			}
			if !strings.Contains(err.Error(), tc.loc) {
				t.Fatalf("error %q does not contain %q", err, tc.loc)
			}
			if !strings.Contains(err.Error(), "duplicate") && !errors.Is(err, errBoom) {
				t.Fatalf("expected cause in chain, got %v", err)
			}
		})
	}
}

func TestWalkPropagatesReceiverErrors(t *testing.T) {
	node := elem.Name("div").New()(elem.Text("a"), elem.Name("br").Void())
	for _, failOn := range []string{"open div", "text", "open br", "close"} {
		t.Run(failOn, func(t *testing.T) {
			r := recorder{failOn: failOn}
			err := walk.Walk(context.Background(), node, &r)
			if !errors.Is(err, errReceiver) {
				t.Fatalf("expected receiver error, got %v", err)
			}
		})
	}
}

func TestWalkChecksContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	var r recorder
	err := walk.Walk(ctx, elem.Name("div").New()(elem.Text("never")), &r)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("expected context.Canceled, got %v", err)
	}
	if len(r.events) != 0 {
		t.Fatalf("expected no events, got %q", r.events)
	}
}

func TestWalkRejectsInvalidObjects(t *testing.T) {
	var r recorder
	err := walk.Walk(context.Background(), elem.Obj{Kind: elem.KindRaw, Tag: "div"}, &r)
	if !errors.Is(err, elem.ErrInvalidObj) {
		t.Fatalf("expected ErrInvalidObj, got %v", err)
	}
	err = walk.Walk(context.Background(), elem.Name("div").New(attr.Obj{Key: "id"}), &r)
	if !errors.Is(err, attr.ErrInvalidObj) {
		t.Fatalf("expected attr.ErrInvalidObj, got %v", err)
	}
}
