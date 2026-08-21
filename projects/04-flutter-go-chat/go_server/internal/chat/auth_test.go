package chat

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var testSecret = []byte("test-secret-do-not-use-in-prod")

func signToken(t *testing.T, secret []byte, sub string, expiresAt time.Time) string {
	t.Helper()
	claims := jwt.MapClaims{
		"sub": sub,
		"exp": expiresAt.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(secret)
	if err != nil {
		t.Fatalf("signing test token: %v", err)
	}
	return signed
}

// TestValidateTokenAcceptsValidToken is the exercise's happy path.
// Skipped until the TODO in auth.go is filled in.
func TestValidateTokenAcceptsValidToken(t *testing.T) {
	t.Skip("TODO(exercise): implement ValidateToken, then remove this Skip")

	token := signToken(t, testSecret, "user-42", time.Now().Add(time.Hour))

	userID, err := ValidateToken(token, testSecret)
	if err != nil {
		t.Fatalf("ValidateToken() error = %v, want nil", err)
	}
	if userID != "user-42" {
		t.Errorf("userID = %q, want %q", userID, "user-42")
	}
}

// TestValidateTokenRejectsBadSignature proves a token signed with a
// different secret is rejected. Skipped until the TODO is filled in.
func TestValidateTokenRejectsBadSignature(t *testing.T) {
	t.Skip("TODO(exercise): implement ValidateToken, then remove this Skip")

	wrongSecret := []byte("a-completely-different-secret")
	token := signToken(t, wrongSecret, "user-42", time.Now().Add(time.Hour))

	if _, err := ValidateToken(token, testSecret); err != ErrInvalidToken {
		t.Errorf("ValidateToken() error = %v, want ErrInvalidToken", err)
	}
}

// TestValidateTokenRejectsExpiredToken proves an expired token is
// rejected, not silently accepted. Skipped until the TODO is filled in.
func TestValidateTokenRejectsExpiredToken(t *testing.T) {
	t.Skip("TODO(exercise): implement ValidateToken, then remove this Skip")

	token := signToken(t, testSecret, "user-42", time.Now().Add(-time.Hour))

	if _, err := ValidateToken(token, testSecret); err != ErrInvalidToken {
		t.Errorf("ValidateToken() error = %v, want ErrInvalidToken", err)
	}
}

// TestValidateTokenRejectsGarbage proves a non-JWT string doesn't panic
// and is rejected cleanly. Skipped until the TODO is filled in.
func TestValidateTokenRejectsGarbage(t *testing.T) {
	t.Skip("TODO(exercise): implement ValidateToken, then remove this Skip")

	if _, err := ValidateToken("not-a-jwt-at-all", testSecret); err != ErrInvalidToken {
		t.Errorf("ValidateToken() error = %v, want ErrInvalidToken", err)
	}
}
