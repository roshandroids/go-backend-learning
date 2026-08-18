/// Dart equivalent of go/store.go's Store.
/// There is no footgun to demonstrate here: Dart's Map and List fields are
/// never "unset" the way a Go zero-value map is nil — either they're
/// initialized in the constructor (as below) or the field is nullable and
/// you're forced to null-check it. Go gives you a THIRD state Dart doesn't
/// have: a non-nil, non-empty-literal, perfectly typed map that still
/// panics on write until you initialize it.
class Task {
  final String id;
  final String title;
  final bool done;

  Task(this.id, this.title, {this.done = false});
}

class Store {
  final Map<String, Task> _tasks = {}; // always initialized, never null

  void add(Task t) => _tasks[t.id] = t;

  Task? get(String id) => _tasks[id];

  List<String> titles() => _tasks.values.map((t) => t.title).toList();
}
