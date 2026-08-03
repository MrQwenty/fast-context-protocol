# Maintainer Repository Setup

Some open-source repository controls are stored in files. Others must be enabled in GitHub settings by a repository administrator. This checklist records the intended configuration for CGP.

## GitHub Pages

1. Open **Settings → Pages**.
2. Under **Build and deployment**, set **Source** to **GitHub Actions**.
3. Open **Actions → Documentation** and run the workflow manually if no deployment starts.
4. Confirm the `github-pages` environment is created.
5. Restrict deployments to the default branch when environment protection is available.

Expected project-site URL:

```text
https://mrqwenty.github.io/fast-context-protocol/
```

The workflow is defined in `.github/workflows/docs.yml`.

## GitHub Sponsors

The repository sponsor button is configured by `.github/FUNDING.yml`:

```yaml
github:
  - MrQwenty
```

Confirm the GitHub Sponsors profile is approved and visible at:

```text
https://github.com/sponsors/MrQwenty
```

## Private vulnerability reporting

1. Open **Settings → Code security and analysis** or **Settings → Advanced Security**.
2. Enable **Private vulnerability reporting**.
3. Verify that non-maintainers can see **Report a vulnerability** under the Security tab.
4. Configure security notification recipients.

The public reporting policy is in `SECURITY.md`.

## Branch protection or ruleset

Protect the default branch `master` with a repository ruleset or branch protection rule.

Recommended minimum:

- require a pull request before merging;
- require conversation resolution;
- require status checks;
- block force pushes;
- block branch deletion;
- require Code Owner review once additional maintainers exist;
- dismiss stale approvals after new commits;
- require the branch to be up to date before merge;
- require signed commits when the contributor workflow supports them.

Suggested required checks after they have run at least once:

- CI / test;
- CodeQL / Analyze;
- Documentation / build;
- dependency review, if enabled.

A single-maintainer project may temporarily allow the lead maintainer to bypass review requirements for emergency and bootstrap changes. Any bypass should be used sparingly.

## Actions permissions

Open **Settings → Actions → General** and confirm:

- GitHub Actions is enabled;
- verified GitHub-maintained actions are allowed;
- workflow permissions are read-only by default unless a workflow declares narrower write permissions;
- pull requests from forks require approval before executing workflows that could access sensitive resources;
- secrets are never exposed to untrusted pull-request workflows.

## Security features

Enable where available:

- dependency graph;
- Dependabot alerts;
- Dependabot security updates;
- secret scanning;
- push protection;
- CodeQL default setup or the repository workflow;
- private vulnerability reporting.

Do not run duplicate CodeQL configurations unless intentionally comparing them.

## Issues and Discussions

Issues should remain enabled for structured bug reports and proposals.

GitHub Discussions is recommended for:

- general questions;
- design exploration before an RFC;
- announcements;
- implementation showcases;
- community support that is not actionable repository work.

If Discussions is enabled, update `SUPPORT.md` and `.github/ISSUE_TEMPLATE/config.yml` with the correct category links.

## Repository metadata

Recommended About-section values:

**Description**

```text
Open protocol and runtime for governed AI context: privacy, trust, provenance, policy and verifiable receipts.
```

**Website**

```text
https://mrqwenty.github.io/fast-context-protocol/
```

**Topics**

```text
ai ai-governance context-protocol privacy security llm mcp interoperability go open-source
```

## Merge policy

Recommended merge methods:

- squash merge for normal contributions;
- merge commits only for release or intentionally preserved branch history;
- disable rebase merge if it conflicts with signing or audit requirements.

Delete merged contributor branches where appropriate.

## Release settings

Before the first tagged release:

- create a release checklist issue;
- configure artifact signing;
- generate checksums and an SBOM;
- verify `CITATION.cff`;
- update `CHANGELOG.md`;
- review `RELEASING.md`;
- verify the website deployment;
- verify the sponsor and security-reporting links.
