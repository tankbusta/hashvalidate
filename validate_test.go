package hashvalidate

import (
	"strings"
	"testing"

	"github.com/tankbusta/hashvalidate/hashes"
	"github.com/tankbusta/hashvalidate/tokenizer"
)

func TestValidateEmbeddedHashes(t *testing.T) {
	types := hashes.GetTypes()

	for _, hashType := range types {
		t.Run(hashType.Name(), func(subT *testing.T) {
			if err := ValidateHash(hashType.Type(), hashType.Example()); err != nil {
				subT.Fatalf("Failed to validate %s: %s", hashType.Name(), err.Error())
			}
		})
	}
}

func TestValidatorErrors(t *testing.T) {
	for _, test := range []struct {
		Name                string
		Input               string
		ExpectedErrContains string
		Tokens              []tokenizer.Token
	}{
		{
			Name:                "SignatureMismatch",
			Input:               "$foo!",
			ExpectedErrContains: "signature mismatch",
			Tokens: []tokenizer.Token{
				{
					Attributes: tokenizer.VerifySignature,
					Signatures: []tokenizer.Signature{
						{
							Expected: "$foo$",
						},
					},
				},
			},
		},
		{
			Name:                "MinLenMismatch",
			Input:               "foo",
			ExpectedErrContains: "failed minlen validation",
			Tokens: []tokenizer.Token{
				{
					Attributes: tokenizer.VerifyLength,
					LengthMin:  4,
				},
			},
		},
		{
			Name:                "MaxLenMismatch",
			Input:               "foo",
			ExpectedErrContains: "failed maxlen validation",
			Tokens: []tokenizer.Token{
				{
					Attributes: tokenizer.VerifyLength,
					LengthMax:  1,
				},
			},
		},
		{
			Name:                "Digits",
			Input:               "1foo",
			ExpectedErrContains: "failed digit validation",
			Tokens: []tokenizer.Token{
				{
					Attributes: tokenizer.VerifyDigit,
				},
			},
		},
		{
			Name:                "HexMismatch",
			Input:               "izs3b52063cd84097a65d1633f5c74f5",
			ExpectedErrContains: "failed hex validation",
			Tokens: []tokenizer.Token{
				{
					Attributes: tokenizer.VerifyLength | tokenizer.VerifyHex,
					LengthMin:  32,
					LengthMax:  32,
				},
			},
		},
		{
			Name:                "FloatMismatch",
			Input:               "123.4578",
			ExpectedErrContains: "failed float validation",
			Tokens: []tokenizer.Token{
				{
					Attributes: tokenizer.VerifyFloat,
				},
			},
		},
		{
			Name:                "Base64StdAlphabetMismatch",
			Input:               "abcde.",
			ExpectedErrContains: "failed base64a validation",
			Tokens: []tokenizer.Token{
				{
					Attributes: tokenizer.VerifyBase64A,
				},
			},
		},
		{
			Name:                "Base64BMismatch",
			Input:               "abcde_",
			ExpectedErrContains: "failed base64b validation",
			Tokens: []tokenizer.Token{
				{
					Attributes: tokenizer.VerifyBase64B,
				},
			},
		},
		{
			Name:                "Base64UrlSafeAlphabetMismatch",
			Input:               "abcde.",
			ExpectedErrContains: "failed base64c (urlsafe) validation",
			Tokens: []tokenizer.Token{
				{
					Attributes: tokenizer.VerifyBase64C,
				},
			},
		},
		{
			Name:                "TokenLenMismatch",
			Input:               "$test$",
			ExpectedErrContains: "Token length mismatch",
			Tokens: []tokenizer.Token{
				{
					Attributes: tokenizer.FixedLength,
					Separator:  "$",
					Length:     3,
				},
			},
		},
	} {
		t.Run(test.Name, func(subT *testing.T) {
			err := ValidateHashWithTokens(test.Input, test.Tokens)
			if err == nil {
				subT.Fatalf("Expected error in %s but got nil", test.Name)
			}

			if !strings.Contains(err.Error(), test.ExpectedErrContains) {
				subT.Fatalf("Expected error to contain `%s` but got `%s` ", test.ExpectedErrContains, err.Error())
			}
		})
	}
}

func TestValidateHashInvalidID(t *testing.T) {
	if err := ValidateHash(99999999, "8743b52063cd84097a65d1633f5c74f5"); err == nil {
		t.Fatal("Expected ValidateHash to fail")
	}
}
