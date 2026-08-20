import 'dart:async';

/// Dart equivalent of go/pipeline.go — but the philosophical gap is the
/// point, not the syntax. Dart's isolates share NO memory (message
/// passing only); Go's goroutines share memory by default and need
/// explicit synchronization. There is no Dart bug class equivalent to a
/// Go data race on a shared map/slice, because Dart's event loop is
/// single-threaded per isolate.
Future<List<int>> squareAll(List<int> values) {
  // Future.wait is the closest Dart analogue to Go's fan-out/fan-in via
  // sync.WaitGroup — but each Future here just runs on the same event
  // loop, cooperatively, not on a separate OS thread.
  final futures = values.map((v) => Future(() => v * v));
  return Future.wait(futures);
}

/// Cancellation via a Stream subscription — Dart's closest analogue to
/// context.Context cancellation. Note this is NOT the same idea as a
/// channel: cancelling a subscription just stops delivery to THIS
/// listener; it doesn't propagate a cancellation signal into whatever
/// produced the stream, the way ctx.Done() propagates forward through an
/// explicit call chain in Go.
void cancellationSketch() {
  final controller = StreamController<int>();
  final subscription = controller.stream.listen((value) {
    print('got $value');
  });

  // ... later, e.g. when a screen is disposed:
  subscription.cancel();
}
