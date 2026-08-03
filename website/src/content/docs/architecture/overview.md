---
title: Architecture overview
description: Protocol, gateway and evidence layers.
---

CGP is designed as three interoperable parts:

1. **Protocol** — context requests, plans, nodes, policies and receipts.
2. **Gateway** — local or sidecar runtime for privacy, planning, trust, policy and routing.
3. **Evidence layer** — portable records that reconstruct decisions without retaining raw secrets.

```text
Sources: files / DB / APIs / MCP
                 │
                 ▼
        Local Privacy Gateway
                 │
                 ▼
          Context Planner
                 │
       ┌─────────┴─────────┐
       ▼                   ▼
 Trust Firewall       Policy Engine
       └─────────┬─────────┘
                 ▼
          Provider Router
       ┌─────────┼─────────┐
       ▼         ▼         ▼
     Cloud      EU       Local
                 │
                 ▼
          Evidence receipts
```

## Implemented today

- HTTP/JSON discovery and resolution baseline;
- token and byte budget handling;
- content-addressed nodes;
- inline, reference and fetch delivery;
- receipt ingestion;
- local document privacy gateway;
- baseline conformance runner.

## Target architecture

- resumable priority streaming;
- signed policy and provider manifests;
- cross-provider composition;
- purpose and retention propagation;
- trust and prompt-injection firewall;
- provider routing by region, retention and training use;
- regulatory policy packs and audit exports.
