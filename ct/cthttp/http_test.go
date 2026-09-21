package cthttp_test

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/protolambda/mustbe/assertion"

	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/core/attr"
	"github.com/protolambda/chord/core/elem"
	"github.com/protolambda/chord/ct"
	"github.com/protolambda/chord/ct/cthttp"
	"github.com/protolambda/chord/html/group/div"
	"github.com/protolambda/chord/html/meta"
	"github.com/protolambda/chord/html/section"
	"github.com/protolambda/chord/html/text"
)

func mustPass(t *testing.T, a assertion.Assertion) {
	t.Helper()
	if err := a.Check(t.Context()); err != nil {
		t.Fatalf("%s: unexpected failure: %v", a, err)
	}
}

func mustFail(t *testing.T, a assertion.Assertion, target error, contains ...string) error {
	t.Helper()
	err := a.Check(t.Context())
	if err == nil {
		t.Fatalf("%s: expected failure", a)
	}
	if !errors.Is(err, target) {
		t.Fatalf("%s: expected %v in chain, got: %v", a, target, err)
	}
	for _, c := range contains {
		if !strings.Contains(err.Error(), c) {
			t.Fatalf("%s: expected %q in error:\n%v", a, c, err)
		}
	}
	return err
}

// accountHandler renders a full page, or only the content fragment for
// HTMX requests, and JSON for the API path.
func accountHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/account", func(w http.ResponseWriter, r *http.Request) {
		content := div.Div(attr.ID("content"))(section.H1()(text.Text("Account")))
		var page elem.Node = meta.HTML()(section.Body()(content))
		if r.Header.Get("HX-Request") == "true" {
			page = content
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Add("Set-Cookie", "a=1")
		w.Header().Add("Set-Cookie", "b=2")
		var out strings.Builder
		if err := core.Render(r.Context(), page, &out); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		_, _ = io.WriteString(w, out.String())
	})
	mux.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{"account":"x"}`)
	})
	return mux
}

func TestServeCapturesResponse(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/account?tab=1", nil)
	res := cthttp.Serve(accountHandler(), req)

	mustPass(t, res)
	mustPass(t, res.Status(http.StatusOK))
	mustPass(t, res.Header("Content-Type", ct.Prefix("text/html")))
	mustPass(t, res.Header("Set-Cookie", ct.Prefix("b=")))
	mustPass(t, res.HeaderAbsent("X-Missing"))
	mustPass(t, res.Body(ct.Contains("<h1>Account</h1>")))

	page := res.HTML()
	mustPass(t, page)
	mustPass(t, page.Find(ct.Tag("body")).Find(ct.Role("heading", ct.Named("Account"))))
	mustPass(t, res.HTML().Find(ct.Tag("html")).Count(1))

	mustFail(t, res.Status(http.StatusNotFound), ct.ErrMismatch, "expected status 404 for GET /account?tab=1, got 200", "body:")
	mustFail(t, res.Header("Content-Type", ct.Exact("text/plain")), ct.ErrMismatch, `got ["text/html; charset=utf-8"]`)
	mustFail(t, res.Header("X-Missing", ct.Exact("x")), ct.ErrMismatch, "header is absent")
	mustFail(t, res.HeaderAbsent("Set-Cookie"), ct.ErrMismatch, `got ["a=1" "b=2"]`)
	mustFail(t, res.Body(ct.Contains("nope")), ct.ErrMismatch, "expected body of GET /account?tab=1 to match contains(\"nope\")")

	if res.StatusCode() != http.StatusOK || res.Headers().Get("Content-Type") == "" || len(res.Bytes()) == 0 {
		t.Fatal("accessors should expose the captured response")
	}
	if got, want := res.String(), "captured(GET /account?tab=1)"; got != want {
		t.Fatalf("String(): got %s, want %s", got, want)
	}
	if got, want := res.Status(200).String(), "status(GET /account?tab=1, 200)"; got != want {
		t.Fatalf("String(): got %s, want %s", got, want)
	}
	if got, want := page.String(), "loads(GET /account?tab=1 html)"; got != want {
		t.Fatalf("String(): got %s, want %s", got, want)
	}
}

func TestHTMLFragmentForPartialResponse(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/account", nil)
	req.Header.Set("HX-Request", "true")
	res := cthttp.Serve(accountHandler(), req)

	fragment := res.HTMLFragment("body")
	mustPass(t, fragment.Find(ct.ID("content")).Matches(ct.HasChild(ct.Tag("h1"))))
	mustPass(t, fragment.Find(ct.Tag("html")).None())
	mustPass(t, res.HTML().Find(ct.Tag("html")).Find(ct.ID("content")))
	if got, want := fragment.String(), "loads(GET /account html fragment in body)"; got != want {
		t.Fatalf("String(): got %s, want %s", got, want)
	}
}

func TestNonHTMLResponse(t *testing.T) {
	res := cthttp.Serve(accountHandler(), httptest.NewRequest(http.MethodGet, "/api", nil))

	mustPass(t, res.Body(ct.Exact(`{"account":"x"}`)))
	mustFail(t, res.HTML(), ct.ErrLoad, `content type "application/json" is not text/html`, `{\"account\":\"x\"}`)
	mustFail(t, res.HTML().Find(ct.Tag("h1")).None(), ct.ErrLoad)
}

func TestBodyLimitIsACaptureFailure(t *testing.T) {
	res := cthttp.Serve(accountHandler(), httptest.NewRequest(http.MethodGet, "/account", nil), cthttp.WithMaxBody(10))

	mustFail(t, res, ct.ErrLoad, "body exceeds 10 bytes")
	// Operational failures take precedence over every derived assertion.
	mustFail(t, res.Status(http.StatusOK), ct.ErrLoad)
	mustFail(t, res.Header("Content-Type", ct.Prefix("text/html")), ct.ErrLoad)
	mustFail(t, res.HTML(), ct.ErrLoad, "body exceeds")
	if len(res.Bytes()) != 10 {
		t.Fatalf("expected the captured prefix, got %d bytes", len(res.Bytes()))
	}
}

type trackedBody struct {
	io.Reader
	closed   bool
	closeErr error
}

func (b *trackedBody) Close() error {
	b.closed = true
	return b.closeErr
}

func TestFromResponseAdoptsBody(t *testing.T) {
	body := &trackedBody{Reader: strings.NewReader("<p>hi</p>")}
	req := httptest.NewRequest(http.MethodPost, "/submit", nil)
	res := cthttp.FromResponse(&http.Response{
		StatusCode: http.StatusCreated,
		Header:     http.Header{"Content-Type": {"text/html"}},
		Body:       body,
		Request:    req,
	})

	if !body.closed {
		t.Fatal("body must be closed after capture")
	}
	mustPass(t, res)
	mustPass(t, res.Status(http.StatusCreated))
	mustPass(t, res.HTML().Find(ct.Tag("p"), ct.Text("hi")))
	if got, want := res.String(), "captured(POST /submit)"; got != want {
		t.Fatalf("String(): got %s, want %s", got, want)
	}

	errClose := errors.New("close failed")
	failing := cthttp.FromResponse(&http.Response{
		StatusCode: http.StatusOK,
		Body:       &trackedBody{Reader: strings.NewReader("x"), closeErr: errClose},
	})
	err := mustFail(t, failing, ct.ErrLoad, "close body")
	if !errors.Is(err, errClose) {
		t.Fatalf("expected close error in chain, got %v", err)
	}
	if got, want := failing.String(), "captured(response)"; got != want {
		t.Fatalf("String(): got %s, want %s", got, want)
	}

	noBody := cthttp.FromResponse(&http.Response{StatusCode: http.StatusNoContent})
	mustPass(t, noBody.Status(http.StatusNoContent))
	mustPass(t, noBody.Body(ct.Exact("")))
}

func TestServeUsesRequestContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	req := httptest.NewRequestWithContext(ctx, http.MethodGet, "/account", nil)
	res := cthttp.Serve(accountHandler(), req)
	// The handler renders with the request context, so cancellation is an
	// application error, not a capture failure.
	mustPass(t, res)
	mustPass(t, res.Status(http.StatusInternalServerError))
	mustPass(t, res.Body(ct.Contains("context canceled")))
}
