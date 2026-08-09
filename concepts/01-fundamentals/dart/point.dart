import 'dart:math' as math;

/// Dart equivalent of go/point.go's Point type.
/// Unlike Go, this is a reference type by default — any method that
/// mutates a field mutates the single shared instance, no pointer needed.
class Point {
  final int x;
  final int y;

  Point(this.x, this.y);

  double distanceTo(Point other) {
    final dx = (x - other.x).toDouble();
    final dy = (y - other.y).toDouble();
    return math.sqrt(dx * dx + dy * dy);
  }
}

/// A mutable variant to mirror Go's pointer-receiver Translate.
class MutablePoint {
  int x;
  int y;

  MutablePoint(this.x, this.y);

  void translate(int dx, int dy) {
    x += dx;
    y += dy;
  }
}
