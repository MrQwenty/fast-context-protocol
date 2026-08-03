---
title: Project status
description: Implemented, experimental and planned capabilities.
---

CGP is pre-alpha. Checkmarks indicate repository artifacts, not production readiness or regulatory certification.

| Capability | Status |
|---|---|
| Discovery and context resolution | Implemented reference baseline |
| Token and byte budgets | Implemented reference baseline |
| Content-addressed nodes | Implemented reference baseline |
| Inline, reference and fetch delivery | Implemented reference baseline |
| Receipt ingestion | Implemented reference baseline |
| Conformance runner | Implemented reference baseline |
| Local privacy gateway | Experimental implementation |
| Redaction and pseudonymization | Experimental implementation |
| Per-document anonymization | Experimental implementation |
| AES-256-GCM reversible vault | Experimental implementation |
| OpenXML, local PDF and OCR extraction | Experimental implementation |
| Resumable priority streaming | Planned |
| MCP adapter | Planned |
| Zero-config provider proxy | Planned |
| Context Trust Firewall | Planned |
| EU policy compiler and packs | Planned |
| Provider Trust Router | Planned |
| Deletion propagation | Planned |
| SDK ecosystem | Planned |

## Stable release gates

- independent interoperability implementation;
- canonical conformance corpus;
- threat model and independent security review;
- fuzzing and adversarial document tests;
- production authorization and tenant isolation;
- durable receipt and deletion semantics;
- reproducible performance and privacy benchmarks;
- key-management and incident-response documentation;
- formal review of regulatory claims;
- backward-compatibility policy.
