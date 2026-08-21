/// Dart's web_socket_channel gives you a Stream/Sink pair for a WebSocket
/// connection — genuinely one of the closer Dart<->Go conceptual matches,
/// since a WebSocket connection IS naturally a duplex stream of messages
/// both sides. The gap is entirely on the server side: your Flutter
/// client manages exactly ONE connection; a Go server may be managing
/// thousands, and must never let one slow client block another — that's
/// what go/hub.go's backpressure select exists for, with nothing
/// equivalent needed here.
///
/// import 'package:web_socket_channel/web_socket_channel.dart';
///
/// final channel = WebSocketChannel.connect(Uri.parse('ws://localhost:8080/ws'));
/// channel.stream.listen(
///   (message) => print('received: $message'),
///   onDone: () => print('disconnected'),
///   onError: (e) => print('error: $e'),
/// );
/// channel.sink.add('hello');
