package modules

import (
	"testing"

	"github.com/tankbusta/hashvalidate/tokenizer"
)

func Test_BCrypt3200Type(t *testing.T) {
	hashType := new(bcrypt3200)

	tokens, err := tokenizer.Tokenize(
		hashType.Example(),
		hashType.Tokens(),
	)
	if err != nil {
		t.Fatalf("Unexpected error validating bcrypt2 hashes: %s", err)
	}

	compareTokens(
		t,
		[]tokenizerComparison{
			{
				Buffer: "$2a$",
				Length: 4,
			},
			{
				Buffer: "05",
				Length: 2,
			},
			{
				Buffer: "MBCzKhG1KhezLh.0LRa0Ku",
				Length: 22,
			},
			{
				Buffer: "w12nLJtpHy6DIaU.JAnqJUDYspHC.Ou",
				Length: 31,
			},
		},
		tokens,
	)
}
