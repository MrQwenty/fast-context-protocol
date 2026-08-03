# Support

Context Governance Protocol is currently pre-alpha. Community support is provided on a best-effort basis.

## Before asking for help

1. Read the [documentation](https://mrqwenty.github.io/fast-context-protocol/).
2. Search existing issues for the same problem.
3. Run the relevant checks:

```bash
go test -race ./...
go vet ./...
```

For documentation development:

```bash
cd website
npm install
npm run build
```

## Where to ask

- **Reproducible bug:** open a Bug Report issue.
- **Protocol or feature proposal:** open a Protocol Proposal or Feature Request.
- **Documentation problem:** open a Documentation issue.
- **General usage question:** open a Question / Support issue.
- **Security vulnerability:** do **not** open a public issue; follow [`SECURITY.md`](SECURITY.md).

## What support does not include

The maintainers cannot provide legal opinions, compliance certification, production security guarantees, or private implementation consulting through public issues.

Commercial support, managed policy updates, certification, and service-level agreements may be introduced separately. GitHub Sponsors supports continued open-source maintenance but does not create a support contract.

## Response expectations

There is currently no guaranteed response time. Clear reproductions, minimal examples, logs with secrets removed, and proposed fixes significantly improve the chance of a useful response.
