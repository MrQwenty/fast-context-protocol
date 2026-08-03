# Changelog

All notable changes to Context Governance Protocol are documented in this file.

The format follows [Keep a Changelog](https://keepachangelog.com/en/1.1.0/). CGP uses Semantic Versioning for implementation releases while the wire protocol remains explicitly versioned.

## [Unreleased]

## [0.1.0] - 2026-08-03

First public pre-alpha release of the CGP reference implementation.

### Added

- HTTP discovery, context resolution, content fetch and receipt ingestion.
- Context plans with explicit token and byte budgets.
- Content-addressed context nodes with provenance, confidence, freshness and sensitivity metadata.
- Inline, reference and fetch delivery modes with baseline cache reuse.
- Local privacy gateway supporting redaction, pseudonymization and per-document anonymization.
- Residual leak scanning and fail-closed privacy policy enforcement.
- AES-256-GCM encrypted vault for reversible pseudonym mappings.
- Local extraction paths for text, OpenXML, PDF and OCR-backed images.
- Privacy receipts containing input/output digests, policy metadata and residual-risk results without original sensitive values.
- Baseline provider conformance runner and canonical fixtures.
- Reference binaries: `fcpd`, `fcpctl`, `fcpconform` and `fcpprivacy`.
- Framework-style technical documentation and a dedicated product website prepared for Vercel.
- Donations page and GitHub Sponsors integration.
- Open-source governance, contribution, support, funding and security policies.
- CodeQL, Dependabot, dependency review and documentation CI.
- Automated progressive release pipeline with cross-platform archives and SHA-256 checksums.

### Changed

- Public technical name adopted as **Context Governance Protocol (CGP)**.
- Existing `fcp` wire paths, headers and binaries retained as pre-alpha compatibility identifiers.
- Documentation website separated into a product landing, documentation hub and community pages.

### Security

- Original documents remain inside the local privacy boundary by design.
- Unsupported or insufficiently inspectable documents can be blocked instead of forwarded.
- Security policy covers context leakage, cache poisoning, provenance forgery, prompt injection, parser abuse and re-identification risk.
- Privacy, security and legal claims are explicitly limited to evidence supported by the implementation.

### Compatibility

- Current discovery endpoint: `/.well-known/fcp`.
- Current protocol namespace: `/fcp/v0.1/...`.
- Current version header: `FCP-Version: 0.1`.
- Current media type: `application/fcp+json`.
- Future `cgp` aliases will be additive and conformance-tested before any deprecation.

### Known limitations

- CGP is pre-alpha and public interfaces may still change.
- Provider trust routing, signed regulatory policy packs, purpose/deletion propagation and the Context Trust Firewall remain roadmap work.
- Automated anonymization reduces exposure but cannot guarantee that contextual or quasi-identifier re-identification is impossible.
- CGP assists governance and evidence collection; it does not automatically certify compliance with the EU AI Act, GDPR or sector-specific law.

### Validation

- `go vet ./...`
- `go test -race ./...`
- Cross-platform builds for Linux, macOS and Windows on AMD64 and ARM64.
- Astro production build for the documentation website.
- Vercel deployment check.

## Release policy

Pre-alpha builds may contain breaking changes. Every tagged release must document:

- protocol compatibility;
- schema changes;
- migration requirements;
- security and privacy implications;
- known limitations;
- validation performed.

[Unreleased]: https://github.com/MrQwenty/fast-context-protocol/compare/v0.1.0...HEAD
[0.1.0]: https://github.com/MrQwenty/fast-context-protocol/releases/tag/v0.1.0
