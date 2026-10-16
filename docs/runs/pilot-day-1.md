# Pilot day 1 — adopt agk

Date: 2026-10-16 (from `date -u`). Item 1, PR #1.

## Preconditions

- Fork default branch: `master`
- `go test ./...` on a clean clone: `ok github.com/rs/cors`, `ok github.com/rs/cors/internal`; 17 top-level tests `--- PASS`, 0 `--- FAIL` (`go test ./... -v`)
- agk `main` SHA: `94a807d697f7180ab246bd3c5198e16136527b0d`; `kit/manifest.json` `kit_version` 0.11.0; `scripts/agk/adopt.py` present
- Setup needed before adopt: `git remote set-head origin --auto` (adopt refused without `refs/remotes/origin/HEAD`; result `origin/HEAD set to master`)

## Zero-touch proof

`git ls-files -s > ../before.txt` before adopt (29 tracked files); `git ls-files -s` compared after adopt and the step 4 edits
(adopt adds untracked files; nothing is staged, so the comparison reads the tracked set).

Pre-existing files that differ: **none (0)**. The brief expected exactly one (`.gitignore`, by an appended block). This
repository had no `.gitignore` (`git ls-files | grep -i gitignore` printed nothing), so adopt printed `create .gitignore`
and made a new file; no tracked file changed. `.github/workflows/go.yml` is untouched.

## Adopt output, dry run (verbatim)

```
create .claude/agents/drift-auditor.md
create .claude/hooks/stop_check.py
create .claude/settings.json
create .github/CODEOWNERS
create .github/ISSUE_TEMPLATE/proposal.md
create .github/ISSUE_TEMPLATE/question.md
create .github/ISSUE_TEMPLATE/report.md
create .github/PULL_REQUEST_TEMPLATE.md
create .github/workflows/drift-audit.yml
create .github/workflows/governance.yml
create .github/workflows/metrics.yml
create .gitignore
create CLAUDE.md
create agk.mk
create agk.toml
create config/neutrality.txt
create config/protected-paths.txt
create config/reviewers.yaml
create constitution/principles.md
create constitution/vision.md
create docs/adr/0001-adopt-agk.md
create docs/adr/TEMPLATE.md
create docs/briefs/TEMPLATE.md
create docs/build-state.md
create docs/requirements.md
create docs/runs/decisions-log.md
create docs/specs/TEMPLATE.md
create scripts/agk/check_adrs.py
create scripts/agk/check_build_state.py
create scripts/agk/check_kit_tests.py
create scripts/agk/check_neutrality.py
create scripts/agk/check_pr_body.py
create scripts/agk/check_protected_changes.py
create scripts/agk/check_requirements.py
create scripts/agk/check_specs.py
create scripts/agk/common.py
create scripts/agk/config.py
create scripts/agk/doctor.py
create scripts/agk/drift_audit.py
create scripts/agk/metrics.py
create scripts/agk/reviewers.py
create scripts/agk/traceability.py
create tests/kit/conftest.py
create tests/kit/drift_fixtures.py
create tests/kit/fixtures/doctor/auto-merge-off/api/repo.json
create tests/kit/fixtures/doctor/auto-merge-off/case.json
create tests/kit/fixtures/doctor/bypass-actor/api/ruleset-1.json
create tests/kit/fixtures/doctor/bypass-actor/case.json
create tests/kit/fixtures/doctor/bypass-on-disabled-ruleset/api/ruleset-2.json
create tests/kit/fixtures/doctor/bypass-on-disabled-ruleset/api/rulesets.json
create tests/kit/fixtures/doctor/bypass-on-disabled-ruleset/case.json
create tests/kit/fixtures/doctor/default-branch-mismatch/case.json
create tests/kit/fixtures/doctor/default-branch-mismatch/tree/agk.toml
create tests/kit/fixtures/doctor/delete-branch-off/api/repo.json
create tests/kit/fixtures/doctor/delete-branch-off/case.json
create tests/kit/fixtures/doctor/disabled-holds-every-rule/api/ruleset-1.json
create tests/kit/fixtures/doctor/disabled-holds-every-rule/api/ruleset-2.json
create tests/kit/fixtures/doctor/disabled-holds-every-rule/api/rulesets.json
create tests/kit/fixtures/doctor/disabled-holds-every-rule/case.json
create tests/kit/fixtures/doctor/drift-disabled/api/workflow-drift-audit.json
create tests/kit/fixtures/doctor/drift-disabled/case.json
create tests/kit/fixtures/doctor/drift-no-schedule/case.json
create tests/kit/fixtures/doctor/drift-no-schedule/tree/.github/workflows/drift-audit.yml
create tests/kit/fixtures/doctor/excluded-default-branch/api/ruleset-1.json
create tests/kit/fixtures/doctor/excluded-default-branch/case.json
create tests/kit/fixtures/doctor/excluded-default-ref/api/ruleset-1.json
create tests/kit/fixtures/doctor/excluded-default-ref/case.json
create tests/kit/fixtures/doctor/merge-commits-allowed/api/ruleset-1.json
create tests/kit/fixtures/doctor/merge-commits-allowed/case.json
create tests/kit/fixtures/doctor/no-admin-empty-list/api/repo.json
create tests/kit/fixtures/doctor/no-admin-empty-list/api/rulesets.json
create tests/kit/fixtures/doctor/no-admin-empty-list/case.json
create tests/kit/fixtures/doctor/no-admin/api/ruleset-1.json
create tests/kit/fixtures/doctor/no-admin/case.json
create tests/kit/fixtures/doctor/no-deletion/api/ruleset-1.json
create tests/kit/fixtures/doctor/no-deletion/case.json
create tests/kit/fixtures/doctor/no-dismiss-stale/api/ruleset-1.json
create tests/kit/fixtures/doctor/no-dismiss-stale/case.json
create tests/kit/fixtures/doctor/no-force-push-block/api/ruleset-1.json
create tests/kit/fixtures/doctor/no-force-push-block/case.json
create tests/kit/fixtures/doctor/no-governance-check/api/ruleset-1.json
create tests/kit/fixtures/doctor/no-governance-check/case.json
create tests/kit/fixtures/doctor/no-linear-history/api/ruleset-1.json
create tests/kit/fixtures/doctor/no-linear-history/case.json
create tests/kit/fixtures/doctor/no-origin-head/case.json
create tests/kit/fixtures/doctor/no-ruleset/api/ruleset-1.json
create tests/kit/fixtures/doctor/no-ruleset/api/rulesets.json
create tests/kit/fixtures/doctor/no-ruleset/case.json
create tests/kit/fixtures/doctor/ok/api/repo.json
create tests/kit/fixtures/doctor/ok/api/ruleset-1.json
create tests/kit/fixtures/doctor/ok/api/rulesets.json
create tests/kit/fixtures/doctor/ok/api/workflow-drift-audit.json
create tests/kit/fixtures/doctor/ok/case.json
create tests/kit/fixtures/doctor/ok/kit/.claude/hooks/x.txt
create tests/kit/fixtures/doctor/ok/kit/kit/manifest.json
create tests/kit/fixtures/doctor/ok/kit/scripts/agk/y.txt
create tests/kit/fixtures/doctor/ok/tree/.claude/hooks/x.txt
create tests/kit/fixtures/doctor/ok/tree/.github/CODEOWNERS
create tests/kit/fixtures/doctor/ok/tree/.github/workflows/drift-audit.yml
create tests/kit/fixtures/doctor/ok/tree/agk.toml
create tests/kit/fixtures/doctor/ok/tree/config/reviewers.yaml
create tests/kit/fixtures/doctor/ok/tree/scripts/agk/y.txt
create tests/kit/fixtures/doctor/placeholder-codeowners/case.json
create tests/kit/fixtures/doctor/placeholder-codeowners/tree/.github/CODEOWNERS
create tests/kit/fixtures/doctor/placeholder-reviewers/case.json
create tests/kit/fixtures/doctor/placeholder-reviewers/tree/config/reviewers.yaml
create tests/kit/fixtures/doctor/stricter-ok/api/ruleset-1.json
create tests/kit/fixtures/doctor/stricter-ok/case.json
create tests/kit/fixtures/doctor/target-all/api/ruleset-1.json
create tests/kit/fixtures/doctor/target-all/case.json
create tests/kit/fixtures/doctor/vendored-edited/case.json
create tests/kit/fixtures/doctor/vendored-edited/tree/.claude/hooks/x.txt
create tests/kit/fixtures/doctor/vendored-missing/case.json
create tests/kit/fixtures/doctor/version-mismatch/case.json
create tests/kit/fixtures/doctor/version-mismatch/tree/agk.toml
create tests/kit/fixtures/doctor/weak-enforcement/api/ruleset-1.json
create tests/kit/fixtures/doctor/weak-enforcement/api/rulesets.json
create tests/kit/fixtures/doctor/weak-enforcement/case.json
create tests/kit/fixtures/doctor/wrong-target/api/ruleset-1.json
create tests/kit/fixtures/doctor/wrong-target/case.json
create tests/kit/fixtures/metrics/cleared-other-commit/checks-0067000000000000000000000000000000000000.json
create tests/kit/fixtures/metrics/cleared-other-commit/comments-103.json
create tests/kit/fixtures/metrics/cleared-other-commit/commits-103.json
create tests/kit/fixtures/metrics/cleared-other-commit/files-103.json
create tests/kit/fixtures/metrics/cleared-other-commit/jobs-10030.json
create tests/kit/fixtures/metrics/cleared-other-commit/pull-103.json
create tests/kit/fixtures/metrics/cleared-other-commit/pulls.json
create tests/kit/fixtures/metrics/closed-unmerged/checks-0069000000000000000000000000000000000000.json
create tests/kit/fixtures/metrics/closed-unmerged/comments-105.json
create tests/kit/fixtures/metrics/closed-unmerged/commits-105.json
create tests/kit/fixtures/metrics/closed-unmerged/files-105.json
create tests/kit/fixtures/metrics/closed-unmerged/jobs-10050.json
create tests/kit/fixtures/metrics/closed-unmerged/pull-105.json
create tests/kit/fixtures/metrics/closed-unmerged/pulls.json
create tests/kit/fixtures/metrics/cost-classes/checks-006f000000000000000000000000000000000000.json
create tests/kit/fixtures/metrics/cost-classes/checks-0070000000000000000000000000000000000000.json
create tests/kit/fixtures/metrics/cost-classes/checks-0071000000000000000000000000000000000000.json
create tests/kit/fixtures/metrics/cost-classes/checks-0072000000000000000000000000000000000000.json
create tests/kit/fixtures/metrics/cost-classes/checks-0073000000000000000000000000000000000000.json
create tests/kit/fixtures/metrics/cost-classes/comments-111.json
create tests/kit/fixtures/metrics/cost-classes/comments-112.json
create tests/kit/fixtures/metrics/cost-classes/comments-113.json
create tests/kit/fixtures/metrics/cost-classes/comments-114.json
create tests/kit/fixtures/metrics/cost-classes/comments-115.json
create tests/kit/fixtures/metrics/cost-classes/commits-111.json
create tests/kit/fixtures/metrics/cost-classes/commits-112.json
create tests/kit/fixtures/metrics/cost-classes/commits-113.json
create tests/kit/fixtures/metrics/cost-classes/commits-114.json
create tests/kit/fixtures/metrics/cost-classes/commits-115.json
create tests/kit/fixtures/metrics/cost-classes/files-111.json
create tests/kit/fixtures/metrics/cost-classes/files-112.json
create tests/kit/fixtures/metrics/cost-classes/files-113.json
create tests/kit/fixtures/metrics/cost-classes/files-114.json
create tests/kit/fixtures/metrics/cost-classes/files-115.json
create tests/kit/fixtures/metrics/cost-classes/jobs-10110.json
create tests/kit/fixtures/metrics/cost-classes/jobs-10120.json
create tests/kit/fixtures/metrics/cost-classes/jobs-10130.json
create tests/kit/fixtures/metrics/cost-classes/jobs-10140.json
create tests/kit/fixtures/metrics/cost-classes/jobs-10150.json
create tests/kit/fixtures/metrics/cost-classes/pull-111.json
create tests/kit/fixtures/metrics/cost-classes/pull-112.json
create tests/kit/fixtures/metrics/cost-classes/pull-113.json
create tests/kit/fixtures/metrics/cost-classes/pull-114.json
create tests/kit/fixtures/metrics/cost-classes/pull-115.json
create tests/kit/fixtures/metrics/cost-classes/pulls.json
create tests/kit/fixtures/metrics/empty-week/pulls.json
create tests/kit/fixtures/metrics/files/checks-0079000000000000000000000000000000000000.json
create tests/kit/fixtures/metrics/files/comments-121.json
create tests/kit/fixtures/metrics/files/commits-121.json
create tests/kit/fixtures/metrics/files/files-121.json
create tests/kit/fixtures/metrics/files/jobs-10210.json
create tests/kit/fixtures/metrics/files/pull-121.json
create tests/kit/fixtures/metrics/files/pulls.json
create tests/kit/fixtures/metrics/fix-lists-cleared/checks-0066000000000000000000000000000000000000.json
create tests/kit/fixtures/metrics/fix-lists-cleared/checks-006600cccccccccccccccccccccccccccccccccc.json
create tests/kit/fixtures/metrics/fix-lists-cleared/checks-006601cccccccccccccccccccccccccccccccccc.json
create tests/kit/fixtures/metrics/fix-lists-cleared/comments-102.json
create tests/kit/fixtures/metrics/fix-lists-cleared/commits-102.json
create tests/kit/fixtures/metrics/fix-lists-cleared/files-102.json
create tests/kit/fixtures/metrics/fix-lists-cleared/jobs-10020.json
create tests/kit/fixtures/metrics/fix-lists-cleared/jobs-10021.json
create tests/kit/fixtures/metrics/fix-lists-cleared/jobs-10022.json
create tests/kit/fixtures/metrics/fix-lists-cleared/pull-102.json
create tests/kit/fixtures/metrics/fix-lists-cleared/pulls.json
create tests/kit/fixtures/metrics/governance-split/checks-0083000000000000000000000000000000000000.json
create tests/kit/fixtures/metrics/governance-split/checks-008300cccccccccccccccccccccccccccccccccc.json
create tests/kit/fixtures/metrics/governance-split/checks-008301cccccccccccccccccccccccccccccccccc.json
create tests/kit/fixtures/metrics/governance-split/comments-131.json
create tests/kit/fixtures/metrics/governance-split/commits-131.json
create tests/kit/fixtures/metrics/governance-split/files-131.json
create tests/kit/fixtures/metrics/governance-split/jobs-10310.json
create tests/kit/fixtures/metrics/governance-split/jobs-10311.json
create tests/kit/fixtures/metrics/governance-split/jobs-10312.json
create tests/kit/fixtures/metrics/governance-split/pull-131.json
create tests/kit/fixtures/metrics/governance-split/pulls.json
create tests/kit/fixtures/metrics/governance-unsplit/checks-0084000000000000000000000000000000000000.json
create tests/kit/fixtures/metrics/governance-unsplit/checks-008400cccccccccccccccccccccccccccccccccc.json
create tests/kit/fixtures/metrics/governance-unsplit/comments-132.json
create tests/kit/fixtures/metrics/governance-unsplit/commits-132.json
create tests/kit/fixtures/metrics/governance-unsplit/files-132.json
create tests/kit/fixtures/metrics/governance-unsplit/jobs-10320.json
create tests/kit/fixtures/metrics/governance-unsplit/jobs-10321.json
create tests/kit/fixtures/metrics/governance-unsplit/pull-132.json
create tests/kit/fixtures/metrics/governance-unsplit/pulls.json
create tests/kit/fixtures/metrics/incomplete/checks-00a1000000000000000000000000000000000000.json
create tests/kit/fixtures/metrics/incomplete/comments-161.json
create tests/kit/fixtures/metrics/incomplete/commits-161.json
create tests/kit/fixtures/metrics/incomplete/files-161.json
create tests/kit/fixtures/metrics/incomplete/jobs-10610.json
create tests/kit/fixtures/metrics/incomplete/pull-161.json
create tests/kit/fixtures/metrics/incomplete/pulls.json
create tests/kit/fixtures/metrics/missing-comments/checks-0097000000000000000000000000000000000000.json
create tests/kit/fixtures/metrics/missing-comments/commits-151.json
create tests/kit/fixtures/metrics/missing-comments/files-151.json
create tests/kit/fixtures/metrics/missing-comments/jobs-10510.json
create tests/kit/fixtures/metrics/missing-comments/pull-151.json
create tests/kit/fixtures/metrics/missing-comments/pulls.json
create tests/kit/fixtures/metrics/no-verdict-merged/checks-0065000000000000000000000000000000000000.json
create tests/kit/fixtures/metrics/no-verdict-merged/checks-006500cccccccccccccccccccccccccccccccccc.json
create tests/kit/fixtures/metrics/no-verdict-merged/comments-101.json
create tests/kit/fixtures/metrics/no-verdict-merged/commits-101.json
create tests/kit/fixtures/metrics/no-verdict-merged/files-101.json
create tests/kit/fixtures/metrics/no-verdict-merged/jobs-10010.json
create tests/kit/fixtures/metrics/no-verdict-merged/jobs-10011.json
create tests/kit/fixtures/metrics/no-verdict-merged/pull-101.json
create tests/kit/fixtures/metrics/no-verdict-merged/pulls.json
create tests/kit/fixtures/metrics/verdict-rejected/checks-0068000000000000000000000000000000000000.json
create tests/kit/fixtures/metrics/verdict-rejected/comments-104.json
create tests/kit/fixtures/metrics/verdict-rejected/commits-104.json
create tests/kit/fixtures/metrics/verdict-rejected/files-104.json
create tests/kit/fixtures/metrics/verdict-rejected/jobs-10040.json
create tests/kit/fixtures/metrics/verdict-rejected/pull-104.json
create tests/kit/fixtures/metrics/verdict-rejected/pulls.json
create tests/kit/fixtures/metrics/verdict-suffix/checks-00ab000000000000000000000000000000000000.json
create tests/kit/fixtures/metrics/verdict-suffix/comments-171.json
create tests/kit/fixtures/metrics/verdict-suffix/commits-171.json
create tests/kit/fixtures/metrics/verdict-suffix/files-171.json
create tests/kit/fixtures/metrics/verdict-suffix/jobs-10710.json
create tests/kit/fixtures/metrics/verdict-suffix/pull-171.json
create tests/kit/fixtures/metrics/verdict-suffix/pulls.json
create tests/kit/fixtures/metrics/window-edges/checks-008d000000000000000000000000000000000000.json
create tests/kit/fixtures/metrics/window-edges/checks-008e000000000000000000000000000000000000.json
create tests/kit/fixtures/metrics/window-edges/checks-008f000000000000000000000000000000000000.json
create tests/kit/fixtures/metrics/window-edges/checks-0090000000000000000000000000000000000000.json
create tests/kit/fixtures/metrics/window-edges/comments-141.json
create tests/kit/fixtures/metrics/window-edges/comments-142.json
create tests/kit/fixtures/metrics/window-edges/comments-143.json
create tests/kit/fixtures/metrics/window-edges/comments-144.json
create tests/kit/fixtures/metrics/window-edges/commits-141.json
create tests/kit/fixtures/metrics/window-edges/commits-142.json
create tests/kit/fixtures/metrics/window-edges/commits-143.json
create tests/kit/fixtures/metrics/window-edges/commits-144.json
create tests/kit/fixtures/metrics/window-edges/files-141.json
create tests/kit/fixtures/metrics/window-edges/files-142.json
create tests/kit/fixtures/metrics/window-edges/files-143.json
create tests/kit/fixtures/metrics/window-edges/files-144.json
create tests/kit/fixtures/metrics/window-edges/jobs-10410.json
create tests/kit/fixtures/metrics/window-edges/jobs-10420.json
create tests/kit/fixtures/metrics/window-edges/jobs-10430.json
create tests/kit/fixtures/metrics/window-edges/jobs-10440.json
create tests/kit/fixtures/metrics/window-edges/pull-141.json
create tests/kit/fixtures/metrics/window-edges/pull-142.json
create tests/kit/fixtures/metrics/window-edges/pull-143.json
create tests/kit/fixtures/metrics/window-edges/pull-144.json
create tests/kit/fixtures/metrics/window-edges/pulls.json
create tests/kit/pytest.ini
create tests/kit/test_agk_config.py
create tests/kit/test_agk_mk.py
create tests/kit/test_check_adrs.py
create tests/kit/test_check_build_state.py
create tests/kit/test_check_kit_tests.py
create tests/kit/test_check_neutrality.py
create tests/kit/test_check_pr_body.py
create tests/kit/test_check_protected_changes.py
create tests/kit/test_check_requirements.py
create tests/kit/test_check_specs.py
create tests/kit/test_doctor.py
create tests/kit/test_drift_audit.py
create tests/kit/test_drift_workflow.py
create tests/kit/test_metrics.py
create tests/kit/test_metrics_workflow.py
create tests/kit/test_reviewers.py
create tests/kit/test_scanner_hygiene.py
create tests/kit/test_stop_hook.py
create tests/kit/test_workflow.py
dry run: 285 files would change; nothing was written
notice: config/reviewers.yaml maps `maintainers` to the placeholder @REPLACE-ME
Apply by hand on GitHub to the default branch `master` (adopt configures nothing):
The ruleset for the default branch is exactly: Active; no bypass; target the default branch; restrict deletions; require linear history; require a pull request (0 approvals, dismiss stale approvals, squash merge only); block force pushes; required status check `governance`.
Repository settings: allow auto-merge; auto-delete head branches.

Scanner exclusions (print only; the kit writes none of this). The kit's code under scripts/agk, tests/kit and
.claude/hooks is the kit's, not yours. If a scanner you run reads the whole tree, exclude those paths in your own
configuration:
  bandit:  bandit -r . -x ./scripts/agk,./tests/kit,./.claude/hooks
  flake8:  flake8 --extend-exclude scripts/agk,tests/kit,.claude/hooks
  pylint:  [tool.pylint.main] ignore-paths = ['^scripts/agk/.*$', '^tests/kit/.*$'] and ignore = ['.claude']
  mypy:    [tool.mypy] exclude = ['^scripts/agk/', '^tests/kit/', '^\.claude/hooks/']

Plugin settings entries (print only; the kit merges no JSON). To let a team enable the agk plugin, add these two keys to
`.claude/settings.json` (a protected path class: the change needs an ADR) and have each collaborator trust the folder;
each collaborator installs the plugin on their own machine, and a cloud session does not load it. The repository's
checks and CI work without the plugin.
  "extraKnownMarketplaces": {"agk": {"source": {"source": "github", "repo": "gbvraju/agk", "ref": "v0.11.0"}}}
  "enabledPlugins": {"agk@agk": true}
The `ref` field of a `github` source here is a branch or tag; the docs show it for the source types table but the
`extraKnownMarketplaces` section lists only `repo` (verify).

notice: the workflows drift-audit.yml and metrics.yml write GitHub issues with the workflow token (`issues: write`); delete a workflow file if the repository must not allow that.
```

## Adopt output, real run (verbatim)

```
create .claude/agents/drift-auditor.md
create .claude/hooks/stop_check.py
create .claude/settings.json
create .github/CODEOWNERS
create .github/ISSUE_TEMPLATE/proposal.md
create .github/ISSUE_TEMPLATE/question.md
create .github/ISSUE_TEMPLATE/report.md
create .github/PULL_REQUEST_TEMPLATE.md
create .github/workflows/drift-audit.yml
create .github/workflows/governance.yml
create .github/workflows/metrics.yml
create .gitignore
create CLAUDE.md
create agk.mk
create agk.toml
create config/neutrality.txt
create config/protected-paths.txt
create config/reviewers.yaml
create constitution/principles.md
create constitution/vision.md
create docs/adr/0001-adopt-agk.md
create docs/adr/TEMPLATE.md
create docs/briefs/TEMPLATE.md
create docs/build-state.md
create docs/requirements.md
create docs/runs/decisions-log.md
create docs/specs/TEMPLATE.md
create scripts/agk/check_adrs.py
create scripts/agk/check_build_state.py
create scripts/agk/check_kit_tests.py
create scripts/agk/check_neutrality.py
create scripts/agk/check_pr_body.py
create scripts/agk/check_protected_changes.py
create scripts/agk/check_requirements.py
create scripts/agk/check_specs.py
create scripts/agk/common.py
create scripts/agk/config.py
create scripts/agk/doctor.py
create scripts/agk/drift_audit.py
create scripts/agk/metrics.py
create scripts/agk/reviewers.py
create scripts/agk/traceability.py
create tests/kit/conftest.py
create tests/kit/drift_fixtures.py
create tests/kit/fixtures/doctor/auto-merge-off/api/repo.json
create tests/kit/fixtures/doctor/auto-merge-off/case.json
create tests/kit/fixtures/doctor/bypass-actor/api/ruleset-1.json
create tests/kit/fixtures/doctor/bypass-actor/case.json
create tests/kit/fixtures/doctor/bypass-on-disabled-ruleset/api/ruleset-2.json
create tests/kit/fixtures/doctor/bypass-on-disabled-ruleset/api/rulesets.json
create tests/kit/fixtures/doctor/bypass-on-disabled-ruleset/case.json
create tests/kit/fixtures/doctor/default-branch-mismatch/case.json
create tests/kit/fixtures/doctor/default-branch-mismatch/tree/agk.toml
create tests/kit/fixtures/doctor/delete-branch-off/api/repo.json
create tests/kit/fixtures/doctor/delete-branch-off/case.json
create tests/kit/fixtures/doctor/disabled-holds-every-rule/api/ruleset-1.json
create tests/kit/fixtures/doctor/disabled-holds-every-rule/api/ruleset-2.json
create tests/kit/fixtures/doctor/disabled-holds-every-rule/api/rulesets.json
create tests/kit/fixtures/doctor/disabled-holds-every-rule/case.json
create tests/kit/fixtures/doctor/drift-disabled/api/workflow-drift-audit.json
create tests/kit/fixtures/doctor/drift-disabled/case.json
create tests/kit/fixtures/doctor/drift-no-schedule/case.json
create tests/kit/fixtures/doctor/drift-no-schedule/tree/.github/workflows/drift-audit.yml
create tests/kit/fixtures/doctor/excluded-default-branch/api/ruleset-1.json
create tests/kit/fixtures/doctor/excluded-default-branch/case.json
create tests/kit/fixtures/doctor/excluded-default-ref/api/ruleset-1.json
create tests/kit/fixtures/doctor/excluded-default-ref/case.json
create tests/kit/fixtures/doctor/merge-commits-allowed/api/ruleset-1.json
create tests/kit/fixtures/doctor/merge-commits-allowed/case.json
create tests/kit/fixtures/doctor/no-admin-empty-list/api/repo.json
create tests/kit/fixtures/doctor/no-admin-empty-list/api/rulesets.json
create tests/kit/fixtures/doctor/no-admin-empty-list/case.json
create tests/kit/fixtures/doctor/no-admin/api/ruleset-1.json
create tests/kit/fixtures/doctor/no-admin/case.json
create tests/kit/fixtures/doctor/no-deletion/api/ruleset-1.json
create tests/kit/fixtures/doctor/no-deletion/case.json
create tests/kit/fixtures/doctor/no-dismiss-stale/api/ruleset-1.json
create tests/kit/fixtures/doctor/no-dismiss-stale/case.json
create tests/kit/fixtures/doctor/no-force-push-block/api/ruleset-1.json
create tests/kit/fixtures/doctor/no-force-push-block/case.json
create tests/kit/fixtures/doctor/no-governance-check/api/ruleset-1.json
create tests/kit/fixtures/doctor/no-governance-check/case.json
create tests/kit/fixtures/doctor/no-linear-history/api/ruleset-1.json
create tests/kit/fixtures/doctor/no-linear-history/case.json
create tests/kit/fixtures/doctor/no-origin-head/case.json
create tests/kit/fixtures/doctor/no-ruleset/api/ruleset-1.json
create tests/kit/fixtures/doctor/no-ruleset/api/rulesets.json
create tests/kit/fixtures/doctor/no-ruleset/case.json
create tests/kit/fixtures/doctor/ok/api/repo.json
create tests/kit/fixtures/doctor/ok/api/ruleset-1.json
create tests/kit/fixtures/doctor/ok/api/rulesets.json
create tests/kit/fixtures/doctor/ok/api/workflow-drift-audit.json
create tests/kit/fixtures/doctor/ok/case.json
create tests/kit/fixtures/doctor/ok/kit/.claude/hooks/x.txt
create tests/kit/fixtures/doctor/ok/kit/kit/manifest.json
create tests/kit/fixtures/doctor/ok/kit/scripts/agk/y.txt
create tests/kit/fixtures/doctor/ok/tree/.claude/hooks/x.txt
create tests/kit/fixtures/doctor/ok/tree/.github/CODEOWNERS
create tests/kit/fixtures/doctor/ok/tree/.github/workflows/drift-audit.yml
create tests/kit/fixtures/doctor/ok/tree/agk.toml
create tests/kit/fixtures/doctor/ok/tree/config/reviewers.yaml
create tests/kit/fixtures/doctor/ok/tree/scripts/agk/y.txt
create tests/kit/fixtures/doctor/placeholder-codeowners/case.json
create tests/kit/fixtures/doctor/placeholder-codeowners/tree/.github/CODEOWNERS
create tests/kit/fixtures/doctor/placeholder-reviewers/case.json
create tests/kit/fixtures/doctor/placeholder-reviewers/tree/config/reviewers.yaml
create tests/kit/fixtures/doctor/stricter-ok/api/ruleset-1.json
create tests/kit/fixtures/doctor/stricter-ok/case.json
create tests/kit/fixtures/doctor/target-all/api/ruleset-1.json
create tests/kit/fixtures/doctor/target-all/case.json
create tests/kit/fixtures/doctor/vendored-edited/case.json
create tests/kit/fixtures/doctor/vendored-edited/tree/.claude/hooks/x.txt
create tests/kit/fixtures/doctor/vendored-missing/case.json
create tests/kit/fixtures/doctor/version-mismatch/case.json
create tests/kit/fixtures/doctor/version-mismatch/tree/agk.toml
create tests/kit/fixtures/doctor/weak-enforcement/api/ruleset-1.json
create tests/kit/fixtures/doctor/weak-enforcement/api/rulesets.json
create tests/kit/fixtures/doctor/weak-enforcement/case.json
create tests/kit/fixtures/doctor/wrong-target/api/ruleset-1.json
create tests/kit/fixtures/doctor/wrong-target/case.json
create tests/kit/fixtures/metrics/cleared-other-commit/checks-0067000000000000000000000000000000000000.json
create tests/kit/fixtures/metrics/cleared-other-commit/comments-103.json
create tests/kit/fixtures/metrics/cleared-other-commit/commits-103.json
create tests/kit/fixtures/metrics/cleared-other-commit/files-103.json
create tests/kit/fixtures/metrics/cleared-other-commit/jobs-10030.json
create tests/kit/fixtures/metrics/cleared-other-commit/pull-103.json
create tests/kit/fixtures/metrics/cleared-other-commit/pulls.json
create tests/kit/fixtures/metrics/closed-unmerged/checks-0069000000000000000000000000000000000000.json
create tests/kit/fixtures/metrics/closed-unmerged/comments-105.json
create tests/kit/fixtures/metrics/closed-unmerged/commits-105.json
create tests/kit/fixtures/metrics/closed-unmerged/files-105.json
create tests/kit/fixtures/metrics/closed-unmerged/jobs-10050.json
create tests/kit/fixtures/metrics/closed-unmerged/pull-105.json
create tests/kit/fixtures/metrics/closed-unmerged/pulls.json
create tests/kit/fixtures/metrics/cost-classes/checks-006f000000000000000000000000000000000000.json
create tests/kit/fixtures/metrics/cost-classes/checks-0070000000000000000000000000000000000000.json
create tests/kit/fixtures/metrics/cost-classes/checks-0071000000000000000000000000000000000000.json
create tests/kit/fixtures/metrics/cost-classes/checks-0072000000000000000000000000000000000000.json
create tests/kit/fixtures/metrics/cost-classes/checks-0073000000000000000000000000000000000000.json
create tests/kit/fixtures/metrics/cost-classes/comments-111.json
create tests/kit/fixtures/metrics/cost-classes/comments-112.json
create tests/kit/fixtures/metrics/cost-classes/comments-113.json
create tests/kit/fixtures/metrics/cost-classes/comments-114.json
create tests/kit/fixtures/metrics/cost-classes/comments-115.json
create tests/kit/fixtures/metrics/cost-classes/commits-111.json
create tests/kit/fixtures/metrics/cost-classes/commits-112.json
create tests/kit/fixtures/metrics/cost-classes/commits-113.json
create tests/kit/fixtures/metrics/cost-classes/commits-114.json
create tests/kit/fixtures/metrics/cost-classes/commits-115.json
create tests/kit/fixtures/metrics/cost-classes/files-111.json
create tests/kit/fixtures/metrics/cost-classes/files-112.json
create tests/kit/fixtures/metrics/cost-classes/files-113.json
create tests/kit/fixtures/metrics/cost-classes/files-114.json
create tests/kit/fixtures/metrics/cost-classes/files-115.json
create tests/kit/fixtures/metrics/cost-classes/jobs-10110.json
create tests/kit/fixtures/metrics/cost-classes/jobs-10120.json
create tests/kit/fixtures/metrics/cost-classes/jobs-10130.json
create tests/kit/fixtures/metrics/cost-classes/jobs-10140.json
create tests/kit/fixtures/metrics/cost-classes/jobs-10150.json
create tests/kit/fixtures/metrics/cost-classes/pull-111.json
create tests/kit/fixtures/metrics/cost-classes/pull-112.json
create tests/kit/fixtures/metrics/cost-classes/pull-113.json
create tests/kit/fixtures/metrics/cost-classes/pull-114.json
create tests/kit/fixtures/metrics/cost-classes/pull-115.json
create tests/kit/fixtures/metrics/cost-classes/pulls.json
create tests/kit/fixtures/metrics/empty-week/pulls.json
create tests/kit/fixtures/metrics/files/checks-0079000000000000000000000000000000000000.json
create tests/kit/fixtures/metrics/files/comments-121.json
create tests/kit/fixtures/metrics/files/commits-121.json
create tests/kit/fixtures/metrics/files/files-121.json
create tests/kit/fixtures/metrics/files/jobs-10210.json
create tests/kit/fixtures/metrics/files/pull-121.json
create tests/kit/fixtures/metrics/files/pulls.json
create tests/kit/fixtures/metrics/fix-lists-cleared/checks-0066000000000000000000000000000000000000.json
create tests/kit/fixtures/metrics/fix-lists-cleared/checks-006600cccccccccccccccccccccccccccccccccc.json
create tests/kit/fixtures/metrics/fix-lists-cleared/checks-006601cccccccccccccccccccccccccccccccccc.json
create tests/kit/fixtures/metrics/fix-lists-cleared/comments-102.json
create tests/kit/fixtures/metrics/fix-lists-cleared/commits-102.json
create tests/kit/fixtures/metrics/fix-lists-cleared/files-102.json
create tests/kit/fixtures/metrics/fix-lists-cleared/jobs-10020.json
create tests/kit/fixtures/metrics/fix-lists-cleared/jobs-10021.json
create tests/kit/fixtures/metrics/fix-lists-cleared/jobs-10022.json
create tests/kit/fixtures/metrics/fix-lists-cleared/pull-102.json
create tests/kit/fixtures/metrics/fix-lists-cleared/pulls.json
create tests/kit/fixtures/metrics/governance-split/checks-0083000000000000000000000000000000000000.json
create tests/kit/fixtures/metrics/governance-split/checks-008300cccccccccccccccccccccccccccccccccc.json
create tests/kit/fixtures/metrics/governance-split/checks-008301cccccccccccccccccccccccccccccccccc.json
create tests/kit/fixtures/metrics/governance-split/comments-131.json
create tests/kit/fixtures/metrics/governance-split/commits-131.json
create tests/kit/fixtures/metrics/governance-split/files-131.json
create tests/kit/fixtures/metrics/governance-split/jobs-10310.json
create tests/kit/fixtures/metrics/governance-split/jobs-10311.json
create tests/kit/fixtures/metrics/governance-split/jobs-10312.json
create tests/kit/fixtures/metrics/governance-split/pull-131.json
create tests/kit/fixtures/metrics/governance-split/pulls.json
create tests/kit/fixtures/metrics/governance-unsplit/checks-0084000000000000000000000000000000000000.json
create tests/kit/fixtures/metrics/governance-unsplit/checks-008400cccccccccccccccccccccccccccccccccc.json
create tests/kit/fixtures/metrics/governance-unsplit/comments-132.json
create tests/kit/fixtures/metrics/governance-unsplit/commits-132.json
create tests/kit/fixtures/metrics/governance-unsplit/files-132.json
create tests/kit/fixtures/metrics/governance-unsplit/jobs-10320.json
create tests/kit/fixtures/metrics/governance-unsplit/jobs-10321.json
create tests/kit/fixtures/metrics/governance-unsplit/pull-132.json
create tests/kit/fixtures/metrics/governance-unsplit/pulls.json
create tests/kit/fixtures/metrics/incomplete/checks-00a1000000000000000000000000000000000000.json
create tests/kit/fixtures/metrics/incomplete/comments-161.json
create tests/kit/fixtures/metrics/incomplete/commits-161.json
create tests/kit/fixtures/metrics/incomplete/files-161.json
create tests/kit/fixtures/metrics/incomplete/jobs-10610.json
create tests/kit/fixtures/metrics/incomplete/pull-161.json
create tests/kit/fixtures/metrics/incomplete/pulls.json
create tests/kit/fixtures/metrics/missing-comments/checks-0097000000000000000000000000000000000000.json
create tests/kit/fixtures/metrics/missing-comments/commits-151.json
create tests/kit/fixtures/metrics/missing-comments/files-151.json
create tests/kit/fixtures/metrics/missing-comments/jobs-10510.json
create tests/kit/fixtures/metrics/missing-comments/pull-151.json
create tests/kit/fixtures/metrics/missing-comments/pulls.json
create tests/kit/fixtures/metrics/no-verdict-merged/checks-0065000000000000000000000000000000000000.json
create tests/kit/fixtures/metrics/no-verdict-merged/checks-006500cccccccccccccccccccccccccccccccccc.json
create tests/kit/fixtures/metrics/no-verdict-merged/comments-101.json
create tests/kit/fixtures/metrics/no-verdict-merged/commits-101.json
create tests/kit/fixtures/metrics/no-verdict-merged/files-101.json
create tests/kit/fixtures/metrics/no-verdict-merged/jobs-10010.json
create tests/kit/fixtures/metrics/no-verdict-merged/jobs-10011.json
create tests/kit/fixtures/metrics/no-verdict-merged/pull-101.json
create tests/kit/fixtures/metrics/no-verdict-merged/pulls.json
create tests/kit/fixtures/metrics/verdict-rejected/checks-0068000000000000000000000000000000000000.json
create tests/kit/fixtures/metrics/verdict-rejected/comments-104.json
create tests/kit/fixtures/metrics/verdict-rejected/commits-104.json
create tests/kit/fixtures/metrics/verdict-rejected/files-104.json
create tests/kit/fixtures/metrics/verdict-rejected/jobs-10040.json
create tests/kit/fixtures/metrics/verdict-rejected/pull-104.json
create tests/kit/fixtures/metrics/verdict-rejected/pulls.json
create tests/kit/fixtures/metrics/verdict-suffix/checks-00ab000000000000000000000000000000000000.json
create tests/kit/fixtures/metrics/verdict-suffix/comments-171.json
create tests/kit/fixtures/metrics/verdict-suffix/commits-171.json
create tests/kit/fixtures/metrics/verdict-suffix/files-171.json
create tests/kit/fixtures/metrics/verdict-suffix/jobs-10710.json
create tests/kit/fixtures/metrics/verdict-suffix/pull-171.json
create tests/kit/fixtures/metrics/verdict-suffix/pulls.json
create tests/kit/fixtures/metrics/window-edges/checks-008d000000000000000000000000000000000000.json
create tests/kit/fixtures/metrics/window-edges/checks-008e000000000000000000000000000000000000.json
create tests/kit/fixtures/metrics/window-edges/checks-008f000000000000000000000000000000000000.json
create tests/kit/fixtures/metrics/window-edges/checks-0090000000000000000000000000000000000000.json
create tests/kit/fixtures/metrics/window-edges/comments-141.json
create tests/kit/fixtures/metrics/window-edges/comments-142.json
create tests/kit/fixtures/metrics/window-edges/comments-143.json
create tests/kit/fixtures/metrics/window-edges/comments-144.json
create tests/kit/fixtures/metrics/window-edges/commits-141.json
create tests/kit/fixtures/metrics/window-edges/commits-142.json
create tests/kit/fixtures/metrics/window-edges/commits-143.json
create tests/kit/fixtures/metrics/window-edges/commits-144.json
create tests/kit/fixtures/metrics/window-edges/files-141.json
create tests/kit/fixtures/metrics/window-edges/files-142.json
create tests/kit/fixtures/metrics/window-edges/files-143.json
create tests/kit/fixtures/metrics/window-edges/files-144.json
create tests/kit/fixtures/metrics/window-edges/jobs-10410.json
create tests/kit/fixtures/metrics/window-edges/jobs-10420.json
create tests/kit/fixtures/metrics/window-edges/jobs-10430.json
create tests/kit/fixtures/metrics/window-edges/jobs-10440.json
create tests/kit/fixtures/metrics/window-edges/pull-141.json
create tests/kit/fixtures/metrics/window-edges/pull-142.json
create tests/kit/fixtures/metrics/window-edges/pull-143.json
create tests/kit/fixtures/metrics/window-edges/pull-144.json
create tests/kit/fixtures/metrics/window-edges/pulls.json
create tests/kit/pytest.ini
create tests/kit/test_agk_config.py
create tests/kit/test_agk_mk.py
create tests/kit/test_check_adrs.py
create tests/kit/test_check_build_state.py
create tests/kit/test_check_kit_tests.py
create tests/kit/test_check_neutrality.py
create tests/kit/test_check_pr_body.py
create tests/kit/test_check_protected_changes.py
create tests/kit/test_check_requirements.py
create tests/kit/test_check_specs.py
create tests/kit/test_doctor.py
create tests/kit/test_drift_audit.py
create tests/kit/test_drift_workflow.py
create tests/kit/test_metrics.py
create tests/kit/test_metrics_workflow.py
create tests/kit/test_reviewers.py
create tests/kit/test_scanner_hygiene.py
create tests/kit/test_stop_hook.py
create tests/kit/test_workflow.py
installed 285 files
notice: config/reviewers.yaml maps `maintainers` to the placeholder @REPLACE-ME
Apply by hand on GitHub to the default branch `master` (adopt configures nothing):
The ruleset for the default branch is exactly: Active; no bypass; target the default branch; restrict deletions; require linear history; require a pull request (0 approvals, dismiss stale approvals, squash merge only); block force pushes; required status check `governance`.
Repository settings: allow auto-merge; auto-delete head branches.

Scanner exclusions (print only; the kit writes none of this). The kit's code under scripts/agk, tests/kit and
.claude/hooks is the kit's, not yours. If a scanner you run reads the whole tree, exclude those paths in your own
configuration:
  bandit:  bandit -r . -x ./scripts/agk,./tests/kit,./.claude/hooks
  flake8:  flake8 --extend-exclude scripts/agk,tests/kit,.claude/hooks
  pylint:  [tool.pylint.main] ignore-paths = ['^scripts/agk/.*$', '^tests/kit/.*$'] and ignore = ['.claude']
  mypy:    [tool.mypy] exclude = ['^scripts/agk/', '^tests/kit/', '^\.claude/hooks/']

Plugin settings entries (print only; the kit merges no JSON). To let a team enable the agk plugin, add these two keys to
`.claude/settings.json` (a protected path class: the change needs an ADR) and have each collaborator trust the folder;
each collaborator installs the plugin on their own machine, and a cloud session does not load it. The repository's
checks and CI work without the plugin.
  "extraKnownMarketplaces": {"agk": {"source": {"source": "github", "repo": "gbvraju/agk", "ref": "v0.11.0"}}}
  "enabledPlugins": {"agk@agk": true}
The `ref` field of a `github` source here is a branch or tag; the docs show it for the source types table but the
`extraKnownMarketplaces` section lists only `repo` (verify).

notice: the workflows drift-audit.yml and metrics.yml write GitHub issues with the workflow token (`issues: write`); delete a workflow file if the repository must not allow that.
```

## Step 4 configuration

- `agk.toml`: `project_tests = "go test ./..."`, `test_dirs = ["."]`, `test_patterns = ["*_test.go"]`; `experimental_reviewer` left unset
- `config/reviewers.yaml`: `@REPLACE-ME` replaced with `@gbvraju`
- `python3 scripts/agk/reviewers.py --codeowners --write`: `wrote agk block of .github/CODEOWNERS`

## Make targets

(pytest and ruff 0.16.x from a throwaway venv at `../agk-venv`, outside the repository)

- `make -f agk.mk agk-governance`: exit 0; ok protected-changes, requirements (3), adrs (1), specs, neutrality (0 tokens), build-state, kit-tests, reviewers, reviewers (.github/CODEOWNERS matches)
- `make -f agk.mk agk-lint`: `All checks passed!`
- `make -f agk.mk agk-test`: `380 passed, 2 skipped in 37.60s`
- `make -f agk.mk agk-project-test`: `ok github.com/rs/cors`, `ok github.com/rs/cors/internal`

## Doctor (`python3 scripts/agk/doctor.py --kit-root ../agk-kit`, verbatim)

```
agk doctor 0.11.0
FAIL ruleset-missing error no ruleset targets the default branch
  fix: create the ruleset of ADR-0001 decision 3: active; no bypass; target the default branch; restrict deletions; require linear history; require a pull request (dismiss stale approvals, no merge commits); block force pushes; required status check `governance`; classic branch protection is not read; use a ruleset (ADR-0001 decision 3)
?? ruleset-enforcement error cannot check (no-ruleset)
  to check: no ruleset targets the default branch; fix ruleset-missing or ruleset-target
?? ruleset-bypass error cannot check (no-ruleset)
  to check: no ruleset targets the default branch; fix ruleset-missing or ruleset-target
?? ruleset-target error cannot check (no-ruleset)
  to check: no ruleset targets the default branch; fix ruleset-missing or ruleset-target
?? ruleset-deletion error cannot check (no-ruleset)
  to check: no ruleset targets the default branch; fix ruleset-missing or ruleset-target
?? ruleset-linear-history error cannot check (no-ruleset)
  to check: no ruleset targets the default branch; fix ruleset-missing or ruleset-target
?? ruleset-force-push error cannot check (no-ruleset)
  to check: no ruleset targets the default branch; fix ruleset-missing or ruleset-target
?? ruleset-pull-request error cannot check (no-ruleset)
  to check: no ruleset targets the default branch; fix ruleset-missing or ruleset-target
?? required-check-governance error cannot check (no-ruleset)
  to check: no ruleset targets the default branch; fix ruleset-missing or ruleset-target
FAIL repo-auto-merge warn repository setting allow_auto_merge is off
  fix: allow auto-merge
FAIL repo-delete-branch warn repository setting delete_branch_on_merge is off
  fix: auto-delete head branches
ok placeholder-team-reviewers error config/reviewers.yaml has no placeholder team
FAIL placeholder-team-codeowners error .github/CODEOWNERS still names the placeholder team
  fix: replace the placeholder with a real team or user, then run `python3 scripts/agk/reviewers.py --codeowners --write`
ok origin-head warn refs/remotes/origin/HEAD exists
ok drift-workflow-file error .github/workflows/drift-audit.yml has a schedule
?? drift-workflow-state error cannot check (no-admin)
  to check: use an account with admin rights on the repository
ok kit-version warn the kit version equals version in agk.toml
ok vendored-missing error every vendored file is present
ok vendored-differs error every vendored file equals the kit's copy
ok default-branch-config warn default_branch in agk.toml equals the repository's default branch
summary: 7 ok, 4 fail, 9 cannot-check
```

Expected `FAIL`/cannot-check lines: ruleset, required check, auto-merge, delete-branch (maintainers apply after merge; the
`drift-workflow-state` check needs an admin account). Unexpected: `placeholder-team-codeowners`, see Q-PL-1.

## Questions

- Q-PL-1: `.github/CODEOWNERS` has two blocks. Adopt created the file whole and unmarked (with `@REPLACE-ME`), and
  `reviewers.py --codeowners --write` then appended a marked `# >>> agk` block with `@gbvraju` because the file had no
  markers. The later block wins in GitHub's matching, but doctor's `placeholder-team-codeowners` fails. Left as
  generated, because the brief forbids editing installed kit files beyond step 4. Is the fix to delete the unmarked
  first block by hand, or to change agk (adopt writes CODEOWNERS as a marked block, as it does for `.gitignore`)?
- Q-PL-2: the brief expected `append .gitignore`; the fork has no `.gitignore`, so adopt created it. Confirm that this
  is acceptable for the zero-touch claim (no tracked file changed).
