/// Dart equivalent of go/validate.go.
/// Functions-as-values and closures work the same way conceptually — the
/// difference shows up later, in Stage 5, when Go middleware composes
/// http.Handler the same way `all` composes Validators here, with no
/// framework/decorator syntax involved.
class Task {
  final String title;
  Task(this.title);
}

/// A function type alias — Dart's equivalent of Go's `type Validator func(...)`.
typedef Validator = String? Function(Task);

String? requireTitle(Task t) => t.title.isEmpty ? 'title must not be empty' : null;

/// Closure composing multiple validators, short-circuiting on the first error.
Validator all(List<Validator> validators) {
  return (Task t) {
    for (final v in validators) {
      final err = v(t);
      if (err != null) return err;
    }
    return null;
  };
}
