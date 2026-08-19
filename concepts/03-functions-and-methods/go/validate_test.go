package validate

import "testing"

func TestRequireTitle(t *testing.T) {
	if err := RequireTitle(Task{Title: "x"}); err != nil {
		t.Errorf("RequireTitle(non-empty) = %v, want nil", err)
	}
	if err := RequireTitle(Task{}); err != ErrEmptyTitle {
		t.Errorf("RequireTitle(empty) = %v, want ErrEmptyTitle", err)
	}
}

// TestAllShortCircuits is the exercise: All must run validators in order
// and stop at the first error. Skipped until the TODO in validate.go is
// filled in.
func TestAllShortCircuits(t *testing.T) {
	t.Skip("TODO(exercise): implement All, then remove this Skip")

	var calls []string
	first := func(Task) error { calls = append(calls, "first"); return ErrEmptyTitle }
	second := func(Task) error { calls = append(calls, "second"); return nil }

	v := All(first, second)
	err := v(Task{Title: "x"})

	if err != ErrEmptyTitle {
		t.Errorf("All() = %v, want ErrEmptyTitle", err)
	}
	if len(calls) != 1 || calls[0] != "first" {
		t.Errorf("calls = %v, want short-circuit after [first]", calls)
	}
}

// TestAllPassesWhenEveryValidatorPasses covers the all-pass path.
// Skipped until the TODO in validate.go is filled in.
func TestAllPassesWhenEveryValidatorPasses(t *testing.T) {
	t.Skip("TODO(exercise): implement All, then remove this Skip")

	v := All(RequireTitle)
	if err := v(Task{Title: "x"}); err != nil {
		t.Errorf("All() = %v, want nil", err)
	}
}
