# Contributing

FCP is currently a design-stage protocol. Changes should preserve a narrow core and include an interoperability rationale.

## Development

```bash
gofmt -w .
go vet ./...
go test -race ./...
```

## Protocol changes

Substantive protocol changes should add or update an RFC in `docs/spec/`. Each proposal should state the problem, wire-level behavior, compatibility impact, security implications, and rejected alternatives.

## Commit style

Use focused commits with imperative subjects, for example:

```text
spec: define receipt semantics
server: enforce sensitivity policy
```
