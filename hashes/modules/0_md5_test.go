package modules

import (
	"testing"

	"github.com/tankbusta/hashvalidate/tokenizer"
)

func Test_MD50Type(t *testing.T) {
	hashType := new(md50)

	tokens, err := tokenizer.Tokenize(
		hashType.Example(),
		hashType.Tokens(),
	)
	if err != nil {
		t.Fatalf("Unexpected error validating MD5 hashes: %s", err)
	}

	compareTokens(
		t,
		[]tokenizerComparison{
			{
				Buffer: hashType.Example(),
				Length: 32,
			},
		},
		tokens,
	)
}
