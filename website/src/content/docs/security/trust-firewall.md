---
title: Context Trust Firewall
description: Keep untrusted data from becoming executable authority.
---

A document, webpage, email or tool result is data. It must not silently become a privileged instruction.

## Trust labels

```text
SYSTEM_POLICY       highest authority
DEVELOPER_POLICY    application authority
USER_INSTRUCTION    authorized user intent
TRUSTED_DATA        verified information
UNTRUSTED_DATA      external or retrieved content
GENERATED_DATA      model-produced content
```

## Normative target

Untrusted context must never be permitted to:

- increase its own authority;
- grant tool permissions;
- alter system or compliance policy;
- disable logging;
- change approved destinations;
- request disclosure of protected context.

## Planned controls

- source-level and span-level taint tracking;
- prompt-injection detection;
- capability firewall for tools;
- output exfiltration checks;
- provenance-preserving transformations;
- signed and explainable deny decisions;
- adversarial conformance fixtures.
