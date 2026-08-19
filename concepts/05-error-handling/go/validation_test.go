package validation

import (
	"errors"
	"testing"
)

// TestValidateAge is table-driven — structurally identical to a
// parameterized flutter_test group, just without framework magic.
// Skipped until the TODO in validation.go is filled in.
func TestValidateAge(t *testing.T) {
	t.Skip("TODO(exercise): implement ValidateAge, then remove this Skip")

	tests := []struct {
		name    string
		age     int
		wantErr bool
	}{
		{"valid adult age", 30, false},
		{"negative age", -1, true},
		{"unreasonably large age", 200, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateAge(tt.age)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateAge(%d) error = %v, wantErr %v", tt.age, err, tt.wantErr)
			}
		})
	}
}

// TestDescribeErrorExtractsField proves errors.As pulls the field name
// out of a *ValidationError. Skipped until the TODO is filled in.
func TestDescribeErrorExtractsField(t *testing.T) {
	t.Skip("TODO(exercise): implement ValidateAge and DescribeError, then remove this Skip")

	err := ValidateAge(-1)
	got := DescribeError(err)
	want := "invalid age: must not be negative"

	if got != want {
		t.Errorf("DescribeError() = %q, want %q", got, want)
	}
}

// TestDescribeErrorFallsBackForOtherErrors proves the non-ValidationError
// path still produces a sane message. Skipped until the TODO is filled in.
func TestDescribeErrorFallsBackForOtherErrors(t *testing.T) {
	t.Skip("TODO(exercise): implement DescribeError, then remove this Skip")

	got := DescribeError(errors.New("boom"))
	want := "unexpected error: boom"

	if got != want {
		t.Errorf("DescribeError() = %q, want %q", got, want)
	}
}
