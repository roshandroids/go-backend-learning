---
slug: roadmap
title: Roadmap
role: milestones
updated: "2026-08-10T00:16:32"
---

# Roadmap

**Authoritative sequencing lives in `ROADMAP.md` (11 stages, WebSocket ladder
1-17) and, beyond this repo, in [[external-curriculum-docs]].** This page does
not restate that checklist — it tracks *actual demonstrated progress* against
it, which the checkboxes in `ROADMAP.md`/`README.md` don't yet reflect (both
are still 100% unchecked despite one exercise being done).

## Learning progression (concept-by-concept, evidence-based)

Scale: Not started → Planned → Learning → Practiced → Applied → Comfortable → Production-oriented.

| Concept | Status | Evidence |
|---|---|---|
| Go syntax & program structure | Practiced | `point.go` package + struct |
| Value vs pointer receivers | Practiced | [[value-vs-pointer-receivers]] — implemented + tested, one exercise variant still open |
| Types & structs (beyond receivers) | Planned | `concepts/02-types-and-structs/` is README-only |
| Functions & methods | Planned | `concepts/03-functions-and-methods/` is README-only |
| Interfaces | Planned | `concepts/04-interfaces/` is README-only |
| Error handling | Planned | `concepts/05-error-handling/` is README-only |
| Concurrency | Planned | `concepts/06-concurrency/` is README-only |
| HTTP servers | Planned | `concepts/07-http/` is README-only |
| REST API design | Planned | `concepts/08-rest-api-design/` is README-only |
| PostgreSQL | Planned | `concepts/09-postgresql/` is README-only |
| WebSockets | Planned | `concepts/10-websockets/` is README-only; also `projects/03-chat-server` |
| Testing (Go) | Learning | One test file exists (`point_test.go`); too little evidence yet to call this "Practiced" as a general skill |
| CI/CD | Not started (as a *learned* skill) | Workflows exist and run, but were scaffolded wholesale from a template, not authored/debugged by hand yet |
| Docker | Not started | No Dockerfile anywhere |
| Auth, GraphQL, Redis, distributed systems | Not started | Not present anywhere in repo; only appear in the external curriculum's later stages |

## Current state → next steps

1. **Current state:** Stage 1 (Go fundamentals) in progress. Point-struct
   receiver exercise done; its Scale/ScaleInPlace variant is an explicit open
   TODO from the 2026-08-09 journal.
2. **Immediate next:** finish the Scale/ScaleInPlace variant, then move to
   `concepts/02-types-and-structs`.
3. **Engineering objective in parallel:** `projects/01-cli` is the simplest
   unstarted project and doesn't depend on any unresolved decision — a
   reasonable next "build" target once Stage 1-3 concepts are covered, per its
   own README's stated scope.
4. **Open decision blocking Project 3/4:** WebSocket library choice
   (`gorilla/websocket` vs `nhooyr.io/websocket`) — see
   [[websocket-server-projects]]. Resolve before starting real code there.
5. **Known gap to fix eventually (not urgent):** [[ci-workflow-shallow-glob-gap]].

## What this brain will NOT do

It will not invent a roadmap beyond what's evidenced in `ROADMAP.md` and the
external curriculum docs, and it will not mark a concept "Comfortable" or
"Applied" without repository evidence (a project actually using it).
