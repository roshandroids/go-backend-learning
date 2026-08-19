/// Dart equivalent of go/validation.go.
/// Dart's exceptions are unchecked and can cross many stack frames
/// invisibly; you can't tell from `validateAge`'s signature whether it
/// throws. Go's `error` return value makes that visible at every call site.
class ValidationException implements Exception {
  final String field;
  final String message;
  ValidationException(this.field, this.message);

  @override
  String toString() => '$field: $message';
}

void validateAge(int age) {
  if (age < 0) {
    throw ValidationException('age', 'must not be negative');
  }
  if (age > 150) {
    throw ValidationException('age', 'must be realistic');
  }
}

String describeError(Object error) {
  if (error is ValidationException) {
    return 'invalid ${error.field}: ${error.message}';
  }
  return 'unexpected error: $error';
}

// Usage:
// try {
//   validateAge(-1);
// } catch (e) {
//   print(describeError(e));
// }
