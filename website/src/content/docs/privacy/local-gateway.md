---
title: Local Privacy Gateway
description: Transform documents before external inference or storage.
---

Sensitive content should be transformed **before** it reaches an external model, vector store, telemetry system or remote log.

```text
Original file
    │
    ▼
Local extraction / OCR
    │
    ▼
Discard metadata and unsupported embedded content
    │
    ▼
Detector ensemble
    │
    ├── redact
    ├── pseudonymize
    └── anonymize
    │
    ▼
Independent residual leak scan
    │
    ├── PASS → sanitized text + PrivacyReceipt
    └── FAIL → block external transmission
```

## Supported inputs

The current reference gateway can process:

- UTF-8 text, Markdown, CSV, JSON, XML, YAML and HTML;
- DOCX, PPTX and XLSX through local OpenXML inspection;
- PDF through a local `pdftotext` adapter;
- images through local `tesseract` OCR.

Unsupported or insufficiently inspectable inputs can fail closed.

## Custom detection

Organization-specific dictionaries and allow lists can identify project names, client names and internal identifiers that generic detectors cannot know.
