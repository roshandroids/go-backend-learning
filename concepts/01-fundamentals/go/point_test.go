package fundamentals

import "testing"

func TestDistanceTo(t *testing.T) {
	a := Point{X: 0, Y: 0}
	b := Point{X: 3, Y: 4}

	got := a.DistanceTo(b)
	want := 5.0

	if got != want {
		t.Errorf("DistanceTo() = %v, want %v", got, want)
	}
}

func TestTranslateMutatesInPlace(t *testing.T) {
	p := Point{X: 1, Y: 1}
	p.Translate(2, 3)

	if p.X != 3 || p.Y != 4 {
		t.Errorf("Translate() = %+v, want {3 4}", p)
	}
}
