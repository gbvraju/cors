# ADR-0003: Upgrade setup-go so the existing test matrix runs again

- Status: Accepted
- Date: 2026-10-16
- Deciders: maintainers
- Related: ADR-0001, ADR-0002
- Requirements: REQ-003
- Protected classes: .github/

## Context

The project's own workflow `.github/workflows/go.yml` (`Test`, job `test (1.23)`) fails on the fork. The adoption did
not change that file (ADR-0001 lists it as an existing protected file left as it is).

Investigation, read-only, on 2026-10-16:

- The job fails at step 2, "Set up Go 1.23" (`actions/setup-go@v1`). Steps 3 to 5 (checkout, dependencies, `go test`)
  are skipped, so no project test runs. The log of job 117060671478 (run 38926907374, push, `2be2515`) ends with:
  ```
  15:48:33.3208362Z ##[group]Run actions/setup-go@v1
  15:48:33.3210258Z   go-version: 1.23
  15:51:33.8570568Z Cleaning up orphan processes
  ```
  The step prints nothing for exactly 3 minutes and the job is then killed, with no error message. The only warning is
  that `actions/setup-go@v1` targets Node.js 20 and is forced onto Node.js 24.
- The same 3-minute failure of that step occurs on all 7 `Test` runs checked: `b877b06`, `626a5fc`, `7478fb2`,
  `688738a`, `2be2515`, `44bba85` and `1d9b30f`. The `coverage` job, which does not use setup-go, passes.
- The fork has no workflow run on `5d83440` (its first run is the adoption push), so there is no CI result from before
  the adoption. The upstream repository's runs were not read.
- Locally, with Go 1.23.12, the job's exact commands (`go get -v -t -d ./...`, `go test ./...`) pass on `5d83440` and
  on `1d9b30f`: `ok github.com/rs/cors` and `ok github.com/rs/cors/internal`.

Conclusion: the failure is not caused by the adoption and is not a test failure. The probable cause is the old action
`setup-go@v1` hanging while it fetches Go (verify).

## Decision

1. In `.github/workflows/go.yml`, replace `actions/setup-go@v1` with the current major version, `v7`, pinned to its
   full commit SHA with the tag in a comment, as agk's own workflows pin their actions:
   `actions/setup-go@b7ad1dad31e06c5925ef5d2fc7ad053ef454303e   # v7.0.0`. The SHA was resolved with
   `git ls-remote https://github.com/actions/setup-go`.
2. Nothing else in the workflow changes: same Go versions, same steps, same commands.

## Consequences

- The existing test matrix can run again on every push, so REQ-003 (the project's tests must actually run) holds in CI.
- The action is pinned to a commit, so a moved tag cannot change what CI runs.
- The change touches one protected path class, `.github/`, so it needs this ADR.

## Residual

- 2026-10-16: if `Test / test (1.23)` still fails after the upgrade, the cause is not setup-go@v1 and the log of the
  new run decides the next step.
