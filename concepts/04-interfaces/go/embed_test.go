package notify

import "testing"

// TestDogAnnounceUsesAnimalSpeak is a Level 1 (Predict) exercise, already
// implemented so you can check your prediction by running it — no TODO
// gap here. Before reading the assertions: does Dog's Speak() override
// Animal's from inside Announce()? (No — embedding gives no dynamic
// dispatch. See embed.go for why.)
func TestDogAnnounceUsesAnimalSpeak(t *testing.T) {
	d := Dog{Animal{Name: "Rex"}}

	if got, want := d.Speak(), "Rex barks"; got != want {
		t.Errorf("d.Speak() = %q, want %q", got, want)
	}

	// The trap: Announce is defined on Animal and calls a.Speak() on
	// ITSELF — it has no idea Dog embeds it or "overrode" Speak.
	if got, want := d.Announce(), "Announcing: Rex makes a sound"; got != want {
		t.Errorf("d.Announce() = %q, want %q (embedding has no dynamic dispatch)", got, want)
	}
}
