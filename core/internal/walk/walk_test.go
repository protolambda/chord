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
	if el.Literal {
		b.WriteString(" literal")
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
func (r *recorder) Doctype(v string) error { return r.record("doctype " + v) }
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

func TestWalkRejectsNilNodes(t *testing.T) {
	var nilFn elem.Fn
	var nilScope elem.Scope
	var nilAttrFn attr.Fn
	tests := map[string]struct {
		node elem.Node
		loc  string
		want error
	}{
		"nil root":  {nil, "at [0]", elem.ErrNilNode},
		"nil child": {elem.Name("div").New()(elem.Text("a"), nil), "at div[0]/[1]", elem.ErrNilNode},
		"nil in bundle": {
			elem.Name("ul").New()(elem.Bundle{elem.Name("li").New(), nil}),
			"at ul[0]/[1]", elem.ErrNilNode,
		},
		"fn returns nil": {
			elem.Name("p").New()(elem.Fn(func(context.Context) (elem.Node, error) { return nil, nil })),
			"at p[0]/[0]", elem.ErrNilNode,
		},
		"nil fn":    {elem.Name("p").New()(nilFn), "at p[0]/[0]", elem.ErrNilNode},
		"nil scope": {elem.Name("p").New()(nilScope), "at p[0]/[0]", elem.ErrNilNode},
		"nil attribute": {
			elem.Name("div").New()(elem.Name("input").Void(attr.ID("x"), nil)),
			"at div[0]/input[0]: attribute", attr.ErrNilNode,
		},
		"attribute fn returns nil": {
			elem.Name("input").Void(attr.Fn(func(context.Context) (attr.Node, error) { return nil, nil })),
			"at input[0]: attribute", attr.ErrNilNode,
		},
		"nil attribute fn": {elem.Name("input").Void(nilAttrFn), "at input[0]: attribute", attr.ErrNilNode},
		"nil in attribute bundle": {
			elem.Name("input").Void(attr.Bundle{attr.Class("a"), nil}),
			"at input[0]: attribute", attr.ErrNilNode,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			var r recorder
			err := walk.Walk(context.Background(), tc.node, &r)
			if !errors.Is(err, tc.want) {
				t.Fatalf("expected %v, got %v", tc.want, err)
			}
			if !strings.Contains(err.Error(), tc.loc) {
				t.Fatalf("error %q does not contain %q", err, tc.loc)
			}
		})
	}
}

func TestWalkTreatsNilSequencesAsEmpty(t *testing.T) {
	node := elem.Name("div").New(attr.Seq(nil), attr.Obj{Kind: attr.KindBundle})(
		elem.Seq(nil),
		elem.Obj{Kind: elem.KindFragment},
	)
	expectEvents(t, events(t, node), []string{`open div`, `close`})
}

type scopeKey struct{}

func scopeText() elem.Node {
	return elem.Fn(func(ctx context.Context) (elem.Node, error) {
		v, _ := ctx.Value(scopeKey{}).(string)
		return elem.Text(v), nil
	})
}

func scopeAttr() attr.Node {
	return attr.Fn(func(ctx context.Context) (attr.Node, error) {
		v, _ := ctx.Value(scopeKey{}).(string)
		return attr.ID(v), nil
	})
}

func withScope(ctx context.Context, v string) context.Context {
	return context.WithValue(ctx, scopeKey{}, v)
}

// scoped is a node that returns obj with a Context derived from the context
// of its evaluation, as a custom node that scopes a context does.
type scoped struct {
	derive func(context.Context) context.Context
	obj    elem.Obj
}

func (n scoped) Eval(ctx context.Context) (elem.Obj, error) {
	obj := n.obj
	obj.Context = n.derive(ctx)
	return obj, nil
}

func TestWalkScopesObjectContext(t *testing.T) {
	toScoped := func(ctx context.Context) context.Context { return withScope(ctx, "scoped") }
	node := elem.Bundle{
		scoped{derive: toScoped, obj: elem.Obj{Kind: elem.KindFragment, Children: slices.Values([]elem.Node{
			scopeText(),
			elem.Name("b").New(scopeAttr())(scopeText()),
		})}},
		scopeText(),
		scoped{derive: toScoped, obj: elem.Obj{Kind: elem.KindElement, Tag: "i",
			Attribs:  attr.Seq(slices.Values([]attr.Node{scopeAttr()})),
			Children: slices.Values([]elem.Node{scopeText()}),
		}},
		scopeText(),
	}
	var r recorder
	if err := walk.Walk(withScope(context.Background(), "root"), node, &r); err != nil {
		t.Fatalf("walk: %v", err)
	}
	expectEvents(t, r.events, []string{
		`text scoped`,
		`open b id="scoped"`, `text scoped`, `close`,
		`text root`,
		`open i id="scoped"`, `text scoped`, `close`,
		`text root`,
	})
}

func TestWalkUsesParentContextForLiteralsWithoutContext(t *testing.T) {
	literal := elem.Obj{Kind: elem.KindElement, Tag: "p", Children: slices.Values([]elem.Node{scopeText()})}
	var r recorder
	if err := walk.Walk(withScope(context.Background(), "parent"), literal, &r); err != nil {
		t.Fatalf("walk: %v", err)
	}
	expectEvents(t, r.events, []string{`open p`, `text parent`, `close`})
}

// A Context that does not derive from the context of the walk was recorded
// outside of it: by an evaluation before the walk, or in another walk. The
// walker uses the parent's context instead, so that a retained object does
// not carry values or a canceled context into another render.
func TestWalkIgnoresContextFromOutsideTheWalk(t *testing.T) {
	canceled, cancel := context.WithCancel(withScope(context.Background(), "canceled"))
	cancel()
	retained, err := elem.Name("i").New(scopeAttr())(scopeText()).Eval(withScope(context.Background(), "retained"))
	if err != nil {
		t.Fatalf("eval: %v", err)
	}
	var inOtherWalk elem.Obj
	capture := elem.Fn(func(ctx context.Context) (elem.Node, error) {
		obj, err := elem.Name("u").New()(scopeText()).Eval(withScope(ctx, "other walk"))
		inOtherWalk = obj
		return obj, err
	})
	if err := walk.Walk(context.Background(), capture, &recorder{}); err != nil {
		t.Fatalf("walk: %v", err)
	}
	node := elem.Bundle{
		elem.Obj{Kind: elem.KindFragment, Context: withScope(context.Background(), "unrelated"),
			Children: slices.Values([]elem.Node{scopeText()})},
		elem.Obj{Kind: elem.KindElement, Tag: "b", Context: canceled,
			Children: slices.Values([]elem.Node{scopeText()})},
		retained,
		inOtherWalk,
	}
	var r recorder
	if err := walk.Walk(withScope(context.Background(), "walk"), node, &r); err != nil {
		t.Fatalf("walk: %v", err)
	}
	expectEvents(t, r.events, []string{
		`text walk`,
		`open b`, `text walk`, `close`,
		`open i id="walk"`, `text walk`, `close`,
		`open u`, `text walk`, `close`,
	})
}

func TestWalkChecksRootContextBelowDetachedContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	node := scoped{derive: context.WithoutCancel, obj: elem.Obj{Kind: elem.KindFragment, Children: slices.Values([]elem.Node{
		elem.Name("div").New(),
	})}}
	var r recorder
	if err := walk.Walk(ctx, node, &r); !errors.Is(err, context.Canceled) {
		t.Fatalf("expected context.Canceled, got %v", err)
	}
}

// liveScopeText reports the scope value of its context, and whether the
// context is canceled.
func liveScopeText() elem.Node {
	return elem.Fn(func(ctx context.Context) (elem.Node, error) {
		v, _ := ctx.Value(scopeKey{}).(string)
		if err := ctx.Err(); err != nil {
			return elem.Text(v + " " + err.Error()), nil
		}
		return elem.Text(v + " live"), nil
	})
}

// A scope context that is done when the walk enters it was canceled by its
// node when Eval returned: its values scope the subtree, with the
// cancellation of the enclosing scope instead of its own.
func TestWalkUsesValuesOfCanceledScopedContext(t *testing.T) {
	canceled := func(ctx context.Context) context.Context {
		scoped, cancel := context.WithCancel(withScope(ctx, "canceled"))
		cancel()
		return scoped
	}
	node := elem.Name("div").New()(
		scoped{derive: canceled, obj: elem.Obj{Kind: elem.KindFragment, Children: slices.Values([]elem.Node{
			elem.Name("span").New(scopeAttr())(liveScopeText()),
			scoped{derive: canceled, obj: elem.Obj{Kind: elem.KindElement, Tag: "b",
				Children: slices.Values([]elem.Node{liveScopeText()})}},
		})}},
		liveScopeText(),
	)
	var r recorder
	if err := walk.Walk(withScope(context.Background(), "root"), node, &r); err != nil {
		t.Fatalf("walk: %v", err)
	}
	expectEvents(t, r.events, []string{
		`open div`,
		`open span id="canceled"`, `text canceled live`, `close`,
		`open b`, `text canceled live`, `close`,
		`text root live`,
		`close`,
	})
}

// A scope context that ends during the walk does not stop it: nodes below it
// see the canceled context and decide, static content is still delivered.
func TestWalkContinuesAfterScopedContextEnds(t *testing.T) {
	var cancel context.CancelFunc
	derive := func(ctx context.Context) context.Context {
		var scoped context.Context
		scoped, cancel = context.WithCancel(withScope(ctx, "scoped"))
		return scoped
	}
	node := scoped{derive: derive, obj: elem.Obj{Kind: elem.KindFragment, Children: slices.Values([]elem.Node{
		liveScopeText(),
		elem.Fn(func(context.Context) (elem.Node, error) {
			cancel()
			return elem.Noop(), nil
		}),
		elem.Name("p").New()(liveScopeText()),
	})}}
	var r recorder
	if err := walk.Walk(context.Background(), node, &r); err != nil {
		t.Fatalf("walk: %v", err)
	}
	expectEvents(t, r.events, []string{
		`text scoped live`,
		`open p`, `text scoped context canceled`, `close`,
	})
}

func TestWalkRefusesTextThatEndsScriptOrStyle(t *testing.T) {
	script := elem.Name("script").New()
	style := elem.Name("style").New()
	tests := map[string]struct {
		node elem.Node
		loc  string
	}{
		"script end tag":        {script(elem.Text(`var s = "</script><img src=x onerror=alert(1)>";`)), "at script[0]/[0]"},
		"script end tag case":   {script(elem.Text(`"</ScRiPt "`)), "at script[0]/[0]"},
		"script comment open":   {script(elem.Text(`x <!-- y`)), "at script[0]/[0]"},
		"split over text nodes": {script(elem.Text("a</scr"), elem.Bundle{elem.Noop(), elem.Text("ipt>")}), "at script[0]/[1]"},
		"split comment open":    {script(elem.Text("<!-"), elem.Text("-")), "at script[0]/[1]"},
		"completes raw content": {script(elem.Raw("</scr"), elem.Text("ipt")), "at script[0]/[1]"},
		"style end tag":         {style(elem.Text(`a::after { content: "</style>"; }`)), "at style[0]/[0]"},
		"uppercase style":       {elem.Name("STYLE").New()(elem.Text("</STYLE")), "at STYLE[0]/[0]"},
		"container end tag": {
			elem.Name("noscript").New()(style(elem.Text("</NOSCRIPT><img>"))),
			"at noscript[0]/style[0]/[0]",
		},
		"outer container end tag": {
			elem.Name("textarea").New()(elem.Name("div").New()(elem.Name("title").New()(script(elem.Text("</textarea"))))),
			"at textarea[0]/div[0]/title[0]/script[0]/[0]",
		},
		"comment open in a script container": {script(style(elem.Text("<!--"))), "at script[0]/style[0]/[0]"},
		"container end tag split over text nodes": {
			elem.Name("noframes").New()(style(elem.Text("</nofr"), elem.Text("ames"))),
			"at noframes[0]/style[0]/[1]",
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			var r recorder
			err := walk.Walk(context.Background(), tc.node, &r)
			if !errors.Is(err, elem.ErrUnsafeText) {
				t.Fatalf("expected ErrUnsafeText, got %v", err)
			}
			if !strings.Contains(err.Error(), tc.loc) {
				t.Fatalf("error %q does not contain %q", err, tc.loc)
			}
		})
	}
}

func TestWalkAcceptsSafeScriptAndStyleText(t *testing.T) {
	tests := map[string]elem.Node{
		"style end tag in script": elem.Name("script").New()(elem.Text(`"</style>"`)),
		"escaped end tag":         elem.Name("script").New()(elem.Text(`"<\/script>" + "\x3C!--"`)),
		"script end tag in style": elem.Name("style").New()(elem.Text(`/* </script> <!-- */`)),
		"separate elements": elem.Bundle{
			elem.Name("script").New()(elem.Text("</scr")),
			elem.Name("script").New()(elem.Text("ipt>")),
		},
		"separated by an element": elem.Name("script").New()(elem.Text("</scr"), elem.Name("b").New(), elem.Text("ipt>")),
		"text in a child element": elem.Name("script").New()(elem.Name("b").New()(elem.Text("</script>"))),
		"script in svg":           elem.Name("svg").New()(elem.Name("script").New()(elem.Text("</script>"))),
		"style in math":           elem.Name("math").New()(elem.Name("style").New()(elem.Text("</style>"))),
		"raw content":             elem.Name("script").New()(elem.Raw("<!-- trusted -->")),
		"other end tag in a container": elem.Name("noscript").New()(
			elem.Name("style").New()(elem.Text("</script> </noscrip <!--")),
		),
		"container end tag outside the container": elem.Bundle{
			elem.Name("noscript").New(),
			elem.Name("style").New()(elem.Text("</noscript>")),
		},
		"text in plaintext": elem.Name("plaintext").New()(elem.Name("script").New()(elem.Text("</plaintext>"))),
		"style in select":   elem.Name("select").New()(elem.Name("style").New()(elem.Text("</select><!--"))),
		"after a frameset": elem.Bundle{
			elem.Name("frameset").New(),
			elem.Name("script").New()(elem.Text("</script><!--")),
		},
	}
	for name, node := range tests {
		t.Run(name, func(t *testing.T) {
			var r recorder
			if err := walk.Walk(context.Background(), node, &r); err != nil {
				t.Fatalf("walk: %v", err)
			}
		})
	}
}

func TestWalkMarksLiteralElements(t *testing.T) {
	node := elem.Bundle{
		elem.Name("script").New()(elem.Name("b").New()),
		elem.Name("Style").New(),
		elem.Name("svg").New()(elem.Name("script").New(), elem.Name("foreignObject").New()(elem.Name("style").New())),
		elem.Name("script").Void(),
		elem.Name("\u017fcript").New(), // long s: not a script element for browsers
		elem.Name("noscript").New()(elem.Name("style").New()),
		elem.Name("SELECT").New()(elem.Name("option").New()(elem.Name("style").New(), elem.Name("script").New())),
		elem.Name("style").New(),
		elem.Name("div").New()(elem.Name("frameset").New()(elem.Name("script").New())),
		elem.Name("script").New(),
	}
	expectEvents(t, events(t, node), []string{
		`open script literal`, `open b`, `close`, `close`,
		`open Style literal`, `close`,
		`open svg`, `open script`, `close`, `open foreignObject`, `open style`, `close`, `close`, `close`,
		`open script void`, `close`,
		"open \u017fcript", `close`,
		`open noscript`, `open style literal`, `close`, `close`,
		`open SELECT`, `open option`, `open style`, `close`, `open script literal`, `close`, `close`, `close`,
		`open style literal`, `close`,
		`open div`, `open frameset`, `open script`, `close`, `close`, `close`,
		`open script`, `close`,
	})
}

func TestWalkDoctype(t *testing.T) {
	node := elem.Bundle{elem.Doctype(), elem.Name("html").New()}
	expectEvents(t, events(t, node), []string{`doctype html`, `open html`, `close`})

	var r recorder
	err := walk.Walk(context.Background(), elem.Name("html").New()(elem.Text("a"), elem.Doctype()), &r)
	if err == nil || !strings.Contains(err.Error(), `at html[0]/[1]: doctype inside element "html"`) {
		t.Fatalf("expected a located doctype error, got %v", err)
	}
	err = walk.Walk(context.Background(), elem.Obj{Kind: elem.KindDoctype, Data: "xhtml"}, &r)
	if !errors.Is(err, elem.ErrInvalidObj) {
		t.Fatalf("expected ErrInvalidObj, got %v", err)
	}
}
