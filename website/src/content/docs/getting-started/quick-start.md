---
title: Quick start
description: Run the reference provider, resolve context and sanitize a document locally.
---

## Requirements

- Go 1.23 or newer
- optional `pdftotext` for local PDF extraction
- optional `tesseract` for local image OCR

## Clone and validate

```bash
git clone https://github.com/MrQwenty/fast-context-protocol.git
cd fast-context-protocol

go test -race ./...
go vet ./...
```

## Run the reference provider

```bash
go run ./cmd/fcpd \
  -listen :8080 \
  -catalog examples/basic-provider/context.json
```

The current compatibility namespace is still `fcp` while the public technical name migrates to CGP.

## Resolve context

```bash
go run ./cmd/fcpctl \
  -endpoint http://localhost:8080 \
  -intent code.review \
  -target pull-request:482 \
  -max-tokens 4000 \
  -max-latency-ms 80
```

## Sanitize a document locally

```bash
go run ./cmd/fcpprivacy \
  -input examples/privacy/sample.txt \
  -output /tmp/sample.sanitized.txt \
  -report /tmp/sample.privacy.json \
  -custom-terms examples/privacy/custom-terms.txt \
  -mode anonymize
```

The command emits sanitized content and a `PrivacyReceipt`. Original detected values are not included in the receipt.

## Validate a provider

```bash
go run ./cmd/fcpconform -endpoint http://localhost:8080
```

The conformance runner exits non-zero when baseline discovery, version negotiation, budget enforcement or error semantics do not conform.
