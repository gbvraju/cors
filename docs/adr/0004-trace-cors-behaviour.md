# ADR-0004: A preflight from an allowed origin is a requirement of this repository

- Status: Accepted
- Date: 2026-10-16
- Deciders: maintainers
- Related: ADR-0001
- Requirements: REQ-101
- Protected classes: docs/requirements.md

## Context

ADR-0001 installed agk. Until now no requirement of the project itself was traced to a test, so the traceability
matrix, the verdict convention and the metrics had no project input. The library already has the behaviour that a
preflight request from an allowed origin gets the matching `Access-Control-Allow-Origin` header, and that one from a
disallowed origin does not. The existing tests cover both (`TestSpec`, cases `MaxAge`, `AllowedMethod` and
`DisallowedOrigin`; `TestHandlePreflightInvalidOriginAbortion`; `TestOptionsSuccessStatusCodeDefault` for the 204
status), but none is tied to a requirement.

## Decision

1. This behaviour is a requirement of this repository: `REQ-101` in `docs/requirements.md`, status `planned`.
   AC1 is the allowed-origin preflight (status 204 and the matching header); AC2 is the disallowed-origin
   preflight (no `Access-Control-Allow-Origin` header).
2. `trace_test.go` carries one test per acceptance criterion, marked `// REQ-101 AC1` and `// REQ-101 AC2`. The
   tests assert existing behaviour only. No product source file changes.
3. The project requirement numbers start at `REQ-101`, as `docs/requirements.md` says.
4. `REQ-101` is done on merge: the behaviour exists and `trace_test.go` proves AC1 and AC2, so its status line
   changes from `planned` to `done`, on the maintainers' instruction.

## Consequences

- `make -f agk.mk agk-traceability` lists `REQ-101 -> trace_test.go`.
- Changing this behaviour later breaks a traced test and has to go through the requirement.
- The status of `REQ-101` stays `planned`; only the maintainers change a status line.

## Residual

- 2026-10-16: `REQ-001` to `REQ-003` have no test marker yet; they are not part of this change.
