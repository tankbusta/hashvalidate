package tokenizer

import (
	"errors"
	"fmt"
	"strings"
)

// Tokenize TODO
func Tokenize(input string, tokens []Token) ([]Token, error) {
	left := len(input)

	tokenIdx := 0
	// Handle all but the last token inside this loop
	for tokenIdx = 0; tokenIdx < len(tokens)-1; tokenIdx++ {
		if tokens[tokenIdx].Attributes.Has(FixedLength) {
			len := tokens[tokenIdx].Length
			if left < len {
				return nil, errors.New("Token length exception")
			}

			tokens[tokenIdx].Buffer = input[:len]

			// Move the input up
			left -= len
			input = input[len:]
		} else {
			if tokens[tokenIdx].Attributes.Has(OptionalRounds) {
				return nil, fmt.Errorf("Cannot validate %s: OptionalRounds not implemented", input)
			}

			var next int

			if tokens[tokenIdx].Attributes.Has(SeparatorFarthest) {
				next = strings.LastIndex(input, tokens[tokenIdx].Separator)
			} else {
				next = strings.Index(input, tokens[tokenIdx].Separator)
			}

			if next == -1 {
				return nil, fmt.Errorf("Separator Unmatched for %s", input)
			}

			tokens[tokenIdx].Length = next
			tokens[tokenIdx].Buffer = input[:next]

			// Move the input up
			input = input[next+1:] // +1 separator
			left -= next + 1       // +1 separator
		}
	}

	// Handle the last token
	if tokens[tokenIdx].Attributes.Has(FixedLength) {
		len := tokens[tokenIdx].Length

		if left != len {
			return nil, fmt.Errorf("Token length missmatch on %d", tokenIdx)
		}
	}

	tokens[tokenIdx].Length = left
	tokens[tokenIdx].Buffer = input

	return tokens, nil
}
