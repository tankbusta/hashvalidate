package modules

import (
	"fmt"

	"github.com/tankbusta/hashvalidate/hashes"
	"github.com/tankbusta/hashvalidate/tokenizer"
)

type jwts struct {
	typ  int
	algo string
}

func init() {
	hashes.Register(16511, jwts{typ: 16511, algo: "HS256"})
	hashes.Register(16512, jwts{typ: 16512, algo: "HS384"})
	hashes.Register(16513, jwts{typ: 16513, algo: "HS512"})
}

func (s jwts) Name() string { return fmt.Sprintf("JWT (JSON Web Token) Algo %s", s.algo) }

func (s jwts) Example() string {
	return "eyJhbGciOiJIUzI1NiJ9.eyIzNDM2MzQyMCI6NTc2ODc1NDd9.f1nXZ3V_Hrr6ee-AFCTLaHRnrkiKmio2t3JqwL32guY"
}

func (s jwts) Type() int { return s.typ }

func (s jwts) Tokens() []tokenizer.Token {
	return []tokenizer.Token{
		{
			LengthMin:  1,
			LengthMax:  2047,
			Separator:  ".",
			Attributes: tokenizer.VerifyLength | tokenizer.VerifyBase64C,
		},
		{
			LengthMin:  0,
			LengthMax:  2047,
			Separator:  ".",
			Attributes: tokenizer.VerifyLength | tokenizer.VerifyBase64C,
		},
		{
			LengthMin:  43,
			LengthMax:  86,
			Attributes: tokenizer.VerifyLength | tokenizer.VerifyBase64C,
		},
	}
}
