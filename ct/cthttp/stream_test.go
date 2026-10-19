package cthttp_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"testing/synctest"
	"time"

	"github.com/protolambda/chord/ct"
	"github.com/protolambda/chord/ct/cthttp"
)

// events is a server-sent event stream that sends n events and then waits
// for the client to go away. It reports the error of its first failed
// write, or nil when the client went away first, on writeErr.
func events(n int, writeErr chan<- error) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.WriteHeader(http.StatusOK)
		for i := 1; ; i++ {
			if i > n {
				<-r.Context().Done()
				writeErr <- nil
				return
			}
			if _, err := fmt.Fprintf(w, "data: %d\n\n", i); err != nil {
				writeErr <- err
				return
			}
			w.(http.Flusher).Flush()
		}
	})
}

func TestServeStreamUntil(t *testing.T) {
	writeErr := make(chan error, 1)
	req := httptest.NewRequest(http.MethodGet, "/events", nil)
	res := cthttp.ServeStream(events(100, writeErr), req, cthttp.WithUntil(ct.Contains("data: 2\n\n")))

	mustPass(t, res.Captured())
	mustPass(t, res.Status(http.StatusOK))
	mustPass(t, res.HeaderOnce("Content-Type", ct.Exact("text/event-stream")))
	mustPass(t, res.Body(ct.Exact("data: 1\n\ndata: 2\n\n")))
	if !res.Flushed() {
		t.Fatal("the stream was flushed")
	}
	if got, want := res.Describe(), "GET /events"; got != want {
		t.Fatalf("Describe(): got %s, want %s", got, want)
	}
	// The handler returned before ServeStream did, and its next write failed.
	select {
	case err := <-writeErr:
		if err == nil {
			t.Fatal("the write after the capture should have failed")
		}
	default:
		t.Fatal("the handler should have returned")
	}
}

func TestServeStreamTimeout(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		writeErr := make(chan error, 1)
		start := time.Now()
		res := cthttp.ServeStream(events(1, writeErr), httptest.NewRequest(http.MethodGet, "/events", nil),
			cthttp.WithTimeout(time.Minute))

		if elapsed := time.Since(start); elapsed != time.Minute {
			t.Fatalf("the capture took %v, want the timeout", elapsed)
		}
		// The end of the timeout is not a failure; the handler saw the
		// request context end.
		mustPass(t, res.Captured())
		mustPass(t, res.Body(ct.Exact("data: 1\n\n")))
		if err := <-writeErr; err != nil {
			t.Fatalf("the handler should have seen the cancellation, got %v", err)
		}
	})
}

func TestServeStreamDefaultTimeout(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		start := time.Now()
		res := cthttp.ServeStream(events(0, make(chan error, 1)), httptest.NewRequest(http.MethodGet, "/events", nil),
			cthttp.WithTimeout(-1))
		if elapsed := time.Since(start); elapsed != cthttp.DefaultStreamTimeout {
			t.Fatalf("the capture took %v, want %v", elapsed, cthttp.DefaultStreamTimeout)
		}
		mustPass(t, res.Body(ct.Exact("")))
	})
}

// A handler that returns ends the capture, as Serve would capture it.
func TestServeStreamHandlerReturns(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Before", "1")
		_, _ = io.WriteString(w, "<li>one</li>")
		w.Header().Set("X-After", "1")
	})
	res := cthttp.ServeStream(handler, httptest.NewRequest(http.MethodGet, "/list", nil))

	mustPass(t, res.Captured())
	mustPass(t, res.Status(http.StatusOK))
	mustPass(t, res.Header("Content-Type", ct.Exact("text/plain; charset=utf-8")))
	mustPass(t, res.Header("X-Before", ct.Exact("1")))
	mustPass(t, res.HeaderAbsent("X-After"))
	mustPass(t, res.HTMLFragment("ul").Find(ct.Role("listitem")).Texts("one"))
	if res.Flushed() {
		t.Fatal("the handler did not flush")
	}

	quiet := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Location", "/elsewhere")
	})
	res = cthttp.ServeStream(quiet, httptest.NewRequest(http.MethodGet, "/", nil))
	mustPass(t, res.Status(http.StatusOK))
	mustPass(t, res.Header("Location", ct.Exact("/elsewhere")))
}

func TestServeStreamBodyLimit(t *testing.T) {
	writeErr := make(chan error, 1)
	res := cthttp.ServeStream(events(100, writeErr), httptest.NewRequest(http.MethodGet, "/events", nil),
		cthttp.WithMaxBody(12))

	mustFail(t, res.Captured(), ct.ErrLoad, "GET /events: body exceeds 12 bytes")
	if got, want := res.BodyString(), "data: 1\n\ndat"; got != want {
		t.Fatalf("body: got %q, want %q", got, want)
	}
	if err := <-writeErr; err == nil {
		t.Fatal("the write over the limit should have failed")
	}
}

func TestServeStreamHandlerDoesNotReturn(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		release := make(chan struct{})
		returned := make(chan struct{})
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer close(returned)
			w.Header().Set("Content-Type", "text/event-stream")
			w.(http.Flusher).Flush()
			<-release // ignores the request context
		})
		start := time.Now()
		res := cthttp.ServeStream(handler, httptest.NewRequest(http.MethodGet, "/events", nil),
			cthttp.WithTimeout(time.Second))

		if elapsed := time.Since(start); elapsed != 2*time.Second {
			t.Fatalf("ServeStream took %v, want the timeout twice", elapsed)
		}
		mustFail(t, res.Captured(), ct.ErrLoad, "the handler did not return within 1s after the request context was canceled")
		// The header that the handler flushed is known all the same.
		if got := res.Headers().Get("Content-Type"); got != "text/event-stream" {
			t.Fatalf("Content-Type: got %q", got)
		}
		close(release)
		<-returned
	})
}

func TestServeStreamPanics(t *testing.T) {
	boom := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("boom")
	})
	res := cthttp.ServeStream(boom, httptest.NewRequest(http.MethodGet, "/", nil))
	err := mustFail(t, res.Captured(), ct.ErrLoad, "the handler panicked: boom")
	if !strings.Contains(err.Error(), "stream_test.go") {
		t.Fatalf("the failure should include the stack of the panic:\n%v", err)
	}

	abort := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic(http.ErrAbortHandler)
	})
	res = cthttp.ServeStream(abort, httptest.NewRequest(http.MethodGet, "/", nil))
	mustFail(t, res.Captured(), ct.ErrLoad, "the handler aborted the response")

	invalid := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(42)
	})
	res = cthttp.ServeStream(invalid, httptest.NewRequest(http.MethodGet, "/", nil))
	mustFail(t, res.Captured(), ct.ErrLoad, "invalid WriteHeader code 42")
}

// Options of ServeStream do not change Serve.
func TestServeIgnoresStreamOptions(t *testing.T) {
	res := cthttp.Serve(accountHandler(), httptest.NewRequest(http.MethodGet, "/account", nil),
		cthttp.WithUntil(ct.Contains("<html>")), cthttp.WithTimeout(time.Nanosecond))
	mustPass(t, res.Captured())
	mustPass(t, res.Body(ct.Contains("<h1>Account</h1>")))
}
