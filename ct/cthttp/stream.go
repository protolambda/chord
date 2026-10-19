package cthttp

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"
	"sync"
	"time"

	"github.com/protolambda/chord/ct"
)

// errCaptureEnded is the error of the writes of a [ServeStream] handler
// after the capture ended, as a write to a closed connection fails.
var errCaptureEnded = errors.New("cthttp: the stream capture ended")

// ServeStream invokes a streaming handler, such as a server-sent event
// stream, and captures its response for a bounded time: until the handler
// returns, the body captured so far matches [WithUntil], or the timeout of
// [WithTimeout] ([DefaultStreamTimeout]) passes, whichever comes first.
// Ending at the timeout is not a failure: the result holds what the handler
// sent until then, for assertions such as [Result.Body]. A body over the
// limit of [WithMaxBody] ends the capture as a failure.
//
// When the capture ends before the handler returns, ServeStream cancels the
// request context, as a client that goes away, makes the handler's further
// writes fail, and waits for the handler to return, at most the timeout
// again. A handler that does not return then, or that panics, is a capture
// failure; a handler that does not return keeps running.
//
// The status and header are the ones that the handler sent with its first
// write, flush or WriteHeader, or when it returned; as with [Serve], the
// first write sets a sniffed Content-Type when the handler did not set one.
// The handler runs in its own goroutine, and the waits use timers only, so
// ServeStream also works inside a testing/synctest bubble, in fake time.
func ServeStream(handler http.Handler, request *http.Request, opts ...Option) *Result {
	o := newOptions(opts)
	ctx, cancel := context.WithCancel(request.Context())
	defer cancel()
	w := &streamWriter{
		header:   make(http.Header),
		maxBody:  o.maxBody,
		until:    o.until,
		hasUntil: o.hasUntil,
		captured: make(chan struct{}),
	}
	done := make(chan struct{})
	var panicErr error // written before done is closed
	go func() {
		defer close(done)
		defer func() {
			switch v := recover(); {
			case v == nil:
			case v == http.ErrAbortHandler:
				panicErr = errors.New("the handler aborted the response")
			default:
				panicErr = fmt.Errorf("the handler panicked: %v\n%s", v, debug.Stack())
			}
		}()
		handler.ServeHTTP(w, request.WithContext(ctx))
	}()

	timer := time.NewTimer(o.timeout)
	defer timer.Stop()
	returned := false
	select {
	case <-done:
		returned = true
	case <-w.captured:
	case <-timer.C:
	}
	w.end()
	cancel()
	var handlerErr error
	if !returned {
		timer.Reset(o.timeout)
		select {
		case <-done:
			returned = true
		case <-timer.C:
			handlerErr = fmt.Errorf("the handler did not return within %v after the request context was canceled", o.timeout)
		}
	}
	if returned {
		handlerErr = panicErr
	}

	r := &Result{desc: fmt.Sprintf("%s %s", request.Method, request.URL.RequestURI())}
	w.mu.Lock()
	defer w.mu.Unlock()
	switch {
	case w.wroteHeader:
		r.status, r.header = w.status, w.snapshot
	case returned:
		// The handler is done with its header: net/http sends it as it is.
		r.status, r.header = http.StatusOK, w.header.Clone()
	}
	r.body = w.body
	r.flushed = w.flushed
	r.err = errors.Join(w.err, handlerErr)
	return r
}

// streamWriter records a streaming response. The handler goroutine writes
// it; ServeStream reads it under mu, and the header map only after the
// handler returned.
type streamWriter struct {
	header http.Header // owned by the handler

	maxBody  int64
	until    ct.ValueMatch
	hasUntil bool
	captured chan struct{} // closed when the capture should end early

	mu          sync.Mutex
	wroteHeader bool
	status      int
	snapshot    http.Header // the header when it was written
	body        []byte
	flushed     bool
	ended       bool  // writes fail
	err         error // a capture failure
}

var _ http.Flusher = (*streamWriter)(nil)

func (w *streamWriter) Header() http.Header {
	return w.header
}

func (w *streamWriter) WriteHeader(code int) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.writeHeader(code)
}

// writeHeader records the header once, as httptest.ResponseRecorder does.
// It must be called with mu held, on the handler goroutine.
func (w *streamWriter) writeHeader(code int) {
	if w.wroteHeader {
		return
	}
	if code < 100 || code > 999 {
		panic(fmt.Sprintf("invalid WriteHeader code %v", code))
	}
	w.wroteHeader = true
	w.status = code
	w.snapshot = w.header.Clone()
}

func (w *streamWriter) Write(p []byte) (int, error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.ended {
		return 0, errCaptureEnded
	}
	if !w.wroteHeader {
		if _, ok := w.header["Content-Type"]; !ok && w.header.Get("Transfer-Encoding") == "" {
			w.header.Set("Content-Type", http.DetectContentType(p))
		}
		w.writeHeader(http.StatusOK)
	}
	if room := w.maxBody - int64(len(w.body)); int64(len(p)) > room {
		w.body = append(w.body, p[:room]...)
		w.err = fmt.Errorf("body exceeds %d bytes", w.maxBody)
		w.stop()
		return int(room), errCaptureEnded
	}
	w.body = append(w.body, p...)
	if w.hasUntil && w.until.Match(string(w.body)) {
		w.stop()
	}
	return len(p), nil
}

func (w *streamWriter) Flush() {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.writeHeader(http.StatusOK)
	w.flushed = true
}

// stop ends the capture early, from a write. It must be called with mu
// held, before the capture ended.
func (w *streamWriter) stop() {
	w.ended = true
	close(w.captured)
}

// end ends the capture: further writes fail.
func (w *streamWriter) end() {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.ended = true
}
