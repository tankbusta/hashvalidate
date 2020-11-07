package modules

import (
	"testing"

	"github.com/tankbusta/hashvalidate/hashes"
	"github.com/tankbusta/hashvalidate/tokenizer"
)

type tokenizerComparison struct {
	Length int
	Buffer string
}

func checkBasicAPIs(t *testing.T, typ hashes.IHashType) {
	if typ.Type() < 0 {
		t.Fatal("Hash Type must be greater than or equal to 0")
	}

	if typ.Name() == "" {
		t.Fatal("Hash Name must not be an empty string")
	}
}

func compareTokens(t *testing.T, expected []tokenizerComparison, actual []tokenizer.Token) {
	if len(expected) != len(actual) {
		t.Fatal("Expected tokens does not equal the actual tokens")
	}

	for i := 0; i < len(expected); i++ {
		if expected[i].Length != actual[i].Length {
			t.Fatalf("Token %d does not have the correct length: got %d expected %d", i, actual[i].Length, expected[i].Length)
		}

		if expected[i].Buffer != actual[i].Buffer {
			t.Fatalf("Tokens at %d are not equal: got %s expected %s", i, actual[i].Buffer, expected[i].Buffer)
		}
	}
}
