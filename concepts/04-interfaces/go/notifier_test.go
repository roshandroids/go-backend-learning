package notify

import (
	"errors"
	"testing"
)

type stubNotifier struct {
	err error
}

func (s stubNotifier) Notify(message string) error { return s.err }

// TestEmailAndSMSSatisfyNotifierImplicitly compiles ONLY because both
// types structurally match the notifier interface — neither declares
// "implements notifier" anywhere. This is the paradigm shift: Go checks
// this at compile time, from usage, not from an explicit declaration.
func TestEmailAndSMSSatisfyNotifierImplicitly(t *testing.T) {
	var _ notifier = EmailNotifier{}
	var _ notifier = SMSNotifier{}
}

// TestAlertReturnsFirstError is the exercise. Skipped until the TODO in
// notifier.go is filled in.
func TestAlertReturnsFirstError(t *testing.T) {
	t.Skip("TODO(exercise): implement Alert, then remove this Skip")

	wantErr := errors.New("boom")
	err := Alert("hi", stubNotifier{}, stubNotifier{err: wantErr}, stubNotifier{})

	if err != wantErr {
		t.Errorf("Alert() = %v, want %v", err, wantErr)
	}
}

// TestAlertPassesWhenAllSucceed covers the all-succeed path. Skipped
// until the TODO in notifier.go is filled in.
func TestAlertPassesWhenAllSucceed(t *testing.T) {
	t.Skip("TODO(exercise): implement Alert, then remove this Skip")

	err := Alert("hi", stubNotifier{}, stubNotifier{})
	if err != nil {
		t.Errorf("Alert() = %v, want nil", err)
	}
}
