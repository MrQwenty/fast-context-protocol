# Roadmap

## 0.1 — Core draft

- [x] protocol scope and terminology
- [x] discovery document
- [x] context request and plan
- [x] token and byte budgets
- [x] content-addressed nodes
- [x] inline, reference, and fetch delivery
- [x] minimal Go provider and client
- [x] initial JSON Schema
- [x] conformance runner baseline
- [ ] canonical fixture suite
- [ ] interoperable second implementation

## Privacy extension — FCP-0005

- [x] local-only privacy processing contract
- [x] redaction, scoped pseudonymization, and per-document anonymization
- [x] encrypted reversible vault using AES-256-GCM
- [x] structured and contextual sensitive-data detectors
- [x] organization-specific dictionaries and allow lists
- [x] post-transform leak scan and fail-closed policy gate
- [x] PrivacyReceipt, policy digest, and input/output integrity digests
- [x] UTF-8 and OpenXML document extraction
- [x] local PDF extraction and image OCR adapters
- [x] privacy smoke test in CI
- [ ] multilingual local NER detector plugin
- [ ] quasi-identifier generalization and re-identification risk benchmark
- [ ] domain policy profiles for healthcare, legal, finance, HR, and government
- [ ] sandboxed extraction worker with resource limits

## 0.2 — Context synchronization

- [ ] context patches and invalidation
- [ ] subscriptions and resumable streams
- [ ] tokenizer profiles
- [ ] signed manifests and verification
- [ ] cross-provider plan composition

## 0.3 — Ecosystem

- [ ] MCP resource adapter
- [ ] TypeScript SDK
- [ ] Python SDK
- [ ] TTFC and token-efficiency benchmark
- [ ] public RFC process
