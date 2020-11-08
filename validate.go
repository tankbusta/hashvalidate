package hashvalidate

import (
	"fmt"

	"github.com/tankbusta/hashvalidate/hashes"
	"github.com/tankbusta/hashvalidate/internal"
	"github.com/tankbusta/hashvalidate/tokenizer"

	// Load embedded hash types
	_ "github.com/tankbusta/hashvalidate/hashes/modules"
)

// ValidateHash ensures the provided hash tokenizes and validates properly based on the provided hash type.
// An error is returned if the hash fails
func ValidateHash(hashTypeID int, hash string) error {
	validator, err := hashes.Open(hashTypeID)
	if err != nil {
		return err
	}

	if customTokens, ok := validator.(hashes.IDynamicToken); ok {
		tokens, err := customTokens.GenerateTokens(hash, validator.Tokens())
		if err != nil {
			return err
		}

		return ValidateHashWithTokens(hash, tokens)
	}

	return ValidateHashWithTokens(hash, validator.Tokens())
}

// ValidateHashWithTokens ensures the provided hash and list of tokens validates correctly.
// An error is returned if the hash fails
func ValidateHashWithTokens(hash string, tokens []tokenizer.Token) error {
	tokens, err := tokenizer.Tokenize(hash, tokens)
	if err != nil {
		return err
	}

	// Verify Tokens
	for i, token := range tokens {
		tokenBuff := token.Buffer

		// Signature Verification
		if token.Attributes.Has(tokenizer.VerifySignature) {
			matched := false

		VerifySig:
			for _, sig := range token.Signatures {
				if tokenBuff == sig.Expected {
					matched = true
					break VerifySig
				}
			}

			if !matched {
				return fmt.Errorf("Token %d failed validation: signature mismatch", i)
			}
		}

		// Min/Max Length Verification
		if token.Attributes.Has(tokenizer.VerifyLength) {
			if token.Length < token.LengthMin {
				return fmt.Errorf("Token %d failed minlen validation: length mismatch got %d expected %d", i, token.Length, token.LengthMin)
			}

			if token.Length > token.LengthMax {
				return fmt.Errorf("Token %d failed maxlen validation: length mismatch got %d expected %d", i, token.Length, token.LengthMax)
			}
		}

		// All Digit Verification
		if token.Attributes.Has(tokenizer.VerifyDigit) && !internal.IsValidDigit(tokenBuff) {
			return fmt.Errorf("Token %d failed digit validation", i)
		}

		// Float Verification
		if token.Attributes.Has(tokenizer.VerifyFloat) && !internal.IsValidHex(tokenBuff) {
			return fmt.Errorf("Token %d failed float validation", i)
		}

		// All Hex Characters Verification
		if token.Attributes.Has(tokenizer.VerifyHex) && !internal.IsValidHex(tokenBuff) {
			return fmt.Errorf("Token %d failed hex validation", i)
		}

		// All Characters standard base64 alphabet
		if token.Attributes.Has(tokenizer.VerifyBase64A) && !internal.IsValidBase64StdAlphabet(tokenBuff) {
			return fmt.Errorf("Token %d failed base64a validation", i)
		}

		if token.Attributes.Has(tokenizer.VerifyBase64B) && !internal.IsValidBase64BAlphabet(tokenBuff) {
			return fmt.Errorf("Token %d failed base64b validation", i)
		}

		// urlsafe base64 alphabet verification
		if token.Attributes.Has(tokenizer.VerifyBase64C) && !internal.IsValidBase64UrlSafeAlphabet(tokenBuff) {
			return fmt.Errorf("Token %d failed base64c (urlsafe) validation", i)
		}
	}

	return nil
}
