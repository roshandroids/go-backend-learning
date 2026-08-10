---
id: websocket-server-projects
title: "Project 3 + 4: WebSocket chat server, standalone then Flutter-integrated"
category: project
status: active
tags: [websockets, projects, open-decision]
created: "2026-08-10T00:11:15"
updated: "2026-08-10T00:15:02"
---

<!-- compiled_truth -->
**Goal:** Project 3 builds a standalone Go WebSocket chat server (Hub pattern,
rooms, backpressure, heartbeats — WebSocket challenge ladder rungs 1-12).
Project 4 extends Project 3's Hub with persistence (Postgres) and JWT auth, and
pairs it with a Flutter client (`flutter_app/`) — ladder rungs 1-17 including a
later Redis-backed multi-instance stage.

**Sequencing constraint (stated intent, not yet exercised):** Project 3 is meant
to prove the Hub pattern out standalone before Flutter integration happens in
Project 4 — i.e. don't start wiring the Flutter client until 03-chat-server's
Hub, rooms, and backpressure handling work on their own.

**Current status: both are stubs.** `projects/03-chat-server/cmd/server/main.go`
and `projects/04-flutter-go-chat/go_server/cmd/server/main.go` are each a single
`fmt.Println("TODO: ...")`; `internal/hub/` (project 3) and `internal/` (project
4's go_server) are empty, held open by `.gitkeep`. `flutter_app/lib/main.dart` is
a bare `MaterialApp` with placeholder text. No dependencies are in either
go.mod yet.

**Open decision — not yet made:** project 3's README lists the WebSocket
library as `github.com/gorilla/websocket` **or** `nhooyr.io/websocket` — this is
an explicit either/or in the source, not a settled choice. Whichever is picked
for Project 3 should carry forward into Project 4's `go_server` (same Hub
lineage). Resolve this before writing real code in 03-chat-server, not
per-project.

**Numbering inconsistency worth knowing about:** `flutter_app/pubspec.yaml`
describes itself as "Flutter client for ... Project 7," while the folder name
and README call it Project 4. Two numbering schemes coexist in the repo —
folder/README numbering (`01`-`04`) is the one this brain and `ROADMAP.md`
follow.


## Timeline

- time: 2026-08-10T00:11:15
  kind: decision
  summary: "Created this page: Project 3 + 4: WebSocket chat server, standalone then Flutter-integrated"
  source: "projects/03-chat-server/README.md, projects/04-flutter-go-chat/README.md"
  affects: [websocket-server-projects]

- time: 2026-08-10T00:15:02
  kind: decision
  summary: "Captured shared scope, sequencing constraint, and the still-open websocket library choice"
  source: "projects/03-chat-server/README.md, projects/04-flutter-go-chat/README.md, ROADMAP.md"
  affects: [websocket-server-projects]
