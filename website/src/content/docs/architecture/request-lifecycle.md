---
title: Request lifecycle
description: From local input to approved model context and evidence.
---

```text
1. Application declares intent and constraints
2. Local input is extracted and sanitized
3. PrivacyReceipt records the transformation
4. Resolver selects context under budget
5. Trust and policy gates evaluate the plan
6. Router selects an allowed processing destination
7. Only approved context reaches the model
8. Receipts record delivery, route and outcome
9. Future invalidations update caches and derivatives
```

The target enforcement vocabulary is deliberately limited:

| Decision | Meaning |
|---|---|
| `ALLOW` | The request satisfies applicable technical policy. |
| `SANITIZE` | Transform sensitive content before continuing. |
| `ROUTE_LOCAL` | External processing is not permitted; use a local model. |
| `REQUIRE_HUMAN` | A qualified reviewer must make or confirm the decision. |
| `DENY` | The operation cannot proceed under the active policy. |
