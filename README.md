# Fast Context Protocol (FCP)

**FCP is an experimental, model-agnostic protocol for discovering, negotiating, compiling, transporting, updating, and verifying context for AI systems under explicit latency, token, freshness, privacy, and quality constraints.**

> Status: pre-alpha / FCP 0.1 design draft. The wire format and APIs are expected to change.

## Why FCP

Existing integration protocols make tools and resources available to AI applications. FCP focuses on the next problem: delivering the smallest sufficient, freshest, and verifiable context for a specific inference or agent operation.

FCP is designed around:

- explicit token, byte, and latency budgets;
- content-addressed context nodes;
- progressive delivery ordered by utility;
- cache-aware references and delta updates;
- provenance, sensitivity, and freshness metadata;
- model-agnostic context compilation;
- separation between control plane and data plane.

FCP is not intended to replace tool execution protocols. It can sit beside them or adapt their resource outputs into a context graph.

## Repository layout

```text
cmd/fcpd/                 Minimal reference server
cmd/fcpctl/               Minimal command-line client
internal/protocol/        FCP 0.1 Go types
internal/server/          Resolver and HTTP transport
spec/schema/              JSON Schema
examples/basic-provider/  Example context catalogue
docs/spec/FCP-0001.md     Initial protocol specification
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

## FCP 0.1 lifecycle

```text
Discover -> Resolve -> Plan -> Fetch/Stream -> Receipt -> Patch/Invalidate
```

The current reference implementation covers discovery, resolution, planning, inline delivery, cached references, and receipts. Streaming, subscriptions, signatures, and binary framing are specified as planned extensions.

## Design principles

1. **Budget is part of the request.** Context selection must respect declared constraints.
2. **Identity is content-based.** Immutable content is addressed by a cryptographic digest.
3. **Metadata is first-class.** Freshness, provenance, sensitivity, and transformations travel with content.
4. **Delivery is progressive.** Critical context arrives before supporting context.
5. **Reuse beats retransmission.** Consumers declare known context and receive references or deltas.
6. **Policy is enforceable.** Providers may exclude content that violates authorization or handling rules.
7. **Interoperability before cleverness.** FCP 0.1 uses ordinary HTTP and JSON before adding optimized transports.

## License

Apache License 2.0. See [LICENSE](LICENSE).
