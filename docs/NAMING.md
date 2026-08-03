# Project naming: Context Governance Protocol (CGP)

## Decision

**Context Governance Protocol (CGP)** is the public technical name of the project initially developed as Fast Context Protocol (FCP).

The name is intentionally literal:

- **Context** — the information selected and delivered to an AI system;
- **Governance** — privacy, security, purpose, trust, jurisdiction, oversight and evidence;
- **Protocol** — the interoperable contract across applications, agents and providers.

The full name should be used in public positioning. `CGP` is the compact technical acronym.

## Why the name changed

`Fast Context Protocol` overemphasized transport speed and the acronym `FCP` is heavily reused. The project now covers a broader and more defensible problem: governing which context an AI system may receive and proving how that decision was made.

The previously proposed working name `Ctxora` has been rejected and must not be used in new documentation, packages or interfaces.

A preliminary repository-name screen is not legal clearance. Before public commercial use, trademark, organization, domain and package-registry checks remain necessary.

## Current compatibility surface

The reference implementation still uses the original pre-alpha namespace:

```text
Repository          fast-context-protocol
Wire paths          /fcp/v0.1/...
Discovery           /.well-known/fcp
Header              FCP-Version
Media type          application/fcp+json
Binaries            fcpd, fcpctl, fcpconform, fcpprivacy
Specification IDs   FCP-NNNN
```

## Migration phases

### N0 — Public technical name

- use Context Governance Protocol and CGP in documentation;
- retain all current `fcp` interfaces;
- distinguish implemented behavior from planned architecture.

### N1 — Additive aliases

- add a `cgp` CLI;
- add `/.well-known/cgp` discovery;
- add `/cgp/v0.x` route aliases;
- accept `CGP-Version` alongside `FCP-Version`;
- verify equivalent behavior through compatibility tests.

### N2 — Versioned namespace

- publish a compatibility RFC;
- introduce CGP schemas, headers and media types in a declared version;
- provide automatic configuration migration;
- publish a deprecation policy before changing legacy support.

### N3 — Stable ecosystem

- rename the repository only after redirects and package behavior are verified;
- publish signed releases and checksums;
- complete formal naming clearance;
- establish protocol governance independent from one implementation.

## Naming principles

The project name must remain:

- clear and pronounceable internationally;
- descriptive of the real protocol boundary;
- suitable for a CLI, package namespace and organization;
- broad enough for context planning, privacy, policy, routing and evidence;
- compatible with formal legal and registry checks.

## Historical identifiers

Published `FCP-NNNN` document identifiers remain stable historical references. Renumbering specifications would create unnecessary ambiguity. Future governance can decide when new proposals adopt a CGP-native identifier sequence.
