# Architecture

## Components

- **Provider catalogue** stores typed context nodes and immutable content digests.
- **Resolver** filters by intent, target, policy, confidence, and freshness, then ranks candidates by utility per token.
- **Planner** chooses inline, reference, or fetch delivery while enforcing token and byte budgets.
- **HTTP transport** exposes discovery, resolve, fetch, receipt, and health endpoints.
- **Consumer CLI** demonstrates protocol negotiation and plan retrieval.

## Trust boundaries

Context content is untrusted. Protocol metadata must never allow source content to alter authorization, routing, delivery mode, or resolver policy. Production providers should isolate tenant caches, authenticate fetches, sign manifests where appropriate, and avoid metadata disclosure for unauthorized nodes.

## Future modules

- graph indexes and semantic retrieval;
- provider federation;
- tokenizer profiles;
- streaming and cancellation;
- delta codecs;
- signed provenance;
- policy engines;
- conformance test suite and benchmark corpus.
