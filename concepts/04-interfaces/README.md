# Interfaces: implicit satisfaction, consumer-defined interfaces, embedding

## Dart concept
An `abstract class Notifier` defined up front, implemented explicitly by
`EmailNotifier`/`SMSNotifier` (`implements Notifier`); and an
`Animal`/`Dog extends Animal` pair where overriding `speak()` really does
change what `announce()` calls, via Dart's real dynamic dispatch.

## Dart implementation
See `dart/notify.dart`.

## Go equivalent
Two exercises:
1. **`notifier`** — an interface defined at the *consumer* (`Alert`),
   unexported, narrow. `EmailNotifier`/`SMSNotifier` satisfy it implicitly
   — no `implements` keyword, no declaration anywhere that they do.
2. **`Animal`/`Dog`** — embedding, and the trap: `Dog` embeds `Animal` and
   defines its own `Speak()`, but this does *not* override `Animal`'s
   `Speak()` from the inside. `Animal.Announce()` calls `a.Speak()` on
   itself and has no idea `Dog` exists.

## Go implementation
See `go/notifier.go` + `go/notifier_test.go`, and `go/embed.go` +
`go/embed_test.go`.

## Important differences
- Go's proverb: *accept interfaces, return structs*. `notifier` exists
  only because `Alert` needs it — there's no exported `Notifier` interface
  sitting next to its only implementations "for testability."
- Interface satisfaction is checked structurally, at compile time,
  wherever a value is used as the interface — never declared at the type's
  own definition.
- Embedding promotes fields/methods syntactically (`dog.Name`, `dog.Speak()`
  both work) but has **no dynamic dispatch** — there is no Go equivalent to
  Dart's `@override` changing behavior seen from inside the base type's own
  methods. If you need real polymorphism, you need an interface, not
  embedding.

## Exercises
**Exercise 1 (Level 2 — Complete):** `Alert` has a `// TODO` gap — call
`Notify` on each notifier in order, returning the first error. Remove the
`t.Skip(...)` calls in `notifier_test.go` once implemented.

**Exercise 2 (Level 1 — Predict):** before running `embed_test.go`,
predict what `d.Announce()` returns for `Dog{Animal{Name: "Rex"}}` — does
it use `Dog`'s `Speak()` or `Animal`'s? This one's already implemented;
the exercise is the prediction, not the code.
