# Functions & Methods: first-class functions, closures, methods

## Dart concept
A `typedef Validator = String? Function(Task)` and an `all(...)` function
that composes a list of validators into one, short-circuiting on the first
error — functions as values, same as Go.

## Dart implementation
See `dart/validate.dart`.

## Go equivalent
A named function type `Validator func(Task) error`, plain functions that
satisfy it with no declaration at the definition site, and an `All`
function that returns a **closure** composing them — the same shape
Stage 5's HTTP middleware (`func(http.Handler) http.Handler`) will reuse.

## Go implementation
See `go/validate.go` and `go/validate_test.go`.

## Important differences
- `Validator` is a type built from a function signature, not an interface
  — any function with a matching signature satisfies it automatically.
- `All`'s returned function is a closure: it captures `validators` from
  `All`'s parameter list and keeps referring to it on every call, long
  after `All` itself has returned — conceptually identical to a Dart
  closure capturing a variable from its enclosing scope.
- There's no framework/decorator syntax for composition — it's just a
  function that takes functions and returns a function. This is exactly
  how Go HTTP middleware works later (Stage 5); nothing new is introduced.

## Exercise (Level 2 — Complete)
`All` has a `// TODO` gap: implement it so it runs each `Validator` against
the `Task`, in order, returning the **first** non-nil error immediately
(short-circuit) — or `nil` if every validator passes. Remove the
`t.Skip(...)` calls in `validate_test.go` once implemented and passing.
