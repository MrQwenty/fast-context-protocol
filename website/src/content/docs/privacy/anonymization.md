---
title: Anonymization model
description: Understand redaction, pseudonymization, anonymization and residual risk.
---

## Redaction

Replaces a detected value with a typed marker:

```text
mario@example.com → [REDACTED:EMAIL]
```

## Pseudonymization

Creates a stable surrogate within a declared scope:

```text
Mario Rossi → <PERSON:8A03F1C99210>
```

Reversible mode stores the mapping in a local AES-256-GCM encrypted vault. Pseudonymized data remains linkable and must still be treated as sensitive.

## Per-document anonymization

Creates unlinkable tokens for a single document and does not retain a mapping:

```text
Mario Rossi → <PERSON_001_A18F2C>
```

## Residual risk

Removing direct identifiers does not guarantee that combinations such as age, profession, city, diagnosis and unusual events cannot identify a person. The receipt therefore exposes a residual-risk signal and pass/fail result rather than claiming unconditional legal anonymity.
