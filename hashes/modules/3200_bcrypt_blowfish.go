package modules

import (
	"github.com/tankbusta/hashvalidate/hashes"
	"github.com/tankbusta/hashvalidate/tokenizer"
)

type bcrypt3200 struct{}

func init() {
	hashes.Register(3200, bcrypt3200{})
}

func (s bcrypt3200) Name() string { return "bcrypt $2*$, Blowfish (Unix)" }

func (s bcrypt3200) Example() string {
	return "$2a$05$MBCzKhG1KhezLh.0LRa0Kuw12nLJtpHy6DIaU.JAnqJUDYspHC.Ou"
}

func (s bcrypt3200) Type() int { return 3200 }

func (s bcrypt3200) Tokens() []tokenizer.Token {
	return []tokenizer.Token{
		{
			Length:     4,
			Attributes: tokenizer.FixedLength | tokenizer.VerifySignature,
			Signatures: []tokenizer.Signature{
				{
					Expected: "$2a$",
				},
				{
					Expected: "$2b$",
				},
				{
					Expected: "$2x$",
				},
				{
					Expected: "$2y$",
				},
			},
		},
		{
			LengthMin:  2,
			LengthMax:  2,
			Separator:  "$",
			Attributes: tokenizer.VerifyLength | tokenizer.VerifyDigit,
		},
		{
			Length:     22,
			Separator:  "$",
			Attributes: tokenizer.FixedLength | tokenizer.VerifyBase64B,
		},
		{
			Length:     31,
			Attributes: tokenizer.FixedLength | tokenizer.VerifyBase64B,
		},
	}
}
