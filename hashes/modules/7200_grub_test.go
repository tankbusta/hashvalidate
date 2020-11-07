package modules

import (
	"testing"

	"github.com/tankbusta/hashvalidate/tokenizer"
)

func Test_Grub7200Type(t *testing.T) {
	hashType := new(grub27200)
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
				Buffer: "grub.pbkdf2.sha512.",
				Length: 19,
			},
			{
				Buffer: "1024",
				Length: 4,
			},
			{
				Buffer: "03510507805003756325721848020561235456073188241051876082416068104377357018503082587026352628170170411053726157658716047762755750",
				Length: 128,
			},
			{
				Buffer: "aac26b18c2b0c44bcf56514d46aabd52eea097d9c95122722087829982e9dd957b2b641cb1e015d4df16a84d0571e96cf6d3de6361431bdeed4ddb0940f2425b",
				Length: 128,
			},
		},
		tokens,
	)
}
