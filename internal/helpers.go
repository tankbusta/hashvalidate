package internal

type verifyByteFnc func(byte) bool

func verifyString(s string, verifyfnc verifyByteFnc) bool {
	for i := 0; i < len(s); i++ {
		if !verifyfnc(s[i]) {
			return false
		}
	}

	return true
}

func isRuneHex(c byte) bool {
	switch {
	case (c >= '0') && (c <= '9'),
		(c >= 'A') && (c <= 'F'),
		(c >= 'a') && (c <= 'f'):
		return true
	}

	return false
}

// IsValidHex returns true if the string contains all hexadecimal characters
func IsValidHex(s string) bool {
	return verifyString(s, isRuneHex)
}

func isRuneDigit(c byte) bool {
	return (c >= '0') && (c <= '9')
}

// IsValidDigit returns true if the string contains all digits
func IsValidDigit(s string) bool {
	return verifyString(s, isRuneDigit)
}

func isRuneFlotable(c byte) bool {
	return ((c >= '0') && (c <= '9') || c == '.')
}

// IsValidFloat returns true if the string is a float
func IsValidFloat(s string) bool {
	return verifyString(s, isRuneFlotable)
}

func isBase64StdAlphabet(c byte) bool {
	switch {
	case (c >= '0') && (c <= '9'),
		(c >= 'A') && (c <= 'Z'),
		(c >= 'a') && (c <= 'z'):
		return true
	}

	return false
}

func isBase64UrlSafeAlphabet(c byte) bool {
	switch {
	case (c >= '0') && (c <= '9'),
		(c >= 'A') && (c <= 'Z'),
		(c >= 'a') && (c <= 'z'),
		(c == '_' || c == '-' || c == '='):
		return true
	}

	return false
}

func isBase64BAlphabet(c byte) bool {
	switch {
	case (c >= '0') && (c <= '9'),
		(c >= 'A') && (c <= 'Z'),
		(c >= 'a') && (c <= 'z'),
		(c == '.' || c == '/' || c == '='):
		return true
	}

	return false
}

// IsValidBase64StdAlphabet returns true if the string contains all characters that fall within the standard base64 alphabet
func IsValidBase64StdAlphabet(s string) bool {
	return verifyString(s, isBase64StdAlphabet)
}

// IsValidBase64UrlSafeAlphabet returns true if the string contains all characters that fall within the base64 url safe alphabet
func IsValidBase64UrlSafeAlphabet(s string) bool {
	return verifyString(s, isBase64UrlSafeAlphabet)
}

// IsValidBase64BAlphabet TODO. Not sure what base64 alphabet has periods?
func IsValidBase64BAlphabet(s string) bool {
	return verifyString(s, isBase64BAlphabet)
}
