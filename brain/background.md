---
slug: background
title: Project background
role: project background
updated: "2026-08-10T00:20:49"
---

# Project background

`go-backend-learning` is a **learning repository**, not a production service. It
documents a Flutter/Dart engineer's progression into backend development with
Go — real code, real tests, real commits, written incrementally rather than
copied from tutorials (per the repo's own `README.md`).

## What this repo actually is right now

The repo was scaffolded in a single day (2026-08-08: all 12 commits, per
`git log --date=short`) following an external strategy document (see
[[external-curriculum-docs]]). The most recent commit (`61f2477`) adds the
journal entry file `docs/journal/2026-08-09.md` — note the filename date
(2026-08-09) is one day after the commit's own git date (2026-08-08); this is
an existing inconsistency in the source data, not a brain error, and is called
out here so it isn't mistaken for a real second day of work. As of that
commit, exactly **one** learning exercise has real, tested code: value vs
pointer receivers on a `Point` struct — see [[value-vs-pointer-receivers]].
Everything else — nine more `concepts/` topic folders and all four `projects/`
apps — is directory scaffolding, `go.mod` files, and template files with
`TODO` stubs.

This is the expected state for a repo one day old, not a stall or a problem.
This brain's job is to keep that honest as work continues, rather than let
documentation drift ahead of (or behind) what's actually implemented.

## Two layers of "why"

1. **Curriculum why** — the stage/project sequencing, teaching methodology,
   and Flutter-vs-Go mental-model framing come from documents that live
   *outside* this repo, in the parent `Backend_stack/` workspace, and are not
   version controlled. See [[external-curriculum-docs]] for what they contain
   and the risk of relying on them.
2. **Repo-structure why** — decisions made specifically for *this* repo's
   layout (e.g. module boundaries) are recorded as brain decision pages, e.g.
   [[concepts-vs-projects-module-split]].

## Who this brain is for

A future Claude/Cursor session picking this repo back up, so it can answer
"what's actually been learned/built" without re-deriving it from a fresh
`git log` and file-by-file read every time — and so it doesn't mistake a
learning-only scaffolding choice for a production recommendation.
