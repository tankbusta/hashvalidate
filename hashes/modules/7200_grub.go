package modules

import (
	"github.com/tankbusta/hashvalidate/hashes"
	"github.com/tankbusta/hashvalidate/tokenizer"
)

type grub27200 struct{}

func init() {
	hashes.Register(7200, grub27200{})
}

func (s grub27200) Name() string { return "GRUB 2" }

func (s grub27200) Example() string {
	return "grub.pbkdf2.sha512.1024.03510507805003756325721848020561235456073188241051876082416068104377357018503082587026352628170170411053726157658716047762755750.aac26b18c2b0c44bcf56514d46aabd52eea097d9c95122722087829982e9dd957b2b641cb1e015d4df16a84d0571e96cf6d3de6361431bdeed4ddb0940f2425b"
}

func (s grub27200) Type() int { return 7200 }

func (s grub27200) Tokens() []tokenizer.Token {
	return []tokenizer.Token{
		{
			Length:     19,
			Attributes: tokenizer.FixedLength | tokenizer.VerifySignature,
			Signatures: []tokenizer.Signature{
				{
					Expected: "grub.pbkdf2.sha512.",
				},
			},
		},
		{
			LengthMin:  1,
			LengthMax:  6,
			Separator:  ".",
			Attributes: tokenizer.VerifyLength | tokenizer.VerifyDigit,
		},
		{
			LengthMin:  0,
			LengthMax:  256,
			Separator:  ".",
			Attributes: tokenizer.VerifyLength | tokenizer.VerifyHex,
		},
		{
			LengthMin:  128,
			LengthMax:  128,
			Attributes: tokenizer.VerifyLength | tokenizer.VerifyHex,
		},
	}
}
