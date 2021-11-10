package pkce

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

const (
	DefaultLength = 32
	MinLength     = 32
	MaxLength     = 96

	MethodPlain = "plain"
	MethodS256  = "S256"

	ParamCodeVerifier        = "code_verifier"
	ParamCodeChallenge       = "code_challenge"
	ParamCodeChallengeMethod = "code_challenge_method"
)

func NewCodeVerifier() string {
	v, err := NewCodeVerifierWithLength(DefaultLength)
	if err != nil {
		panic(err)
	}
	return v
}

func NewCodeVerifierWithLength(n int) (string, error) {
	if n < MinLength || n > MaxLength {
		return "", fmt.Errorf("invalid length: %v", n)
	}
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return NewCodeVerifierFromBytes(b), nil
}

func NewCodeVerifierFromBytes(b []byte) string {
	return base64.RawURLEncoding.EncodeToString(b)
}

func CodeChallengeS256(v string) string {
	h := sha256.Sum256([]byte(v))
	return base64.RawURLEncoding.EncodeToString(h[:])
}
