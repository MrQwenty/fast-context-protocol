# Security Policy

FCP is pre-alpha and must not yet be used as a security boundary in production.

Please report vulnerabilities privately through GitHub's security advisory feature once enabled for the repository. Do not include secrets, customer data, personal data, or exploit details in public issues.

The highest-priority areas are authorization bypass, cross-tenant context leakage, cache poisoning, digest confusion, provenance forgery, replay, prompt-injection propagation, privacy detector evasion, re-identification, reversible-vault key leakage, unsafe document parsing, malicious Office/PDF payloads, OCR confusion, and denial of service through oversized or adversarial context.

## Privacy gateway requirements

- Original documents must remain local.
- Unsupported or uninspectable documents must fail closed.
- Extraction should run in a sandbox with no network access and constrained CPU, memory, file size, and execution time.
- Reversible mappings must use authenticated encryption and separate key storage.
- Logs, errors, receipts, telemetry, and traces must not contain original detected values.
- Pseudonymized content remains sensitive and must not be described as anonymous.
- Automated anonymization cannot guarantee that indirect or contextual re-identification is impossible.
- Detector models, dictionaries, and policies are part of the trusted computing base and should be versioned and signed.
