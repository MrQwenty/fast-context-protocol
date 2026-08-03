# Context Governance Protocol Roadmap

Context Governance Protocol (CGP) is the public technical name of the project initially developed as Fast Context Protocol (FCP). Checkmarks indicate repository artifacts, not production readiness or regulatory certification.

## 0.1 — Core protocol

- [x] protocol scope and terminology
- [x] discovery document
- [x] context request and plan
- [x] token and byte budgets
- [x] content-addressed nodes
- [x] inline, reference and fetch delivery
- [x] minimal Go provider and client
- [x] initial JSON Schema
- [x] baseline conformance runner
- [ ] canonical fixture suite
- [ ] independent interoperable implementation

## Privacy — FCP-0005

- [x] local-only privacy contract
- [x] redaction, scoped pseudonymization and per-document anonymization
- [x] AES-256-GCM reversible vault
- [x] custom dictionaries and allow lists
- [x] residual leak scan and fail-closed gate
- [x] PrivacyReceipt and integrity digests
- [x] text, OpenXML, local PDF and local OCR extraction
- [x] privacy smoke test in CI
- [ ] multilingual local NER
- [ ] quasi-identifier generalization
- [ ] re-identification benchmark
- [ ] sector policy profiles
- [ ] sandboxed extraction workers

## Documentation and adoption

- [x] framework-style documentation architecture
- [x] landing page and concept guides
- [x] privacy, governance, security and MCP documentation
- [x] API, receipt, naming and project-status reference
- [x] GitHub Pages build and deployment workflow
- [ ] live Pages deployment after repository setting is enabled
- [ ] versioned documentation per protocol release
- [ ] searchable API schema explorer
- [ ] multilingual documentation

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
- [ ] Anthropic, Gemini, Azure OpenAI and Ollama adapters
- [ ] `cgp init --profile eu-safe`
- [ ] `cgp run -- <application>`
- [ ] `cgp doctor`, `cgp explain` and `cgp audit`
- [ ] TypeScript, Python and public Go SDKs
- [ ] TTFC, TTUC, token-efficiency and privacy benchmarks
- [ ] public RFC process

## 0.4 — Trust and governance

- [ ] FCP-0006 EU Compliance Envelope
- [ ] FCP-0008 Context Trust Firewall and taint tracking
- [ ] FCP-0009 purpose, retention and data-rights propagation
- [ ] FCP-0010 AI Bill of Materials and evidence export
- [ ] FCP-0011 signed regulatory policy packs
- [ ] FCP-0012 transparency and human oversight
- [ ] FCP-0013 compliance simulator
- [ ] FCP-0014 EU-only provider routing
- [ ] FCP-0015 explainable policy decisions

These controls are intended to accelerate compliance and collect evidence. They must not be represented as automatic legal certification.

## Naming migration

See [`docs/NAMING.md`](docs/NAMING.md).

- [x] adopt Context Governance Protocol as the public technical name
- [x] reject the Ctxora working name
- [x] preserve `fcp` wire names and binaries during pre-alpha
- [x] document the compatibility strategy
- [ ] add the `cgp` CLI as an alias
- [ ] add `/.well-known/cgp` and `/cgp/v0.x` aliases
- [ ] add `CGP-Version` compatibility headers
- [ ] publish a versioned migration RFC
- [ ] complete trademark, domain, package and organization-name clearance
- [ ] rename the repository only after redirects and package behavior are verified

No existing namespace will be removed through an undocumented breaking change.

## Open-source model

- [x] Apache-2.0 reference implementation
- [x] open protocol, schema and conformance direction
- [x] documented open-core commercial boundary
- [ ] contributor governance and maintainer policy
- [ ] trademark and conformance-mark policy
- [ ] security disclosure and release-signing process
- [ ] independent implementation grants

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
