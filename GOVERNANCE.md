# Project Governance

Context Governance Protocol (CGP) is an open-source, pre-alpha protocol and reference implementation. This document explains how technical and community decisions are made.

## Principles

The project is governed around five principles:

1. **Interoperability before implementation preference.** Protocol decisions must be usable by independent implementations.
2. **Security and privacy before convenience.** Unsafe defaults are not accepted merely to reduce integration effort.
3. **Evidence before claims.** Performance, privacy, security, and regulatory claims require reproducible evidence.
4. **Backward compatibility by explicit policy.** Breaking changes require a documented migration path.
5. **Open discussion, accountable decisions.** Major decisions are discussed publicly unless they involve embargoed security information.

## Roles

### Users

Anyone who uses CGP, reads the specifications, or provides feedback.

### Contributors

Anyone who contributes code, documentation, tests, design feedback, issue triage, research, or community support.

### Maintainers

Maintainers review contributions, manage releases, enforce project policies, and protect protocol compatibility. Maintainers are listed in [`MAINTAINERS.md`](MAINTAINERS.md).

### Lead maintainer

During the pre-alpha stage, the lead maintainer is **Matteo Pelosi (`@MrQwenty`)**. The lead maintainer is responsible for final decisions when consensus cannot be reached, release integrity, security embargoes, and appointing additional maintainers.

This is an initial governance model, not a permanent claim of unilateral control. The project should move toward a multi-maintainer technical steering model as independent contributors and implementations emerge.

## Decision process

### Routine changes

Bug fixes, documentation improvements, tests, and implementation changes that do not alter the protocol may be accepted through normal pull-request review.

### Protocol changes

Changes to wire formats, semantics, conformance requirements, security boundaries, or compatibility guarantees require an RFC in `docs/spec/`.

An RFC must include:

- problem statement;
- goals and non-goals;
- terminology;
- normative behavior;
- compatibility impact;
- privacy and security implications;
- operational and regulatory considerations;
- test and conformance requirements;
- rejected alternatives;
- migration plan.

Protocol decisions aim for rough consensus supported by implementation evidence. Silence is not treated as consensus.

### Security decisions

Security vulnerabilities may be handled privately under [`SECURITY.md`](SECURITY.md). Embargoed fixes may be merged and released before full public discussion when disclosure would create material risk.

## Maintainer appointment

New maintainers may be appointed based on sustained, high-quality participation, sound technical judgment, respectful collaboration, and demonstrated care for compatibility, privacy, and security.

Maintainer status may be removed for prolonged inactivity, repeated policy violations, compromised accounts, conflicts of interest that cannot be managed, or conduct that damages project safety.

## Conflicts of interest

Maintainers must disclose material conflicts of interest when evaluating protocol decisions, provider integrations, certification criteria, or commercial services. A conflicted maintainer should recuse themselves where practical.

## Releases

Release artifacts must be reproducible where practical, linked to a source commit, accompanied by release notes, and validated by the conformance and security checks appropriate to the release stage.

## Amendments

Changes to this governance document require a pull request and explicit maintainer approval. Substantial governance changes should remain open for public review before adoption.
