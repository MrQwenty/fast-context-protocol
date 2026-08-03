---
title: Content-addressed Context Graph
description: Represent context as typed, verifiable nodes and relationships rather than prompt strings.
---

CGP treats context as typed, immutable, content-addressed nodes.

```text
Context root
├── User intent
├── Sanitized document
│   ├── Relevant clause
│   │   └── Original provenance
│   └── Supporting clause
│       └── Original provenance
├── Applicable policy
└── Decision receipt
```

Each node can carry:

- SHA-256 identity;
- type and media type;
- byte and token estimates;
- priority and confidence;
- freshness deadline;
- sensitivity;
- provenance;
- relationships;
- relevant intents and targets;
- privacy transformation state.

## Benefits

- deterministic cache reuse;
- deduplication across providers;
- integrity validation;
- selective fetch;
- future delta updates;
- dependency-aware deletion;
- reproducible Context Plans.
