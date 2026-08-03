---
title: EU governance direction
description: Technical controls and evidence aligned with European AI and data-governance obligations.
---

CGP is being designed as a **compliance accelerator and evidence system**, not an automatic legal certification service.

## Target profile

```yaml
profile: eu-safe

privacy:
  mode: anonymize
  fail_closed: true

providers:
  processing_region: eu-only
  retention: zero
  training_use: deny
  local_fallback: true

governance:
  classify_ai_act: true
  transparency_marking: true
  human_review: automatic
  audit_receipts: true
```

## Planned governance primitives

- AI Act role and risk classification support;
- prohibited-practice preflight;
- transparency payloads;
- meaningful human-oversight requirements;
- GDPR purpose and legal-basis binding;
- retention, restriction, objection and deletion propagation;
- EU/EEA provider-routing constraints;
- no-training and zero-retention requirements;
- signed, versioned regulatory policy packs;
- AI Bill of Materials;
- audit and incident evidence export;
- CI compliance simulation.

## Official legal sources

- [EU Artificial Intelligence Act — Regulation (EU) 2024/1689](https://eur-lex.europa.eu/eli/reg/2024/1689/oj/eng)
- [GDPR — Regulation (EU) 2016/679](https://eur-lex.europa.eu/eli/reg/2016/679/oj/eng)
- [European Commission AI Act policy portal](https://digital-strategy.ec.europa.eu/en/policies/regulatory-framework-ai)
- [EDPB: anonymisation and pseudonymisation](https://www.edpb.europa.eu/topics/ai-and-technology/anonymisation-pseudonymisation_en)
- [NIS2 — Directive (EU) 2022/2555](https://eur-lex.europa.eu/eli/dir/2022/2555/oj/eng)
- [Cyber Resilience Act — Regulation (EU) 2024/2847](https://eur-lex.europa.eu/eli/reg/2024/2847/oj/eng)

Applicability depends on the organization, role, use case, sector, jurisdiction, contracts and actual operational controls. Qualified legal review remains necessary.
