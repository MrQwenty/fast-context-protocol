---
title: Security reporting
description: How to report vulnerabilities and understand CGP's current security status.
---

CGP is pre-alpha and must not yet be treated as a complete production security or compliance boundary.

## Report privately

Use [GitHub private vulnerability reporting](https://github.com/MrQwenty/fast-context-protocol/security/advisories/new).

Do not create a public issue for vulnerabilities involving:

- context or tenant leakage;
- unsafe provider routing;
- receipt or provenance forgery;
- prompt-injection authority escalation;
- privacy-detector bypass;
- vault-key exposure;
- malicious document parsing;
- credential or personal-data exposure.

Use synthetic data in every reproduction.

## Security model

CGP aims to make the following boundaries explicit:

- original documents remain local under local-only privacy policies;
- unsupported documents fail closed under fail-closed policies;
- untrusted context cannot increase its own authority;
- provider routing must respect declared policy capabilities;
- receipts must be integrity-protected;
- sensitive values must not enter logs or telemetry.

Several of these controls remain roadmap work and are not production guarantees.

## Authoritative policy

Read the complete [Security Policy](https://github.com/MrQwenty/fast-context-protocol/blob/master/SECURITY.md).
