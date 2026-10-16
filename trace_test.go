package cors

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// The tests in this file trace one documented behaviour of the handler to the
// requirement REQ-101 in docs/requirements.md. They assert behaviour the
// library already has.

// REQ-101 AC1
func TestTracePreflightAllowedOrigin(t *testing.T) {
	s := New(Options{
		AllowedOrigins: []string{"http://foobar.com"},
		AllowedMethods: []string{"GET"},
	})

	req, _ := http.NewRequest("OPTIONS", "http://example.com/foo", nil)
	req.Header.Add("Origin", "http://foobar.com")
	req.Header.Add("Access-Control-Request-Method", "GET")

	res := httptest.NewRecorder()
	s.Handler(testHandler).ServeHTTP(res, req)

	assertResponse(t, res, http.StatusNoContent)
	assertHeaders(t, res.Header(), http.Header{
		"Vary":                         {"Origin, Access-Control-Request-Method, Access-Control-Request-Headers"},
		"Access-Control-Allow-Origin":  {"http://foobar.com"},
		"Access-Control-Allow-Methods": {"GET"},
	})
}

// REQ-101 AC2
func TestTracePreflightDisallowedOrigin(t *testing.T) {
	s := New(Options{
		AllowedOrigins: []string{"http://foobar.com"},
		AllowedMethods: []string{"GET"},
	})

	req, _ := http.NewRequest("OPTIONS", "http://example.com/foo", nil)
	req.Header.Add("Origin", "http://barbaz.com")
	req.Header.Add("Access-Control-Request-Method", "GET")

	res := httptest.NewRecorder()
	s.Handler(testHandler).ServeHTTP(res, req)

	if got := res.Header().Values("Access-Control-Allow-Origin"); len(got) != 0 {
		t.Errorf("Response header %q = %q, want none", "Access-Control-Allow-Origin", got)
	}
	assertHeaders(t, res.Header(), http.Header{
		"Vary": {"Origin, Access-Control-Request-Method, Access-Control-Request-Headers"},
	})
}
