package cors

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

// The tests in this file trace the requirement REQ-103 in docs/requirements.md:
// an OptionsSuccessStatus below 100 or above 999, which net/http rejects,
// falls back to the default 204.

// assertPreflightStatus sends a preflight request through the three entry
// points of the handler and checks that each answers with want. A panic is
// reported as a failure of the subtest instead of ending the test binary.
func assertPreflightStatus(t *testing.T, options Options, want int) {
	t.Helper()
	s := New(options)

	req, _ := http.NewRequest("OPTIONS", "http://example.com/foo", nil)
	req.Header.Add("Access-Control-Request-Method", "GET")

	serve := func(t *testing.T, f func(res http.ResponseWriter)) {
		t.Helper()
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("preflight with OptionsSuccessStatus %d panicked: %v", options.OptionsSuccessStatus, r)
			}
		}()
		res := httptest.NewRecorder()
		f(res)
		assertResponse(t, res, want)
	}

	t.Run("Handler", func(t *testing.T) {
		serve(t, func(res http.ResponseWriter) {
			s.Handler(testHandler).ServeHTTP(res, req)
		})
	})
	t.Run("HandlerFunc", func(t *testing.T) {
		serve(t, func(res http.ResponseWriter) {
			s.HandlerFunc(res, req)
		})
	})
	t.Run("Negroni", func(t *testing.T) {
		serve(t, func(res http.ResponseWriter) {
			s.ServeHTTP(res, req, testHandler)
		})
	})
}

// REQ-103 AC1
func TestOptionsSuccessStatusBelowRangeFallsBack(t *testing.T) {
	for _, status := range []int{-5, 99} {
		t.Run(strconv.Itoa(status), func(t *testing.T) {
			assertPreflightStatus(t, Options{OptionsSuccessStatus: status}, http.StatusNoContent)
		})
	}
}

// REQ-103 AC2
func TestOptionsSuccessStatusAboveRangeFallsBack(t *testing.T) {
	for _, status := range []int{1000, 5000} {
		t.Run(strconv.Itoa(status), func(t *testing.T) {
			assertPreflightStatus(t, Options{OptionsSuccessStatus: status}, http.StatusNoContent)
		})
	}
}

// REQ-103 AC3
func TestOptionsSuccessStatusInRangeIsKept(t *testing.T) {
	for _, status := range []int{http.StatusOK, 299, 999} {
		t.Run(strconv.Itoa(status), func(t *testing.T) {
			assertPreflightStatus(t, Options{OptionsSuccessStatus: status}, status)
		})
	}
}
