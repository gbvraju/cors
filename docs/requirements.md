# Requirements

Each requirement is a `### REQ-nnn` section with a `- Status:` line (planned | in-progress | done).
Only the Status line may change without an ADR. `done` requires at least one test marker `REQ-nnn` in a
`test_*.py` file under `tests/`.

`REQ-001`..`REQ-099` are reserved for requirements defined by the kit. Add this project's requirements from
`REQ-101`.

### REQ-001 — Protected paths change only with an ADR in the same PR

- Acceptance: AC1 check fails without an ADR that names the path class; AC2 Status-only change in requirements exempt; AC3 message names the path class
- Status: planned

### REQ-002 — Agents never merge or bypass checks

- Acceptance: AC1 branch rules documented in ADR-0001; AC2 settings deny force-push and hard reset; AC3 stop hook blocks on red governance
- Status: planned

### REQ-003 — Every requirement is traceable to tests

- Acceptance: AC1 matrix generated; AC2 done requires a marker; AC3 matrix never committed
- Status: planned

### REQ-101 — A preflight request from an allowed origin gets the matching Access-Control-Allow-Origin header

- Acceptance: AC1 an `OPTIONS` preflight from `http://foobar.com` with `Access-Control-Request-Method: GET`, to a handler that allows that origin and `GET`, gets status 204 and `Access-Control-Allow-Origin: http://foobar.com`; AC2 the same preflight from `http://barbaz.com`, which is not allowed, gets no `Access-Control-Allow-Origin` header
- Status: done

### REQ-102 — A preflight request for a method that is not allowed gets no allow headers

- Acceptance: AC1 an `OPTIONS` preflight from `http://foobar.com` with `Access-Control-Request-Method: PUT`, to a handler that allows that origin and the methods `PUT` and `DELETE`, gets `Access-Control-Allow-Methods: PUT` and `Access-Control-Allow-Origin: http://foobar.com`; AC2 the same preflight with `Access-Control-Request-Method: PATCH`, which is not allowed, gets neither an `Access-Control-Allow-Methods` nor an `Access-Control-Allow-Origin` header
- Status: done

### REQ-103 — An OptionsSuccessStatus that net/http rejects falls back to 204

- Acceptance: AC1 a handler configured with `OptionsSuccessStatus` below 100 (`-5` or `99`) answers a preflight request with status 204, through `Handler`, `HandlerFunc` and `ServeHTTP`, instead of panicking; AC2 the same for a value above 999 (`1000` or `5000`); AC3 a value from 100 to 999 is kept (`200`, `299` and `999` are used as given)
- Status: done
