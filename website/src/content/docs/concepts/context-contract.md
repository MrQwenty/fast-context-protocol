---
title: Context Contract
description: Declare what an AI operation needs and the constraints that govern delivery.
---

A **Context Contract** turns implicit prompt-building assumptions into an explicit request.

```json
{
  "requestId": "req_01K...",
  "intent": {
    "type": "legal-document-summary",
    "target": "document:contract-42"
  },
  "consumer": {
    "modelFamily": "generic",
    "contextWindow": 128000,
    "modalities": ["text"]
  },
  "budget": {
    "maxTokens": 8000,
    "maxBytes": 524288,
    "maxLatencyMs": 100
  },
  "requirements": {
    "freshnessMs": 5000,
    "minimumConfidence": 0.85,
    "includeProvenance": true,
    "acceptSensitivity": ["public", "internal", "sanitized"]
  },
  "knownContext": ["sha256:already-cached-node"]
}
```

## Why it matters

Without a contract, each application independently decides:

- which sources to retrieve;
- how much context to inject;
- whether data is fresh;
- whether personal information may leave the device;
- whether the provider satisfies jurisdiction and retention requirements;
- what evidence is retained.

CGP makes those decisions inspectable and eventually enforceable.

## Context Plan

A provider responds with a ranked plan rather than an unbounded resource dump:

```json
{
  "requestId": "req_01K...",
  "planId": "plan_01K...",
  "protocolVersion": "0.1",
  "contextRoot": "sha256:root...",
  "complete": true,
  "estimatedTokens": 6240,
  "estimatedBytes": 182400,
  "estimatedLatencyMs": 34,
  "chunks": [
    {
      "rank": 1,
      "delivery": "inline",
      "score": 0.98,
      "reason": "required for the requested intent",
      "node": {
        "id": "sha256:...",
        "type": "document.clause",
        "contentType": "text/plain",
        "tokenEstimate": 830,
        "priority": 1,
        "confidence": 0.97,
        "sensitivity": "sanitized"
      }
    }
  ]
}
```
