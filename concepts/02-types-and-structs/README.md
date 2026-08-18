# Types & Structs: zero values, arrays vs slices, maps

## Dart concept
A `Task` class with a `Map<String, Task>` field, always initialized in the
constructor — Dart gives you no way to have a typed `Map` that exists but
silently panics on write.

## Dart implementation
See `dart/store.dart`.

## Go equivalent
A `Store` struct wrapping a `map[string]Task`, demonstrating **zero values**
for both structs and maps, and the asymmetry between them: a nil map is
always safe to *read*, never safe to *write*, while a nil slice is safe for
both reading and `append`.

## Go implementation
See `go/store.go` and `go/store_test.go`.

## Important differences
- `Task{}` is immediately usable the moment it's declared — no constructor,
  no "late" field, no null-check required. This is true of every Go struct.
- A struct field of map type has a zero value of `nil`, not an empty map
  literal — unlike Dart, where `Map<String, Task> _tasks = {}` is always a
  real, empty, writable map from the constructor onward.
- Reading a nil map (`m[key]`) is safe and returns the zero value with
  `ok: false` — but writing to it (`m[key] = v`) panics at runtime. Slices
  don't have this asymmetry: `append` on a nil slice works fine.
- There's no Dart equivalent to this "read-safe, write-unsafe" middle state
  — Dart's collections are either genuinely initialized or genuinely `null`.

## Exercise (Level 2 — Complete)
`Store.Add` has a `// TODO` gap: fix it so a zero-value `Store` (`var s
Store`, no constructor call) can have `Add` called on it without panicking,
by lazily initializing the nil map on first write. Remove the `t.Skip(...)`
in `TestAddOnZeroValueStore` once it's implemented and passing.
