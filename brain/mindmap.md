---
slug: mindmap
title: Feature mindmap
role: feature mindmap
updated: "2026-08-10T00:16:14"
---

# Feature mindmap

```mermaid
mindmap
  root((go-backend-learning))
    Learning axis
      Go fundamentals
        Value vs pointer receivers ✅ practiced
        Types & structs — planned
        Functions & methods — planned
        Interfaces — planned
        Error handling — planned
        Concurrency — planned
      Backend fundamentals
        HTTP servers — planned
        REST API design — planned
        PostgreSQL — planned
        WebSockets — planned
      Dart to Go bridge
        Reference vs value semantics
        Each concept paired with a Dart equivalent
    Building axis
      Project 1 — CLI — stub
      Project 2 — REST API — stub
      Project 3 — Chat server — stub, open lib choice
      Project 4 — Flutter+Go chat — stub, depends on Project 3's Hub
    Governance
      concepts vs projects module split
      ADR 0001
      External curriculum docs (unversioned, one dir up)
```

The two axes — **learning** (concept-by-concept, in `concepts/`) and
**building** (project-by-project, in `projects/`) — are meant to reinforce each
other per the external curriculum's design: a concept is learned in isolation,
then later applied inside a project. As of now only the learning axis has any
real content (one concept), and the building axis is entirely unstarted.
