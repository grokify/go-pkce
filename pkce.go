package pkce

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

const (
	LengthDefault = 32
	LengthMin     = 32
	LengthMax     = 96

	MethodPlain = "plain"
	MethodS256  = "S256"

	ParamCodeVerifier        = "code_verifier"
	ParamCodeChallenge       = "code_challenge"
	ParamCodeChallengeMethod = "code_challenge_method"
)

func NewCodeVerifier(n int) (string, error) {
	if n < 0 {
		n = LengthDefault
	}
	if n < LengthMin || n > LengthMax {
		return "", fmt.Errorf("invalid length: %v", n)
	}
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return NewCodeVerifierBytes(b), nil
}

func NewCodeVerifierBytes(b []byte) string {
	return base64.RawURLEncoding.EncodeToString(b)
}

func CodeChallengeS256(v string) string {
	h := sha256.Sum256([]byte(v))
	return base64.RawURLEncoding.EncodeToString(h[:])
}
