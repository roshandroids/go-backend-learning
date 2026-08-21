package chat

import "errors"

// ErrInvalidToken covers every validation failure — bad signature,
// expired, malformed. The WS handshake caller only needs to know
// "reject the connection," not which specific jwt-library error fired.
var ErrInvalidToken = errors.New("invalid token")

// ValidateToken parses and validates a JWT signed with secret (HS256),
// returning the "sub" claim as the authenticated user ID. This is the
// WS-handshake auth check for WS ladder rung 14 — once wired into
// projects/03's ServeWS, a middleware-equivalent step validates the
// token from the request (query param or Authorization header) before
// upgrading, then the resulting user ID would be injected into
// context.Context for downstream code to read — the same
// context-propagation idiom as concepts/07-http, applied to auth
// instead of logging.
//
// TODO(exercise, Level 2 — Complete): use jwt.ParseWithClaims (from
// github.com/golang-jwt/jwt/v5) with jwt.MapClaims and a keyfunc that
// returns secret (reject any signing method other than HS256 in the
// keyfunc). On any parse/validation error, return "", ErrInvalidToken
// — don't leak the underlying jwt-library error to the caller.
// Otherwise extract the "sub" claim as a string; if it's missing or not
// a string, also return "", ErrInvalidToken. On success, return the
// user ID and a nil error.
func ValidateToken(tokenString string, secret []byte) (string, error) {
	// TODO
	return "", nil
}
