package modules

import (
	"testing"

	"github.com/tankbusta/hashvalidate/tokenizer"
)

func Test_Kerberos13100Type(t *testing.T) {
	hashType := new(kerberos_13100)
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
				Buffer: "$krb5tgs$23$",
				Length: 12,
			},
			{
				Buffer: "*user$realm$test/spn*$",
				Length: 22,
			},
			{
				Buffer: "b548e10f5694ae018d7ad63c257af7dc",
				Length: 32,
			},
			{
				Buffer: "35e8e45658860bc31a859b41a08989265f4ef8afd75652ab4d7a30ef151bf6350d879ae189a8cb769e01fa573c6315232b37e4bcad9105520640a781e5fd85c09615e78267e494f433f067cc6958200a82f70627ce0eebc2ac445729c2a8a0255dc3ede2c4973d2d93ac8c1a56b26444df300cb93045d05ff2326affaa3ae97f5cd866c14b78a459f0933a550e0b6507bf8af27c2391ef69fbdd649dd059a4b9ae2440edd96c82479645ccdb06bae0eead3b7f639178a90cf24d9a",
				Length: 374,
			},
		},
		tokens,
	)
}
