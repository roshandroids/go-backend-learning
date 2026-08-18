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

// TestScaleDoesNotMutate proves the value-receiver Scale is the wrong
// choice: the caller's original Point is unchanged after the call.
// Skipped until the exercise TODO in point.go is filled in.
func TestScaleDoesNotMutate(t *testing.T) {
	t.Skip("TODO(exercise): implement Scale, then remove this Skip")

	p := Point{X: 2, Y: 3}
	p.Scale(2)

	if p.X != 2 || p.Y != 3 {
		t.Errorf("Scale() mutated original = %+v, want unchanged {2 3}", p)
	}
}

// TestScaleInPlaceMutates proves the pointer-receiver fix works.
// Skipped until the exercise TODO in point.go is filled in.
func TestScaleInPlaceMutates(t *testing.T) {
	t.Skip("TODO(exercise): implement ScaleInPlace, then remove this Skip")

	p := Point{X: 2, Y: 3}
	p.ScaleInPlace(2)

	if p.X != 4 || p.Y != 6 {
		t.Errorf("ScaleInPlace() = %+v, want {4 6}", p)
	}
}
