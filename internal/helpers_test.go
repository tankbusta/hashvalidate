package internal

import "testing"

func TestIsValidHex(t *testing.T) {
	if !IsValidHex("deadbeef") {
		t.Fatal("expected deadbeef to return true")
	}

	if !IsValidHex("deadb3ef") {
		t.Fatal("expected deadb3ef to return true")
	}

	if IsValidHex("zoo") {
		t.Fatal("expected zoo to return false")
	}

	if !IsValidHex("123456") {
		t.Fatal("expected 123456 to return true")
	}
}

func TestIsValidDigit(t *testing.T) {
	if !IsValidDigit("1234567890") {
		t.Fatal("expected 1234567890 to return true")
	}

	if IsValidDigit("deadbeef") {
		t.Fatal("expected deadbeef to return false")
	}

	if IsValidDigit("1234a56") {
		t.Fatal("expected 1234a56 to return false")
	}
}

func TestIsValidFloat(t *testing.T) {
	if !IsValidFloat("1234.57890") {
		t.Fatal("expected 1234.57890 to return true")
	}

	if IsValidDigit("1234.dead") {
		t.Fatal("expected 1234.dead to return false")
	}
}
