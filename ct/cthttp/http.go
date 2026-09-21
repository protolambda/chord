// Package cthttp runs HTTP handlers and turns responses into ct subjects.
//
// [Serve] invokes a handler synchronously with a recorder and captures the
// response; [FromResponse] adopts a response obtained elsewhere. Both read
// the body once, up to a limit, so that several assertions and the HTML
// parser never compete for it. The [Result] is itself an assertion that the
// capture succeeded, and every assertion derived from it reports a capture
// failure rather than a mismatch.
//
//	req := httptest.NewRequest(http.MethodGet, "/account", nil)
//	res := cthttp.Serve(handler, req)
//
//	t.Must(res.Status(http.StatusOK))
//	t.Must(res.Header("Content-Type", ct.Prefix("text/html")))
//	page := res.HTML()
//	t.Must(page.Find(ct.Role("heading", ct.Named("Account"))))
//
// Failure messages include a bounded excerpt of the body. Keep secrets out
// of test fixtures, or use [Result.Bytes] with custom assertions instead.
package cthttp

import (
	"context"
	"errors"
	"fmt"
	"io"
	"mime"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/protolambda/mustbe/assertion"

	"github.com/protolambda/chord/ct"
	"github.com/protolambda/chord/ct/cthtml"
)

// DefaultMaxBody is the default body capture limit in bytes.
const DefaultMaxBody = 8 << 20

const excerptLength = 200

// Option configures body capture.
type Option func(*options)

type options struct {
	maxBody int64
}

// WithMaxBody limits the number of body bytes that are captured. A larger
// body is a capture failure.
func WithMaxBody(n int64) Option {
	return func(o *options) {
		o.maxBody = n
	}
}

// Serve invokes the handler with the request, records the response, and
// captures it. The handler runs synchronously before Serve returns.
func Serve(handler http.Handler, request *http.Request, opts ...Option) *Result {
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, request)
	res := rec.Result()
	r := capture(fmt.Sprintf("%s %s", request.Method, request.URL.RequestURI()), res, opts)
	r.flushed = rec.Flushed
	return r
}

// FromResponse captures a response. It adopts the body: the body is read up
// to the limit and closed, and read and close errors are retained as
// capture failures.
func FromResponse(response *http.Response, opts ...Option) *Result {
	desc := "response"
	if response.Request != nil && response.Request.URL != nil {
		desc = fmt.Sprintf("%s %s", response.Request.Method, response.Request.URL.RequestURI())
	}
	return capture(desc, response, opts)
}

func capture(desc string, res *http.Response, opts []Option) *Result {
	o := options{maxBody: DefaultMaxBody}
	for _, opt := range opts {
		opt(&o)
	}
	r := &Result{
		desc:    desc,
		status:  res.StatusCode,
		header:  res.Header.Clone(),
		trailer: res.Trailer.Clone(),
	}
	if res.Body == nil {
		return r
	}
	body, readErr := io.ReadAll(io.LimitReader(res.Body, o.maxBody+1))
	closeErr := res.Body.Close()
	if readErr != nil {
		readErr = fmt.Errorf("read body: %w", readErr)
	} else if int64(len(body)) > o.maxBody {
		readErr = fmt.Errorf("body exceeds %d bytes", o.maxBody)
		body = body[:o.maxBody]
	}
	if closeErr != nil {
		closeErr = fmt.Errorf("close body: %w", closeErr)
	}
	r.body = body
	r.err = errors.Join(readErr, closeErr)
	return r
}

// Result is a captured response. It is an assertion that the capture
// succeeded, and the root of response assertions.
type Result struct {
	desc    string
	status  int
	header  http.Header
	trailer http.Header
	body    []byte
	flushed bool
	err     error
}

var _ assertion.Assertion = (*Result)(nil)

// String describes the assertion: the response was captured.
func (r *Result) String() string {
	return fmt.Sprintf("captured(%s)", r.desc)
}

// Check reports a capture failure, wrapped in [ct.ErrLoad].
func (r *Result) Check(context.Context) error {
	if r.err != nil {
		return fmt.Errorf("%w: %s: %w", ct.ErrLoad, r.desc, r.err)
	}
	return nil
}

// StatusCode returns the response status code.
func (r *Result) StatusCode() int { return r.status }

// Headers returns a copy of the response headers.
func (r *Result) Headers() http.Header { return r.header.Clone() }

// Trailers returns a copy of the response trailers.
func (r *Result) Trailers() http.Header { return r.trailer.Clone() }

// Bytes returns a copy of the captured body.
func (r *Result) Bytes() []byte { return append([]byte(nil), r.body...) }

// Flushed reports whether the handler flushed the response. It is only
// known for results of [Serve].
func (r *Result) Flushed() bool { return r.flushed }

// Status asserts the response status code.
func (r *Result) Status(expected int) assertion.Assertion {
	return check{
		desc:   fmt.Sprintf("status(%s, %d)", r.desc, expected),
		result: r,
		run: func() error {
			if r.status != expected {
				return fmt.Errorf("%w: expected status %d for %s, got %d\nbody: %s",
					ct.ErrMismatch, expected, r.desc, r.status, r.excerpt())
			}
			return nil
		},
	}
}

// Header asserts that at least one value of the header satisfies the matcher.
func (r *Result) Header(name string, expected ct.ValueMatch) assertion.Assertion {
	return check{
		desc:   fmt.Sprintf("header(%s, %q, %s)", r.desc, name, expected),
		result: r,
		run: func() error {
			values := r.header.Values(name)
			for _, v := range values {
				if expected.Match(v) {
					return nil
				}
			}
			if len(values) == 0 {
				return fmt.Errorf("%w: expected header %s of %s to match %s, header is absent",
					ct.ErrMismatch, name, r.desc, expected)
			}
			return fmt.Errorf("%w: expected header %s of %s to match %s, got %q",
				ct.ErrMismatch, name, r.desc, expected, values)
		},
	}
}

// HeaderAbsent asserts that the header is not present.
func (r *Result) HeaderAbsent(name string) assertion.Assertion {
	return check{
		desc:   fmt.Sprintf("headerAbsent(%s, %q)", r.desc, name),
		result: r,
		run: func() error {
			if values := r.header.Values(name); len(values) > 0 {
				return fmt.Errorf("%w: expected header %s of %s to be absent, got %q",
					ct.ErrMismatch, name, r.desc, values)
			}
			return nil
		},
	}
}

// Body asserts that the body satisfies the matcher.
func (r *Result) Body(expected ct.ValueMatch) assertion.Assertion {
	return check{
		desc:   fmt.Sprintf("body(%s, %s)", r.desc, expected),
		result: r,
		run: func() error {
			if !expected.Match(string(r.body)) {
				return fmt.Errorf("%w: expected body of %s to match %s, got %d bytes: %s",
					ct.ErrMismatch, r.desc, expected, len(r.body), r.excerpt())
			}
			return nil
		},
	}
}

// HTML returns a subject for the body parsed as a complete HTML document.
// Loading fails when the capture failed or the content type is not text/html.
func (r *Result) HTML(opts ...ct.Option) *ct.Subject {
	return ct.From(&htmlSource{result: r}, opts...)
}

// HTMLFragment returns a subject for the body parsed as the content of an
// element with the given tag, for example a "div" for an HTMX partial.
// Loading fails when the capture failed or the content type is not text/html.
func (r *Result) HTMLFragment(contextTag string, opts ...ct.Option) *ct.Subject {
	return ct.From(&htmlSource{result: r, context: contextTag}, opts...)
}

func (r *Result) excerpt() string {
	body := strings.TrimSpace(string(r.body))
	if len(body) > excerptLength {
		return fmt.Sprintf("%q…", body[:excerptLength])
	}
	return fmt.Sprintf("%q", body)
}

// check is an assertion over a captured result. A capture failure takes
// precedence over the assertion's own condition.
type check struct {
	desc   string
	result *Result
	run    func() error
}

var _ assertion.Assertion = check{}

func (c check) String() string {
	return c.desc
}

func (c check) Check(ctx context.Context) error {
	if err := c.result.Check(ctx); err != nil {
		return err
	}
	return c.run()
}

// htmlSource parses a captured body through cthtml after validating it.
type htmlSource struct {
	result  *Result
	context string
}

var _ ct.Source = (*htmlSource)(nil)

func (s *htmlSource) String() string {
	if s.context != "" {
		return fmt.Sprintf("%s html fragment in %s", s.result.desc, s.context)
	}
	return s.result.desc + " html"
}

func (s *htmlSource) Load(ctx context.Context) (ct.Document, error) {
	if s.result.err != nil {
		return nil, s.result.err
	}
	contentType := s.result.header.Get("Content-Type")
	mediaType, _, err := mime.ParseMediaType(contentType)
	if err != nil || mediaType != "text/html" {
		return nil, fmt.Errorf("content type %q is not text/html; status %d, body: %s",
			contentType, s.result.status, s.result.excerpt())
	}
	if s.context != "" {
		return cthtml.FragmentBytes(s.context, s.result.body).Load(ctx)
	}
	return cthtml.PageBytes(s.result.body).Load(ctx)
}
