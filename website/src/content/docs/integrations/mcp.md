---
title: CGP and MCP
description: Use MCP for connectivity and CGP for context governance.
---

CGP is not intended to replace MCP tool and resource connectivity. It can consume MCP resources and govern the context compiled from them.

```text
MCP GitHub ─┐
MCP Notion ─┼──► CGP ─► Deduplicate ─► Sanitize ─► Verify
MCP DB ─────┤                                  │
Files/API ──┘                                  ▼
                                      Apply budget and policy
                                                │
                                                ▼
                                      Compile for the model
```

| Concern | MCP | CGP |
|---|---|---|
| Primary purpose | Connect models to tools and resources | Govern, optimize and prove context delivery |
| Tool execution | Core concern | Deliberately outside the protocol core |
| Resource discovery | Yes | Can consume it |
| Token/byte/latency contract | Host-defined | Protocol primitive |
| Local anonymization | Not a core responsibility | First-class gateway |
| Purpose and jurisdiction | Application responsibility | Governance target |
| Receipts | Implementation-specific | Core evidence abstraction |

**MCP:** What is available?  
**CGP:** What is the minimum safe and authorized context this model should receive now?
