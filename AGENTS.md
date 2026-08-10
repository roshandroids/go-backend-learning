<!-- BEGIN brain.md -->
## Project Brain

This project keeps a **Project Brain**: a persistent memory layer of its durable decisions, requirements, and constraints. Read `./BRAIN.md` for the full read/write contract.

Use it actively:
- Before any task or discussion, load the relevant brain context with the `brain` CLI's read commands.
- Whenever a decision, requirement, constraint, or durable insight surfaces — in discussion or in code — record it with the `brain` CLI before moving on; don't wait to be asked.
- All reads and writes go through the `brain` CLI — never hand-edit brain files.

The brain skills (`brain-setup`, `brain-page`, `brain-ingest`, `brain-bootstrap`) are installed in your global skills directory.
<!-- END brain.md -->

## Project Brain workflow for this repo

This repo is a **learning project**, not a production service. The brain's job is to
preserve *why* things are the way they are and *how far the learning has progressed*,
so a future session doesn't re-derive context or re-litigate settled calls.

**Before substantial work:**
1. Read `BRAIN.md`, then `brain read-root background`, `brain read-root architecture`,
   `brain read-root roadmap` for current state and direction.
2. `brain list-pages` and read any page relevant to the task (e.g. a concept folder,
   a project module, a rejected approach).
3. Decide whether the task is **learning-oriented** (a `concepts/*` exercise, exploring
   an idea) or **engineering-oriented** (building out a `projects/*` module toward
   something runnable). Don't propose production-grade abstractions for a learning
   exercise, and don't leave a `projects/*` module at exercise-quality.
4. Cross-check any brain claim about "what exists" against the actual source/tests —
   the brain records intent and reasoning, not a live mirror of the code.

**During development:**
- Do not update the brain for routine code changes (a function written, a test added
  that doesn't change direction).
- Do update the brain when: a concept moves to a new learning stage (e.g. planned to
  practiced), an architectural decision is made or changed, an approach is explicitly
  rejected, or a stated assumption turns out to be wrong.
- Always distinguish, in any new brain content: **current truth** (what the code does
  today) vs **reason** (why, often "chosen to learn X") vs **production consideration**
  (what a real system would additionally need) — do not collapse these into one
  recommendation.
- `ROADMAP.md` and the per-project READMEs remain the source of truth for curriculum
  sequencing and per-project status checklists; the brain reconciles with them, it does
  not fork a competing roadmap.

**Before finishing a work session:**
1. Run the relevant Go checks: `gofmt -l .` from repo root; `go vet ./... && go test -race ./...`
   from inside `concepts/` and inside each `projects/*` module with a `go.mod` (there is
   no root `go.mod` — root-level `go test ./...` will fail).
2. If durable knowledge changed, write it via the `brain` CLI (never hand-edit).
3. `brain reindex` and `brain lint-links` (the pre-commit hook also runs these).
4. Review what changed in `brain/` before committing, same as reviewing any other diff.
