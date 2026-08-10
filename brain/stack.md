---
slug: stack
title: Tech stack
role: tech-stack choices
updated: "2026-08-10T00:15:53"
---

# Tech stack

## In use today

- **Go 1.22** — every `go.mod` in the repo (`concepts/`, all four `projects/*`)
  pins `go 1.22`. All five modules currently have **zero dependencies** —
  standard library only.
- **Module paths are still placeholders** — every `go.mod` reads
  `github.com/YOUR_USERNAME/go-backend-learning/...`; the actual GitHub remote
  is `roshandroids/go-backend-learning`. Not yet reconciled.
- **Dart/Flutter** — `concepts/*/dart/` mirrors each Go concept for comparison
  (only `01-fundamentals` has real content); `projects/04-flutter-go-chat/flutter_app`
  is a Flutter app scaffold (`flutter_go_chat`, Flutter SDK `>=3.0.0 <4.0.0`).
- **CI** — GitHub Actions only: `gofmt -l .` and `go vet && go test -race`
  per module. No linter beyond `gofmt` (no `golangci-lint` config exists).

## Named in READMEs but not yet added to any go.mod (aspirational, not installed)

- `projects/03-chat-server`, `projects/04-flutter-go-chat/go_server`: either
  `github.com/gorilla/websocket` or `nhooyr.io/websocket` — explicitly an
  open either/or, see [[websocket-server-projects]].
- `projects/04-flutter-go-chat/go_server`: `database/sql` + a Postgres driver,
  `golang-jwt` for auth.
- `projects/04-flutter-go-chat/flutter_app/pubspec.yaml` already declares
  `web_socket_channel: ^2.4.0` and `uuid: ^4.0.0` — these ARE in the dependency
  manifest, just unused in code so far (`lib/main.dart` is a placeholder screen).

## Not present at all

Docker, docker-compose, any database, Redis, GraphQL, any auth implementation,
observability tooling. Don't assume any of these exist just because they're
mentioned in the external curriculum roadmap ([[external-curriculum-docs]]) or
in a project README's "Technologies" section — those describe *intent*, not
current state.
