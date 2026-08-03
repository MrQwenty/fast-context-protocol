# Ctxora Roadmap

Ctxora is the working public name of the project originally developed as Fast Context Protocol (FCP). Checkmarks indicate repository artifacts, not production readiness or regulatory certification.

## 0.1 — Core protocol

- [x] protocol scope and terminology
- [x] discovery document
- [x] context request and plan
- [x] token and byte budgets
- [x] content-addressed nodes
- [x] inline, reference, and fetch delivery
- [x] minimal Go provider and client
- [x] initial JSON Schema
- [x] baseline conformance runner
- [ ] canonical fixture suite
- [ ] independent interoperable implementation

## Privacy — FCP-0005

- [x] local-only privacy contract
- [x] redaction, scoped pseudonymization, and per-document anonymization
- [x] AES-256-GCM reversible vault
- [x] custom dictionaries and allow lists
- [x] residual leak scan and fail-closed gate
- [x] PrivacyReceipt and integrity digests
- [x] text, OpenXML, local PDF, and local OCR extraction
- [x] privacy smoke test in CI
- [ ] multilingual local NER
- [ ] quasi-identifier generalization
- [ ] re-identification benchmark
- [ ] sector policy profiles
- [ ] sandboxed extraction workers

## 0.2 — Synchronization and integrity

- [ ] context patches and invalidation
- [ ] subscriptions and resumable priority streams
- [ ] tokenizer profiles
- [ ] signed manifests and receipts
- [ ] cross-provider plan composition
- [ ] dependency-aware deletion propagation

## 0.3 — Simple integration

- [ ] MCP resource adapter
- [ ] OpenAI-compatible local proxy
- [ ] Anthropic, Gemini, Azure OpenAI, and Ollama adapters
- [ ] `ctxora init --profile eu-safe`
- [ ] `ctxora run -- <application>`
- [ ] `ctxora doctor`, `ctxora explain`, and `ctxora audit`
- [ ] TypeScript, Python, and public Go SDKs
- [ ] TTFC, TTUC, token-efficiency, and privacy benchmarks
- [ ] public RFC process

## 0.4 — Trust and governance

- [ ] FCP-0006 EU Compliance Envelope
- [ ] FCP-0008 Context Trust Firewall and taint tracking
- [ ] FCP-0009 purpose, retention, and data-rights propagation
- [ ] FCP-0010 AI Bill of Materials and evidence export
- [ ] FCP-0011 signed regulatory policy packs
- [ ] FCP-0012 transparency and human oversight
- [ ] FCP-0013 compliance simulator
- [ ] FCP-0014 EU-only provider routing
- [ ] FCP-0015 explainable policy decisions

These controls are intended to accelerate compliance and collect evidence. They must not be presented as automatic legal certification.

## Naming migration

See [`docs/NAMING.md`](docs/NAMING.md).

- [x] adopt Ctxora as the working public brand
- [x] preserve `fcp` wire names and binaries during pre-alpha
- [x] document the compatibility strategy
- [ ] add the `ctxora` CLI as an alias
- [ ] add `/.well-known/ctxora` and `/ctxora/v0.x` aliases
- [ ] add `Ctxora-Version` compatibility headers
- [ ] publish a versioned migration RFC
- [ ] complete trademark, domain, package, and organization-name clearance
- [ ] rename the repository only after redirect and package behavior are verified

No existing namespace will be removed through an undocumented breaking change.

## Stable-release gates

- independent interoperability tests
- canonical conformance corpus
- threat model and independent security review
- fuzzing and adversarial document tests
- production authorization and tenant isolation
- durable receipt and deletion semantics
- reproducible performance and privacy benchmarks
- documented key management and incident response
- formal review of regulatory claims
- backward-compatibility policy
