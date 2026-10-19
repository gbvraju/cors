# ADR-0008: A request from an allowed origin gets the configured exposed headers

- Status: Proposed
- Date: 2026-10-19
- Deciders: maintainers
- Related: ADR-0004, ADR-0006
- Requirements: REQ-104
- Protected classes: docs/requirements.md

## Context

ADR-0004 and ADR-0006 made two preflight behaviours requirements (`REQ-101`, `REQ-102`). A documented behaviour of the
actual request has no requirement: `ExposedHeaders` lists the response headers a browser may read, and the handler
sends them in `Access-Control-Expose-Headers` in canonical form when the origin is allowed, and sends no such header
when none is configured. The existing tests cover both (`TestSpec`, case `ExposedHeader`;
`TestAccessControlExposeHeadersPresence`), but none is tied to a requirement.

## Decision

1. This behaviour is a requirement of this repository: `REQ-104` in `docs/requirements.md`, status `planned`.
   AC1 is the allowed-origin request with `ExposedHeaders` `X-Header-1` and `x-header-2` (the header value is
   `X-Header-1, X-Header-2`); AC2 is the same request with no `ExposedHeaders` (no `Access-Control-Expose-Headers`
   header).
2. `trace_expose_test.go` carries one test per acceptance criterion, marked `// REQ-104 AC1` and
   `// REQ-104 AC2`. The tests assert existing behaviour only. No product source file changes, so the pull request
   carries `Behaviour: unchanged`.

## Consequences

- `make -f agk.mk agk-traceability` lists `REQ-104 -> trace_expose_test.go`.
- Changing this behaviour later breaks a traced test and has to go through the requirement.
- The status of `REQ-104` stays `planned` until the maintainers change it.

## Residual

- 2026-10-19: none.
