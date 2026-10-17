# ADR-0005: Upgrade agk from 0.11.0 to 0.16.0

- Status: Proposed
- Date: 2026-10-17
- Deciders: maintainers
- Related: ADR-0001
- Requirements: REQ-001
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
   (`ac_coverage = "warn"`, `proof_timeout = 900`, `inbox_hours = 48`) and does not write them; they are not added here.
2. The new workflow `.github/workflows/product-report.yml` (Mondays 07:13 UTC and by hand) is installed. Like the
   drift-audit and metrics workflows, it writes issues with the workflow token.
3. `.github/PULL_REQUEST_TEMPLATE.md` is replaced by the 0.16.0 template. No product file, test of the project, or
   project workflow is changed.

## Consequences

- The checks, the doctor and the report scripts are those of 0.16.0; the kit's tests under `tests/kit/` grow with them.
- `doctor.py` reads the installed version and the vendored files against the 0.16.0 kit.
- The weekly product report can run; its first output is a dry run, not a decision.

## Residual

- 2026-10-17: whether GitHub disables the scheduled product-report workflow after inactivity is unconfirmed (verify);
  if the scheduled run does not happen, the maintainers dispatch it by hand.
