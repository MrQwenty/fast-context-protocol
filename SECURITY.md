# Security Policy

Context Governance Protocol (CGP) is pre-alpha and must not yet be treated as a complete production security or compliance boundary.

## Supported versions

| Version | Supported |
|---|---|
| `master` pre-alpha | Best-effort security fixes |
| Tagged pre-alpha releases | Best effort until superseded |
| Older snapshots and forks | Not supported by the project |

A stable support window will be published before `v1.0.0`.

## Reporting a vulnerability

Report vulnerabilities privately through [GitHub private vulnerability reporting](https://github.com/MrQwenty/fast-context-protocol/security/advisories/new).

Do not disclose the issue in a public issue, pull request, discussion, social post, benchmark, or proof of concept before coordinated disclosure.

Include, where possible:

- affected commit or version;
- affected component;
- threat model and preconditions;
- reproducible steps using synthetic data;
- impact assessment;
- proposed remediation or mitigation;
- whether exploitation has been observed;
- preferred credit information.

Never include real secrets, personal data, customer documents, production credentials, or regulated datasets.

## Response process

The maintainers will attempt to:

1. acknowledge a complete report;
2. validate severity and affected versions;
3. coordinate remediation and regression tests;
4. prepare releases or mitigations;
5. agree on a disclosure date with the reporter where practical;
6. publish an advisory with credit, unless anonymity is requested.

There is currently no contractual response-time guarantee. Reports indicating active exploitation, cross-tenant leakage, remote code execution, credential exposure, or systematic privacy bypass receive the highest priority.

## Priority threat areas

High-priority areas include:

- authorization bypass;
- cross-tenant or cross-purpose context leakage;
- unsafe provider routing;
- cache poisoning and digest confusion;
- provenance or receipt forgery;
- replay and downgrade attacks;
- prompt-injection propagation and authority escalation;
- context exfiltration through tools or outputs;
- privacy-detector evasion and re-identification;
- reversible-vault key leakage;
- unsafe Office, PDF, archive, image, or OCR processing;
- path traversal and decompression bombs;
- malicious policy packs or signature bypass;
- denial of service through oversized or adversarial context;
- logging, telemetry, or error paths that expose protected values.

## Privacy gateway requirements

- Original documents must remain local unless an explicit, authorized policy permits otherwise.
- Unsupported or uninspectable documents must fail closed under fail-closed policies.
- Extraction should run in a sandbox without network access and with CPU, memory, file-size, recursion, and execution-time limits.
- Reversible mappings must use authenticated encryption and separate key storage.
- Logs, errors, receipts, telemetry, and traces must not contain original detected values.
- Pseudonymized content remains sensitive and must not be described as anonymous.
- Automated anonymization cannot guarantee that indirect or contextual re-identification is impossible.
- Detector models, dictionaries, policy packs, and trust roots are part of the trusted computing base and should be versioned, reviewed, and signed.

## Safe-harbor intent

Good-faith research that follows this policy, avoids privacy harm and service disruption, and provides reasonable time for remediation will not be treated by the project maintainers as malicious activity. This statement cannot bind third parties or replace applicable law.

## Public security documentation

Security architecture, threat models, hardening guidance, and resolved advisories should be documented publicly once disclosure is safe. Embargoed details remain private until coordinated disclosure.
