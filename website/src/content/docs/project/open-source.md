---
title: Open-source strategy
description: Recommended open-core boundary for adoption, trust and sustainable development.
---

## Recommendation

CGP should be developed as an **open protocol with an open reference core and commercial governance services**.

## Open and auditable

The following components should remain open source:

- protocol specifications and schemas;
- conformance suite and canonical fixtures;
- reference provider and gateway core;
- privacy transformations and receipt verification;
- MCP and provider adapter interfaces;
- Go, Python and TypeScript SDKs;
- baseline EU policy vocabulary and rule format;
- threat model and security documentation.

This makes the interoperability contract inspectable and allows independent implementations to prove compatibility.

## Commercial differentiation

A sustainable company can charge for:

- managed signed regulatory-policy update channels;
- sector and country implementation packs;
- enterprise identity, tenant isolation and administration;
- managed evidence vault and retention controls;
- certified provider and data-source connectors;
- confidential-compute and EU-only routing services;
- advanced observability, incident workflows and audit export;
- certification programs, support and service-level agreements.

## Licensing direction

- **Code:** Apache License 2.0 is appropriate for the current reference implementation because it permits broad adoption and includes an explicit patent grant.
- **Specifications:** keep them openly readable and contribution-friendly; a dedicated specification licence can be evaluated before standardization.
- **Trademark:** retain control of the CGP name and conformance marks so incompatible implementations cannot present themselves as certified.

## Why not fully closed

A closed protocol would make it difficult to become an interoperability standard and would weaken trust in privacy, security and compliance claims.

## Why not give away every service

Regulatory maintenance, provider attestations, evidence retention, enterprise operations and certified integrations require continuous work. Those operational capabilities form a credible commercial layer without hiding the protocol itself.
