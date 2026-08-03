---
title: Receipts
description: Portable evidence for privacy, context planning, routing, policy and deletion.
---

Receipts make runtime decisions inspectable without attaching original sensitive values.

## PrivacyReceipt

```json
{
  "receiptId": "privacy:...",
  "mode": "anonymize",
  "inputDigest": "sha256:...",
  "outputDigest": "sha256:...",
  "processor": "fcp-local-privacy-gateway",
  "localOnly": true,
  "detected": 12,
  "transformed": 12,
  "unresolved": 0,
  "residualRisk": 0,
  "passed": true
}
```

## Planned receipt families

- **Plan Receipt** — selected, reused, omitted and rejected context nodes;
- **Route Receipt** — selected provider and policy-based rejections;
- **Policy Receipt** — decision, rule, policy version, evidence and remediation;
- **Human Oversight Receipt** — reviewer, competence, evidence and override;
- **Deletion Receipt** — source and derivative invalidation acknowledgements;
- **Incident Receipt** — preserved evidence and notification timers.

Receipts are intended to become signed, portable and independently verifiable.
