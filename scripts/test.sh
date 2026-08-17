#!/usr/bin/env bash
# Runs the checks CLAUDE.md asks for before ending a work session, across
# every Go module in the repo (there is no root go.mod by design — see ADR
# 0001 / brain page concepts-vs-projects-module-split).
set -euo pipefail
cd "$(dirname "$0")/.."

echo "== gofmt =="
files=$(gofmt -l .)
if [ -n "$files" ]; then
  echo "$files"
  echo "gofmt found unformatted files" >&2
  exit 1
fi

echo "== concepts =="
(cd concepts && go vet ./... && go test -race ./...)

find projects -name go.mod | while read -r mod; do
  d=$(dirname "$mod")
  echo "== $d =="
  (cd "$d" && go vet ./... && go test -race ./...)
done

echo "all checks passed"
