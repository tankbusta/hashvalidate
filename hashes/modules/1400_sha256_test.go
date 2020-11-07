package modules

import (
	"testing"

	"github.com/tankbusta/hashvalidate/tokenizer"
)

func Test_SHA256Type(t *testing.T) {
	hashType := new(sha256_1400)
	checkBasicAPIs(t, hashType)

	tokens, err := tokenizer.Tokenize(
		hashType.Example(),
		hashType.Tokens(),
	)
	if err != nil {
		t.Fatalf("Unexpected error validating %s hashes: %s", hashType.Name(), err)
	}

	compareTokens(
		t,
		[]tokenizerComparison{
			{
				Buffer: hashType.Example(),
				Length: 64,
			},
		},
		tokens,
	)
}
