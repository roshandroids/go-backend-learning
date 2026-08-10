---
id: concepts-vs-projects-module-split
title: concepts/ single module vs projects/ per-module go.mod
category: decision
status: active
tags: [go, modules, architecture]
created: "2026-08-10T00:11:14"
updated: "2026-08-10T00:11:27"
---

<!-- compiled_truth -->
**Decision:** `concepts/` is one shared Go module (`concepts/go.mod`) covering all
numbered topic folders (`01-fundamentals` .. `10-websockets`). Each entry under
`projects/` gets its own independent `go.mod`.

**Reason:** Concept folders are throwaway teaching snippets that will never be
deployed — one module keeps `go test ./...` trivial across all of them. Projects
are real, independently runnable things (some later containerized) — separate
modules mirror how real services are structured and prevent one project's
dependency bump from touching another's `go.sum`.

**Rejected alternatives (from the ADR):**
- One repo-wide module — would force every concept folder to share every
  project's dependencies (Postgres driver, JWT lib, etc.) even when irrelevant.
- One module per concept folder (10+ modules) — unnecessary overhead for
  snippets this small.

**Learning vs production framing:** this is a **learning-repo-specific** structural
choice, not a production recommendation. A production monorepo would more likely
use a single module with internal package boundaries, or a proper workspace
(`go.work`); the per-project-module split here exists to teach dependency
isolation, not because it's the only correct way to structure a real service.

**Practical consequence:** there is no root `go.mod`. `go test ./...` from the
repo root fails. Must `cd` into `concepts/` or a `projects/*` dir first — this is
exactly what `.github/workflows/go-test.yml` does.

See also: [[ci-workflow-shallow-glob-gap]].


## Timeline

- time: 2026-08-10T00:11:14
  kind: decision
  summary: "Created this page: concepts/ single module vs projects/ per-module go.mod"
  source: docs/decisions/0001-repo-structure.md
  affects: [concepts-vs-projects-module-split]

- time: 2026-08-10T00:11:27
  kind: decision
  summary: "Captured ADR 0001 decision and rejected alternatives so future agents don't re-litigate module boundary"
  source: docs/decisions/0001-repo-structure.md
  affects: [concepts-vs-projects-module-split]
