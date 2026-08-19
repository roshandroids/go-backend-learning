# Error Handling: error values, wrapping, errors.Is/As, panic vs error

## Dart concept
`validateAge` throws a `ValidationException` (unchecked, invisible from
the function signature); a `try`/`catch` at the call site inspects it
with `is ValidationException` to format a specific message, falling back
to a generic one otherwise.

## Dart implementation
See `dart/validation.dart`.

## Go equivalent
`ValidateAge` returns a `*ValidationError` (as the built-in `error`
interface) instead of throwing — visible in the function's own signature.
`DescribeError` uses `errors.As` to extract the concrete type, the closest
Go gets to `catch (e) { if (e is X) ... }`, but as an explicit call rather
than implicit stack unwinding.

## Go implementation
See `go/validation.go` and `go/validation_test.go`.

## Important differences
- A Go function's signature tells you it can fail (`error` return value)
  — there's no equivalent to Dart's invisible `throws`.
- No exception hierarchy to `catch` broadly — `errors.As` checks for one
  specific error type at a time.
- `panic`/`recover` is **not** a try/catch substitute. `ValidateAge`
  returning an error for bad input, rather than panicking, is the correct
  choice — panic is reserved for programmer bugs, not expected failure
  modes like "validation failed."

## Exercise (Level 2 — Complete)
Two `// TODO` gaps in `validation.go`:
1. `ValidateAge` — return a `*ValidationError` for out-of-range ages (see
   the TODO comment for exact field/message text the tests expect).
2. `DescribeError` — use `errors.As` to extract a `*ValidationError` and
   format it specially; fall back to a generic message otherwise.

Remove the `t.Skip(...)` calls in `validation_test.go` once both are
implemented and passing.
