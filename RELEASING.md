# Releasing CGP

This document defines the intended release process for Context Governance Protocol.

## Release stages

- `0.x.y-alpha` — design and implementation may change substantially;
- `0.x.y-beta` — feature-complete candidate with active interoperability testing;
- `0.x.y-rc.N` — release candidate with only critical fixes expected;
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

## Pre-release checklist

1. Confirm the working tree is clean and the intended commit is reviewed.
2. Update `CHANGELOG.md`.
3. Run:

```bash
gofmt -w $(find . -name '*.go' -type f)
go vet ./...
go test -race ./...
go build ./cmd/fcpd ./cmd/fcpctl ./cmd/fcpconform ./cmd/fcpprivacy
```

4. Build documentation:

```bash
cd website
npm install
npm run build
```

5. Run conformance and privacy smoke tests.
6. Review schemas and fixtures for compatibility.
7. Review dependency and CodeQL results.
8. Confirm no secrets, private fixtures, generated credentials, or local artifacts are included.
9. Document known limitations and migration steps.
10. Create a signed tag where signing infrastructure is available.

## Release notes

Release notes must include:

- summary;
- protocol and schema changes;
- compatibility impact;
- migration guidance;
- privacy and security changes;
- known issues;
- validation performed;
- artifact checksums;
- contributor acknowledgements.

## Security releases

Security releases may follow an embargoed process under `SECURITY.md`. Public notes should provide sufficient mitigation guidance without exposing users before patched artifacts are available.

## Artifact integrity

Stable releases should provide:

- source tag;
- checksums;
- reproducible build instructions;
- software bill of materials;
- provenance attestation;
- signed artifacts where practical.

## Deprecation

No public protocol surface should be removed without:

- a published compatibility notice;
- a replacement or migration path;
- an announced deprecation window;
- conformance tests for old and new behavior during the transition.
