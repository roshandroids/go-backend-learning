---
id: external-curriculum-docs
title: "External curriculum docs live outside this repo, unversioned"
category: reference
status: active
tags: [curriculum, roadmap, external]
created: "2026-08-10T00:11:14"
updated: "2026-08-10T00:11:57"
---

<!-- compiled_truth -->
**What exists outside this repo:** the parent workspace directory
(`Backend_stack/`, one level above `go-backend-learning/`) contains the actual
curriculum documents this repo's structure, roadmap, and journal templates were
generated from:

- `docs/learning/go-learning-roadmap.md` — full curriculum: 11 stages, 9 numbered
  projects, an explicit thesis that Go is "a reaction against" Flutter/Clean
  Architecture instincts, and a list of places where Flutter mental models
  actively mislead (over-interfacing, DI containers, inheritance-via-embedding,
  Streams≠channels, try/catch instincts, null-safety≠zero-values,
  mixins-as-generics).
- `docs/learning/go-mentorship-protocol.md` — the teaching loop this repo's
  exercises follow: Explain → Tiny Example → Predict → Implement → Break →
  Debug → Compare with Dart → Apply to Project, plus a 7-level exercise
  difficulty scale (Predict → Complete → Fix → Refactor → Design → Build →
  Production Failure).
- `docs/repo-strategy/go-repo-strategy.md` — the authoritative doc this repo's
  folder layout, commit conventions, branching, and version-tag scheme were
  built from; its "First 10 Commits" script matches this repo's git log almost
  commit-for-commit.
- `docs/archive/go-repo-strategy-draft-v1.md` — explicitly marked superseded,
  kept for history only.

**Why this is a risk worth recording:** the parent `Backend_stack/` directory is
**not a git repository** — none of these governing documents are version
controlled. `ROADMAP.md` in this repo links to them by relative path
(`../docs/learning/...`) but a future agent working from a clone of just
`go-backend-learning/` (e.g. on another machine, or from the GitHub remote)
will find those links dangling. This repo's own git history and README/ROADMAP
are self-sufficient for understanding *what was built*, but not for the full
stage sequencing or teaching methodology behind *why the curriculum is ordered
this way*.

**How to apply:** when reasoning about "what stage comes next" or "how should
an exercise be scaffolded," prefer reading the external docs directly if
available locally; do not assume they'll always be reachable, and do not copy
their content into this brain wholesale (they're the authoritative source —
duplicating them here would create a competing, driftable copy).


## Timeline

- time: 2026-08-10T00:11:14
  kind: decision
  summary: "Created this page: External curriculum docs live outside this repo, unversioned"
  source: "../docs/learning/, ../docs/repo-strategy/ (parent Backend_stack workspace)"
  affects: [external-curriculum-docs]

- time: 2026-08-10T00:11:57
  kind: decision
  summary: Flagged that the authoritative curriculum lives outside this repo and outside git
  source: "audit of parent Backend_stack/ workspace, 2026-08-10"
  affects: [external-curriculum-docs]
