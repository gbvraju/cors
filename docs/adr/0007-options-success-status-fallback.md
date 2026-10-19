# ADR-0007: An out-of-range OptionsSuccessStatus falls back to 204

- Status: Proposed
- Date: 2026-10-19
- Deciders: maintainers
- Related: ADR-0004
- Requirements: REQ-103
- Protected classes: docs/requirements.md

## Context

`Options.OptionsSuccessStatus` is the status code of a successful preflight response; zero means the default 204.
Any other value is passed unchanged to `WriteHeader`. `net/http` panics on a code below 100 or above 999
(`invalid WriteHeader code`). Checked on 2026-10-19 with `-5`, `99`, `1000` and `5000`: every preflight handled by
`Handler`, `HandlerFunc` or `ServeHTTP` panics. On a real server the panic closes the connection (the client sees an
EOF) and the server logs `http: panic serving ...: invalid WriteHeader code 1000`, so a browser cannot complete the
preflight and every cross-origin call fails with no CORS-specific error. The mistake is silent at construction: `New`
accepts the value and the first preflight is the first sign of it.

## Decision

1. A value below 100 or above 999 is replaced by the default 204 in `New`. The zero value takes the same path. A
   value from 100 to 999 is used as given, so nothing that works today changes.
2. The product change is one condition in `New` in `cors.go` and a sentence in the `OptionsSuccessStatus` comment. It is
   covered by `REQ-103` (`AC1` below 100, `AC2` above 999, `AC3` a value from 100 to 999 is kept), with tests in
   `status_test.go` written first and failing on the code before the change (12 failing cases for `AC1` and `AC2`,
   9 passing for `AC3`).
3. Design choice: fall back, do not fail. `New` returns a `*Cors` with no error, and the library is used as
   middleware that is built once at start-up; the handler that results keeps serving with the documented default
   status instead of dropping the connection.

## Alternatives considered

- **Reject the value at construction time** (panic in `New`, or return an error). Rejected: `New` has no error
  result, so returning an error changes the signature of the central function and breaks every caller; panicking in
  `New` turns a failure that today shows at the first preflight into a failure at start-up for every deployment that
  currently runs with such a value, which is a breaking change for a mistake the library can absorb. Falling back
  keeps the zero-value rule ("not set means 204") and extends it to values that cannot be sent.

## Consequences

- A preflight no longer panics for an out-of-range `OptionsSuccessStatus`; it answers 204.
- The misconfiguration is no longer reported to the caller. A debug line is not added (the smallest change).
- A value from 100 to 999 that is not a success status (for example 404) is still used as given; it is out of scope.
- The product source file `cors.go` changes; the protected class touched is `docs/requirements.md` only.

## Residual

- 2026-10-19: whether to log a warning through `Options.Logger` when the fallback is taken is undecided.
