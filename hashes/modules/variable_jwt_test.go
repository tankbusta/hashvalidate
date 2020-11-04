package modules

import (
	"testing"

	"github.com/tankbusta/hashvalidate/tokenizer"
)

func Test_VariableJWT(t *testing.T) {
	hashType := &jwts{typ: 16511}

	tokens, err := tokenizer.Tokenize(
		hashType.Example(),
		hashType.Tokens(),
	)
	if err != nil {
		t.Fatalf("Unexpected error validating JWT hashes: %s", err)
	}

	compareTokens(
		t,
		[]tokenizerComparison{
			{
				Buffer: "eyJhbGciOiJIUzI1NiJ9",
				Length: 20,
			},
			{
				Buffer: "eyIzNDM2MzQyMCI6NTc2ODc1NDd9",
				Length: 28,
			},
			{
				Buffer: "f1nXZ3V_Hrr6ee-AFCTLaHRnrkiKmio2t3JqwL32guY",
				Length: 43,
			},
		},
		tokens,
	)
}
