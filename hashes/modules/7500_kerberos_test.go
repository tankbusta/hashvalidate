package modules

import (
	"testing"

	"github.com/tankbusta/hashvalidate/tokenizer"
)

func Test_Kerberos7500Type(t *testing.T) {
	hashType := new(kerberos_7500)
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
				Buffer: "$krb5pa$23$",
				Length: 11,
			},
			{
				Buffer: "user",
				Length: 4,
			},
			{
				Buffer: "realm",
				Length: 5,
			},
			{
				Buffer: "salt",
				Length: 4,
			},
			{
				Buffer: "5cbb0c882a2b26956e81644edbdb746326f4f5f0e947144fb3095dffe4b4b03e854fc1d6",
				Length: 72,
			},
			{
				Buffer: "31323632303636373330383333353630",
				Length: 32,
			},
		},
		tokens,
	)
}
