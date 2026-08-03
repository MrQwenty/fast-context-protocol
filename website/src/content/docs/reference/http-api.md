---
title: HTTP API
description: Current pre-alpha HTTP and JSON compatibility surface.
---

The current implementation still exposes the historical `fcp` compatibility namespace.

## Discovery

```http
GET /.well-known/fcp
```

## Resolve context

```http
POST /fcp/v0.1/context/resolve
FCP-Version: 0.1
Content-Type: application/json
```

The response media type is currently:

```text
application/fcp+json
```

## Fetch a node

```http
GET /fcp/v0.1/context/{sha256:digest}
```

## Submit a receipt

```http
POST /fcp/v0.1/receipts
```

## Compatibility direction

CGP aliases will be introduced additively and tested against the legacy namespace before any deprecation is considered. Historical `FCP-NNNN` specification identifiers remain stable references.
