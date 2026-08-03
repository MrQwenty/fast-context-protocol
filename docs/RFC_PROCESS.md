# CGP RFC Process

Protocol evolution is documented through numbered RFCs in `docs/spec/`.

## When an RFC is required

An RFC is required for changes to:

- wire messages, fields, headers, media types, or endpoints;
- context-node identity and graph semantics;
- capability negotiation or versioning;
- conformance requirements;
- privacy, trust, authorization, or policy semantics;
- receipt, signature, deletion, or invalidation behavior;
- interoperability guarantees;
- compatibility or deprecation policy.

Implementation details that do not change observable protocol behavior usually do not require an RFC.

## States

An RFC may be:

- **Draft** — under active design and review;
- **Experimental** — implemented for interoperability testing but not stable;
- **Accepted** — approved as normative for a protocol version;
- **Superseded** — replaced by a newer RFC;
- **Withdrawn** — abandoned without becoming normative;
- **Deprecated** — still recognized but scheduled for removal under a compatibility policy.

## Numbering

Historical documents use `FCP-NNNN`. The project retains those identifiers to avoid breaking references while the public technical name moves to CGP. A future RFC may define a CGP-native numbering scheme.

## Proposal workflow

1. Open a Protocol Proposal issue.
2. Discuss problem, scope, prior art, and interoperability need.
3. Create a draft RFC in `docs/spec/`.
4. Add schemas, fixtures, and reference behavior where practical.
5. Collect implementation and security review.
6. Resolve compatibility and migration concerns.
7. Maintainers record the decision and RFC state.

## Review criteria

Reviewers evaluate:

- whether the problem belongs in the protocol;
- independent implementability;
- deterministic and testable behavior;
- complexity added to clients and providers;
- privacy and security consequences;
- failure and downgrade behavior;
- compatibility with deployed versions;
- performance and operational impact;
- regulatory claims and evidence requirements;
- whether a simpler extension or application-layer solution exists.

## Normative language

RFCs should use **MUST**, **MUST NOT**, **SHOULD**, **SHOULD NOT**, and **MAY** only for requirements that can be independently observed or tested.

## Acceptance

Acceptance requires maintainer approval and enough implementation evidence to establish that the proposal is coherent and testable. Significant RFCs should have at least two interoperable implementations before being considered stable.

## Changes after acceptance

Editorial corrections may be merged without changing semantics. Normative changes require a new RFC or an explicitly versioned revision with compatibility analysis.
