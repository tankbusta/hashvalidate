package modules

import (
	"testing"

	"github.com/tankbusta/hashvalidate/tokenizer"
)

func Test_Kerberos19800Type(t *testing.T) {
	hashType := new(kerberos_19800)
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
				Buffer: "$krb5pa$17$",
				Length: 11,
			},
			{
				Buffer: "hashcat",
				Length: 7,
			},
			{
				Buffer: "HASHCATDOMAIN.COM",
				Length: 17,
			},
			{
				Buffer: "a17776abe5383236c58582f515843e029ecbff43706d177651b7b6cdb2713b17597ddb35b1c9c470c281589fd1d51cca125414d19e40e333",
				Length: 112,
			},
		},
		tokens,
	)
}
