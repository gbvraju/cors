# ADR-0002: CODEOWNERS written without markers by adopt

- Status: Proposed
- Date: 2026-10-16
- Deciders: maintainers
- Related: ADR-0001
- Requirements: REQ-005
- Protected classes: .github/

## Context

ADR-0001 decision 1 says one marked agk block is appended to an existing `.github/CODEOWNERS` and that an absent
one is created. On 2026-10-16 `adopt.py` (agk 94a807d, kit 0.11.0) created `.github/CODEOWNERS` whole and
without the `# >>> agk` ... `# <<< agk` markers, with every line owned by the placeholder `@REPLACE-ME`. Running
`python3 scripts/agk/reviewers.py --codeowners --write` then found no markers and appended a marked block with
`@gbvraju` after the unmarked lines. The file held two blocks, and `doctor.py` failed
`placeholder-team-codeowners`. This is question Q-PL-1 in `docs/runs/pilot-day-1.md`. It is a kit defect and is
tracked and fixed in agk, not in this repository. Adopt does create an absent `.gitignore` with markers; the
maintainers accepted that behaviour (Q-PL-2).

## Decision

1. The stale unmarked block at the top of `.github/CODEOWNERS` (two comment lines and fourteen `@REPLACE-ME`
   lines) is removed on the maintainers' instruction. Only the marked agk block remains.
2. The change touches one protected path class, `.github/`. No kit file under `scripts/agk/` or `tests/kit/` is
   edited here; the defect is fixed in agk and arrives through a later kit upgrade.

## Consequences

- `reviewers.py --check-codeowners` and `doctor.py` report no placeholder team.
- `.github/CODEOWNERS` again matches `config/reviewers.yaml`, so reviewer groups own their path classes (REQ-005 AC2).

## Residual

- 2026-10-16: the adopt defect is open in agk; until it is fixed, a repository adopted without a `.github/CODEOWNERS`
  shows the same two blocks.
