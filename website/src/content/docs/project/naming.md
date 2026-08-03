---
title: Naming and compatibility
description: Why the project is called Context Governance Protocol and how the legacy FCP namespace will migrate.
---

## Context Governance Protocol

The name describes the project in three words:

- **Context** — the governed asset delivered to an AI system;
- **Governance** — privacy, trust, purpose, policy, jurisdiction and evidence;
- **Protocol** — the interoperable contract across applications and providers.

The full name should be used in public positioning; **CGP** is the compact technical acronym.

The previously proposed working name `Ctxora` was rejected and is not part of the project naming strategy.

## Current compatibility surface

```text
Repository          fast-context-protocol
Wire paths          /fcp/v0.1/...
Discovery           /.well-known/fcp
Header              FCP-Version
Media type          application/fcp+json
Binaries            fcpd, fcpctl, fcpconform, fcpprivacy
Specification IDs   FCP-NNNN
```

## Migration policy

1. introduce `cgp` CLI and protocol aliases additively;
2. verify equivalent behavior through conformance tests;
3. publish a compatibility RFC and deprecation policy;
4. rename repository and packages only after redirects and migration paths are verified;
5. retain historical specification identifiers to avoid ambiguous references.
