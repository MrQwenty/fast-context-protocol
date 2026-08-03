# Contributing to CGP

Thank you for contributing to **Context Governance Protocol (CGP)**.

CGP is a pre-alpha protocol and reference implementation. Contributions must preserve interoperability, privacy, security, and factual accuracy. Planned features must not be presented as implemented, and technical controls must not be presented as automatic legal certification.

By participating, you agree to follow [`CODE_OF_CONDUCT.md`](CODE_OF_CONDUCT.md).

## Before contributing

- Search existing issues and pull requests.
- Read the [documentation](https://mrqwenty.github.io/fast-context-protocol/).
- For substantial work, open an issue before implementation.
- Report vulnerabilities privately under [`SECURITY.md`](SECURITY.md).
- Never submit secrets, customer data, personal data, production documents, or proprietary test fixtures.

## Ways to contribute

Contributions include:

- bug fixes;
- protocol design and RFC review;
- interoperability fixtures;
- privacy and security hardening;
- conformance tests;
- performance benchmarks;
- documentation and examples;
- SDKs and integrations;
- issue triage and community support;
- regulatory-source corrections and evidence-format research.

## Development setup

### Go core

Requirements:

- Go 1.23 or newer;
- optional `pdftotext` for PDF extraction tests;
- optional `tesseract` for OCR tests.

Run:

```bash
gofmt -w $(find . -name '*.go' -type f)
go vet ./...
go test -race ./...
go build ./cmd/fcpd ./cmd/fcpctl ./cmd/fcpconform ./cmd/fcpprivacy
```

### Documentation website

Requirements:

- Node.js 22 or newer.

Run:

```bash
cd website
npm install
npm run dev
npm run build
```

## Choosing the correct contribution path

### Implementation change

Open a Bug Report or Feature Request and submit a focused pull request.

### Protocol change

Open a Protocol Proposal. Changes to wire formats, semantics, capability negotiation, conformance, trust boundaries, privacy state, or compatibility require an RFC in `docs/spec/`.

### Security or privacy vulnerability

Use GitHub private vulnerability reporting. Do not create a public issue or public proof of concept before coordinated disclosure.

### Regulatory documentation

Use primary sources from EU institutions, supervisory authorities, standards bodies, or official national authorities. Clearly distinguish:

- binding law;
- delegated or implementing acts;
- official guidance;
- standards;
- project interpretation;
- planned controls.

## RFC requirements

An RFC must include:

1. problem statement;
2. goals and non-goals;
3. terminology;
4. normative behavior using clear MUST, SHOULD, and MAY language;
5. wire examples and error semantics;
6. compatibility impact;
7. privacy, security, and abuse analysis;
8. operational and regulatory considerations;
9. conformance requirements;
10. rejected alternatives;
11. migration and rollout plan.

The governance process is defined in [`GOVERNANCE.md`](GOVERNANCE.md).

## Pull requests

Pull requests should:

- solve one coherent problem;
- include tests for new behavior;
- preserve existing compatibility unless an approved RFC says otherwise;
- update schemas, fixtures, and documentation together;
- state fail-open or fail-closed behavior explicitly;
- avoid unrelated formatting or refactoring;
- complete the pull-request template.

Draft pull requests are welcome for early technical feedback.

## Tests and quality gates

At minimum, run:

```bash
gofmt -w $(find . -name '*.go' -type f)
go vet ./...
go test -race ./...
```

Protocol changes should also add:

- positive conformance cases;
- malformed and adversarial cases;
- backward-compatibility cases;
- deterministic fixtures where applicable;
- privacy and security regression tests.

Documentation changes should run:

```bash
cd website
npm run build
```

## Commit messages

Use focused imperative messages with a component prefix:

```text
protocol: define receipt verification semantics
privacy: block uninspectable OpenXML relationships
server: enforce declared token budget
docs: explain provider routing constraints
ci: add conformance smoke test
```

## Developer Certificate of Origin

By submitting a contribution, you certify that you have the right to submit it under the project's Apache-2.0 license and that the contribution does not knowingly include third-party material without compatible licensing and attribution.

A formal DCO bot may be introduced before stable releases. Until then, maintainers may request a `Signed-off-by` line for substantial or externally sponsored contributions.

## Review and acceptance

Maintainers may request changes for correctness, maintainability, interoperability, safety, evidence quality, scope, or compatibility. Acceptance is not guaranteed solely because a contribution is technically functional.

## Getting help

See [`SUPPORT.md`](SUPPORT.md) and use the repository's structured issue forms.
