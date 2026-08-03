---
title: Repository layout
description: Find the protocol, server, privacy gateway, schemas and specifications.
---

```text
cmd/fcpd/                  Reference provider
cmd/fcpctl/                Command-line client
cmd/fcpconform/            Provider conformance runner
cmd/fcpprivacy/            Local document privacy gateway
internal/privacy/          Detection, extraction, transformation and vault
internal/protocol/         Wire types
internal/server/           Resolver and HTTP transport
spec/schema/               Core and privacy JSON Schemas
examples/basic-provider/   Example context catalogue
examples/privacy/          Privacy fixtures
docs/spec/FCP-0001.md      Core protocol draft
docs/spec/FCP-0005.md      Privacy gateway draft
website/                   Framework-style documentation site
```
