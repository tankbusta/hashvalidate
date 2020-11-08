package hashes

import (
	"testing"

	"github.com/tankbusta/hashvalidate/tokenizer"
)

type dummyType struct{}

func (s dummyType) Name() string { return "Dummy Hash" }

func (s dummyType) Example() string {
	return "foobar"
}

func (s dummyType) Type() int { return 9999999 }

func (s dummyType) Tokens() []tokenizer.Token {
	return []tokenizer.Token{
		{
			LengthMin:  6,
			LengthMax:  6,
			Attributes: tokenizer.VerifyLength,
		},
	}
}

func TestRegister(t *testing.T) {
	exampleType := dummyType{}

	Register(exampleType.Type(), exampleType)
	if len(GetTypes()) == 0 {
		t.Error("GetTypes should be at least 1")
	}

	// Register it again for a panic
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic on second registration")
		}
	}()

	Register(exampleType.Type(), exampleType)
}

func TestRegisterNil(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic on nil type")
		}
	}()

	Register(99999998, nil)
}

func TestOpen(t *testing.T) {
	exampleType := dummyType{}

	Register(99999997, exampleType)

	if _, err := Open(99999997); err != nil {
		t.Errorf("Unexpected err on open: %s", err.Error())
	}

	if _, err := Open(-1); err == nil {
		t.Errorf("Expected err on open with invalid hash type: %s", err.Error())
	}
}
