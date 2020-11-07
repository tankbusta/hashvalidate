package modules

import (
	"github.com/tankbusta/hashvalidate/hashes"
	"github.com/tankbusta/hashvalidate/tokenizer"
)

type sha1_100 struct{}

func init() {
	hashes.Register(100, sha1_100{})
}

func (s sha1_100) Name() string { return "SHA1" }

func (s sha1_100) Example() string {
	return "b89eaac7e61417341b710b727768294d0e6a277b"
}

func (s sha1_100) Type() int { return 100 }

func (s sha1_100) Tokens() []tokenizer.Token {
	return []tokenizer.Token{
		{
			LengthMin:  40,
			LengthMax:  40,
			Attributes: tokenizer.VerifyLength | tokenizer.VerifyHex,
		},
	}
}
