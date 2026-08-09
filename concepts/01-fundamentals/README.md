# 01 — Go Fundamentals: Structs, Methods, Receivers

## Dart concept
A `Point` class with an immutable `distanceTo` method and, separately, a
mutable class for in-place updates — Dart objects are always reference
types, so "does this mutate the original" is never ambiguous.

## Dart implementation
See `dart/point.dart`.

## Go equivalent
A `Point` struct with two methods that demonstrate the one Go concept with
no Dart analogue at all: **value vs pointer receivers**.
`DistanceTo` uses a value receiver (operates on a copy); `Translate` uses a
pointer receiver (mutates the original).

## Go implementation
See `go/point.go` and `go/point_test.go`.

## Important differences
- Go structs are value types by default — passing one to a function copies it.
- There is no constructor keyword; `Point{X: 1, Y: 2}` is the whole story.
- Export is capitalization (`Point`, `X`, `DistanceTo`), not a keyword.
- Choosing value vs pointer receiver is a real design decision in Go with
  no Dart equivalent to reason from directly.

## Exercise (Level 2 — Complete)
Add a `Scale(factor float64)` method to `Point` using a **value** receiver
(intentionally the wrong choice), write a test proving the mutation does
NOT persist on the caller's original value, then add a `ScaleInPlace`
using a pointer receiver that does persist. Compare both in a short note
at the bottom of this README.
