# Fast Context Protocol (FCP)

**FCP is an experimental, model-agnostic protocol for discovering, negotiating, compiling, transporting, updating, and verifying context for AI systems under explicit latency, token, freshness, privacy, and quality constraints.**

> Status: pre-alpha / FCP 0.1 core draft with experimental privacy extensions. The wire format and APIs are expected to change.

## Why FCP

Existing integration protocols make tools and resources available to AI applications. FCP focuses on the next problem: delivering the smallest sufficient, freshest, verifiable, and policy-compliant context for a specific inference or agent operation.

FCP is designed around:

- explicit token, byte, and latency budgets;
- content-addressed context nodes;
- progressive delivery ordered by utility;
- cache-aware references and delta updates;
- provenance, sensitivity, and freshness metadata;
- model-agnostic context compilation;
- separation between control plane and data plane;
- local, fail-closed document anonymization with verifiable privacy receipts.

FCP is not intended to replace tool execution protocols. It can sit beside them or adapt their resource outputs into a context graph.

## Repository layout

```text
cmd/fcpd/                 Minimal reference server
cmd/fcpctl/               Minimal command-line client
cmd/fcpconform/           Provider conformance runner
cmd/fcpprivacy/           Local document privacy gateway
internal/privacy/         Detection, extraction, transformation, vault, and leak scan
internal/protocol/        FCP wire types
internal/server/          Resolver and HTTP transport
spec/schema/              Core and privacy JSON Schemas
examples/basic-provider/  Example context catalogue
examples/privacy/         Privacy gateway fixtures
docs/spec/FCP-0001.md     Core protocol specification
docs/spec/FCP-0005.md     Privacy gateway and anonymization receipts
```

## Run the reference implementation

```bash
go test ./...
go run ./cmd/fcpd -listen :8080 -catalog examples/basic-provider/context.json
```

In another terminal:

```bash
go run ./cmd/fcpctl \
  -endpoint http://localhost:8080 \
  -intent code.review \
  -target pull-request:482 \
  -max-tokens 4000 \
  -max-latency-ms 80
```

## Sanitize a document before an external LLM

The FCP privacy gateway processes documents locally. The original binary never enters the FCP data plane. Text and OpenXML documents are inspected natively; PDF and image adapters invoke local `pdftotext` and `tesseract` installations. Unsupported or uninspectable documents fail closed rather than being forwarded unchanged.

```bash
go run ./cmd/fcpprivacy \
  -input examples/privacy/sample.txt \
  -output /tmp/sample.sanitized.txt \
  -report /tmp/sample.privacy.json \
  -custom-terms examples/privacy/custom-terms.txt \
  -mode anonymize
```

The command produces only sanitized text and a `PrivacyReceipt` containing input/output hashes, the applied policy, entity counts, residual risk, and pass/fail status. Receipts never contain detected original values.

Supported modes:

- `redact`: replace detected values with typed redaction markers;
- `pseudonymize`: generate stable scoped surrogates for internal workflows;
- `anonymize`: generate per-document unlinkable surrogates without retaining a mapping.

Stable pseudonyms require a local secret:

```bash
export FCP_PRIVACY_SECRET='replace-with-at-least-16-random-bytes'
go run ./cmd/fcpprivacy \
  -input contract.docx \
  -output contract.sanitized.txt \
  -mode pseudonymize \
  -scope workspace-42
```

Reversible pseudonymization additionally requires `-reversible`, a local `-vault` path, and `FCP_PRIVACY_VAULT_KEY`. The mapping is encrypted with AES-256-GCM and is never attached to context sent to a model.

Automated anonymization substantially reduces exposure but cannot prove that contextual or quasi-identifier re-identification is impossible. High-impact use cases require domain-specific detectors, adversarial evaluation, and potentially human review.

## Validate a provider

Run the baseline conformance checks against any FCP endpoint:

```bash
go run ./cmd/fcpconform -endpoint http://localhost:8080
```

The command emits a machine-readable JSON report and exits non-zero when discovery, version negotiation, budget enforcement, or error semantics are non-conforming.

## Lifecycle

```text
Local privacy preflight -> Discover -> Resolve -> Plan -> Fetch/Stream -> Receipt -> Patch/Invalidate
```

## Design principles

1. **Budget is part of the request.** Context selection must respect declared constraints.
2. **Identity is content-based.** Immutable content is addressed by a cryptographic digest.
3. **Metadata is first-class.** Freshness, provenance, sensitivity, privacy state, and transformations travel with content.
4. **Delivery is progressive.** Critical context arrives before supporting context.
5. **Reuse beats retransmission.** Consumers declare known context and receive references or deltas.
6. **Policy is enforceable.** Providers may exclude content that violates authorization or handling rules.
7. **Privacy precedes transport.** Raw documents are sanitized locally before an external model or provider receives context.
8. **Interoperability before cleverness.** The baseline uses ordinary HTTP and JSON before optimized transports.

## License

Apache License 2.0. See [LICENSE](LICENSE).
