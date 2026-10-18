# ADR-0005: Upgrade agk from 0.11.0 to 0.16.0

- Status: Proposed
- Date: 2026-10-17
- Deciders: maintainers
- Related: ADR-0001
- Requirements: REQ-001, REQ-002, REQ-003
- Protected classes: `.github/`, `scripts/agk/`, `tests/kit/`, `agk.toml`

## Context

ADR-0001 installed agk 0.11.0 on 2026-10-16. The kit is now at 0.16.0 (agk `main` at `8cb925c`). The upgrade script
`scripts/agk/upgrade.py` jumps from the recorded version to the kit's in one run. Its dry run on 2026-10-17 listed
308 files to create, 10 to replace and 265 unchanged, and no conflict, so no kit file was edited locally. The upgrade
writes only kit files (`kit/manifest.json` `verbatim` list) and the `version` line of `agk.toml`; the files it changes
fall in four protected classes: `.github/`, `scripts/agk/`, `tests/kit/` and `agk.toml`. The repository's check needs
an ADR in the same pull request that names each class.

## Decision

1. Upgrade the installed kit to 0.16.0 with `python3 <kit>/scripts/agk/upgrade.py .`, and set the `version` line of
   `agk.toml` to `0.16.0`. Nothing else in `agk.toml` changes. The upgrade lists three new template keys
   (`ac_coverage = "warn"`, `proof_timeout = 900`, `inbox_hours = 48`) and does not write them; they are left unset and the defaults apply.
2. The upgrade pull request advances `REQ-001`, `REQ-002` and `REQ-003`, as the adoption pull request did
   (the maintainers' ruling).
3. The new workflow `.github/workflows/product-report.yml` (Mondays 07:13 UTC and by hand) is installed. Like the
   drift-audit and metrics workflows, it writes issues with the workflow token.
4. `.github/PULL_REQUEST_TEMPLATE.md` is replaced by the 0.16.0 template. No product file, test of the project, or
   project workflow is changed.
5. Before this pull request was opened the kit released 0.16.1 (agk `ff99d1e`), which moves the kit's Go fixture
   tree under a `testdata` directory, because the old tree made `go test ./...` fail with `setup failed`. The
   branch is upgraded from 0.16.0 to 0.16.1 with `upgrade.py` (17 files created, 1 replaced, none in conflict), and
   the old tree `tests/kit/fixtures/facts/go-style/tree` is deleted with `git rm`, as `upgrade.py` lists it as
   `removed-in-kit` and never deletes it. The version installed by this pull request is 0.16.1.

## Consequences

- The checks, the doctor and the report scripts are those of 0.16.1; the kit's tests under `tests/kit/` grow with them.
- `doctor.py` reads the installed version and the vendored files against the 0.16.1 kit.
- The weekly product report can run; its first output is a dry run, not a decision.

## Residual

- 2026-10-17: whether GitHub disables the scheduled product-report workflow after inactivity is unconfirmed (verify);
  if the scheduled run does not happen, the maintainers dispatch it by hand.
