package internal

import (
	"encoding/base64"
	"testing"
)

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

func TestIsValidBase64StdAlphabet(t *testing.T) {
	out := base64.StdEncoding.EncodeToString([]byte("hello world"))
	if !IsValidBase64StdAlphabet(out) {
		t.Fatalf("expected %s to return true", out)
	}

	if IsValidBase64StdAlphabet("dGghaXMhM3c1MTNhc2YudGghaXMh.dGghaXMhM3c1MTNhc2YudGghaXMh") {
		t.Fatal("expected dGghaXMhM3c1MTNhc2YudGghaXMh.dGghaXMhM3c1MTNhc2YudGghaXMh to return false")
	}
}

func TestIsValidBase64UrlSafeAlphabet(t *testing.T) {
	urlout := base64.URLEncoding.EncodeToString([]byte("this is a test"))
	if !IsValidBase64UrlSafeAlphabet(urlout) {
		t.Fatalf("expected %s to return true", urlout)
	}

	if IsValidBase64UrlSafeAlphabet("aGV+sbG8gd29ybGQ=") {
		t.Fatal("expected aGV+sbG8gd29ybGQ= to return false")
	}
}

func TestIsValidBase64BAlphabet(t *testing.T) {
	if !IsValidBase64BAlphabet("aGV.sbGgd29ybGQ=") {
		t.Fatal("expected aGV.sbGgd29ybGQ= to return true")
	}

	if IsValidBase64BAlphabet("aGV+sbG8gd29ybGQ=") {
		t.Fatal("expected aGV+sbG8gd29ybGQ= to return false")
	}
}
