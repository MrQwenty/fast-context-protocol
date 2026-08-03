## Summary

Describe what changed and why.

## Change type

- [ ] Bug fix
- [ ] Documentation
- [ ] Reference implementation
- [ ] Protocol or schema change
- [ ] Security or privacy hardening
- [ ] Performance or benchmark
- [ ] Build, CI, or maintenance

## Validation

List the commands and tests you ran.

```text
# Example
go test -race ./...
go vet ./...
```

## Protocol compatibility

- [ ] This change does not alter protocol behavior.
- [ ] This change alters protocol behavior and includes or updates an RFC.
- [ ] Backward-compatibility and migration impact are documented.
- [ ] Conformance fixtures or tests are included where applicable.

## Privacy and security

- [ ] No secrets, personal data, customer data, or private fixtures are included.
- [ ] Threats and abuse cases introduced by this change were considered.
- [ ] Logs, errors, receipts, and telemetry do not expose protected values.
- [ ] Fail-open versus fail-closed behavior is explicit.

## Documentation

- [ ] User-facing behavior is documented.
- [ ] The website or README is updated where necessary.
- [ ] Planned functionality is not described as already implemented.

## Checklist

- [ ] The change is focused and reviewable.
- [ ] New code is tested.
- [ ] Formatting and static checks pass.
- [ ] Commit messages are clear.
- [ ] I have read `CONTRIBUTING.md` and `CODE_OF_CONDUCT.md`.
