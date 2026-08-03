# Project naming: Ctxora

## Decision

**Ctxora** is the recommended working public name for the project originally introduced as **Fast Context Protocol (FCP)**.

The name is pronounced approximately **context-or-a** and is intended to cover the complete ecosystem:

- **Ctxora Protocol** — the interoperable context contract;
- **Ctxora Gateway** — the local or sidecar runtime;
- **`ctxora`** — the future command-line interface;
- **Ctxora Policy Packs** — signed governance rules;
- **Ctxora Receipts** — portable privacy, routing, policy, and context evidence.

## Why the project needs a new name

`FCP` is a short and heavily reused acronym in technology and other industries. It does not communicate the project's broader direction: governed context planning, local privacy, trust enforcement, provider routing, and verifiable evidence.

Ctxora was selected as a working name because:

1. `ctx` is familiar shorthand for context among developers;
2. it is short enough for a CLI and package namespace;
3. it does not restrict the project to transport speed;
4. it can represent the protocol, runtime, policy, and evidence layers together;
5. a preliminary exact-name screen on 3 August 2026 found no matching GitHub repository or clearly established AI context protocol using the name.

The preliminary screen is not legal clearance. Before public commercial use, the project must complete trademark, company-name, domain, package-registry, and jurisdiction-specific checks.

## Compatibility policy

The rename must not silently break existing experiments.

Current compatibility names remain:

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

### N0 — Working brand

- use Ctxora in public-facing documentation;
- preserve all existing protocol and binary names;
- distinguish implemented behavior from planned architecture.

### N1 — Additive aliases

- add a `ctxora` CLI;
- add `/.well-known/ctxora`;
- add `/ctxora/v0.x` route aliases;
- accept `Ctxora-Version` alongside `FCP-Version`;
- test that old and new namespaces produce equivalent behavior.

### N2 — Versioned namespace

- publish a compatibility RFC;
- introduce Ctxora-specific schemas, headers, and media types in a declared version;
- provide automatic configuration migration;
- publish a deprecation policy before changing legacy support.

### N3 — Stable ecosystem

- rename the repository only after redirect and package behavior are verified;
- publish signed releases and checksums;
- complete formal naming clearance;
- establish protocol governance independent from a single implementation.

## Naming principles

A final project name must be:

- distinctive rather than a generic protocol acronym;
- easy to pronounce and type internationally;
- suitable for a CLI, package namespace, and organization name;
- broad enough for context, privacy, policy, routing, and evidence;
- defensible after formal legal and registry checks.

## Historical identifiers

The `FCP-NNNN` document identifiers may remain as historical RFC numbers even after the public brand changes. Renumbering published specifications would create needless ambiguity. Future governance can decide whether new specifications continue the sequence or adopt a Ctxora-native identifier format.
