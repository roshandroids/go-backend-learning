---
slug: flow
title: Key flows
role: key flows
updated: "2026-08-10T00:16:03"
---

# Key flows

There is no running data/request flow anywhere in this repo yet — every
`main.go` and `main.dart` is a placeholder that prints or renders a `TODO`
string. This page will start being meaningful once a project moves past stub
stage; until then, recording a flow here would be fiction.

## First flow expected

`projects/01-cli` (CLI transaction summarizer) is the simplest project and the
most likely candidate for the repo's first real flow:

```mermaid
graph LR
    A["CSV file path (arg)"] --> B["parse CSV"]
    B --> C["compute summary"]
    C --> D["print to stdout"]
```

This is inferred from the stated intent in `projects/01-cli/README.md` and
`main.go`'s TODO comment, not from any implemented code. Update this page with
the *actual* flow once `01-cli` has real logic — don't leave the inferred
diagram standing in as if it were built.
