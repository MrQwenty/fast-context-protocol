---
title: Contributing
description: How to contribute code, specifications, tests, documentation and research to CGP.
---

CGP welcomes contributions to the protocol, reference implementation, privacy gateway, conformance suite, documentation, integrations and research.

## Start with the correct channel

- **Bug:** use the Bug Report issue form.
- **Implementation feature:** use the Feature Request form.
- **Protocol semantics:** use the Protocol Proposal form.
- **Documentation:** use the Documentation form.
- **Security vulnerability:** report privately through GitHub Security Advisories.

## Local validation

```bash
gofmt -w $(find . -name '*.go' -type f)
go vet ./...
go test -race ./...
go build ./cmd/fcpd ./cmd/fcpctl ./cmd/fcpconform ./cmd/fcpprivacy
```

Documentation:

```bash
cd website
npm install
npm run build
```

## Protocol contributions

Changes to wire formats, semantics, capabilities, conformance, compatibility, privacy state, trust boundaries or receipts require an RFC in `docs/spec/`.

An RFC must explain:

- the interoperability problem;
- normative behavior;
- compatibility and migration;
- security and privacy implications;
- conformance tests;
- rejected alternatives.

## Data safety

Never submit real credentials, personal data, customer documents, production logs or proprietary fixtures. Use synthetic examples and sanitize every attachment.

## Project policies

- [Full contributing guide](https://github.com/MrQwenty/fast-context-protocol/blob/master/CONTRIBUTING.md)
- [Governance](https://github.com/MrQwenty/fast-context-protocol/blob/master/GOVERNANCE.md)
- [Code of Conduct](https://github.com/MrQwenty/fast-context-protocol/blob/master/CODE_OF_CONDUCT.md)
- [RFC process](https://github.com/MrQwenty/fast-context-protocol/blob/master/docs/RFC_PROCESS.md)
