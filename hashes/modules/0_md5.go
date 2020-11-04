package modules

import (
	"github.com/tankbusta/hashvalidate/hashes"
	"github.com/tankbusta/hashvalidate/tokenizer"
)

type md50 struct{}

func init() {
	hashes.Register(0, md50{})
}

func (s md50) Name() string { return "MD5" }

func (s md50) Example() string {
	return "8743b52063cd84097a65d1633f5c74f5"
}

func (s md50) Type() int { return 0 }

func (s md50) Tokens() []tokenizer.Token {
	return []tokenizer.Token{
		{
			LengthMin:  32,
			LengthMax:  32,
			Attributes: tokenizer.VerifyLength | tokenizer.VerifyHex,
		},
	}
}
