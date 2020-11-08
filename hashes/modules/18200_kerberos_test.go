package modules

import (
	"testing"

	"github.com/tankbusta/hashvalidate/tokenizer"
)

func Test_Kerberos18200Type(t *testing.T) {
	hashType := new(kerberos_18200)
	checkBasicAPIs(t, hashType)

	dynamicTokens, err := hashType.GenerateTokens(hashType.Example(), hashType.Tokens())
	if err != nil {
		t.Fatalf("Unexpected error generating tokens: %s", err)
	}

	tokens, err := tokenizer.Tokenize(
		hashType.Example(),
		dynamicTokens,
	)
	if err != nil {
		t.Fatalf("Unexpected error validating %s hashes: %s", hashType.Name(), err)
	}

	compareTokens(
		t,
		[]tokenizerComparison{
			{
				Buffer: "$krb5asrep$23$",
				Length: 14,
			},
			{
				Buffer: "user@domain.com:",
				Length: 16,
			},
			{
				Buffer: "3e156ada591263b8aab0965f5aebd837",
				Length: 32,
			},
			{
				Buffer: "007497cb51b6c8116d6407a782ea0e1c5402b17db7afa6b05a6d30ed164a9933c754d720e279c6c573679bd27128fe77e5fea1f72334c1193c8ff0b370fadc6368bf2d49bbfdba4c5dccab95e8c8ebfdc75f438a0797dbfb2f8a1a5f4c423f9bfc1fea483342a11bd56a216f4d5158ccc4b224b52894fadfba3957dfe4b6b8f5f9f9fe422811a314768673e0c924340b8ccb84775ce9defaa3baa0910b676ad0036d13032b0dd94e3b13903cc738a7b6d00b0b3c210d1f972a6c7cae9bd3c959acf7565be528fc179118f28c679f6deeee1456f0781eb8154e18e49cb27b64bf74cd7112a0ebae2102ac",
				Length: 468,
			},
		},
		tokens,
	)
}
