package notify

// Animal/Dog demonstrate the embedding trap: embedding is composition,
// NOT inheritance. There is no dynamic dispatch — a method defined on
// Animal that calls Animal's own Speak() will always call Animal's
// Speak(), even when an outer type embedding Animal defines its own
// Speak() with the same name.
type Animal struct {
	Name string
}

func (a Animal) Speak() string {
	return a.Name + " makes a sound"
}

// Announce is defined on Animal and calls a.Speak() on itself. Announce
// has no idea Dog exists, or that Dog "overrides" Speak — that's the crux
// of the trap.
func (a Animal) Announce() string {
	return "Announcing: " + a.Speak()
}

// Dog embeds Animal (NOT "extends" — there's no such keyword) and defines
// its own Speak(). This looks like an override from a Dart/OO instinct,
// but it does not participate in dynamic dispatch anywhere inside
// Animal's own methods.
type Dog struct {
	Animal
}

func (d Dog) Speak() string {
	return d.Name + " barks"
}
