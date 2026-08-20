/// You've only ever been an HTTP CLIENT in Dart (Dio/http calling a
/// backend) — this file sketches the mirror image using package:shelf,
/// Dart's closest equivalent to net/http, just to show the shape is the
/// same: middleware as function composition, not a widget tree.
///
/// import 'package:shelf/shelf.dart';
///
/// Middleware loggingMiddleware() {
///   return (Handler innerHandler) {
///     return (Request request) async {
///       final start = DateTime.now();
///       final response = await innerHandler(request);
///       print('${request.method} ${request.url} ${DateTime.now().difference(start)}');
///       return response;
///     };
///   };
/// }
///
/// Response healthHandler(Request request) =>
///     Response.ok('{"status":"ok"}', headers: {'content-type': 'application/json'});
///
/// The mental model closest to what you already know isn't Dio — it's a
/// Flutter Widget's build(): a pure function from input (Request) to
/// output (Response). Composing shelf Middleware is literal function
/// wrapping, same as Go's `func(http.Handler) http.Handler` — no
/// decorators, no annotations, no DI container involved either side.
