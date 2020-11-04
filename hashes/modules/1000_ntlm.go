package modules

import (
	"github.com/tankbusta/hashvalidate/hashes"
	"github.com/tankbusta/hashvalidate/tokenizer"
)

type ntlm1000 struct{}

func init() {
	hashes.Register(1000, ntlm1000{})
}

func (s ntlm1000) Name() string { return "NTLM" }

func (s ntlm1000) Example() string {
	return "b4b9b02e6f09a9bd760f388b67351e2b"
}

func (s ntlm1000) Type() int { return 1000 }

func (s ntlm1000) Tokens() []tokenizer.Token {
	return []tokenizer.Token{
		{
			LengthMin:  32,
			LengthMax:  32,
			Attributes: tokenizer.VerifyLength | tokenizer.VerifyHex,
		},
	}
}
