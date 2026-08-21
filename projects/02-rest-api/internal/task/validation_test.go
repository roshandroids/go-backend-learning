package task

import "testing"

// TestValidateCreateRequest is table-driven, same idiom as
// concepts/05-error-handling. Skipped until the TODO in validation.go
// is filled in.
func TestValidateCreateRequest(t *testing.T) {
	t.Skip("TODO(exercise): implement validateCreateRequest, then remove this Skip")

	tests := []struct {
		name    string
		title   string
		wantErr bool
	}{
		{"valid title", "buy milk", false},
		{"empty title", "", true},
		{"whitespace-only title", "   ", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateCreateRequest(CreateRequest{Title: tt.title})
			if (err != nil) != tt.wantErr {
				t.Errorf("validateCreateRequest(%q) error = %v, wantErr %v", tt.title, err, tt.wantErr)
			}
		})
	}
}
