/// Dart equivalent of the error-classification idiom in go/errors.go.
/// The closest Dart analogue to database/sql is Drift (compile-time-
/// checked SQL, typed queries), or sqflite's raw string queries at the
/// lower level — but neither gives you a stdlib-blessed sentinel error
/// for "no rows" the way database/sql's sql.ErrNoRows does; you'd throw
/// or return null instead.
class TaskNotFoundException implements Exception {
  final String id;
  TaskNotFoundException(this.id);
}

/// classifyError's Dart shape: Drift's `getSingleOrNull()` already
/// returns null for "not found" rather than throwing a sentinel — so
/// this translation step barely exists in Dart. That's exactly the
/// contrast worth noticing: Go's error-as-value design makes you handle
/// "not found" explicitly at every call site; Drift/sqflite let you
/// forget to null-check just as easily as forgetting a try/catch.
Object? classifyError(String id, Object? error) {
  if (error == null) return null;
  // A raw sqflite query might throw a generic DatabaseException here
  // instead of a typed "no rows" signal — there's no equivalent to
  // errors.Is(err, sql.ErrNoRows) to branch on.
  return error;
}
