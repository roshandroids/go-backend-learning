---
id: value-vs-pointer-receivers
title: "Go value vs pointer receivers (first implemented concept)"
category: concept
status: active
tags: [go, fundamentals, receivers]
created: "2026-08-10T00:11:14"
updated: "2026-08-10T00:11:41"
---

<!-- compiled_truth -->
**What exists:** `concepts/01-fundamentals/go/point.go` defines a `Point` struct
with a method using a value receiver (`DistanceTo`, non-mutating) and a method
using a pointer receiver (`Translate`, mutates in place). Covered by
`point_test.go` (`TestDistanceTo`, `TestTranslateMutatesInPlace`). A Dart
equivalent lives alongside it at `concepts/01-fundamentals/dart/point.dart`.

**Why this matters (the Dart contrast):** in Dart, objects are always reference
types, so "does this mutate the original" is never ambiguous. In Go, the same
question depends on whether a method's receiver is a value (operates on a copy)
or a pointer (operates on the original) — this is the first concept in the
curriculum that has no clean Dart analogue and requires new intuition.

**Status: Practiced, not yet Comfortable.** One exercise done end-to-end with a
passing test, but the assigned Scale/ScaleInPlace variant (exercising the same
value-vs-pointer choice from a different angle) is still an open TODO for the
next session per the 2026-08-09 journal entry. Do not mark this concept
"Comfortable" or "Applied" until that variant lands, and it hasn't yet appeared
used anywhere in `projects/`.

**This is the only entry in [[concepts-vs-projects-module-split]]'s `concepts/`
module with real implemented code** — the other nine topic folders
(`02-types-and-structs` through `10-websockets`) are README-only scaffolds with
empty `.gitkeep`-held `go/`/`dart/` directories.


## Timeline

- time: 2026-08-10T00:11:14
  kind: decision
  summary: "Created this page: Go value vs pointer receivers (first implemented concept)"
  source: "concepts/01-fundamentals/go/point.go, docs/journal/2026-08-09.md"
  affects: [value-vs-pointer-receivers]

- time: 2026-08-10T00:11:41
  kind: decision
  summary: "Recorded the repo's first and only implemented learning concept, with honest status"
  source: "concepts/01-fundamentals/go/point.go, concepts/01-fundamentals/go/point_test.go, docs/journal/2026-08-09.md"
  affects: [value-vs-pointer-receivers]
