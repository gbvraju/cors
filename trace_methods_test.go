package cors

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// The tests in this file trace one documented behaviour of the handler to the
// requirement REQ-102 in docs/requirements.md. They assert behaviour the
// library already has.

// REQ-102 AC1
func TestTracePreflightAllowedMethod(t *testing.T) {
	s := New(Options{
		AllowedOrigins: []string{"http://foobar.com"},
		AllowedMethods: []string{"PUT", "DELETE"},
	})

	req, _ := http.NewRequest("OPTIONS", "http://example.com/foo", nil)
	req.Header.Add("Origin", "http://foobar.com")
	req.Header.Add("Access-Control-Request-Method", "PUT")

	res := httptest.NewRecorder()
	s.Handler(testHandler).ServeHTTP(res, req)

	assertHeaders(t, res.Header(), http.Header{
		"Vary":                         {"Origin, Access-Control-Request-Method, Access-Control-Request-Headers"},
		"Access-Control-Allow-Origin":  {"http://foobar.com"},
		"Access-Control-Allow-Methods": {"PUT"},
	})
}

// REQ-102 AC2
func TestTracePreflightDisallowedMethod(t *testing.T) {
	s := New(Options{
		AllowedOrigins: []string{"http://foobar.com"},
		AllowedMethods: []string{"PUT", "DELETE"},
	})

	req, _ := http.NewRequest("OPTIONS", "http://example.com/foo", nil)
	req.Header.Add("Origin", "http://foobar.com")
	req.Header.Add("Access-Control-Request-Method", "PATCH")

	res := httptest.NewRecorder()
	s.Handler(testHandler).ServeHTTP(res, req)

	for _, name := range []string{"Access-Control-Allow-Methods", "Access-Control-Allow-Origin"} {
		if got := res.Header().Values(name); len(got) != 0 {
			t.Errorf("Response header %q = %q, want none", name, got)
		}
	}
	assertHeaders(t, res.Header(), http.Header{
		"Vary": {"Origin, Access-Control-Request-Method, Access-Control-Request-Headers"},
	})
}
