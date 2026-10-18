# ADR-0006: A preflight for a method that is not allowed gets no allow headers

- Status: Proposed
- Date: 2026-10-18
- Deciders: maintainers
- Related: ADR-0004
- Requirements: REQ-102
- Protected classes: docs/requirements.md

## Context

ADR-0004 made one preflight behaviour a requirement (`REQ-101`, the allowed origin). A second documented behaviour
has no requirement: `AllowedMethods` decides which preflight requests get the allow headers. A preflight for a
method in the list gets `Access-Control-Allow-Methods` and the matching `Access-Control-Allow-Origin`; one for a
method not in the list gets neither. The existing tests cover both (`TestSpec`, cases `AllowedMethod` and
`DisallowedMethod`), but none is tied to a requirement.

## Decision

1. This behaviour is a requirement of this repository: `REQ-102` in `docs/requirements.md`, status `planned`.
   AC1 is the allowed-method preflight (`PUT` with `AllowedMethods` `PUT` and `DELETE`); AC2 is the disallowed-method
   preflight (`PATCH`: no `Access-Control-Allow-Methods` and no `Access-Control-Allow-Origin` header).
2. `trace_methods_test.go` carries one test per acceptance criterion, marked `// REQ-102 AC1` and
   `// REQ-102 AC2`. The tests assert existing behaviour only. No product source file changes, so the pull request
   carries `Behaviour: unchanged`.

## Consequences

- `make -f agk.mk agk-traceability` lists `REQ-102 -> trace_methods_test.go`.
- Changing this behaviour later breaks a traced test and has to go through the requirement.
- The status of `REQ-102` stays `planned` until the maintainers change it.

## Residual

- 2026-10-18: none.
