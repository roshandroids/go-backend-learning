/// Dart equivalent of go/pagination.go — the same lesson you already
/// know from infinite-scroll pagination in a Flutter ListView, just
/// enforced server-side here instead of client-side.
class Task {
  final String id;
  final String title;
  Task(this.id, this.title);
}

class Page {
  final List<Task> tasks;
  final String nextCursor; // '' means no more pages
  Page(this.tasks, this.nextCursor);
}

class Store {
  final List<Task> _tasks = [];

  void add(Task t) => _tasks.add(t);

  /// Cursor-based, same idea as a Flutter ListView requesting "give me
  /// the next N items after the last one I rendered" — never "give me
  /// items 40-60," which is what offset pagination does and which
  /// silently breaks if items are inserted/removed between requests.
  Page list(String cursor, int limit) {
    var start = 0;
    if (cursor.isNotEmpty) {
      final idx = _tasks.indexWhere((t) => t.id == cursor);
      if (idx != -1) start = idx + 1;
    }
    final end = (start + limit).clamp(0, _tasks.length);
    final page = _tasks.sublist(start, end);
    final nextCursor = end < _tasks.length ? page.last.id : '';
    return Page(page, nextCursor);
  }
}
