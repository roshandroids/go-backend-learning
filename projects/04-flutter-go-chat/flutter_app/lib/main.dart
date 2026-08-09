// Stub entry point for the Flutter+Go chat client.
// TODO: connect to ws://localhost:8080/ws (WebSocket challenge ladder #1),
// following the roadmap's Project 7 spec.
import 'package:flutter/material.dart';

void main() {
  runApp(const ChatApp());
}

class ChatApp extends StatelessWidget {
  const ChatApp({super.key});

  @override
  Widget build(BuildContext context) {
    return const MaterialApp(
      home: Scaffold(
        body: Center(child: Text('TODO: chat UI')),
      ),
    );
  }
}
