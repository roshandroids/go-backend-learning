---
id: ci-workflow-shallow-glob-gap
title: go-test.yml only checks one go.mod nesting level
category: reference
status: active
tags: [ci, gap]
created: "2026-08-10T00:11:15"
updated: "2026-08-10T00:12:13"
---

<!-- compiled_truth -->
**What exists:** `.github/workflows/go-test.yml`'s "Test each project module" step
loops `for d in projects/*/` and checks `$d/go.mod` — this only finds `go.mod`
one directory level below `projects/`.

**The gap:** `projects/04-flutter-go-chat/go_server/go.mod` is nested one level
deeper than that (`projects/04-flutter-go-chat/` itself has no `go.mod`; its
Go code lives in the `go_server/` subdirectory). The current loop's glob never
matches it, so **the Go server half of the flagship project 4 is silently never
tested in CI**, even after real code and tests are added to it.

**Status: known, not fixed.** This was discovered during a documentation/brain
audit, not a development task — fixing CI behavior was out of scope for that
work. The fix is straightforward when someone picks it up: replace the shallow
glob with a recursive find, e.g. `find projects -name go.mod -execdir go vet
./... \; -execdir go test -race ./... \;` (or equivalent), mirroring what a
correct depth-agnostic loop looks like.

**How to apply:** don't assume `go-test.yml` passing means `go_server/` was
tested — check whether this gap has since been fixed before trusting a green
CI run for project 4.


## Timeline

- time: 2026-08-10T00:11:15
  kind: decision
  summary: "Created this page: go-test.yml only checks one go.mod nesting level"
  source: .github/workflows/go-test.yml
  affects: [ci-workflow-shallow-glob-gap]

- time: 2026-08-10T00:12:13
  kind: decision
  summary: "Recorded a real CI gap found during audit; not fixed, per task scope (no behavior changes)"
  source: ".github/workflows/go-test.yml, audit 2026-08-10"
  affects: [ci-workflow-shallow-glob-gap]
