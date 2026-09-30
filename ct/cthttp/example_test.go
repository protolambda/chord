package cthttp_test

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/protolambda/chord/ct"
	"github.com/protolambda/chord/ct/cthttp"
)

// A handler is served once; the result holds the response for several
// assertions, and the body parses as HTML, here a partial that Go sniffs as
// text/plain. In a test, pass each assertion to t.Must.
func ExampleServe() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, `<li><a href="/tasks/1">Task 1</a></li>`)
	})
	res := cthttp.Serve(handler, httptest.NewRequest(http.MethodGet, "/tasks", nil))

	ctx := context.Background()
	fmt.Println(res.Captured().Check(ctx))
	fmt.Println(res.Status(http.StatusOK).Check(ctx))
	fmt.Println(res.Body(ct.Contains("Task 1")).Check(ctx))
	items := res.HTMLFragment("ul").Find(ct.Role("listitem"))
	fmt.Println(items.Texts("Task 1").Check(ctx))

	fmt.Println(res.Describe())
	fmt.Println(res.BodyString())
	// Output:
	// <nil>
	// <nil>
	// <nil>
	// <nil>
	// GET /tasks
	// <li><a href="/tasks/1">Task 1</a></li>
}
