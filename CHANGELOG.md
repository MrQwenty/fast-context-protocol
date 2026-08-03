# Changelog

All notable changes to Context Governance Protocol will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/), and the project intends to follow Semantic Versioning once the public protocol versioning policy is finalized.

## [Unreleased]

### Added

- Context planning reference implementation with explicit token and byte budgets.
- Content-addressed context nodes, plans, fetch behavior, and receipts.
- Local privacy gateway with redaction, pseudonymization, anonymization, residual leak scanning, and encrypted reversible vaults.
- Baseline conformance runner.
- Astro Starlight documentation website.
- Open-source governance, contribution, funding, support, and security automation.

### Changed

- Public technical name adopted as **Context Governance Protocol (CGP)**.
- Existing `fcp` wire paths and binaries retained as pre-alpha compatibility identifiers.

### Security

- Added fail-closed privacy processing requirements.
- Added CodeQL and Dependabot configuration.

## Release policy

Pre-alpha builds may contain breaking changes. Every tagged release must document:

- protocol compatibility;
- schema changes;
- migration requirements;
- security and privacy implications;
- known limitations;
- validation performed.
