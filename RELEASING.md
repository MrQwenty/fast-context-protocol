# Releasing CGP

This document defines the release process for Context Governance Protocol.

## Progressive releases

Every pull request merged into `master` automatically creates the next patch release:

```text
v0.1.0 → v0.1.1 → v0.1.2 → ...
```

The workflow is defined in `.github/workflows/release.yml` and runs only after a pull request is actually merged. Direct pushes to `master` do not create additional releases.

Until CGP reaches a stable compatibility commitment, automatically generated releases are published as GitHub pre-releases.

Each progressive release:

1. resolves the exact merged commit;
2. reuses an existing tag when a workflow is retried;
3. increments the latest normal semantic-version patch tag;
4. runs `go vet` and race-enabled tests;
5. builds all current command-line programs;
6. packages Linux, macOS, and Windows archives for amd64 and arm64;
7. generates `SHA256SUMS`;
8. creates an annotated Git tag;
9. publishes a GitHub Release with generated release notes.

Current packaged commands:

```text
fcpd
fcpctl
fcpconform
fcpprivacy
```

The legacy `fcp` binary names remain during the additive migration to the CGP namespace.

## Initial version

When no semantic-version tag exists, the automation starts at:

```text
v0.1.0
```

Subsequent merged pull requests increment only the patch component. Minor and major version changes remain explicit maintainer decisions because they communicate compatibility changes.

## Manual recovery

The workflow supports `workflow_dispatch`. A maintainer may optionally provide a target commit SHA. The job is idempotent for commits that already have a release tag.

Manual runs are intended for recovery from infrastructure failures, not for creating duplicate releases.

## Release stages

- `0.x.y` Git tags published as GitHub pre-releases — design and implementation may change substantially;
- beta — feature-complete candidate with active interoperability testing;
- release candidate — only critical fixes expected;
- `1.0.0` — first stable compatibility commitment.

## Version surfaces

A release may version several related surfaces:

- reference implementation;
- protocol revision;
- JSON Schemas;
- conformance suite;
- privacy processor;
- website documentation;
- future SDKs and policy packs.

Release notes must state which surfaces changed.

## Validation performed automatically

```bash
go vet ./...
go test -race ./...
```

The release job also performs cross-platform builds with `CGO_ENABLED=0` and `-trimpath`.

Documentation continues to be validated and deployed by `.github/workflows/docs.yml`.

## Release notes

Generated notes are categorized by `.github/release.yml`. Pull requests should carry meaningful labels when possible.

Release notes should make clear:

- protocol and schema changes;
- compatibility impact;
- migration guidance;
- privacy and security changes;
- known issues;
- validation performed;
- artifact checksums.

## Security releases

Security releases may follow an embargoed process under `SECURITY.md`. Public notes should provide sufficient mitigation guidance without exposing users before patched artifacts are available.

## Artifact integrity

Current progressive releases provide:

- source tag;
- compressed platform archives;
- SHA-256 checksums;
- reproducible build commands.

Before stable releases, the project should additionally provide:

- software bill of materials;
- provenance attestations;
- signed artifacts;
- reproducibility verification on an independent runner.

## Deprecation

No public protocol surface should be removed without:

- a published compatibility notice;
- a replacement or migration path;
- an announced deprecation window;
- conformance tests for old and new behavior during the transition.
