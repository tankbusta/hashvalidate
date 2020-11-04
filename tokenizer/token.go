package tokenizer

// TokenAttr indicates how this token should be validated
type TokenAttr uint16

const (
	// FixedLength validates the token is of a fixed length as defined by the `Length` propery of a Token
	FixedLength TokenAttr = 1 << iota
	// SeparatorFarthest instructs the validator to look for the separator at the farthest end of the string
	SeparatorFarthest
	// OptionalRounds TODO
	OptionalRounds
	// VerifySignature instructs the validator to do a string comparison on a given token
	VerifySignature
	// VerifyLength instructs the validator to check the token's min/max length
	VerifyLength
	// VerifyDigit instructs the validator to ensure the token contains all digits
	VerifyDigit
	// VerifyFloat instructs the validator to ensure the token is a float
	VerifyFloat
	// VerifyHex instructs the validator to check that the token contains all hexadecimal values
	VerifyHex
	// VerifyBase64A instructs the validator to verify the token contains all characters that are apart of
	// the standard base64 alphabet
	VerifyBase64A
	// VerifyBase64B TODO
	VerifyBase64B
	// VerifyBase64C instructs the validator to verify the token contains all characters that are apart of
	// a urlsafe base64 alphabet
	VerifyBase64C
)

// Has indicates if the attribute bitmask has an attribute
func (s TokenAttr) Has(attr TokenAttr) bool {
	return s&attr != 0
}

type (
	// Token TODO
	Token struct {
		Separator  string
		Length     int
		LengthMin  int
		LengthMax  int
		Attributes TokenAttr
		Signatures []Signature
		Buffer     string
	}

	// Signature contains an expected value that is expected inside a token of a hash
	Signature struct {
		Expected string
	}
)
