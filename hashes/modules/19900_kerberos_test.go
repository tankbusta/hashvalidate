package modules

import (
	"testing"

	"github.com/tankbusta/hashvalidate/tokenizer"
)

func Test_Kerberos19900Type(t *testing.T) {
	hashType := new(kerberos_19900)
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
				Buffer: "$krb5pa$18$",
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
				Buffer: "96c289009b05181bfd32062962740b1b1ce5f74eb12e0266cde74e81094661addab08c0c1a178882c91a0ed89ae4e0e68d2820b9cce69770",
				Length: 112,
			},
		},
		tokens,
	)
}
