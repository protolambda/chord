package core_test

import (
	"context"
	"errors"
	"io"
	"slices"
	"strings"
	"testing"
	"time"

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
	want := "<div>\n  <pre>line 1\n  line 2<b>bold</b></pre>\n  <p>after</p>\n</div>\n"
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

func TestRenderRejectsNilNodes(t *testing.T) {
	tests := map[string]elem.Node{
		"nil child": elem.Name("div").New()(nil),
		"fn returns nil": elem.Name("div").New()(elem.Fn(func(context.Context) (elem.Node, error) {
			return nil, nil
		})),
	}
	for name, node := range tests {
		t.Run(name, func(t *testing.T) {
			var out strings.Builder
			err := core.Render(context.Background(), node, &out)
			if !errors.Is(err, elem.ErrNilNode) {
				t.Fatalf("expected ErrNilNode, got %v", err)
			}
			if want := "render: at div[0]/[0]: nil element node"; !strings.HasPrefix(err.Error(), want) {
				t.Fatalf("error message:\n got: %s\nwant prefix: %s", err, want)
			}
		})
	}
}

type userKey struct{}

// withUser is a custom node that evaluates its content with a user in the context.
type withUser struct {
	user  string
	inner elem.Node
}

func (n withUser) Eval(ctx context.Context) (elem.Obj, error) {
	return n.inner.Eval(context.WithValue(ctx, userKey{}, n.user))
}

func userText() elem.Node {
	return elem.Fn(func(ctx context.Context) (elem.Node, error) {
		user, _ := ctx.Value(userKey{}).(string)
		return elem.Text("user=" + user), nil
	})
}

func userAttr() attr.Node {
	return attr.Fn(func(ctx context.Context) (attr.Node, error) {
		user, _ := ctx.Value(userKey{}).(string)
		return attr.Data("user", user), nil
	})
}

func TestRenderEvaluatesChildrenWithOverriddenContext(t *testing.T) {
	node := elem.Bundle{
		withUser{user: "alice", inner: elem.Name("div").New(userAttr())(
			elem.Name("p").New()(userText()),
			withUser{user: "bob", inner: elem.Bundle{userText(), elem.Text(";")}},
			userText(),
		)},
		userText(),
	}
	want := `<div data-user="alice"><p>user=alice</p>user=bob;user=alice</div>user=`
	if got := render(t, node); got != want {
		t.Fatalf("rendered output mismatch:\n got: %s\nwant: %s", got, want)
	}
}

// timedUser is a custom node that derives a context with a timeout for its
// own work and cancels it when Eval returns, the usual Go pattern.
type timedUser struct {
	user  string
	inner elem.Node
}

func (n timedUser) Eval(ctx context.Context) (elem.Obj, error) {
	ctx, cancel := context.WithTimeout(context.WithValue(ctx, userKey{}, n.user), time.Minute)
	defer cancel()
	return n.inner.Eval(ctx)
}

// liveText reports whether the context it is evaluated with is canceled.
func liveText() elem.Node {
	return elem.Fn(func(ctx context.Context) (elem.Node, error) {
		if err := ctx.Err(); err != nil {
			return elem.Text("err=" + err.Error()), nil
		}
		return elem.Text("live"), nil
	})
}

// A node that cancels the context it derived when its Eval returns scopes the
// values of that context to its subtree, but not the cancellation: the
// subtree is rendered, and evaluated with a live context.
func TestRenderContextCanceledOnReturnKeepsValues(t *testing.T) {
	node := elem.Name("div").New()(
		timedUser{user: "alice", inner: elem.Name("p").New(userAttr())(
			userText(), elem.Text(";"), liveText(),
			timedUser{user: "bob", inner: elem.Name("b").New()(userText(), elem.Text(";"), liveText())},
		)},
		userText(),
	)
	want := `<div><p data-user="alice">user=alice;live<b>user=bob;live</b></p>user=</div>`
	if got := render(t, node); got != want {
		t.Fatalf("rendered output mismatch:\n got: %s\nwant: %s", got, want)
	}
	fallback := core.Fallback(node, func(_ context.Context, err error) elem.Node {
		return elem.Text("fallback: " + err.Error())
	})
	if got := render(t, fallback); got != want {
		t.Fatalf("rendered output mismatch:\n got: %s\nwant: %s", got, want)
	}
}

// Canceling the render still reaches the evaluation of a subtree whose scope
// context was canceled on return, and stops the walk.
func TestRenderCancelReachesSubtreeOfContextCanceledOnReturn(t *testing.T) {
	ctx, cancel := context.WithCancelCause(context.Background())
	errStop := errors.New("client went away")
	var seen []error
	node := timedUser{user: "alice", inner: elem.Name("div").New()(
		elem.Fn(func(ctx context.Context) (elem.Node, error) {
			derived, stop := context.WithCancel(ctx)
			defer stop()
			seen = append(seen, ctx.Err())
			cancel(errStop)
			seen = append(seen, ctx.Err(), context.Cause(ctx), context.Cause(derived))
			<-ctx.Done()
			return elem.Text("x"), nil
		}),
		elem.Name("p").New(),
	)}
	var out strings.Builder
	err := core.Render(ctx, node, &out)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("expected context.Canceled, got %v", err)
	}
	if want := []error{nil, context.Canceled, errStop, errStop}; !slices.Equal(seen, want) {
		t.Fatalf("context seen by the child: got %v, want %v", seen, want)
	}
}

// An evaluated object that is retained and rendered again gets the context of
// the new render, not the one it was evaluated with.
func TestRenderRetainedObjectUsesRenderContext(t *testing.T) {
	withUserCtx := func(ctx context.Context, user string) context.Context {
		return context.WithValue(ctx, userKey{}, user)
	}
	evaluated, cancel := context.WithCancel(withUserCtx(context.Background(), "carol"))
	pre, err := elem.Name("div").New(userAttr())(userText()).Eval(evaluated)
	if err != nil {
		t.Fatalf("eval: %v", err)
	}
	cancel()
	for _, user := range []string{"dave", "erin"} {
		var out strings.Builder
		if err := core.Render(withUserCtx(context.Background(), user), pre, &out); err != nil {
			t.Fatalf("render: %v", err)
		}
		if want := `<div data-user="` + user + `">user=` + user + `</div>`; out.String() != want {
			t.Fatalf("rendered output mismatch:\n got: %s\nwant: %s", out.String(), want)
		}
	}
	// Evaluated within one render, and rendered in another.
	var inner elem.Obj
	capture := elem.Fn(func(ctx context.Context) (elem.Node, error) {
		obj, err := withUser{user: "frank", inner: elem.Name("p").New()(userText())}.Eval(ctx)
		inner = obj
		return obj, err
	})
	if got := render(t, capture); got != `<p>user=frank</p>` {
		t.Fatalf("rendered output mismatch: %s", got)
	}
	var out strings.Builder
	if err := core.Render(withUserCtx(context.Background(), "gina"), inner, &out); err != nil {
		t.Fatalf("render: %v", err)
	}
	if want := `<p>user=gina</p>`; out.String() != want {
		t.Fatalf("rendered output mismatch:\n got: %s\nwant: %s", out.String(), want)
	}
}

func TestRenderWritesScriptAndStyleTextLiterally(t *testing.T) {
	node := elem.Bundle{
		elem.Name("script").New()(elem.Text(`if (a < b && c) { x = 'y'; }`), elem.Text(` "z"`)),
		elem.Name("STYLE").New()(elem.Text(`a > b { content: "&"; }`)),
		elem.Name("p").New()(elem.Text(`a < b`)),
	}
	want := `<script>if (a < b && c) { x = 'y'; } "z"</script>` +
		`<STYLE>a > b { content: "&"; }</STYLE>` +
		`<p>a &lt; b</p>`
	if got := render(t, node); got != want {
		t.Fatalf("rendered output mismatch:\n got: %s\nwant: %s", got, want)
	}
}

func TestRenderEscapesScriptTextOutsideHTML(t *testing.T) {
	node := elem.Bundle{
		elem.Name("svg").New()(elem.Name("script").New()(elem.Text("a < b"))),
		elem.Name("script").New()(elem.Name("template").New()(elem.Text("<b>"))),
	}
	want := `<svg><script>a &lt; b</script></svg><script><template>&lt;b&gt;</template></script>`
	if got := render(t, node); got != want {
		t.Fatalf("rendered output mismatch:\n got: %s\nwant: %s", got, want)
	}
}

func TestRenderRefusesTextThatEndsScript(t *testing.T) {
	var out strings.Builder
	node := elem.Name("script").New()(elem.Text(`let s = "</script><script>alert(1)</script>";`))
	err := core.Render(context.Background(), node, &out)
	if !errors.Is(err, elem.ErrUnsafeText) {
		t.Fatalf("expected ErrUnsafeText, got %v", err)
	}
	if want := `render: at script[0]/[0]: unsafe script or style text: script text contains "</script"`; err.Error() != want {
		t.Fatalf("error message:\n got: %s\nwant: %s", err, want)
	}
}

func renderIndented(t *testing.T, node elem.Node) string {
	t.Helper()
	var out strings.Builder
	if err := core.Render(context.Background(), node, &out, core.WithIndent()); err != nil {
		t.Fatalf("render: %v", err)
	}
	return out.String()
}

func TestRenderIndentKeepsTextareaScriptAndStyleContent(t *testing.T) {
	node := elem.Name("form").New()(
		elem.Name("textarea").New()(elem.Text("line 1\n  line 2")),
		elem.Name("script").New()(elem.Text("f();\n  g();")),
		elem.Name("style").New()(elem.Text("p { x: y; }")),
	)
	want := "<form><textarea>line 1\n  line 2</textarea><script>f();\n  g();</script><style>p { x: y; }</style></form>\n"
	if got := renderIndented(t, node); got != want {
		t.Fatalf("rendered output mismatch:\n got: %q\nwant: %q", got, want)
	}
}

func TestRenderIndentKeepsInlineContentOnOneLine(t *testing.T) {
	node := elem.Name("div").New()(
		elem.Name("h1").New()(elem.Text("Title")),
		elem.Name("p").New()(
			elem.Text("Read the "),
			elem.Name("a").New(attr.Name("href").Value("/docs"))(elem.Text("docs")),
			elem.Text("."),
		),
		elem.Name("ul").New()(
			elem.Name("li").New()(elem.Name("a").New()(elem.Text("One"))),
			elem.Name("li").New()(elem.Name("span").New()(elem.Text("Two")), elem.Name("br").Void()),
		),
	)
	want := `<div>
  <h1>Title</h1>
  <p>Read the <a href="/docs">docs</a>.</p>
  <ul>
    <li><a>One</a></li>
    <li><span>Two</span><br/></li>
  </ul>
</div>
`
	if got := renderIndented(t, node); got != want {
		t.Fatalf("rendered output mismatch:\n got: %q\nwant: %q", got, want)
	}
}

func TestRenderIndentLayout(t *testing.T) {
	div, p, span := elem.Name("div").New(), elem.Name("p").New(), elem.Name("span").New()
	tests := map[string]struct {
		node elem.Node
		want string
	}{
		"top-level blocks":  {elem.Bundle{div(), div()}, "<div></div>\n<div></div>\n"},
		"top-level inline":  {elem.Bundle{span(), elem.Text("a"), span()}, "<span></span>a<span></span>"},
		"text before block": {elem.Bundle{elem.Text("a"), div()}, "a<div></div>\n"},
		"mixed children":    {div(p(), span(), p()), "<div>\n  <p></p><span></span><p></p>\n</div>\n"},
		"block in inline":   {span(div()), "<span>\n  <div></div>\n</span>"},
		"void block":        {div(p(), elem.Name("hr").Void(), p()), "<div>\n  <p></p>\n  <hr/>\n  <p></p>\n</div>\n"},
		"void inline":       {p(elem.Text("a"), elem.Name("br").Void(), elem.Text("b")), "<p>a<br/>b</p>\n"},
		"uppercase tags":    {elem.Name("DIV").New()(elem.Name("P").New()), "<DIV>\n  <P></P>\n</DIV>\n"},
		"inside pre":        {elem.Name("pre").New()(div(div())), "<pre><div><div></div></div></pre>\n"},
		"comment":           {div(elem.Comment("c"), p()), "<div><!-- c --><p></p>\n</div>\n"},
		"document": {
			elem.Name("html").New()(
				elem.Name("head").New()(
					elem.Name("meta").Void(attr.KV("charset", "utf-8")),
					elem.Name("title").New()(elem.Text("T")),
					elem.Name("script").New()(elem.Text("x()")),
				),
				elem.Name("body").New()(div(elem.Text("a"))),
			),
			`<html>
  <head>
    <meta charset="utf-8"/>
    <title>T</title>
    <script>x()</script>
  </head>
  <body>
    <div>a</div>
  </body>
</html>
`,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			if got := renderIndented(t, tc.node); got != tc.want {
				t.Fatalf("rendered output mismatch:\n got: %q\nwant: %q", got, tc.want)
			}
		})
	}
}

func TestRenderIndentOnlyAddsWhitespaceBetweenTags(t *testing.T) {
	div, p, a := elem.Name("div").New(), elem.Name("p").New(), elem.Name("a").New()
	page := elem.Name("html").New()(
		elem.Name("head").New()(elem.Name("title").New()(elem.Text("T"))),
		elem.Name("body").New()(
			div(elem.Name("h1").New()(elem.Text("Title")), p(elem.Text("x "), a(elem.Text("y")), elem.Text("."))),
			elem.Name("table").New()(elem.Name("tr").New()(elem.Name("td").New()(elem.Text("1")), elem.Name("td").New()(a()))),
			elem.Name("form").New()(elem.Name("label").New()(elem.Text("L")), elem.Name("textarea").New()(elem.Text("v")),
				elem.Name("select").New()(elem.Name("option").New()(elem.Text("o")))),
		),
	)
	compact := render(t, page)
	indented := renderIndented(t, page)
	var stripped strings.Builder
	for i := 0; i < len(indented); i++ {
		if indented[i] == '\n' {
			j := i + 1
			for j < len(indented) && indented[j] == ' ' {
				j++
			}
			if i == 0 || indented[i-1] != '>' || (j < len(indented) && indented[j] != '<') {
				t.Fatalf("whitespace outside of a tag boundary at %d in %q", i, indented)
			}
			i = j - 1
			continue
		}
		stripped.WriteByte(indented[i])
	}
	if stripped.String() != compact {
		t.Fatalf("indented output differs beyond whitespace:\n got: %q\nwant: %q", stripped.String(), compact)
	}
}

func TestRenderDoctype(t *testing.T) {
	page := elem.Bundle{
		elem.Doctype(),
		elem.Name("html").New()(elem.Name("body").New()(elem.Name("p").New()(elem.Text("x")))),
	}
	if got, want := render(t, page), "<!DOCTYPE html><html><body><p>x</p></body></html>"; got != want {
		t.Fatalf("rendered output mismatch:\n got: %q\nwant: %q", got, want)
	}
	want := "<!DOCTYPE html>\n<html>\n  <body>\n    <p>x</p>\n  </body>\n</html>\n"
	if got := renderIndented(t, page); got != want {
		t.Fatalf("rendered output mismatch:\n got: %q\nwant: %q", got, want)
	}
	if got, want := renderIndented(t, elem.Doctype()), "<!DOCTYPE html>\n"; got != want {
		t.Fatalf("rendered output mismatch:\n got: %q\nwant: %q", got, want)
	}
	if got, want := renderIndented(t, elem.Bundle{elem.Text("x"), elem.Doctype()}), "x<!DOCTYPE html>\n"; got != want {
		t.Fatalf("rendered output mismatch:\n got: %q\nwant: %q", got, want)
	}
}
