package modules

import (
	"github.com/tankbusta/hashvalidate/hashes"
	"github.com/tankbusta/hashvalidate/tokenizer"
)

type kerberos_7500 struct{}

func init() {
	hashes.Register(7500, kerberos_7500{})
}

func (s kerberos_7500) Name() string { return "Kerberos 5, etype 23, AS-REQ Pre-Auth" }

func (s kerberos_7500) Example() string {
	return "$krb5pa$23$user$realm$salt$5cbb0c882a2b26956e81644edbdb746326f4f5f0e947144fb3095dffe4b4b03e854fc1d631323632303636373330383333353630"
}

func (s kerberos_7500) Type() int { return 7500 }

func (s kerberos_7500) Tokens() []tokenizer.Token {
	return []tokenizer.Token{
		{
			Length:     11,
			Attributes: tokenizer.FixedLength | tokenizer.VerifySignature,
			Signatures: []tokenizer.Signature{
				{
					Expected: "$krb5pa$23$",
				},
			},
		},
		{
			LengthMin:  0,
			LengthMax:  64,
			Separator:  "$",
			Attributes: tokenizer.VerifyLength,
		},
		{
			LengthMin:  0,
			LengthMax:  64,
			Separator:  "$",
			Attributes: tokenizer.VerifyLength,
		},
		{
			LengthMin:  0,
			LengthMax:  128,
			Separator:  "$",
			Attributes: tokenizer.VerifyLength,
		},
		{
			Length:     72,
			Attributes: tokenizer.FixedLength | tokenizer.VerifyHex,
		},
		{
			Length:     32,
			Attributes: tokenizer.FixedLength | tokenizer.VerifyHex,
		},
	}
}
