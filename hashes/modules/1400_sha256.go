package modules

import (
	"github.com/tankbusta/hashvalidate/hashes"
	"github.com/tankbusta/hashvalidate/tokenizer"
)

type sha256_1400 struct{}

func init() {
	hashes.Register(1400, sha256_1400{})
}

func (s sha256_1400) Name() string { return "SHA-256" }

func (s sha256_1400) Example() string {
	return "127e6fbfe24a750e72930c220a8e138275656b8e5d8f48a98c3c92df2caba935"
}

func (s sha256_1400) Type() int { return 1400 }

func (s sha256_1400) Tokens() []tokenizer.Token {
	return []tokenizer.Token{
		{
			LengthMin:  64,
			LengthMax:  64,
			Attributes: tokenizer.VerifyLength | tokenizer.VerifyHex,
		},
	}
}
