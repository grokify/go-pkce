package pkce

import (
	"encoding/base64"
	"regexp"
	"testing"
)

func TestVerifier(t *testing.T) {
	rx := regexp.MustCompile(`^[0-9A-Za-z_-]*$`)
	cv := NewCodeVerifier()

	if !rx.MatchString(cv) {
		t.Errorf("invalid result for NewCodeVerifier() got [%s]", cv)
	}
	cvDec, err := base64.RawURLEncoding.DecodeString(cv)
	if err != nil {
		t.Errorf("error result for b64.RawURLEncoding.DecodeString(\"%s\") err [%s]", cv, err.Error())
	}
	if len(cvDec) != LengthDefault {
		t.Errorf("invalid length for NewCodeVerifier() want [%d], got [%d][%s]",
			LengthDefault, len(cvDec), cv)
	}
}

var verifierTests = []struct {
	len   int
	isErr bool
}{
	{32, false},
	{96, false},
	{0, true},
	{100, true},
}

func TestVerifierLength(t *testing.T) {
	rx := regexp.MustCompile(`^[0-9A-Za-z_-]*$`)
	for _, tt := range verifierTests {
		cv, err := NewCodeVerifierWithLength(tt.len)
		if err != nil && tt.isErr {
			continue
		}
		if err != nil {
			t.Errorf("invalid result for NewCodeVerifierWithLength(%d) error [%s]",
				tt.len, err.Error())
		}
		if !rx.MatchString(cv) {
			t.Errorf("invalid result for NewCodeVerifierWithLength(%d) got [%s]", tt.len, cv)
		}
		cvDec, err := base64.RawURLEncoding.DecodeString(cv)
		if err != nil {
			t.Errorf("error result for b64.RawURLEncoding.DecodeString(\"%s\") err [%s]", cv, err.Error())
		}
		if len(cvDec) != tt.len {
			t.Errorf("invalid length for NewCodeVerifierWithLength(%d) want [%d], got [%d][%s]",
				tt.len, tt.len, len(cvDec), cv)
		}
	}
}

var challengeTests = []struct {
	verifierBytes  []byte
	verifierString string
	challengeS256  string
}{
	{
		[]byte{116, 24, 223, 180, 151, 153, 224, 37, 79, 250, 96, 125, 216, 173,
			187, 186, 22, 212, 37, 77, 105, 214, 191, 240, 91, 88, 5, 88, 83,
			132, 141, 121},
		"dBjftJeZ4CVP-mB92K27uhbUJU1p1r_wW1gFWFOEjXk",
		"E9Melhoa2OwvFrEMTJguCHaoeK1t8URWbuGJSstw-cM"},
}

// TestChallenge tests code verifier and code challenge functions.
func TestChallenge(t *testing.T) {
	for _, tt := range challengeTests {
		cv := NewCodeVerifierFromBytes(tt.verifierBytes)
		if cv != tt.verifierString {
			t.Errorf("invalid result for NewCodeVerifierFromBytes want [%s], got [%s]",
				tt.verifierString, cv)
		}
		ccS256 := CodeChallengeS256(cv)
		if ccS256 != tt.challengeS256 {
			t.Errorf("invalid result for CodeChallengeS256 want [%s], got [%s]",
				tt.challengeS256, ccS256)
		}
	}
}
