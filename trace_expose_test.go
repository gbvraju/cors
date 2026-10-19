package cors

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// The tests in this file trace one documented behaviour of the handler to the
// requirement REQ-104 in docs/requirements.md. They assert behaviour the
// library already has.

// REQ-104 AC1
func TestTraceActualRequestExposedHeaders(t *testing.T) {
	s := New(Options{
		AllowedOrigins: []string{"http://foobar.com"},
		ExposedHeaders: []string{"X-Header-1", "x-header-2"},
	})

	req, _ := http.NewRequest("GET", "http://example.com/foo", nil)
	req.Header.Add("Origin", "http://foobar.com")

	res := httptest.NewRecorder()
	s.Handler(testHandler).ServeHTTP(res, req)

	assertHeaders(t, res.Header(), http.Header{
		"Vary":                          {"Origin"},
		"Access-Control-Allow-Origin":   {"http://foobar.com"},
		"Access-Control-Expose-Headers": {"X-Header-1, X-Header-2"},
	})
}

// REQ-104 AC2
func TestTraceActualRequestNoExposedHeaders(t *testing.T) {
	s := New(Options{
		AllowedOrigins: []string{"http://foobar.com"},
	})

	req, _ := http.NewRequest("GET", "http://example.com/foo", nil)
	req.Header.Add("Origin", "http://foobar.com")

	res := httptest.NewRecorder()
	s.Handler(testHandler).ServeHTTP(res, req)

	if got := res.Header().Values("Access-Control-Expose-Headers"); len(got) != 0 {
		t.Errorf("Response header %q = %q, want none", "Access-Control-Expose-Headers", got)
	}
	assertHeaders(t, res.Header(), http.Header{
		"Vary":                        {"Origin"},
		"Access-Control-Allow-Origin": {"http://foobar.com"},
	})
}
