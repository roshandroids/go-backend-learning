/// Dart equivalent of go/notifier.go — but idiomatically Dart defines the
/// abstract class first, at the "producer," and implements it explicitly.
/// Go's version (notifier.go) defines the interface at the CONSUMER,
/// unexported, and neither implementation ever says "implements" anything.
abstract class Notifier {
  String notify(String message);
}

class EmailNotifier implements Notifier {
  @override
  String notify(String message) => 'emailed: $message';
}

class SMSNotifier implements Notifier {
  @override
  String notify(String message) => 'texted: $message';
}

/// Animal/Dog show the behavior Go's embedding deliberately does NOT
/// replicate: overriding speak() here DOES change what announce() calls,
/// because Dart has real dynamic dispatch through inheritance.
class Animal {
  final String name;
  Animal(this.name);

  String speak() => '$name makes a sound';

  String announce() => 'Announcing: ${speak()}'; // dynamically dispatches
}

class Dog extends Animal {
  Dog(super.name);

  @override
  String speak() => '$name barks';
}

// dog.announce() => "Announcing: Rex barks" in Dart (dynamic dispatch).
// Go's Dog{Animal{"Rex"}}.Announce() => "Announcing: Rex makes a sound"
// (no dynamic dispatch through embedding). This mismatch IS the trap.
