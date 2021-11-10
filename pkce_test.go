package pkce

import (
	"testing"
)

var pkceTests = []struct {
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

// TestPKCE tests code verifier and code challenge functions.
func TestPKCE(t *testing.T) {
	for _, tt := range pkceTests {
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
